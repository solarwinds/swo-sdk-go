// THESE TESTS ARE MANUALLY CREATED END-TO-END TESTS FOR THE DEM URI API.
//
// URI Create and Update operations are eventually consistent,
// which means these entities take more time than the request allows for creation.
// This requires us to write retry logic to wait for operations to complete,
// preventing us from generating these tests from the tests.arazzo file,
// since there's no way to "wait" for entity creation,
// or a way to implement retry logic in generated tests.

package tests

import (
	"context"
	"net/http"
	"testing"
	"time"

	"github.com/solarwinds/swo-sdk-go/swov1"
	"github.com/solarwinds/swo-sdk-go/swov1/models/components"
	"github.com/solarwinds/swo-sdk-go/swov1/models/operations"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	uriEntityName = TestEntityName + "-uri-e2e"
)

func TestSDK_DemUriCrudLifecycle(t *testing.T) {
	ctx := context.Background()
	s := CreateTestClient("uri-e2e-crud")

	createURIRes, err := s.Dem.CreateURI(ctx, components.DemURI{
		Name:       uriEntityName,
		IPOrDomain: OriginalTestDomain,
		AvailabilityCheckSettings: components.DemURIAvailabilityCheckSettingsInput{
			TestFrom: components.DemTestFrom{
				Type:   components.TypeRegion,
				Values: []string{TestRegion},
			},
			TestIntervalInSeconds: DefaultTestInterval,
			Ping: &components.Ping{
				Enabled: true,
			},
		},
	})
	require.NoError(t, err, "Failed to create test URI")
	require.Equal(t, http.StatusCreated, createURIRes.HTTPMeta.Response.StatusCode, "Create URI returned unexpected status")
	require.NotNil(t, createURIRes.CommonEntityID, "Create response should contain entity ID")

	entityID := createURIRes.CommonEntityID.ID
	t.Logf("Created URI entity with ID: %s", entityID)
	// Ensure cleanup regardless of test outcome
	defer func() {
		_, err := s.Dem.DeleteURI(ctx, operations.DeleteURIRequest{
			EntityID: entityID,
		})
		if err != nil {
			t.Logf("Failed to cleanup URI %s: %v", entityID, err)
		} else {
			t.Logf("Deleted URI entity with ID: %s", entityID)
		}
	}()

	uri := waitForURI(t, ctx, s, entityID, DefaultTimeout)
	verifyURI(t, uri, uriEntityName, OriginalTestDomain)

	updateURIRes, err := s.Dem.UpdateURI(ctx, operations.UpdateURIRequest{
		EntityID: entityID,
		DemURI: components.DemURI{
			Name:       uriEntityName,
			IPOrDomain: UpdatedTestDomain,
			AvailabilityCheckSettings: components.DemURIAvailabilityCheckSettingsInput{
				TestFrom: components.DemTestFrom{
					Type:   components.TypeRegion,
					Values: []string{TestRegion},
				},
				TestIntervalInSeconds: UpdatedTestInterval,
				Ping: &components.Ping{
					Enabled: true,
				},
			},
		},
	})

	if err != nil {
		t.Logf("Update operation encountered error: %v", err)
	} else {
		statusCode := updateURIRes.HTTPMeta.Response.StatusCode
		if statusCode == http.StatusOK || statusCode == http.StatusNoContent {
			t.Logf("Updated URI entity with ID: %s", entityID)
			waitForUpdatedURI(t, ctx, s, entityID, UpdatedTestDomain, UpdatedTestInterval, UpdateTimeout)
		}
	}

	deleteRes, err := s.Dem.DeleteURI(ctx, operations.DeleteURIRequest{
		EntityID: entityID,
	})
	require.NoError(t, err, "Failed to delete test URI")
	require.Equal(t, http.StatusNoContent, deleteRes.HTTPMeta.Response.StatusCode, "Delete URI returned unexpected status")

	getURIRes, err := s.Dem.GetURI(ctx, operations.GetURIRequest{
		EntityID: entityID,
	})
	if err != nil {
		t.Logf("Expected error after deletion: %v", err)
	} else {
		assert.Equal(t, http.StatusNotFound, getURIRes.HTTPMeta.Response.StatusCode, "URI should not be found after deletion")
	}

	t.Log("DEM URI End-to-end test completed.")
}

func waitForURI(t *testing.T, ctx context.Context, s *swov1.Swo, entityID string, timeout time.Duration) *components.DemGetURIResponse {
	var uri *components.DemGetURIResponse

	assert.Eventually(t, func() bool {
		getURIRes, err := s.Dem.GetURI(ctx, operations.GetURIRequest{
			EntityID: entityID,
		})

		if err != nil {
			return false
		}

		if getURIRes.HTTPMeta.Response.StatusCode == http.StatusNotFound {
			return false
		}

		if getURIRes.HTTPMeta.Response.StatusCode != http.StatusOK {
			return false
		}

		uri = getURIRes.DemGetURIResponse
		return true

	}, timeout, DefaultRetryInterval, "URI should become available within %v", timeout)

	return uri
}

func verifyURI(t *testing.T, uri *components.DemGetURIResponse, expectedName, expectedDomain string) {
	assert.Equal(t, expectedName, uri.Name)
	assert.Equal(t, expectedDomain, uri.IPOrDomain)
	assert.Contains(t, ValidURIStatuses, uri.Status)
}

func waitForUpdatedURI(t *testing.T, ctx context.Context, s *swov1.Swo, entityID, expectedDomain string, expectedInterval float64, timeout time.Duration) {
	assert.Eventually(t, func() bool {
		getURIRes, err := s.Dem.GetURI(ctx, operations.GetURIRequest{
			EntityID: entityID,
		})

		if err != nil || getURIRes.HTTPMeta.Response.StatusCode != http.StatusOK {
			return false
		}

		uriData := getURIRes.DemGetURIResponse

		if uriData.IPOrDomain == expectedDomain {
			if uriData.AvailabilityCheckSettings.TestIntervalInSeconds == expectedInterval {
				return true
			}
		}
		return false

	}, timeout, DefaultUpdateRetryInterval, "Updated URI data should be created within %v", timeout)
}
