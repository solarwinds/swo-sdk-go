// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package tests

import (
	"context"
	"github.com/solarwinds/swo-sdk-go/swov1"
	"github.com/solarwinds/swo-sdk-go/swov1/internal/utils"
	"github.com/solarwinds/swo-sdk-go/swov1/models/components"
	"github.com/solarwinds/swo-sdk-go/swov1/models/operations"
	"github.com/solarwinds/swo-sdk-go/swov1/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestMetadata_ListEntityTypes(t *testing.T) {
	ctx := context.Background()

	testHTTPClient := createTestHTTPClient("listEntityTypes")

	s := swov1.New(
		swov1.WithServerURL(utils.GetEnv("TEST_SERVER_URL", "http://localhost:18080")),
		swov1.WithClient(testHTTPClient),
		swov1.WithSecurity(utils.GetEnv("SWO_API_TOKEN", "value")),
	)

	res, err := s.Metadata.ListEntityTypes(ctx)
	require.NoError(t, err)
	assert.Equal(t, 200, res.HTTPMeta.Response.StatusCode)
	assert.NotNil(t, res.Object)
	assert.Equal(t, &operations.ListEntityTypesResponseBody{
		Types: []string{
			"Service",
			"ServiceInstance",
			"KubernetesCluster",
		},
	}, res.Object)

}

func TestMetadata_ListMetricsForEntityType(t *testing.T) {
	ctx := context.Background()

	testHTTPClient := createTestHTTPClient("listMetricsForEntityType")

	s := swov1.New(
		swov1.WithServerURL(utils.GetEnv("TEST_SERVER_URL", "http://localhost:18080")),
		swov1.WithClient(testHTTPClient),
		swov1.WithSecurity(utils.GetEnv("SWO_API_TOKEN", "value")),
	)

	res, err := s.Metadata.ListMetricsForEntityType(ctx, operations.ListMetricsForEntityTypeRequest{
		Type: "<value>",
	})
	require.NoError(t, err)
	assert.Equal(t, 200, res.HTTPMeta.Response.StatusCode)
	assert.NotNil(t, res.Object)
	assert.Equal(t, &operations.ListMetricsForEntityTypeResponseBody{
		Type: "KubernetesCluster",
		Metrics: []components.CommonMetricInfo{
			components.CommonMetricInfo{
				Name:             "composite.custom.system.disk.io.rate",
				DisplayName:      swov1.String("Disk IO rate"),
				Description:      swov1.String("Disk bytes transferred per second"),
				Units:            swov1.String("bytes/s"),
				Formula:          swov1.String("rate(system.disk.io[5m]"),
				LastReportedTime: types.MustNewTimeFromString("2024-11-25T16:38:24Z"),
			},
		},
	}, res.Object)

}
