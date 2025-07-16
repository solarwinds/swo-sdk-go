# CloudAccounts
(*CloudAccounts*)

## Overview

### Available Operations

* [ActivateAwsIntegration](#activateawsintegration) - Activate AWS Integration
* [CreateOrgStructure](#createorgstructure) - Create Organizational Structure
* [UpdateAwsIntegration](#updateawsintegration) - Update AWS Integration
* [ValidateMgmtAccountOnboarding](#validatemgmtaccountonboarding) - Validate Management Account Onboarding

## ActivateAwsIntegration

Activate AWS Integration.

### Example Usage

```go
package main

import(
	"context"
	"os"
	"github.com/solarwinds/swo-sdk-go/swov1"
	"github.com/solarwinds/swo-sdk-go/swov1/models/components"
	"log"
)

func main() {
    ctx := context.Background()

    s := swov1.New(
        swov1.WithSecurity(os.Getenv("SWO_API_TOKEN")),
    )

    res, err := s.CloudAccounts.ActivateAwsIntegration(ctx, components.CloudAccountsAwsActivateIntegrationRequest{
        ManagementAccountID: "<id>",
        AccountID: "<id>",
        Enable: true,
    })
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                      | Type                                                                                                                           | Required                                                                                                                       | Description                                                                                                                    |
| ------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                                          | :heavy_check_mark:                                                                                                             | The context to use for the request.                                                                                            |
| `request`                                                                                                                      | [components.CloudAccountsAwsActivateIntegrationRequest](../../models/components/cloudaccountsawsactivateintegrationrequest.md) | :heavy_check_mark:                                                                                                             | The request object to use for the request.                                                                                     |
| `opts`                                                                                                                         | [][operations.Option](../../models/operations/option.md)                                                                       | :heavy_minus_sign:                                                                                                             | The options for this request.                                                                                                  |

### Response

**[*operations.ActivateAwsIntegrationResponse](../../models/operations/activateawsintegrationresponse.md), error**

### Errors

| Error Type                                                           | Status Code                                                          | Content Type                                                         |
| -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- |
| apierrors.ActivateAwsIntegrationResponseBody                         | 400                                                                  | application/json                                                     |
| apierrors.ActivateAwsIntegrationCloudAccountsResponseBody            | 401                                                                  | application/json                                                     |
| apierrors.ActivateAwsIntegrationCloudAccountsResponseResponseBody    | 404                                                                  | application/json                                                     |
| apierrors.ActivateAwsIntegrationCloudAccountsResponse500ResponseBody | 500                                                                  | application/json                                                     |
| apierrors.APIError                                                   | 4XX, 5XX                                                             | \*/\*                                                                |

## CreateOrgStructure

Create AWS Organizational Structure.

### Example Usage

```go
package main

import(
	"context"
	"os"
	"github.com/solarwinds/swo-sdk-go/swov1"
	"github.com/solarwinds/swo-sdk-go/swov1/models/components"
	"log"
)

func main() {
    ctx := context.Background()

    s := swov1.New(
        swov1.WithSecurity(os.Getenv("SWO_API_TOKEN")),
    )

    res, err := s.CloudAccounts.CreateOrgStructure(ctx, components.CloudAccountsAwsOrganisationalUnitRequest{
        MgmtAccountID: "<id>",
        Structure: []components.CloudAccountsAwsOrganisationalUnit{},
    })
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                    | Type                                                                                                                         | Required                                                                                                                     | Description                                                                                                                  |
| ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                                        | :heavy_check_mark:                                                                                                           | The context to use for the request.                                                                                          |
| `request`                                                                                                                    | [components.CloudAccountsAwsOrganisationalUnitRequest](../../models/components/cloudaccountsawsorganisationalunitrequest.md) | :heavy_check_mark:                                                                                                           | The request object to use for the request.                                                                                   |
| `opts`                                                                                                                       | [][operations.Option](../../models/operations/option.md)                                                                     | :heavy_minus_sign:                                                                                                           | The options for this request.                                                                                                |

### Response

**[*operations.CreateOrgStructureResponse](../../models/operations/createorgstructureresponse.md), error**

### Errors

| Error Type                                                       | Status Code                                                      | Content Type                                                     |
| ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- |
| apierrors.CreateOrgStructureResponseBody                         | 400                                                              | application/json                                                 |
| apierrors.CreateOrgStructureCloudAccountsResponseBody            | 401                                                              | application/json                                                 |
| apierrors.CreateOrgStructureCloudAccountsResponseResponseBody    | 404                                                              | application/json                                                 |
| apierrors.CreateOrgStructureCloudAccountsResponse500ResponseBody | 500                                                              | application/json                                                 |
| apierrors.APIError                                               | 4XX, 5XX                                                         | \*/\*                                                            |

## UpdateAwsIntegration

Update AWS Integration details.

### Example Usage

```go
package main

import(
	"context"
	"os"
	"github.com/solarwinds/swo-sdk-go/swov1"
	"github.com/solarwinds/swo-sdk-go/swov1/models/components"
	"log"
)

func main() {
    ctx := context.Background()

    s := swov1.New(
        swov1.WithSecurity(os.Getenv("SWO_API_TOKEN")),
    )

    res, err := s.CloudAccounts.UpdateAwsIntegration(ctx, components.CloudAccountsAwsUpdateIntegrationRequest{
        ManagementAccountID: "<id>",
        AccountID: "<id>",
        AccountName: "<value>",
        RoleArn: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CloudAccountsAwsUpdateIntegrationResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                  | Type                                                                                                                       | Required                                                                                                                   | Description                                                                                                                |
| -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                                      | :heavy_check_mark:                                                                                                         | The context to use for the request.                                                                                        |
| `request`                                                                                                                  | [components.CloudAccountsAwsUpdateIntegrationRequest](../../models/components/cloudaccountsawsupdateintegrationrequest.md) | :heavy_check_mark:                                                                                                         | The request object to use for the request.                                                                                 |
| `opts`                                                                                                                     | [][operations.Option](../../models/operations/option.md)                                                                   | :heavy_minus_sign:                                                                                                         | The options for this request.                                                                                              |

### Response

**[*operations.UpdateAwsIntegrationResponse](../../models/operations/updateawsintegrationresponse.md), error**

### Errors

| Error Type                                                      | Status Code                                                     | Content Type                                                    |
| --------------------------------------------------------------- | --------------------------------------------------------------- | --------------------------------------------------------------- |
| apierrors.UpdateAwsIntegrationResponseBody                      | 400                                                             | application/json                                                |
| apierrors.UpdateAwsIntegrationCloudAccountsResponseBody         | 401                                                             | application/json                                                |
| apierrors.UpdateAwsIntegrationCloudAccountsResponseResponseBody | 500                                                             | application/json                                                |
| apierrors.APIError                                              | 4XX, 5XX                                                        | \*/\*                                                           |

## ValidateMgmtAccountOnboarding

Validate if the management account is onboarded.

### Example Usage

```go
package main

import(
	"context"
	"os"
	"github.com/solarwinds/swo-sdk-go/swov1"
	"github.com/solarwinds/swo-sdk-go/swov1/models/components"
	"github.com/solarwinds/swo-sdk-go/swov1/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := swov1.New(
        swov1.WithSecurity(os.Getenv("SWO_API_TOKEN")),
    )

    res, err := s.CloudAccounts.ValidateMgmtAccountOnboarding(ctx, operations.ValidateMgmtAccountOnboardingRequest{
        Request: components.CloudAccountsAwsMgmtAccountOnboardingRequest{
            ManagementAccountID: "<id>",
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CloudAccountsAwsMgmtAccountOnboardingResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                          | Type                                                                                                               | Required                                                                                                           | Description                                                                                                        |
| ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                              | :heavy_check_mark:                                                                                                 | The context to use for the request.                                                                                |
| `request`                                                                                                          | [operations.ValidateMgmtAccountOnboardingRequest](../../models/operations/validatemgmtaccountonboardingrequest.md) | :heavy_check_mark:                                                                                                 | The request object to use for the request.                                                                         |
| `opts`                                                                                                             | [][operations.Option](../../models/operations/option.md)                                                           | :heavy_minus_sign:                                                                                                 | The options for this request.                                                                                      |

### Response

**[*operations.ValidateMgmtAccountOnboardingResponse](../../models/operations/validatemgmtaccountonboardingresponse.md), error**

### Errors

| Error Type                                                               | Status Code                                                              | Content Type                                                             |
| ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ |
| apierrors.ValidateMgmtAccountOnboardingResponseBody                      | 400                                                                      | application/json                                                         |
| apierrors.ValidateMgmtAccountOnboardingCloudAccountsResponseBody         | 401                                                                      | application/json                                                         |
| apierrors.ValidateMgmtAccountOnboardingCloudAccountsResponseResponseBody | 500                                                                      | application/json                                                         |
| apierrors.APIError                                                       | 4XX, 5XX                                                                 | \*/\*                                                                    |