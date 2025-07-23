// THESE TESTS ARE MANUALLY CREATED END-TO-END TESTS FOR THE METADATA API.
//
// Entity Create and Update operations are eventually consistent,
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

	"github.com/solarwinds/swo-sdk-go/swov1/models/components"
	"github.com/solarwinds/swo-sdk-go/swov1/models/operations"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	metadataEntityName = TestEntityName + "-metadata-e2e"
)

func TestSDK_MetadataE2E(t *testing.T) {
	ctx := context.Background()
	s := CreateTestClient("metadata-e2e")

	createWebsiteRes, err := s.Dem.CreateWebsite(ctx, components.DemWebsite{
		Name: metadataEntityName,
		URL:  OriginalTestURL,
		AvailabilityCheckSettings: &components.AvailabilityCheckSettings{
			TestFrom: components.DemTestFrom{
				Type:   components.TypeRegion,
				Values: []string{TestRegion},
			},
			TestIntervalInSeconds: DefaultTestInterval,
			Protocols:             []components.DemWebsiteProtocol{components.DemWebsiteProtocolHTTPS},
		},
	})
	require.NoError(t, err, "Failed to create test website for metadata tests")
	require.Equal(t, http.StatusCreated, createWebsiteRes.HTTPMeta.Response.StatusCode, "Create website returned unexpected status")
	require.NotNil(t, createWebsiteRes.CommonEntityID, "Create response should contain entity ID")

	entityID := createWebsiteRes.CommonEntityID.ID
	t.Logf("Created entity with ID: %s", entityID)
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

	waitForEntityAvailability(ctx, t, s, entityID, metadataEntityName)

	res, err := s.Metadata.ListEntityTypes(ctx)
	require.NoError(t, err, "Failed to list entity types")
	require.Equal(t, http.StatusOK, res.HTTPMeta.Response.StatusCode, "List entity types returned unexpected status")
	assert.NotEmpty(t, res.Object.Types, "Entity types should not be empty")

	websiteMetricsRes, err := s.Metadata.ListMetricsForEntityType(ctx, operations.ListMetricsForEntityTypeRequest{
		Type: WebsiteEntityType,
	})
	require.NoError(t, err, "Failed to list metrics for Website entity type")
	require.Equal(t, http.StatusOK, websiteMetricsRes.HTTPMeta.Response.StatusCode, "List metrics for Website returned unexpected status")
	assert.Equal(t, WebsiteEntityType, websiteMetricsRes.Object.Type)
	assert.NotEmpty(t, websiteMetricsRes.Object.Metrics, "Website metrics should not be empty")

	t.Log("Metadata E2E test completed.")
}
