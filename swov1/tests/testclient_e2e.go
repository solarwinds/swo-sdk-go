// THIS FILE WAS MANUALLY CREATED TO SUPPORT THE MANUALLY CREATED E2E TESTS.

package tests

import (
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
)

var ValidWebsiteStatuses = []components.DemGetWebsiteResponseStatus{
	components.DemGetWebsiteResponseStatusUp,
	components.DemGetWebsiteResponseStatusDown,
	components.DemGetWebsiteResponseStatusUnknown,
	components.DemGetWebsiteResponseStatusPaused,
	components.DemGetWebsiteResponseStatusMaintenance,
}

var ValidURIStatuses = []components.Status{
	components.StatusUp,
	components.StatusDown,
	components.StatusUnknown,
	components.StatusPaused,
	components.StatusMaintenance,
}

func CreateTestClient(testName string) *swov1.Swo {
	token := utils.GetEnv("SWO_STAGE_API_TOKEN", "")
	if token == "" {
		panic("The SWO_STAGE_API_TOKEN environment variable is not set. Set a valid SWO_STAGE_API_TOKEN to run the tests.")
	}

	testHTTPClient := createTestHTTPClient(testName)
	return swov1.New(
		swov1.WithServerURL(utils.GetEnv("CUSTOM_API_URL", "https://api.na-01.st-ssp.solarwinds.com")),
		swov1.WithSecurity(token),
		swov1.WithClient(testHTTPClient),
	)
}
