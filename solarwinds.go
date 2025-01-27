// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package swosdkgo

import (
	"context"
	"fmt"
	"github.com/solarwinds/swo-sdk-go/internal/hooks"
	"github.com/solarwinds/swo-sdk-go/internal/utils"
	"github.com/solarwinds/swo-sdk-go/models/components"
	"github.com/solarwinds/swo-sdk-go/retry"
	"net/http"
	"time"
)

// ServerList contains the list of servers available to the SDK
var ServerList = []string{
	// Regional endpoint
	"https://api.na-01.cloud.solarwinds.com",
}

// HTTPClient provides an interface for suplying the SDK with a custom HTTP client
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// String provides a helper function to return a pointer to a string
func String(s string) *string { return &s }

// Bool provides a helper function to return a pointer to a bool
func Bool(b bool) *bool { return &b }

// Int provides a helper function to return a pointer to an int
func Int(i int) *int { return &i }

// Int64 provides a helper function to return a pointer to an int64
func Int64(i int64) *int64 { return &i }

// Float32 provides a helper function to return a pointer to a float32
func Float32(f float32) *float32 { return &f }

// Float64 provides a helper function to return a pointer to a float64
func Float64(f float64) *float64 { return &f }

// Pointer provides a helper function to return a pointer to a type
func Pointer[T any](v T) *T { return &v }

type sdkConfiguration struct {
	Client            HTTPClient
	Security          func(context.Context) (interface{}, error)
	ServerURL         string
	ServerIndex       int
	ServerDefaults    []map[string]string
	Language          string
	OpenAPIDocVersion string
	SDKVersion        string
	GenVersion        string
	UserAgent         string
	RetryConfig       *retry.Config
	Hooks             *hooks.Hooks
	Timeout           *time.Duration
}

func (c *sdkConfiguration) GetServerDetails() (string, map[string]string) {
	if c.ServerURL != "" {
		return c.ServerURL, nil
	}

	return ServerList[c.ServerIndex], c.ServerDefaults[c.ServerIndex]
}

// SolarWinds Observability: SolarWinds Observability REST API
// [Rest API Documentation](https://documentation.solarwinds.com/en/success_center/observability/content/api/api-swagger.htm)
type SolarWinds struct {
	Changeevents *Changeevents
	Dem          *Dem
	Entities     *Entities
	Logs         *Logs
	Metadata     *Metadata
	Metrics      *Metrics
	Tokens       *Tokens

	sdkConfiguration sdkConfiguration
}

type SDKOption func(*SolarWinds)

// WithServerURL allows the overriding of the default server URL
func WithServerURL(serverURL string) SDKOption {
	return func(sdk *SolarWinds) {
		sdk.sdkConfiguration.ServerURL = serverURL
	}
}

// WithTemplatedServerURL allows the overriding of the default server URL with a templated URL populated with the provided parameters
func WithTemplatedServerURL(serverURL string, params map[string]string) SDKOption {
	return func(sdk *SolarWinds) {
		if params != nil {
			serverURL = utils.ReplaceParameters(serverURL, params)
		}

		sdk.sdkConfiguration.ServerURL = serverURL
	}
}

// WithServerIndex allows the overriding of the default server by index
func WithServerIndex(serverIndex int) SDKOption {
	return func(sdk *SolarWinds) {
		if serverIndex < 0 || serverIndex >= len(ServerList) {
			panic(fmt.Errorf("server index %d out of range", serverIndex))
		}

		sdk.sdkConfiguration.ServerIndex = serverIndex
	}
}

// WithRegion allows setting the region variable for url substitution
func WithRegion(region string) SDKOption {
	return func(sdk *SolarWinds) {
		for idx := range sdk.sdkConfiguration.ServerDefaults {
			if _, ok := sdk.sdkConfiguration.ServerDefaults[idx]["region"]; !ok {
				continue
			}

			sdk.sdkConfiguration.ServerDefaults[idx]["region"] = fmt.Sprintf("%v", region)
		}
	}
}

// WithClient allows the overriding of the default HTTP client used by the SDK
func WithClient(client HTTPClient) SDKOption {
	return func(sdk *SolarWinds) {
		sdk.sdkConfiguration.Client = client
	}
}

// WithSecurity configures the SDK to use the provided security details
func WithSecurity(apiToken string) SDKOption {
	return func(sdk *SolarWinds) {
		security := components.Security{APIToken: &apiToken}
		sdk.sdkConfiguration.Security = utils.AsSecuritySource(&security)
	}
}

// WithSecuritySource configures the SDK to invoke the Security Source function on each method call to determine authentication
func WithSecuritySource(security func(context.Context) (components.Security, error)) SDKOption {
	return func(sdk *SolarWinds) {
		sdk.sdkConfiguration.Security = func(ctx context.Context) (interface{}, error) {
			return security(ctx)
		}
	}
}

func WithRetryConfig(retryConfig retry.Config) SDKOption {
	return func(sdk *SolarWinds) {
		sdk.sdkConfiguration.RetryConfig = &retryConfig
	}
}

// WithTimeout Optional request timeout applied to each operation
func WithTimeout(timeout time.Duration) SDKOption {
	return func(sdk *SolarWinds) {
		sdk.sdkConfiguration.Timeout = &timeout
	}
}

// New creates a new instance of the SDK with the provided options
func New(opts ...SDKOption) *SolarWinds {
	sdk := &SolarWinds{
		sdkConfiguration: sdkConfiguration{
			Language:          "go",
			OpenAPIDocVersion: "1.0.2",
			SDKVersion:        "0.2.0",
			GenVersion:        "2.495.1",
			UserAgent:         "speakeasy-sdk/go 0.2.0 2.495.1 1.0.2 github.com/solarwinds/swo-sdk-go",
			ServerDefaults: []map[string]string{
				{
					"region": "na-01",
				},
			},
			Hooks: hooks.New(),
		},
	}
	for _, opt := range opts {
		opt(sdk)
	}

	if sdk.sdkConfiguration.Security == nil {
		var envVarSecurity components.Security
		if utils.PopulateSecurityFromEnv(&envVarSecurity) {
			sdk.sdkConfiguration.Security = utils.AsSecuritySource(envVarSecurity)
		}
	}

	// Use WithClient to override the default client if you would like to customize the timeout
	if sdk.sdkConfiguration.Client == nil {
		sdk.sdkConfiguration.Client = &http.Client{Timeout: 60 * time.Second}
	}

	currentServerURL, _ := sdk.sdkConfiguration.GetServerDetails()
	serverURL := currentServerURL
	serverURL, sdk.sdkConfiguration.Client = sdk.sdkConfiguration.Hooks.SDKInit(currentServerURL, sdk.sdkConfiguration.Client)
	if serverURL != currentServerURL {
		sdk.sdkConfiguration.ServerURL = serverURL
	}

	sdk.Changeevents = newChangeevents(sdk.sdkConfiguration)

	sdk.Dem = newDem(sdk.sdkConfiguration)

	sdk.Entities = newEntities(sdk.sdkConfiguration)

	sdk.Logs = newLogs(sdk.sdkConfiguration)

	sdk.Metadata = newMetadata(sdk.sdkConfiguration)

	sdk.Metrics = newMetrics(sdk.sdkConfiguration)

	sdk.Tokens = newTokens(sdk.sdkConfiguration)

	return sdk
}
