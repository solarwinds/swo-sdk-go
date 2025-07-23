// THESE TESTS ARE MANUALLY CREATED END-TO-END TESTS FOR THE DEM Websites API.
//
// Website Create and Update operations are eventually consistent,
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
	WebsiteAvailabilityTimeout = DefaultTimeout
	WebsiteUpdateTimeout       = UpdateTimeout
	RetryInterval              = DefaultRetryInterval
	UpdateRetryInterval        = DefaultUpdateRetryInterval

	TestWebsiteName = "swo-sdk-dem-websites-e2e-crud-test"
	OriginalURL     = OriginalTestURL
	UpdatedURL      = UpdatedTestURL
)

func TestSDK_DemWebsiteCrudLifecycle(t *testing.T) {
	ctx := context.Background()
	s := CreateTestClient("website-e2e-crud")

	createWebsiteRes, err := s.Dem.CreateWebsite(ctx, components.DemWebsite{
		Name: TestWebsiteName,
		URL:  OriginalURL,
		AvailabilityCheckSettings: &components.AvailabilityCheckSettings{
			TestFrom: components.DemTestFrom{
				Type:   components.TypeRegion,
				Values: []string{TestRegion},
			},
			TestIntervalInSeconds: DefaultTestInterval,
			Protocols:             []components.DemWebsiteProtocol{components.DemWebsiteProtocolHTTPS},
		},
	})
	require.NoError(t, err, "Failed to create test website")
	require.Equal(t, http.StatusCreated, createWebsiteRes.HTTPMeta.Response.StatusCode, "Create website returned unexpected status")
	require.NotNil(t, createWebsiteRes.CommonEntityID, "Create response should contain entity ID")

	entityID := createWebsiteRes.CommonEntityID.ID
	t.Logf("Created website entity with ID: %s", entityID)
	// Cleanup regardless of test outcome
	defer func() {
		_, err := s.Dem.DeleteWebsite(ctx, operations.DeleteWebsiteRequest{
			EntityID: entityID,
		})
		if err != nil {
			t.Logf("Failed to cleanup website %s: %v", entityID, err)
		} else {
			t.Logf("Deleted website entity with ID: %s", entityID)
		}
	}()

	website := WaitForWebsite(t, ctx, s, entityID, WebsiteAvailabilityTimeout)
	VerifyWebsite(t, website, TestWebsiteName, OriginalURL)

	updateWebsiteRes, err := s.Dem.UpdateWebsite(ctx, operations.UpdateWebsiteRequest{
		EntityID: entityID,
		DemWebsite: components.DemWebsite{
			Name: TestWebsiteName,
			URL:  UpdatedURL,
			AvailabilityCheckSettings: &components.AvailabilityCheckSettings{
				TestFrom: components.DemTestFrom{
					Type:   components.TypeRegion,
					Values: []string{TestRegion},
				},
				TestIntervalInSeconds: UpdatedTestInterval,
				Protocols:             []components.DemWebsiteProtocol{components.DemWebsiteProtocolHTTPS},
			},
		},
	})

	if err != nil {
		t.Logf("Update operation encountered error: %v", err)
	} else {
		statusCode := updateWebsiteRes.HTTPMeta.Response.StatusCode
		if statusCode == http.StatusOK || statusCode == http.StatusNoContent {
			t.Logf("Updated website entity with ID: %s", entityID)
			WaitForUpdatedWebsite(t, ctx, s, entityID, UpdatedURL, UpdatedTestInterval, WebsiteUpdateTimeout)

			getRes, err := s.Dem.GetWebsite(ctx, operations.GetWebsiteRequest{
				EntityID: entityID,
			})

			require.NoError(t, err, "Failed to get entity")
			require.Equal(t, http.StatusOK, getRes.HTTPMeta.Response.StatusCode, "Entity returned with unexpected status")
		}
	}

	deleteRes, err := s.Dem.DeleteWebsite(ctx, operations.DeleteWebsiteRequest{
		EntityID: entityID,
	})
	require.NoError(t, err, "Failed to delete test website")
	require.Equal(t, http.StatusOK, deleteRes.HTTPMeta.Response.StatusCode, "Delete website returned unexpected status")

	getWebsiteRes, err := s.Dem.GetWebsite(ctx, operations.GetWebsiteRequest{
		EntityID: entityID,
	})
	if err != nil {
		t.Logf("Expected error after deletion: %v", err)
	} else {
		assert.Equal(t, http.StatusNotFound, getWebsiteRes.HTTPMeta.Response.StatusCode, "Website should not be found after it's deleted")
	}

	t.Log("DEM Websites End-to-end test completed successfully.")
}

func WaitForWebsite(t *testing.T, ctx context.Context, s *swov1.Swo, entityID string, timeout time.Duration) *components.DemGetWebsiteResponse {
	var website *components.DemGetWebsiteResponse

	assert.Eventually(t, func() bool {
		getWebsiteRes, err := s.Dem.GetWebsite(ctx, operations.GetWebsiteRequest{
			EntityID: entityID,
		})

		if err != nil {
			return false
		}

		if getWebsiteRes.HTTPMeta.Response.StatusCode == http.StatusNotFound {
			return false
		}

		if getWebsiteRes.HTTPMeta.Response.StatusCode != http.StatusOK {
			return false
		}

		website = getWebsiteRes.DemGetWebsiteResponse
		assert.Equal(t, entityID, website.ID)
		assert.Equal(t, WebsiteEntityType, website.Type)
		return true

	}, timeout, RetryInterval, "Website should become available within %v", timeout)

	return website
}

func VerifyWebsite(t *testing.T, website *components.DemGetWebsiteResponse, expectedName, expectedURL string) {
	assert.Equal(t, expectedName, website.Name)
	assert.Equal(t, expectedURL, website.URL)
	assert.NotNil(t, website.AvailabilityCheckSettings)
	assert.Contains(t, ValidWebsiteStatuses, website.Status)
}

func WaitForUpdatedWebsite(t *testing.T, ctx context.Context, s *swov1.Swo, entityID, expectedURL string, expectedInterval float64, timeout time.Duration) {
	assert.Eventually(t, func() bool {
		getWebsiteRes, err := s.Dem.GetWebsite(ctx, operations.GetWebsiteRequest{
			EntityID: entityID,
		})

		if err != nil || getWebsiteRes.HTTPMeta.Response.StatusCode != http.StatusOK {
			return false
		}

		website := getWebsiteRes.DemGetWebsiteResponse

		if website.URL == expectedURL {
			if website.AvailabilityCheckSettings != nil && website.AvailabilityCheckSettings.TestIntervalInSeconds == expectedInterval {
				return true
			}
		}
		return false

	}, timeout, UpdateRetryInterval, "Updated website data should be created within %v", timeout)
}
