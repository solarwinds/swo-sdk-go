# Dem

## Overview

### Available Operations

* [ListProbes](#listprobes) - Get a list of existing synthetic probes
* [GetDemSettings](#getdemsettings) - Get DEM settings
* [SetDemSettings](#setdemsettings) - Set DEM settings
* [CreateTransaction](#createtransaction) - Create transaction monitoring configuration
* [GetTransaction](#gettransaction) - Get transaction monitoring configuration
* [UpdateTransaction](#updatetransaction) - Update transaction monitoring configuration
* [DeleteTransaction](#deletetransaction) - Delete transaction
* [PauseTransactionMonitoring](#pausetransactionmonitoring) - Pause monitoring of the transaction
* [UnpauseTransactionMonitoring](#unpausetransactionmonitoring) - Unpause monitoring of the transaction
* [CreateURI](#createuri) - Create URI monitoring configuration
* [GetURI](#geturi) - Get URI monitoring configuration
* [UpdateURI](#updateuri) - Update URI monitoring configuration
* [DeleteURI](#deleteuri) - Delete URI
* [GetURIOutageStatuses](#geturioutagestatuses) - Get outage statuses
* [PauseURIMonitoring](#pauseurimonitoring) - Pause monitoring of the URI
* [GetURITestResults](#geturitestresults) - Get test results
* [UnpauseURIMonitoring](#unpauseurimonitoring) - Unpause monitoring of the URI
* [CreateWebsite](#createwebsite) - Create website monitoring configuration
* [GetWebsite](#getwebsite) - Get website monitoring configuration
* [UpdateWebsite](#updatewebsite) - Update website monitoring configuration
* [DeleteWebsite](#deletewebsite) - Delete website
* [GetWebsiteOutageStatuses](#getwebsiteoutagestatuses) - Get outage statuses
* [PauseWebsiteMonitoring](#pausewebsitemonitoring) - Pause monitoring of a website
* [GetWebsiteTestResults](#getwebsitetestresults) - Get test results
* [UnpauseWebsiteMonitoring](#unpausewebsitemonitoring) - Unpause monitoring of a website

## ListProbes

Get a list of existing synthetic probes

### Example Usage

<!-- UsageSnippet language="go" operationID="listProbes" method="get" path="/v1/dem/probes" -->
```go
package main

import(
	"context"
	"os"
	"github.com/solarwinds/swo-sdk-go/swov1"
	"log"
)

func main() {
    ctx := context.Background()

    s := swov1.New(
        swov1.WithSecurity(os.Getenv("SWO_API_TOKEN")),
    )

    res, err := s.Dem.ListProbes(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.DemListProbesResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.ListProbesResponse](../../models/operations/listprobesresponse.md), error**

### Errors

| Error Type                                | Status Code                               | Content Type                              |
| ----------------------------------------- | ----------------------------------------- | ----------------------------------------- |
| apierrors.CommonUnauthorizedErrorResponse | 401                                       | application/json                          |
| apierrors.CommonInternalErrorResponse     | 500                                       | application/json                          |
| apierrors.APIError                        | 4XX, 5XX                                  | \*/\*                                     |

## GetDemSettings

Get DEM settings

### Example Usage

<!-- UsageSnippet language="go" operationID="getDemSettings" method="get" path="/v1/dem/settings" -->
```go
package main

import(
	"context"
	"os"
	"github.com/solarwinds/swo-sdk-go/swov1"
	"log"
)

func main() {
    ctx := context.Background()

    s := swov1.New(
        swov1.WithSecurity(os.Getenv("SWO_API_TOKEN")),
    )

    res, err := s.Dem.GetDemSettings(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.DemOrganizationSettings != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetDemSettingsResponse](../../models/operations/getdemsettingsresponse.md), error**

### Errors

| Error Type                                | Status Code                               | Content Type                              |
| ----------------------------------------- | ----------------------------------------- | ----------------------------------------- |
| apierrors.CommonUnauthorizedErrorResponse | 401                                       | application/json                          |
| apierrors.CommonInternalErrorResponse     | 500                                       | application/json                          |
| apierrors.APIError                        | 4XX, 5XX                                  | \*/\*                                     |

## SetDemSettings

Set DEM settings

### Example Usage

<!-- UsageSnippet language="go" operationID="setDemSettings" method="put" path="/v1/dem/settings" -->
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

    res, err := s.Dem.SetDemSettings(ctx, components.DemOrganizationSettingsInput{
        AvailabilityOutageConfiguration: &components.DemOutageConfiguration{
            FailingTestLocations: components.FailingTestLocationsAll,
            ConsecutiveForDown: 2,
        },
        TransactionOutageConfiguration: &components.DemOutageConfiguration{
            FailingTestLocations: components.FailingTestLocationsAll,
            ConsecutiveForDown: 2,
        },
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

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [components.DemOrganizationSettingsInput](../../models/components/demorganizationsettingsinput.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `opts`                                                                                             | [][operations.Option](../../models/operations/option.md)                                           | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |

### Response

**[*operations.SetDemSettingsResponse](../../models/operations/setdemsettingsresponse.md), error**

### Errors

| Error Type                                | Status Code                               | Content Type                              |
| ----------------------------------------- | ----------------------------------------- | ----------------------------------------- |
| apierrors.CommonBadRequestErrorResponse   | 400                                       | application/json                          |
| apierrors.CommonUnauthorizedErrorResponse | 401                                       | application/json                          |
| apierrors.CommonInternalErrorResponse     | 500                                       | application/json                          |
| apierrors.APIError                        | 4XX, 5XX                                  | \*/\*                                     |

## CreateTransaction

Create transaction monitoring configuration

### Example Usage

<!-- UsageSnippet language="go" operationID="createTransaction" method="post" path="/v1/dem/transactions" -->
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

    res, err := s.Dem.CreateTransaction(ctx, components.DemTransaction{
        Name: "Solarwinds",
        Description: swov1.Pointer("Opens Solarwinds homepage"),
        TestDefinition: components.DemTransactionTestDefinition{
            TestFrom: components.DemTestFrom{
                Type: components.TypeRegion,
                Values: []string{
                    "NA",
                },
            },
            PlatformOptions: &components.PlatformOptions{
                ProbePlatforms: []components.DemProbePlatform{
                    components.DemProbePlatformAws,
                },
                TestFromAll: swov1.Pointer(true),
            },
            OutageConfiguration: &components.OutageConfiguration{
                FailingTestLocations: components.DemTransactionTestDefinitionFailingTestLocationsAll,
                ConsecutiveForDown: 2,
            },
            TestIntervalInSeconds: 14400,
            WindowSize: components.DemWindowSize{
                Width: 28965,
                Height: 156492,
            },
            Commands: []components.DemTransactionCommand{
                components.DemTransactionCommand{
                    Command: components.DemTransactionCommandNameOpen,
                    Target: swov1.Pointer("https://example.com"),
                },
            },
        },
        Tags: []components.CommonTag{
            components.CommonTag{
                Key: "environment",
                Value: "production",
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CommonEntityID != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                              | Type                                                                   | Required                                                               | Description                                                            |
| ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- |
| `ctx`                                                                  | [context.Context](https://pkg.go.dev/context#Context)                  | :heavy_check_mark:                                                     | The context to use for the request.                                    |
| `request`                                                              | [components.DemTransaction](../../models/components/demtransaction.md) | :heavy_check_mark:                                                     | The request object to use for the request.                             |
| `opts`                                                                 | [][operations.Option](../../models/operations/option.md)               | :heavy_minus_sign:                                                     | The options for this request.                                          |

### Response

**[*operations.CreateTransactionResponse](../../models/operations/createtransactionresponse.md), error**

### Errors

| Error Type                                | Status Code                               | Content Type                              |
| ----------------------------------------- | ----------------------------------------- | ----------------------------------------- |
| apierrors.CommonBadRequestErrorResponse   | 400                                       | application/json                          |
| apierrors.CommonUnauthorizedErrorResponse | 401                                       | application/json                          |
| apierrors.CommonInternalErrorResponse     | 500                                       | application/json                          |
| apierrors.APIError                        | 4XX, 5XX                                  | \*/\*                                     |

## GetTransaction

Get transaction monitoring configuration

### Example Usage

<!-- UsageSnippet language="go" operationID="getTransaction" method="get" path="/v1/dem/transactions/{entityId}" -->
```go
package main

import(
	"context"
	"os"
	"github.com/solarwinds/swo-sdk-go/swov1"
	"github.com/solarwinds/swo-sdk-go/swov1/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := swov1.New(
        swov1.WithSecurity(os.Getenv("SWO_API_TOKEN")),
    )

    res, err := s.Dem.GetTransaction(ctx, operations.GetTransactionRequest{
        EntityID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DemGetTransactionResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.GetTransactionRequest](../../models/operations/gettransactionrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |
| `opts`                                                                               | [][operations.Option](../../models/operations/option.md)                             | :heavy_minus_sign:                                                                   | The options for this request.                                                        |

### Response

**[*operations.GetTransactionResponse](../../models/operations/gettransactionresponse.md), error**

### Errors

| Error Type                                | Status Code                               | Content Type                              |
| ----------------------------------------- | ----------------------------------------- | ----------------------------------------- |
| apierrors.CommonUnauthorizedErrorResponse | 401                                       | application/json                          |
| apierrors.CommonNotFoundErrorResponse     | 404                                       | application/json                          |
| apierrors.CommonInternalErrorResponse     | 500                                       | application/json                          |
| apierrors.APIError                        | 4XX, 5XX                                  | \*/\*                                     |

## UpdateTransaction

Update transaction monitoring configuration

### Example Usage

<!-- UsageSnippet language="go" operationID="updateTransaction" method="put" path="/v1/dem/transactions/{entityId}" -->
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

    res, err := s.Dem.UpdateTransaction(ctx, operations.UpdateTransactionRequest{
        EntityID: "<id>",
        DemTransaction: components.DemTransaction{
            Name: "Solarwinds",
            Description: swov1.Pointer("Opens Solarwinds homepage"),
            TestDefinition: components.DemTransactionTestDefinition{
                TestFrom: components.DemTestFrom{
                    Type: components.TypeRegion,
                    Values: []string{
                        "NA",
                    },
                },
                PlatformOptions: &components.PlatformOptions{
                    ProbePlatforms: []components.DemProbePlatform{
                        components.DemProbePlatformAws,
                    },
                    TestFromAll: swov1.Pointer(true),
                },
                OutageConfiguration: nil,
                TestIntervalInSeconds: 14400,
                WindowSize: components.DemWindowSize{
                    Width: 750299,
                    Height: 72607,
                },
                Commands: []components.DemTransactionCommand{
                    components.DemTransactionCommand{
                        Command: components.DemTransactionCommandNameOpen,
                        Target: swov1.Pointer("https://example.com"),
                    },
                },
            },
            Tags: []components.CommonTag{
                components.CommonTag{
                    Key: "environment",
                    Value: "production",
                },
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CommonEntityID != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.UpdateTransactionRequest](../../models/operations/updatetransactionrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `opts`                                                                                     | [][operations.Option](../../models/operations/option.md)                                   | :heavy_minus_sign:                                                                         | The options for this request.                                                              |

### Response

**[*operations.UpdateTransactionResponse](../../models/operations/updatetransactionresponse.md), error**

### Errors

| Error Type                                | Status Code                               | Content Type                              |
| ----------------------------------------- | ----------------------------------------- | ----------------------------------------- |
| apierrors.CommonBadRequestErrorResponse   | 400                                       | application/json                          |
| apierrors.CommonUnauthorizedErrorResponse | 401                                       | application/json                          |
| apierrors.CommonNotFoundErrorResponse     | 404                                       | application/json                          |
| apierrors.CommonInternalErrorResponse     | 500                                       | application/json                          |
| apierrors.APIError                        | 4XX, 5XX                                  | \*/\*                                     |

## DeleteTransaction

Delete transaction

### Example Usage

<!-- UsageSnippet language="go" operationID="deleteTransaction" method="delete" path="/v1/dem/transactions/{entityId}" -->
```go
package main

import(
	"context"
	"os"
	"github.com/solarwinds/swo-sdk-go/swov1"
	"github.com/solarwinds/swo-sdk-go/swov1/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := swov1.New(
        swov1.WithSecurity(os.Getenv("SWO_API_TOKEN")),
    )

    res, err := s.Dem.DeleteTransaction(ctx, operations.DeleteTransactionRequest{
        EntityID: "<id>",
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

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.DeleteTransactionRequest](../../models/operations/deletetransactionrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `opts`                                                                                     | [][operations.Option](../../models/operations/option.md)                                   | :heavy_minus_sign:                                                                         | The options for this request.                                                              |

### Response

**[*operations.DeleteTransactionResponse](../../models/operations/deletetransactionresponse.md), error**

### Errors

| Error Type                                | Status Code                               | Content Type                              |
| ----------------------------------------- | ----------------------------------------- | ----------------------------------------- |
| apierrors.CommonUnauthorizedErrorResponse | 401                                       | application/json                          |
| apierrors.CommonNotFoundErrorResponse     | 404                                       | application/json                          |
| apierrors.CommonInternalErrorResponse     | 500                                       | application/json                          |
| apierrors.APIError                        | 4XX, 5XX                                  | \*/\*                                     |

## PauseTransactionMonitoring

Pause monitoring of the transaction

### Example Usage

<!-- UsageSnippet language="go" operationID="pauseTransactionMonitoring" method="put" path="/v1/dem/transactions/{entityId}/pauseMonitoring" -->
```go
package main

import(
	"context"
	"os"
	"github.com/solarwinds/swo-sdk-go/swov1"
	"github.com/solarwinds/swo-sdk-go/swov1/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := swov1.New(
        swov1.WithSecurity(os.Getenv("SWO_API_TOKEN")),
    )

    res, err := s.Dem.PauseTransactionMonitoring(ctx, operations.PauseTransactionMonitoringRequest{
        EntityID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CommonEntityID != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                    | Type                                                                                                         | Required                                                                                                     | Description                                                                                                  |
| ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                        | :heavy_check_mark:                                                                                           | The context to use for the request.                                                                          |
| `request`                                                                                                    | [operations.PauseTransactionMonitoringRequest](../../models/operations/pausetransactionmonitoringrequest.md) | :heavy_check_mark:                                                                                           | The request object to use for the request.                                                                   |
| `opts`                                                                                                       | [][operations.Option](../../models/operations/option.md)                                                     | :heavy_minus_sign:                                                                                           | The options for this request.                                                                                |

### Response

**[*operations.PauseTransactionMonitoringResponse](../../models/operations/pausetransactionmonitoringresponse.md), error**

### Errors

| Error Type                                | Status Code                               | Content Type                              |
| ----------------------------------------- | ----------------------------------------- | ----------------------------------------- |
| apierrors.CommonUnauthorizedErrorResponse | 401                                       | application/json                          |
| apierrors.CommonNotFoundErrorResponse     | 404                                       | application/json                          |
| apierrors.CommonInternalErrorResponse     | 500                                       | application/json                          |
| apierrors.APIError                        | 4XX, 5XX                                  | \*/\*                                     |

## UnpauseTransactionMonitoring

Unpause monitoring of the transaction

### Example Usage

<!-- UsageSnippet language="go" operationID="unpauseTransactionMonitoring" method="put" path="/v1/dem/transactions/{entityId}/unpauseMonitoring" -->
```go
package main

import(
	"context"
	"os"
	"github.com/solarwinds/swo-sdk-go/swov1"
	"github.com/solarwinds/swo-sdk-go/swov1/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := swov1.New(
        swov1.WithSecurity(os.Getenv("SWO_API_TOKEN")),
    )

    res, err := s.Dem.UnpauseTransactionMonitoring(ctx, operations.UnpauseTransactionMonitoringRequest{
        EntityID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CommonEntityID != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                        | Type                                                                                                             | Required                                                                                                         | Description                                                                                                      |
| ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                            | :heavy_check_mark:                                                                                               | The context to use for the request.                                                                              |
| `request`                                                                                                        | [operations.UnpauseTransactionMonitoringRequest](../../models/operations/unpausetransactionmonitoringrequest.md) | :heavy_check_mark:                                                                                               | The request object to use for the request.                                                                       |
| `opts`                                                                                                           | [][operations.Option](../../models/operations/option.md)                                                         | :heavy_minus_sign:                                                                                               | The options for this request.                                                                                    |

### Response

**[*operations.UnpauseTransactionMonitoringResponse](../../models/operations/unpausetransactionmonitoringresponse.md), error**

### Errors

| Error Type                                | Status Code                               | Content Type                              |
| ----------------------------------------- | ----------------------------------------- | ----------------------------------------- |
| apierrors.CommonUnauthorizedErrorResponse | 401                                       | application/json                          |
| apierrors.CommonNotFoundErrorResponse     | 404                                       | application/json                          |
| apierrors.CommonInternalErrorResponse     | 500                                       | application/json                          |
| apierrors.APIError                        | 4XX, 5XX                                  | \*/\*                                     |

## CreateURI

Create URI monitoring configuration

### Example Usage

<!-- UsageSnippet language="go" operationID="createUri" method="post" path="/v1/dem/uris" -->
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

    res, err := s.Dem.CreateURI(ctx, components.DemURI{
        Name: "solarwinds.com",
        IPOrDomain: "solarwinds.com",
        AvailabilityCheckSettings: components.DemURIAvailabilityCheckSettingsInput{
            PlatformOptions: &components.DemURIAvailabilityCheckSettingsInputPlatformOptions{
                ProbePlatforms: []components.DemProbePlatform{
                    components.DemProbePlatformAws,
                },
                TestFromAll: swov1.Pointer(true),
            },
            TestFrom: components.DemTestFrom{
                Type: components.TypeRegion,
                Values: []string{
                    "NA",
                },
            },
            TestIntervalInSeconds: 14400,
            OutageConfiguration: nil,
            DNS: &components.DNS{
                Enabled: false,
                Nameserver: "8.8.8.8",
                Port: swov1.Pointer[int](53),
                IPToExpect: "1.2.3.4",
            },
            Ping: &components.Ping{
                Enabled: false,
            },
            TCP: &components.TCP{
                Enabled: true,
                Port: 443,
                StringToSend: swov1.Pointer("GET / HTTP/1.1\r\nHost: solarwinds.com\r\nConnection: close\r\n\r\n"),
                StringToExpect: swov1.Pointer("HTTP/1.1 200 OK"),
            },
            UDP: nil,
        },
        Tags: []components.CommonTag{
            components.CommonTag{
                Key: "environment",
                Value: "production",
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CommonEntityID != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `request`                                                | [components.DemURI](../../models/components/demuri.md)   | :heavy_check_mark:                                       | The request object to use for the request.               |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.CreateURIResponse](../../models/operations/createuriresponse.md), error**

### Errors

| Error Type                                | Status Code                               | Content Type                              |
| ----------------------------------------- | ----------------------------------------- | ----------------------------------------- |
| apierrors.CommonBadRequestErrorResponse   | 400                                       | application/json                          |
| apierrors.CommonUnauthorizedErrorResponse | 401                                       | application/json                          |
| apierrors.CommonInternalErrorResponse     | 500                                       | application/json                          |
| apierrors.APIError                        | 4XX, 5XX                                  | \*/\*                                     |

## GetURI

Get URI monitoring configuration

### Example Usage

<!-- UsageSnippet language="go" operationID="getUri" method="get" path="/v1/dem/uris/{entityId}" -->
```go
package main

import(
	"context"
	"os"
	"github.com/solarwinds/swo-sdk-go/swov1"
	"github.com/solarwinds/swo-sdk-go/swov1/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := swov1.New(
        swov1.WithSecurity(os.Getenv("SWO_API_TOKEN")),
    )

    res, err := s.Dem.GetURI(ctx, operations.GetURIRequest{
        EntityID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DemGetURIResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                            | Type                                                                 | Required                                                             | Description                                                          |
| -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- |
| `ctx`                                                                | [context.Context](https://pkg.go.dev/context#Context)                | :heavy_check_mark:                                                   | The context to use for the request.                                  |
| `request`                                                            | [operations.GetURIRequest](../../models/operations/geturirequest.md) | :heavy_check_mark:                                                   | The request object to use for the request.                           |
| `opts`                                                               | [][operations.Option](../../models/operations/option.md)             | :heavy_minus_sign:                                                   | The options for this request.                                        |

### Response

**[*operations.GetURIResponse](../../models/operations/geturiresponse.md), error**

### Errors

| Error Type                                | Status Code                               | Content Type                              |
| ----------------------------------------- | ----------------------------------------- | ----------------------------------------- |
| apierrors.CommonUnauthorizedErrorResponse | 401                                       | application/json                          |
| apierrors.CommonNotFoundErrorResponse     | 404                                       | application/json                          |
| apierrors.CommonInternalErrorResponse     | 500                                       | application/json                          |
| apierrors.APIError                        | 4XX, 5XX                                  | \*/\*                                     |

## UpdateURI

Update URI monitoring configuration

### Example Usage

<!-- UsageSnippet language="go" operationID="updateUri" method="put" path="/v1/dem/uris/{entityId}" -->
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

    res, err := s.Dem.UpdateURI(ctx, operations.UpdateURIRequest{
        EntityID: "<id>",
        DemURI: components.DemURI{
            Name: "solarwinds.com",
            IPOrDomain: "solarwinds.com",
            AvailabilityCheckSettings: components.DemURIAvailabilityCheckSettingsInput{
                PlatformOptions: &components.DemURIAvailabilityCheckSettingsInputPlatformOptions{
                    ProbePlatforms: []components.DemProbePlatform{
                        components.DemProbePlatformAws,
                    },
                    TestFromAll: swov1.Pointer(true),
                },
                TestFrom: components.DemTestFrom{
                    Type: components.TypeRegion,
                    Values: []string{
                        "NA",
                    },
                },
                TestIntervalInSeconds: 14400,
                OutageConfiguration: &components.DemURIAvailabilityCheckSettingsInputOutageConfiguration{
                    FailingTestLocations: components.DemURIAvailabilityCheckSettingsInputFailingTestLocationsAll,
                    ConsecutiveForDown: 2,
                },
                DNS: &components.DNS{
                    Enabled: false,
                    Nameserver: "8.8.8.8",
                    Port: swov1.Pointer[int](53),
                    IPToExpect: "1.2.3.4",
                },
                Ping: &components.Ping{
                    Enabled: false,
                },
                TCP: &components.TCP{
                    Enabled: true,
                    Port: 443,
                    StringToSend: swov1.Pointer("GET / HTTP/1.1\r\nHost: solarwinds.com\r\nConnection: close\r\n\r\n"),
                    StringToExpect: swov1.Pointer("HTTP/1.1 200 OK"),
                },
                UDP: &components.UDP{
                    Enabled: false,
                    Port: 8888,
                    StringToSend: "GET / HTTP/1.1\r\nHost: solarwinds.com\r\nConnection: close\r\n\r\n",
                    StringToExpect: "HTTP/1.1 200 OK",
                },
            },
            Tags: []components.CommonTag{
                components.CommonTag{
                    Key: "environment",
                    Value: "production",
                },
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CommonEntityID != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                  | Type                                                                       | Required                                                                   | Description                                                                |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `ctx`                                                                      | [context.Context](https://pkg.go.dev/context#Context)                      | :heavy_check_mark:                                                         | The context to use for the request.                                        |
| `request`                                                                  | [operations.UpdateURIRequest](../../models/operations/updateurirequest.md) | :heavy_check_mark:                                                         | The request object to use for the request.                                 |
| `opts`                                                                     | [][operations.Option](../../models/operations/option.md)                   | :heavy_minus_sign:                                                         | The options for this request.                                              |

### Response

**[*operations.UpdateURIResponse](../../models/operations/updateuriresponse.md), error**

### Errors

| Error Type                                | Status Code                               | Content Type                              |
| ----------------------------------------- | ----------------------------------------- | ----------------------------------------- |
| apierrors.CommonBadRequestErrorResponse   | 400                                       | application/json                          |
| apierrors.CommonUnauthorizedErrorResponse | 401                                       | application/json                          |
| apierrors.CommonNotFoundErrorResponse     | 404                                       | application/json                          |
| apierrors.CommonInternalErrorResponse     | 500                                       | application/json                          |
| apierrors.APIError                        | 4XX, 5XX                                  | \*/\*                                     |

## DeleteURI

Delete URI

### Example Usage

<!-- UsageSnippet language="go" operationID="deleteUri" method="delete" path="/v1/dem/uris/{entityId}" -->
```go
package main

import(
	"context"
	"os"
	"github.com/solarwinds/swo-sdk-go/swov1"
	"github.com/solarwinds/swo-sdk-go/swov1/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := swov1.New(
        swov1.WithSecurity(os.Getenv("SWO_API_TOKEN")),
    )

    res, err := s.Dem.DeleteURI(ctx, operations.DeleteURIRequest{
        EntityID: "<id>",
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

| Parameter                                                                  | Type                                                                       | Required                                                                   | Description                                                                |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `ctx`                                                                      | [context.Context](https://pkg.go.dev/context#Context)                      | :heavy_check_mark:                                                         | The context to use for the request.                                        |
| `request`                                                                  | [operations.DeleteURIRequest](../../models/operations/deleteurirequest.md) | :heavy_check_mark:                                                         | The request object to use for the request.                                 |
| `opts`                                                                     | [][operations.Option](../../models/operations/option.md)                   | :heavy_minus_sign:                                                         | The options for this request.                                              |

### Response

**[*operations.DeleteURIResponse](../../models/operations/deleteuriresponse.md), error**

### Errors

| Error Type                                | Status Code                               | Content Type                              |
| ----------------------------------------- | ----------------------------------------- | ----------------------------------------- |
| apierrors.CommonUnauthorizedErrorResponse | 401                                       | application/json                          |
| apierrors.CommonNotFoundErrorResponse     | 404                                       | application/json                          |
| apierrors.CommonInternalErrorResponse     | 500                                       | application/json                          |
| apierrors.APIError                        | 4XX, 5XX                                  | \*/\*                                     |

## GetURIOutageStatuses

Get outage statuses

### Example Usage

<!-- UsageSnippet language="go" operationID="getUriOutageStatuses" method="get" path="/v1/dem/uris/{entityId}/outageStatuses" -->
```go
package main

import(
	"context"
	"os"
	"github.com/solarwinds/swo-sdk-go/swov1"
	"github.com/solarwinds/swo-sdk-go/swov1/types"
	"github.com/solarwinds/swo-sdk-go/swov1/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := swov1.New(
        swov1.WithSecurity(os.Getenv("SWO_API_TOKEN")),
    )

    res, err := s.Dem.GetURIOutageStatuses(ctx, operations.GetURIOutageStatusesRequest{
        EntityID: "<id>",
        StartTime: types.MustTimeFromString("2025-01-07T04:04:57.949Z"),
        EndTime: types.MustTimeFromString("2026-12-19T15:16:56.899Z"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        for {
            // handle items

            res, err = res.Next()

            if err != nil {
                // handle error
            }

            if res == nil {
                break
            }
        }
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.GetURIOutageStatusesRequest](../../models/operations/geturioutagestatusesrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `opts`                                                                                           | [][operations.Option](../../models/operations/option.md)                                         | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |

### Response

**[*operations.GetURIOutageStatusesResponse](../../models/operations/geturioutagestatusesresponse.md), error**

### Errors

| Error Type                                | Status Code                               | Content Type                              |
| ----------------------------------------- | ----------------------------------------- | ----------------------------------------- |
| apierrors.CommonBadRequestErrorResponse   | 400                                       | application/json                          |
| apierrors.CommonUnauthorizedErrorResponse | 401                                       | application/json                          |
| apierrors.CommonInternalErrorResponse     | 500                                       | application/json                          |
| apierrors.APIError                        | 4XX, 5XX                                  | \*/\*                                     |

## PauseURIMonitoring

Pause monitoring of the URI

### Example Usage

<!-- UsageSnippet language="go" operationID="pauseUriMonitoring" method="put" path="/v1/dem/uris/{entityId}/pauseMonitoring" -->
```go
package main

import(
	"context"
	"os"
	"github.com/solarwinds/swo-sdk-go/swov1"
	"github.com/solarwinds/swo-sdk-go/swov1/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := swov1.New(
        swov1.WithSecurity(os.Getenv("SWO_API_TOKEN")),
    )

    res, err := s.Dem.PauseURIMonitoring(ctx, operations.PauseURIMonitoringRequest{
        EntityID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CommonEntityID != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.PauseURIMonitoringRequest](../../models/operations/pauseurimonitoringrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../models/operations/option.md)                                     | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.PauseURIMonitoringResponse](../../models/operations/pauseurimonitoringresponse.md), error**

### Errors

| Error Type                                | Status Code                               | Content Type                              |
| ----------------------------------------- | ----------------------------------------- | ----------------------------------------- |
| apierrors.CommonUnauthorizedErrorResponse | 401                                       | application/json                          |
| apierrors.CommonNotFoundErrorResponse     | 404                                       | application/json                          |
| apierrors.CommonInternalErrorResponse     | 500                                       | application/json                          |
| apierrors.APIError                        | 4XX, 5XX                                  | \*/\*                                     |

## GetURITestResults

Get test results

### Example Usage

<!-- UsageSnippet language="go" operationID="getUriTestResults" method="get" path="/v1/dem/uris/{entityId}/testResults" -->
```go
package main

import(
	"context"
	"os"
	"github.com/solarwinds/swo-sdk-go/swov1"
	"github.com/solarwinds/swo-sdk-go/swov1/types"
	"github.com/solarwinds/swo-sdk-go/swov1/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := swov1.New(
        swov1.WithSecurity(os.Getenv("SWO_API_TOKEN")),
    )

    res, err := s.Dem.GetURITestResults(ctx, operations.GetURITestResultsRequest{
        EntityID: "<id>",
        StartTime: types.MustTimeFromString("2026-04-27T21:55:07.961Z"),
        EndTime: types.MustTimeFromString("2024-02-25T12:00:24.429Z"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        for {
            // handle items

            res, err = res.Next()

            if err != nil {
                // handle error
            }

            if res == nil {
                break
            }
        }
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.GetURITestResultsRequest](../../models/operations/geturitestresultsrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `opts`                                                                                     | [][operations.Option](../../models/operations/option.md)                                   | :heavy_minus_sign:                                                                         | The options for this request.                                                              |

### Response

**[*operations.GetURITestResultsResponse](../../models/operations/geturitestresultsresponse.md), error**

### Errors

| Error Type                                | Status Code                               | Content Type                              |
| ----------------------------------------- | ----------------------------------------- | ----------------------------------------- |
| apierrors.CommonBadRequestErrorResponse   | 400                                       | application/json                          |
| apierrors.CommonUnauthorizedErrorResponse | 401                                       | application/json                          |
| apierrors.CommonInternalErrorResponse     | 500                                       | application/json                          |
| apierrors.APIError                        | 4XX, 5XX                                  | \*/\*                                     |

## UnpauseURIMonitoring

Unpause monitoring of the URI

### Example Usage

<!-- UsageSnippet language="go" operationID="unpauseUriMonitoring" method="put" path="/v1/dem/uris/{entityId}/unpauseMonitoring" -->
```go
package main

import(
	"context"
	"os"
	"github.com/solarwinds/swo-sdk-go/swov1"
	"github.com/solarwinds/swo-sdk-go/swov1/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := swov1.New(
        swov1.WithSecurity(os.Getenv("SWO_API_TOKEN")),
    )

    res, err := s.Dem.UnpauseURIMonitoring(ctx, operations.UnpauseURIMonitoringRequest{
        EntityID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CommonEntityID != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.UnpauseURIMonitoringRequest](../../models/operations/unpauseurimonitoringrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `opts`                                                                                           | [][operations.Option](../../models/operations/option.md)                                         | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |

### Response

**[*operations.UnpauseURIMonitoringResponse](../../models/operations/unpauseurimonitoringresponse.md), error**

### Errors

| Error Type                                | Status Code                               | Content Type                              |
| ----------------------------------------- | ----------------------------------------- | ----------------------------------------- |
| apierrors.CommonUnauthorizedErrorResponse | 401                                       | application/json                          |
| apierrors.CommonNotFoundErrorResponse     | 404                                       | application/json                          |
| apierrors.CommonInternalErrorResponse     | 500                                       | application/json                          |
| apierrors.APIError                        | 4XX, 5XX                                  | \*/\*                                     |

## CreateWebsite

Create website monitoring configuration

### Example Usage

<!-- UsageSnippet language="go" operationID="createWebsite" method="post" path="/v1/dem/websites" -->
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

    res, err := s.Dem.CreateWebsite(ctx, components.DemWebsite{
        Name: "solarwinds.com",
        URL: "https://www.solarwinds.com",
        AvailabilityCheckSettings: &components.AvailabilityCheckSettings{
            PlatformOptions: &components.DemWebsitePlatformOptions{
                ProbePlatforms: []components.DemProbePlatform{
                    components.DemProbePlatformAws,
                },
                TestFromAll: swov1.Pointer(true),
            },
            TestFrom: components.DemTestFrom{
                Type: components.TypeRegion,
                Values: []string{
                    "NA",
                },
            },
            TestIntervalInSeconds: 14400,
            OutageConfiguration: &components.DemWebsiteOutageConfiguration{
                FailingTestLocations: components.DemWebsiteFailingTestLocationsAll,
                ConsecutiveForDown: 2,
            },
            CheckForString: &components.CheckForString{
                Operator: components.OperatorContains,
                Value: "string",
            },
            Protocols: []components.DemWebsiteProtocol{
                components.DemWebsiteProtocolHTTP,
                components.DemWebsiteProtocolHTTPS,
            },
            Ssl: &components.Ssl{
                Enabled: swov1.Pointer(true),
                DaysPriorToExpiration: swov1.Pointer[int](7),
                IgnoreIntermediateCertificates: swov1.Pointer(true),
            },
            CustomHeaders: []components.DemCustomHeaders{
                components.DemCustomHeaders{
                    Name: "string",
                    Value: "string",
                },
            },
            AllowInsecureRenegotiation: swov1.Pointer(true),
            PostData: swov1.Pointer("{\"example\": \"value\"}"),
        },
        Tags: []components.CommonTag{
            components.CommonTag{
                Key: "environment",
                Value: "production",
            },
        },
        Rum: &components.Rum{
            ApdexTimeInSeconds: swov1.Pointer[int](4),
            Spa: true,
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CommonEntityID != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                      | Type                                                           | Required                                                       | Description                                                    |
| -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- |
| `ctx`                                                          | [context.Context](https://pkg.go.dev/context#Context)          | :heavy_check_mark:                                             | The context to use for the request.                            |
| `request`                                                      | [components.DemWebsite](../../models/components/demwebsite.md) | :heavy_check_mark:                                             | The request object to use for the request.                     |
| `opts`                                                         | [][operations.Option](../../models/operations/option.md)       | :heavy_minus_sign:                                             | The options for this request.                                  |

### Response

**[*operations.CreateWebsiteResponse](../../models/operations/createwebsiteresponse.md), error**

### Errors

| Error Type                                | Status Code                               | Content Type                              |
| ----------------------------------------- | ----------------------------------------- | ----------------------------------------- |
| apierrors.CommonBadRequestErrorResponse   | 400                                       | application/json                          |
| apierrors.CommonUnauthorizedErrorResponse | 401                                       | application/json                          |
| apierrors.CommonInternalErrorResponse     | 500                                       | application/json                          |
| apierrors.APIError                        | 4XX, 5XX                                  | \*/\*                                     |

## GetWebsite

Get website monitoring configuration

### Example Usage

<!-- UsageSnippet language="go" operationID="getWebsite" method="get" path="/v1/dem/websites/{entityId}" -->
```go
package main

import(
	"context"
	"os"
	"github.com/solarwinds/swo-sdk-go/swov1"
	"github.com/solarwinds/swo-sdk-go/swov1/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := swov1.New(
        swov1.WithSecurity(os.Getenv("SWO_API_TOKEN")),
    )

    res, err := s.Dem.GetWebsite(ctx, operations.GetWebsiteRequest{
        EntityID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DemGetWebsiteResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |
| `request`                                                                    | [operations.GetWebsiteRequest](../../models/operations/getwebsiterequest.md) | :heavy_check_mark:                                                           | The request object to use for the request.                                   |
| `opts`                                                                       | [][operations.Option](../../models/operations/option.md)                     | :heavy_minus_sign:                                                           | The options for this request.                                                |

### Response

**[*operations.GetWebsiteResponse](../../models/operations/getwebsiteresponse.md), error**

### Errors

| Error Type                                | Status Code                               | Content Type                              |
| ----------------------------------------- | ----------------------------------------- | ----------------------------------------- |
| apierrors.CommonUnauthorizedErrorResponse | 401                                       | application/json                          |
| apierrors.CommonNotFoundErrorResponse     | 404                                       | application/json                          |
| apierrors.CommonInternalErrorResponse     | 500                                       | application/json                          |
| apierrors.APIError                        | 4XX, 5XX                                  | \*/\*                                     |

## UpdateWebsite

Update website monitoring configuration

### Example Usage

<!-- UsageSnippet language="go" operationID="updateWebsite" method="put" path="/v1/dem/websites/{entityId}" -->
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

    res, err := s.Dem.UpdateWebsite(ctx, operations.UpdateWebsiteRequest{
        EntityID: "<id>",
        DemWebsite: components.DemWebsite{
            Name: "solarwinds.com",
            URL: "https://www.solarwinds.com",
            AvailabilityCheckSettings: &components.AvailabilityCheckSettings{
                PlatformOptions: &components.DemWebsitePlatformOptions{
                    ProbePlatforms: []components.DemProbePlatform{
                        components.DemProbePlatformAws,
                    },
                    TestFromAll: swov1.Pointer(true),
                },
                TestFrom: components.DemTestFrom{
                    Type: components.TypeRegion,
                    Values: []string{
                        "NA",
                    },
                },
                TestIntervalInSeconds: 14400,
                OutageConfiguration: &components.DemWebsiteOutageConfiguration{
                    FailingTestLocations: components.DemWebsiteFailingTestLocationsAll,
                    ConsecutiveForDown: 2,
                },
                CheckForString: &components.CheckForString{
                    Operator: components.OperatorContains,
                    Value: "string",
                },
                Protocols: []components.DemWebsiteProtocol{
                    components.DemWebsiteProtocolHTTP,
                    components.DemWebsiteProtocolHTTPS,
                },
                Ssl: &components.Ssl{
                    Enabled: swov1.Pointer(true),
                    DaysPriorToExpiration: swov1.Pointer[int](7),
                    IgnoreIntermediateCertificates: swov1.Pointer(true),
                },
                CustomHeaders: []components.DemCustomHeaders{
                    components.DemCustomHeaders{
                        Name: "string",
                        Value: "string",
                    },
                },
                AllowInsecureRenegotiation: swov1.Pointer(true),
                PostData: swov1.Pointer("{\"example\": \"value\"}"),
            },
            Tags: []components.CommonTag{
                components.CommonTag{
                    Key: "environment",
                    Value: "production",
                },
            },
            Rum: &components.Rum{
                ApdexTimeInSeconds: swov1.Pointer[int](4),
                Spa: true,
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CommonEntityID != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.UpdateWebsiteRequest](../../models/operations/updatewebsiterequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |
| `opts`                                                                             | [][operations.Option](../../models/operations/option.md)                           | :heavy_minus_sign:                                                                 | The options for this request.                                                      |

### Response

**[*operations.UpdateWebsiteResponse](../../models/operations/updatewebsiteresponse.md), error**

### Errors

| Error Type                                | Status Code                               | Content Type                              |
| ----------------------------------------- | ----------------------------------------- | ----------------------------------------- |
| apierrors.CommonBadRequestErrorResponse   | 400                                       | application/json                          |
| apierrors.CommonUnauthorizedErrorResponse | 401                                       | application/json                          |
| apierrors.CommonNotFoundErrorResponse     | 404                                       | application/json                          |
| apierrors.CommonInternalErrorResponse     | 500                                       | application/json                          |
| apierrors.APIError                        | 4XX, 5XX                                  | \*/\*                                     |

## DeleteWebsite

Delete website

### Example Usage

<!-- UsageSnippet language="go" operationID="deleteWebsite" method="delete" path="/v1/dem/websites/{entityId}" -->
```go
package main

import(
	"context"
	"os"
	"github.com/solarwinds/swo-sdk-go/swov1"
	"github.com/solarwinds/swo-sdk-go/swov1/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := swov1.New(
        swov1.WithSecurity(os.Getenv("SWO_API_TOKEN")),
    )

    res, err := s.Dem.DeleteWebsite(ctx, operations.DeleteWebsiteRequest{
        EntityID: "<id>",
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

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.DeleteWebsiteRequest](../../models/operations/deletewebsiterequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |
| `opts`                                                                             | [][operations.Option](../../models/operations/option.md)                           | :heavy_minus_sign:                                                                 | The options for this request.                                                      |

### Response

**[*operations.DeleteWebsiteResponse](../../models/operations/deletewebsiteresponse.md), error**

### Errors

| Error Type                                | Status Code                               | Content Type                              |
| ----------------------------------------- | ----------------------------------------- | ----------------------------------------- |
| apierrors.CommonUnauthorizedErrorResponse | 401                                       | application/json                          |
| apierrors.CommonNotFoundErrorResponse     | 404                                       | application/json                          |
| apierrors.CommonInternalErrorResponse     | 500                                       | application/json                          |
| apierrors.APIError                        | 4XX, 5XX                                  | \*/\*                                     |

## GetWebsiteOutageStatuses

Get outage statuses

### Example Usage

<!-- UsageSnippet language="go" operationID="getWebsiteOutageStatuses" method="get" path="/v1/dem/websites/{entityId}/outageStatuses" -->
```go
package main

import(
	"context"
	"os"
	"github.com/solarwinds/swo-sdk-go/swov1"
	"github.com/solarwinds/swo-sdk-go/swov1/types"
	"github.com/solarwinds/swo-sdk-go/swov1/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := swov1.New(
        swov1.WithSecurity(os.Getenv("SWO_API_TOKEN")),
    )

    res, err := s.Dem.GetWebsiteOutageStatuses(ctx, operations.GetWebsiteOutageStatusesRequest{
        EntityID: "<id>",
        StartTime: types.MustTimeFromString("2026-11-15T10:02:22.738Z"),
        EndTime: types.MustTimeFromString("2026-02-24T15:45:54.358Z"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        for {
            // handle items

            res, err = res.Next()

            if err != nil {
                // handle error
            }

            if res == nil {
                break
            }
        }
    }
}
```

### Parameters

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.GetWebsiteOutageStatusesRequest](../../models/operations/getwebsiteoutagestatusesrequest.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |
| `opts`                                                                                                   | [][operations.Option](../../models/operations/option.md)                                                 | :heavy_minus_sign:                                                                                       | The options for this request.                                                                            |

### Response

**[*operations.GetWebsiteOutageStatusesResponse](../../models/operations/getwebsiteoutagestatusesresponse.md), error**

### Errors

| Error Type                                | Status Code                               | Content Type                              |
| ----------------------------------------- | ----------------------------------------- | ----------------------------------------- |
| apierrors.CommonBadRequestErrorResponse   | 400                                       | application/json                          |
| apierrors.CommonUnauthorizedErrorResponse | 401                                       | application/json                          |
| apierrors.CommonInternalErrorResponse     | 500                                       | application/json                          |
| apierrors.APIError                        | 4XX, 5XX                                  | \*/\*                                     |

## PauseWebsiteMonitoring

Pause monitoring of a website

### Example Usage

<!-- UsageSnippet language="go" operationID="pauseWebsiteMonitoring" method="put" path="/v1/dem/websites/{entityId}/pauseMonitoring" -->
```go
package main

import(
	"context"
	"os"
	"github.com/solarwinds/swo-sdk-go/swov1"
	"github.com/solarwinds/swo-sdk-go/swov1/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := swov1.New(
        swov1.WithSecurity(os.Getenv("SWO_API_TOKEN")),
    )

    res, err := s.Dem.PauseWebsiteMonitoring(ctx, operations.PauseWebsiteMonitoringRequest{
        EntityID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CommonEntityID != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.PauseWebsiteMonitoringRequest](../../models/operations/pausewebsitemonitoringrequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `opts`                                                                                               | [][operations.Option](../../models/operations/option.md)                                             | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |

### Response

**[*operations.PauseWebsiteMonitoringResponse](../../models/operations/pausewebsitemonitoringresponse.md), error**

### Errors

| Error Type                                | Status Code                               | Content Type                              |
| ----------------------------------------- | ----------------------------------------- | ----------------------------------------- |
| apierrors.CommonUnauthorizedErrorResponse | 401                                       | application/json                          |
| apierrors.CommonNotFoundErrorResponse     | 404                                       | application/json                          |
| apierrors.CommonInternalErrorResponse     | 500                                       | application/json                          |
| apierrors.APIError                        | 4XX, 5XX                                  | \*/\*                                     |

## GetWebsiteTestResults

Get test results

### Example Usage

<!-- UsageSnippet language="go" operationID="getWebsiteTestResults" method="get" path="/v1/dem/websites/{entityId}/testResults" -->
```go
package main

import(
	"context"
	"os"
	"github.com/solarwinds/swo-sdk-go/swov1"
	"github.com/solarwinds/swo-sdk-go/swov1/types"
	"github.com/solarwinds/swo-sdk-go/swov1/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := swov1.New(
        swov1.WithSecurity(os.Getenv("SWO_API_TOKEN")),
    )

    res, err := s.Dem.GetWebsiteTestResults(ctx, operations.GetWebsiteTestResultsRequest{
        EntityID: "<id>",
        StartTime: types.MustTimeFromString("2024-05-25T23:02:07.467Z"),
        EndTime: types.MustTimeFromString("2024-06-25T13:36:39.740Z"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        for {
            // handle items

            res, err = res.Next()

            if err != nil {
                // handle error
            }

            if res == nil {
                break
            }
        }
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.GetWebsiteTestResultsRequest](../../models/operations/getwebsitetestresultsrequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `opts`                                                                                             | [][operations.Option](../../models/operations/option.md)                                           | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |

### Response

**[*operations.GetWebsiteTestResultsResponse](../../models/operations/getwebsitetestresultsresponse.md), error**

### Errors

| Error Type                                | Status Code                               | Content Type                              |
| ----------------------------------------- | ----------------------------------------- | ----------------------------------------- |
| apierrors.CommonBadRequestErrorResponse   | 400                                       | application/json                          |
| apierrors.CommonUnauthorizedErrorResponse | 401                                       | application/json                          |
| apierrors.CommonInternalErrorResponse     | 500                                       | application/json                          |
| apierrors.APIError                        | 4XX, 5XX                                  | \*/\*                                     |

## UnpauseWebsiteMonitoring

Unpause monitoring of a website

### Example Usage

<!-- UsageSnippet language="go" operationID="unpauseWebsiteMonitoring" method="put" path="/v1/dem/websites/{entityId}/unpauseMonitoring" -->
```go
package main

import(
	"context"
	"os"
	"github.com/solarwinds/swo-sdk-go/swov1"
	"github.com/solarwinds/swo-sdk-go/swov1/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := swov1.New(
        swov1.WithSecurity(os.Getenv("SWO_API_TOKEN")),
    )

    res, err := s.Dem.UnpauseWebsiteMonitoring(ctx, operations.UnpauseWebsiteMonitoringRequest{
        EntityID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CommonEntityID != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.UnpauseWebsiteMonitoringRequest](../../models/operations/unpausewebsitemonitoringrequest.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |
| `opts`                                                                                                   | [][operations.Option](../../models/operations/option.md)                                                 | :heavy_minus_sign:                                                                                       | The options for this request.                                                                            |

### Response

**[*operations.UnpauseWebsiteMonitoringResponse](../../models/operations/unpausewebsitemonitoringresponse.md), error**

### Errors

| Error Type                                | Status Code                               | Content Type                              |
| ----------------------------------------- | ----------------------------------------- | ----------------------------------------- |
| apierrors.CommonUnauthorizedErrorResponse | 401                                       | application/json                          |
| apierrors.CommonNotFoundErrorResponse     | 404                                       | application/json                          |
| apierrors.CommonInternalErrorResponse     | 500                                       | application/json                          |
| apierrors.APIError                        | 4XX, 5XX                                  | \*/\*                                     |