# Dem
(*Dem*)

## Overview

### Available Operations

* [GetDemSettings](#getdemsettings) - Get DEM settings
* [SetDemSettings](#setdemsettings) - Set DEM settings
* [CreateWebsite](#createwebsite) - Create website monitoring configuration
* [GetWebsite](#getwebsite) - Get website monitoring configuration
* [UpdateWebsite](#updatewebsite) - Update website monitoring configuration
* [DeleteWebsite](#deletewebsite) - Delete website
* [PauseWebsiteMonitoring](#pausewebsitemonitoring) - Pause monitoring of a website
* [UnpauseWebsiteMonitoring](#unpausewebsitemonitoring) - Unpause monitoring of a website

## GetDemSettings

Get DEM settings

### Example Usage

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
    if res.OutageConfiguration != nil {
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

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## SetDemSettings

Set DEM settings

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

    res, err := s.Dem.SetDemSettings(ctx, components.OutageConfiguration{
        FailingTestLocations: components.FailingTestLocationsAll,
        ConsecutiveForDown: 2,
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

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [components.OutageConfiguration](../../models/components/outageconfiguration.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |
| `opts`                                                                           | [][operations.Option](../../models/operations/option.md)                         | :heavy_minus_sign:                                                               | The options for this request.                                                    |

### Response

**[*operations.SetDemSettingsResponse](../../models/operations/setdemsettingsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## CreateWebsite

Create website monitoring configuration

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

    res, err := s.Dem.CreateWebsite(ctx, components.Website{
        Name: "solarwinds.com",
        URL: "https://www.solarwinds.com",
        AvailabilityCheckSettings: &components.AvailabilityCheckSettings{
            CheckForString: &components.CheckForString{
                Operator: components.CheckForStringOperatorContains,
                Value: "string",
            },
            TestIntervalInSeconds: 14400,
            Protocols: []components.WebsiteProtocol{
                components.WebsiteProtocolHTTP,
                components.WebsiteProtocolHTTPS,
            },
            PlatformOptions: &components.ProbePlatformOptions{
                ProbePlatforms: []components.ProbePlatform{
                    components.ProbePlatformAws,
                },
                TestFromAll: swov1.Bool(true),
            },
            TestFrom: components.TestFrom{
                Type: components.ProbeLocationTypeRegion,
                Values: []string{
                    "NA",
                },
            },
            Ssl: &components.Ssl{
                Enabled: swov1.Bool(true),
                DaysPriorToExpiration: swov1.Int(7),
                IgnoreIntermediateCertificates: swov1.Bool(true),
            },
            CustomHeaders: []components.CustomHeaders{
                components.CustomHeaders{
                    Name: "string",
                    Value: "string",
                },
            },
            AllowInsecureRenegotiation: swov1.Bool(true),
            PostData: swov1.String("{\"example\": \"value\"}"),
            OutageConfiguration: &components.WebsiteOutageConfiguration{
                FailingTestLocations: components.WebsiteFailingTestLocationsAll,
                ConsecutiveForDown: 2,
            },
        },
        Tags: []components.Tag{
            components.Tag{
                Key: "environment",
                Value: "production",
            },
        },
        Rum: &components.Rum{
            ApdexTimeInSeconds: swov1.Int(4),
            Spa: true,
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.EntityID != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `request`                                                | [components.Website](../../models/components/website.md) | :heavy_check_mark:                                       | The request object to use for the request.               |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.CreateWebsiteResponse](../../models/operations/createwebsiteresponse.md), error**

### Errors

| Error Type                          | Status Code                         | Content Type                        |
| ----------------------------------- | ----------------------------------- | ----------------------------------- |
| apierrors.CreateWebsiteResponseBody | 400                                 | application/json                    |
| apierrors.APIError                  | 4XX, 5XX                            | \*/\*                               |

## GetWebsite

Get website monitoring configuration

### Example Usage

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
    if res.Object != nil {
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

| Error Type                       | Status Code                      | Content Type                     |
| -------------------------------- | -------------------------------- | -------------------------------- |
| apierrors.GetWebsiteResponseBody | 400                              | application/json                 |
| apierrors.APIError               | 4XX, 5XX                         | \*/\*                            |

## UpdateWebsite

Update website monitoring configuration

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

    res, err := s.Dem.UpdateWebsite(ctx, operations.UpdateWebsiteRequest{
        EntityID: "<id>",
        Website: components.Website{
            Name: "solarwinds.com",
            URL: "https://www.solarwinds.com",
            AvailabilityCheckSettings: &components.AvailabilityCheckSettings{
                CheckForString: &components.CheckForString{
                    Operator: components.CheckForStringOperatorContains,
                    Value: "string",
                },
                TestIntervalInSeconds: 14400,
                Protocols: []components.WebsiteProtocol{
                    components.WebsiteProtocolHTTP,
                    components.WebsiteProtocolHTTPS,
                },
                PlatformOptions: &components.ProbePlatformOptions{
                    ProbePlatforms: []components.ProbePlatform{
                        components.ProbePlatformAws,
                    },
                    TestFromAll: swov1.Bool(true),
                },
                TestFrom: components.TestFrom{
                    Type: components.ProbeLocationTypeRegion,
                    Values: []string{
                        "NA",
                    },
                },
                Ssl: &components.Ssl{
                    Enabled: swov1.Bool(true),
                    DaysPriorToExpiration: swov1.Int(7),
                    IgnoreIntermediateCertificates: swov1.Bool(true),
                },
                CustomHeaders: []components.CustomHeaders{
                    components.CustomHeaders{
                        Name: "string",
                        Value: "string",
                    },
                },
                AllowInsecureRenegotiation: swov1.Bool(true),
                PostData: swov1.String("{\"example\": \"value\"}"),
                OutageConfiguration: &components.WebsiteOutageConfiguration{
                    FailingTestLocations: components.WebsiteFailingTestLocationsAll,
                    ConsecutiveForDown: 2,
                },
            },
            Tags: []components.Tag{
                components.Tag{
                    Key: "environment",
                    Value: "production",
                },
            },
            Rum: &components.Rum{
                ApdexTimeInSeconds: swov1.Int(4),
                Spa: true,
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.EntityID != nil {
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

| Error Type                          | Status Code                         | Content Type                        |
| ----------------------------------- | ----------------------------------- | ----------------------------------- |
| apierrors.UpdateWebsiteResponseBody | 400                                 | application/json                    |
| apierrors.APIError                  | 4XX, 5XX                            | \*/\*                               |

## DeleteWebsite

Delete website

### Example Usage

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
    if res.EntityID != nil {
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

| Error Type                          | Status Code                         | Content Type                        |
| ----------------------------------- | ----------------------------------- | ----------------------------------- |
| apierrors.DeleteWebsiteResponseBody | 400                                 | application/json                    |
| apierrors.APIError                  | 4XX, 5XX                            | \*/\*                               |

## PauseWebsiteMonitoring

Pause monitoring of a website

### Example Usage

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
    if res.EntityID != nil {
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

| Error Type                                   | Status Code                                  | Content Type                                 |
| -------------------------------------------- | -------------------------------------------- | -------------------------------------------- |
| apierrors.PauseWebsiteMonitoringResponseBody | 400                                          | application/json                             |
| apierrors.APIError                           | 4XX, 5XX                                     | \*/\*                                        |

## UnpauseWebsiteMonitoring

Unpause monitoring of a website

### Example Usage

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
    if res.EntityID != nil {
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

| Error Type                                     | Status Code                                    | Content Type                                   |
| ---------------------------------------------- | ---------------------------------------------- | ---------------------------------------------- |
| apierrors.UnpauseWebsiteMonitoringResponseBody | 400                                            | application/json                               |
| apierrors.APIError                             | 4XX, 5XX                                       | \*/\*                                          |