package tests

// THESE TESTS ARE MANUALLY CREATED END-TO-END TESTS FOR THE ENTITIES API.
//
// Entity Create and Update operations are eventually consistent,
// which means these entities take more time than the request allows for creation.
// This requires us to write retry logic to wait for operations to complete,
// preventing us from generating these tests from the tests.arazzo file,
// since there's no way to "wait" for entity creation,
// or a way to implement retry logic in generated tests.

import (
	"context"
	"net/http"
	"testing"

	"github.com/solarwinds/swo-sdk-go/swov1"
	"github.com/solarwinds/swo-sdk-go/swov1/models/components"
	"github.com/solarwinds/swo-sdk-go/swov1/models/operations"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	testEntityName = "swo-sdk-entities-e2e-crud-test"
	testEntityURL  = "https://swo-sdk-entities-test.com"
)

func TestSDK_EntitiesCrudLifecycle(t *testing.T) {
	ctx := context.Background()
	s := CreateTestClient("entities-e2e-crud")

	createWebsiteRes, err := s.Dem.CreateWebsite(ctx, components.DemWebsite{
		Name: testEntityName,
		URL:  testEntityURL,
		AvailabilityCheckSettings: &components.AvailabilityCheckSettings{
			TestFrom: components.DemTestFrom{
				Type:   components.TypeRegion,
				Values: []string{TestRegion},
			},
			TestIntervalInSeconds: DefaultTestInterval,
			Protocols:             []components.DemWebsiteProtocol{components.DemWebsiteProtocolHTTPS},
		},
	})
	require.NoError(t, err, "Failed to create test website for entity operations")
	require.Equal(t, http.StatusCreated, createWebsiteRes.HTTPMeta.Response.StatusCode, "Create website returned unexpected status")
	require.NotNil(t, createWebsiteRes.CommonEntityID, "Create response should contain entity ID")

	entityID := createWebsiteRes.CommonEntityID.ID
	t.Logf("Created entity with ID: %s", entityID)
	// Cleanup regardless of test outcome
	defer func() {
		_, err := s.Dem.DeleteWebsite(ctx, operations.DeleteWebsiteRequest{
			EntityID: entityID,
		})
		if err != nil {
			t.Logf("Failed to cleanup entity %s: %v", entityID, err)
		} else {
			t.Logf("Deleted entity with ID: %s", entityID)
		}
	}()

	waitForEntityAvailability(ctx, t, s, entityID)

	getEntityRes, err := s.Entities.GetEntityByID(ctx, operations.GetEntityByIDRequest{
		ID: entityID,
	})
	require.NoError(t, err, "Failed to get entity by ID")
	require.Equal(t, http.StatusOK, getEntityRes.HTTPMeta.Response.StatusCode, "Get entity returned unexpected status")
	require.NotNil(t, getEntityRes.EntitiesEntity, "Get entity response should contain entity data")

	assert.Equal(t, entityID, getEntityRes.EntitiesEntity.ID, "Entity ID should match")
	assert.Equal(t, WebsiteEntityType, getEntityRes.EntitiesEntity.Type, "Entity type should be Website")
	if getEntityRes.EntitiesEntity.Name != nil {
		assert.Equal(t, testEntityName, *getEntityRes.EntitiesEntity.Name, "Entity name should match")
	}

	listEntitiesRes, err := s.Entities.ListEntities(ctx, operations.ListEntitiesRequest{
		Type:     WebsiteEntityType,
		PageSize: swov1.Int(1),
	})
	require.NoError(t, err, "Failed to list entities")
	require.Equal(t, http.StatusOK, listEntitiesRes.HTTPMeta.Response.StatusCode, "List entities returned unexpected status")
	require.NotNil(t, listEntitiesRes.Object, "List entities response should contain object")
	require.NotNil(t, listEntitiesRes.Object.Entities, "List entities response should contain entities array")

	hasOneEntity := len(listEntitiesRes.Object.Entities) == 1
	assert.True(t, hasOneEntity, "Entities list should have exactly 1 entity")

	updatedDisplayName := "Updated Display Name for " + testEntityName
	updateTags := map[string]*string{
		"environment": swov1.String("test"),
		"team":        swov1.String("sdk-testing"),
		"updated":     swov1.String("true"),
	}

	updateRes, err := s.Entities.UpdateEntityByID(ctx, operations.UpdateEntityByIDRequest{
		ID: entityID,
		EntitiesEntity: components.EntitiesEntityInput{
			DisplayName: &updatedDisplayName,
			Tags:        updateTags,
		},
	})

	if err != nil {
		t.Logf("Update operation encountered error: %v", err)
	} else {
		if updateRes.HTTPMeta.Response.StatusCode == http.StatusAccepted {
			t.Logf("Updated entity with ID: %s", entityID)
			waitForEntityUpdate(ctx, t, s, entityID, updatedDisplayName, updateTags)
			getRes, err := s.Entities.GetEntityByID(ctx, operations.GetEntityByIDRequest{
				ID: entityID,
			})
			require.NoError(t, err, "Failed to get entity")
			require.Equal(t, http.StatusOK, getRes.HTTPMeta.Response.StatusCode, "Entity returned with unexpected status")
			assert.Equal(t, updatedDisplayName, *getRes.EntitiesEntity.DisplayName, "DisplayName should be updated")
		}
	}

	t.Log("Entities End-to-end test completed successfully.")
}

func waitForEntityAvailability(ctx context.Context, t *testing.T, s *swov1.Swo, entityID string) {
	assert.Eventually(t, func() bool {
		getEntityRes, err := s.Entities.GetEntityByID(ctx, operations.GetEntityByIDRequest{
			ID: entityID,
		})

		if err != nil || getEntityRes.HTTPMeta.Response.StatusCode != http.StatusOK {
			return false
		}

		assert.Equal(t, entityID, getEntityRes.EntitiesEntity.ID)
		assert.Equal(t, WebsiteEntityType, getEntityRes.EntitiesEntity.Type)
		return true
	}, DefaultTimeout, DefaultRetryInterval, "Entity %s is available", entityID)
}

func waitForEntityUpdate(ctx context.Context, t *testing.T, s *swov1.Swo, entityID, expectedDisplayName string, expectedTags map[string]*string) {
	assert.Eventually(t, func() bool {
		getEntityRes, err := s.Entities.GetEntityByID(ctx, operations.GetEntityByIDRequest{
			ID: entityID,
		})

		if err != nil || getEntityRes.HTTPMeta.Response.StatusCode != http.StatusOK {
			return false
		}

		entity := getEntityRes.EntitiesEntity

		if entity.DisplayName != nil && *entity.DisplayName == expectedDisplayName {
			tagsUpdated := true
			for key, expectedValue := range expectedTags {
				if entity.Tags == nil {
					tagsUpdated = false
					break
				}
				if actualValue, exists := entity.Tags[key]; !exists {
					tagsUpdated = false
					break
				} else if actualValue == nil || *actualValue != *expectedValue {
					tagsUpdated = false
					break
				}
			}
			return tagsUpdated
		} else {
			return false
		}
	}, UpdateTimeout, DefaultUpdateRetryInterval, "Updated entity data should be created within 2 minutes")
}
