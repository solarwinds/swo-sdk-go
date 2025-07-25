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

func TestEntities_ListEntities(t *testing.T) {
	ctx := context.Background()

	testHTTPClient := createTestHTTPClient("listEntities")

	s := swov1.New(
		swov1.WithServerURL(utils.GetEnv("TEST_SERVER_URL", "http://localhost:18080")),
		swov1.WithClient(testHTTPClient),
		swov1.WithSecurity(utils.GetEnv("SWO_API_TOKEN", "value")),
	)

	res, err := s.Entities.ListEntities(ctx, operations.ListEntitiesRequest{
		Type: "<value>",
	})
	require.NoError(t, err)
	assert.Equal(t, 200, res.HTTPMeta.Response.StatusCode)
	assert.NotNil(t, res.Object)
	assert.Equal(t, &operations.ListEntitiesResponseBody{
		Entities: []components.EntitiesEntity{
			components.EntitiesEntity{
				ID:            "e-1234567890",
				Type:          "SyslogHost",
				Name:          swov1.String("syslog-host-1"),
				DisplayName:   swov1.String("SyslogTest"),
				CreatedTime:   types.MustNewTimeFromString("2024-11-25T16:38:24Z"),
				UpdatedTime:   types.MustNewTimeFromString("2024-12-01T16:38:24Z"),
				LastSeenTime:  types.MustTimeFromString("2024-11-25T16:38:24Z"),
				InMaintenance: false,
				Healthscore: &components.Healthscore{
					Score:    100,
					Category: components.CategoryGood,
				},
				Tags: map[string]*string{
					"gg.tk.token":  swov1.String("test"),
					"kfi.tk.token": swov1.String("qa-test"),
				},
				Attributes: map[string]any{
					"protocols": []any{
						"HTTP",
					},
					"features": []any{
						"rum",
					},
					"isAvailabilityCheckPaused": false,
					"extensions": map[string]any{
						"has_extension": true,
					},
				},
			},
		},
		PageInfo: components.CommonPageInfo{
			PrevPage: "<value>",
			NextPage: "<value>",
		},
	}, res.Object)

}

func TestEntities_GetEntityByID(t *testing.T) {
	ctx := context.Background()

	testHTTPClient := createTestHTTPClient("getEntityById")

	s := swov1.New(
		swov1.WithServerURL(utils.GetEnv("TEST_SERVER_URL", "http://localhost:18080")),
		swov1.WithClient(testHTTPClient),
		swov1.WithSecurity(utils.GetEnv("SWO_API_TOKEN", "value")),
	)

	res, err := s.Entities.GetEntityByID(ctx, operations.GetEntityByIDRequest{
		ID: "<id>",
	})
	require.NoError(t, err)
	assert.Equal(t, 200, res.HTTPMeta.Response.StatusCode)
	assert.NotNil(t, res.EntitiesEntity)
	assert.Equal(t, &components.EntitiesEntity{
		ID:            "e-1234567890",
		Type:          "SyslogHost",
		Name:          swov1.String("syslog-host-1"),
		DisplayName:   swov1.String("SyslogTest"),
		CreatedTime:   types.MustNewTimeFromString("2024-11-25T16:38:24Z"),
		UpdatedTime:   types.MustNewTimeFromString("2024-12-01T16:38:24Z"),
		LastSeenTime:  types.MustTimeFromString("2024-11-25T16:38:24Z"),
		InMaintenance: false,
		Healthscore: &components.Healthscore{
			Score:    100,
			Category: components.CategoryGood,
		},
		Tags: map[string]*string{
			"gg.tk.token":  swov1.String("test"),
			"kfi.tk.token": swov1.String("qa-test"),
		},
		Attributes: map[string]any{
			"protocols": []any{
				"HTTP",
			},
			"features": []any{
				"rum",
			},
			"isAvailabilityCheckPaused": false,
			"extensions": map[string]any{
				"has_extension": true,
			},
		},
	}, res.EntitiesEntity)

}

func TestEntities_UpdateEntityByID(t *testing.T) {
	ctx := context.Background()

	testHTTPClient := createTestHTTPClient("updateEntityById")

	s := swov1.New(
		swov1.WithServerURL(utils.GetEnv("TEST_SERVER_URL", "http://localhost:18080")),
		swov1.WithClient(testHTTPClient),
		swov1.WithSecurity(utils.GetEnv("SWO_API_TOKEN", "value")),
	)

	res, err := s.Entities.UpdateEntityByID(ctx, operations.UpdateEntityByIDRequest{
		ID: "<id>",
		EntitiesEntity: components.EntitiesEntityInput{
			DisplayName: swov1.String("SyslogTest"),
			Tags: map[string]*string{
				"gg.tk.token":  swov1.String("test"),
				"kfi.tk.token": swov1.String("qa-test"),
			},
		},
	})
	require.NoError(t, err)
	assert.Equal(t, 202, res.HTTPMeta.Response.StatusCode)

}
