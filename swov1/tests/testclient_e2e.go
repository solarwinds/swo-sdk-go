// THIS FILE WAS MANUALLY CREATED TO SUPPORT THE MANUALLY CREATED E2E TESTS.

package tests

import (
	"fmt"
	"time"

	"github.com/solarwinds/swo-sdk-go/swov1"
	"github.com/solarwinds/swo-sdk-go/swov1/internal/utils"
	"github.com/solarwinds/swo-sdk-go/swov1/models/components"
)

const (
	TestRegion = "NA"

	DefaultTestInterval = 300
	UpdatedTestInterval = 600

	DefaultTimeout             = 3 * time.Minute
	UpdateTimeout              = 4 * time.Minute
	DefaultRetryInterval       = 5 * time.Second
	DefaultUpdateRetryInterval = 10 * time.Second

	TestEntityName = "swo-sdk-e2e-crud-test"
	TestEntityURL  = "https://swo-sdk-e2e-test.com"

	WebsiteEntityType = "Website"
	URIEntityType     = "Uri"

	OriginalTestDomain = "create-sdk-e2e-test.com"
	UpdatedTestDomain  = "update-sdk-e2e-test.com"

	OriginalTestURL = "https://create-sdk-e2e-test.com"
	UpdatedTestURL  = "https://update-sdk-e2e-test.com"

	publicApiUrlEnvName   = "PUBLIC_SWO_API_STAGE_URL"
	publicApiTokenEnvName = "SWO_STAGE_API_TOKEN"
)

var ValidWebsiteStatuses = []components.DemGetWebsiteResponseStatus{
	components.DemGetWebsiteResponseStatusUp,
	components.DemGetWebsiteResponseStatusDown,
	components.DemGetWebsiteResponseStatusUnknown,
	components.DemGetWebsiteResponseStatusPaused,
	components.DemGetWebsiteResponseStatusMaintenance,
}

var ValidURIStatuses = []components.DemGetURIResponseStatus{
	components.StatusUp,
	components.StatusDown,
	components.StatusUnknown,
	components.StatusPaused,
	components.StatusMaintenance,
}

func CreateTestClient(testName string) *swov1.Swo {
	publicApiUrl := requireEnv(publicApiUrlEnvName)
	publicApiToken := requireEnv(publicApiTokenEnvName)

	testHTTPClient := createTestHTTPClient(testName)
	return swov1.New(
		swov1.WithServerURL(publicApiUrl),
		swov1.WithSecurity(publicApiToken),
		swov1.WithClient(testHTTPClient),
	)
}

func requireEnv(name string) string {
	val := utils.GetEnv(name, "")
	if val == "" {
		panic(fmt.Sprintf("The %s environment variable is not set. Set a valid %s to run the tests.", name, name))
	}
	return val
}
