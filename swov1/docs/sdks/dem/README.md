# Dem
(*Dem*)

## Overview

### Available Operations

* [CreateWebsiteMonitor](#createwebsitemonitor) - Create website monitoring configuration
* [GetWebsiteMonitor](#getwebsitemonitor) - Get website monitoring configuration
* [UpdateWebsiteMonitor](#updatewebsitemonitor) - Update website monitoring configuration
* [DeleteWebsiteMonitor](#deletewebsitemonitor) - Delete website
* [PauseWebsiteMonitor](#pausewebsitemonitor) - Pause monitoring of a website
* [UnpauseWebsiteMonitor](#unpausewebsitemonitor) - Unpause monitoring of a website

## CreateWebsiteMonitor

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

    res, err := s.Dem.CreateWebsiteMonitor(ctx, components.Website{
        Name: "solarwinds.com",
        URL: "https://www.solarwinds.com",
        AvailabilityCheckSettings: &components.AvailabilityCheckSettings{
            CheckForString: &components.CheckForString{
                Operator: components.OperatorContains,
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
            Ssl: &components.SslMonitoring{
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
    if res.Object != nil {
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

**[*operations.CreateWebsiteMonitorResponse](../../models/operations/createwebsitemonitorresponse.md), error**

### Errors

| Error Type                                 | Status Code                                | Content Type                               |
| ------------------------------------------ | ------------------------------------------ | ------------------------------------------ |
| apierrors.CreateWebsiteMonitorResponseBody | 400                                        | application/json                           |
| apierrors.APIError                         | 4XX, 5XX                                   | \*/\*                                      |

## GetWebsiteMonitor

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

    res, err := s.Dem.GetWebsiteMonitor(ctx, operations.GetWebsiteMonitorRequest{
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

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.GetWebsiteMonitorRequest](../../models/operations/getwebsitemonitorrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `opts`                                                                                     | [][operations.Option](../../models/operations/option.md)                                   | :heavy_minus_sign:                                                                         | The options for this request.                                                              |

### Response

**[*operations.GetWebsiteMonitorResponse](../../models/operations/getwebsitemonitorresponse.md), error**

### Errors

| Error Type                              | Status Code                             | Content Type                            |
| --------------------------------------- | --------------------------------------- | --------------------------------------- |
| apierrors.GetWebsiteMonitorResponseBody | 400                                     | application/json                        |
| apierrors.APIError                      | 4XX, 5XX                                | \*/\*                                   |

## UpdateWebsiteMonitor

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

    res, err := s.Dem.UpdateWebsiteMonitor(ctx, operations.UpdateWebsiteMonitorRequest{
        EntityID: "<id>",
        Website: components.Website{
            Name: "solarwinds.com",
            URL: "https://www.solarwinds.com",
            AvailabilityCheckSettings: &components.AvailabilityCheckSettings{
                CheckForString: &components.CheckForString{
                    Operator: components.OperatorContains,
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
                Ssl: &components.SslMonitoring{
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
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.UpdateWebsiteMonitorRequest](../../models/operations/updatewebsitemonitorrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `opts`                                                                                           | [][operations.Option](../../models/operations/option.md)                                         | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |

### Response

**[*operations.UpdateWebsiteMonitorResponse](../../models/operations/updatewebsitemonitorresponse.md), error**

### Errors

| Error Type                                 | Status Code                                | Content Type                               |
| ------------------------------------------ | ------------------------------------------ | ------------------------------------------ |
| apierrors.UpdateWebsiteMonitorResponseBody | 400                                        | application/json                           |
| apierrors.APIError                         | 4XX, 5XX                                   | \*/\*                                      |

## DeleteWebsiteMonitor

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

    res, err := s.Dem.DeleteWebsiteMonitor(ctx, operations.DeleteWebsiteMonitorRequest{
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

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.DeleteWebsiteMonitorRequest](../../models/operations/deletewebsitemonitorrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `opts`                                                                                           | [][operations.Option](../../models/operations/option.md)                                         | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |

### Response

**[*operations.DeleteWebsiteMonitorResponse](../../models/operations/deletewebsitemonitorresponse.md), error**

### Errors

| Error Type                                 | Status Code                                | Content Type                               |
| ------------------------------------------ | ------------------------------------------ | ------------------------------------------ |
| apierrors.DeleteWebsiteMonitorResponseBody | 400                                        | application/json                           |
| apierrors.APIError                         | 4XX, 5XX                                   | \*/\*                                      |

## PauseWebsiteMonitor

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

    res, err := s.Dem.PauseWebsiteMonitor(ctx, operations.PauseWebsiteMonitorRequest{
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

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.PauseWebsiteMonitorRequest](../../models/operations/pausewebsitemonitorrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../models/operations/option.md)                                       | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.PauseWebsiteMonitorResponse](../../models/operations/pausewebsitemonitorresponse.md), error**

### Errors

| Error Type                                | Status Code                               | Content Type                              |
| ----------------------------------------- | ----------------------------------------- | ----------------------------------------- |
| apierrors.PauseWebsiteMonitorResponseBody | 400                                       | application/json                          |
| apierrors.APIError                        | 4XX, 5XX                                  | \*/\*                                     |

## UnpauseWebsiteMonitor

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

    res, err := s.Dem.UnpauseWebsiteMonitor(ctx, operations.UnpauseWebsiteMonitorRequest{
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

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.UnpauseWebsiteMonitorRequest](../../models/operations/unpausewebsitemonitorrequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `opts`                                                                                             | [][operations.Option](../../models/operations/option.md)                                           | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |

### Response

**[*operations.UnpauseWebsiteMonitorResponse](../../models/operations/unpausewebsitemonitorresponse.md), error**

### Errors

| Error Type                                  | Status Code                                 | Content Type                                |
| ------------------------------------------- | ------------------------------------------- | ------------------------------------------- |
| apierrors.UnpauseWebsiteMonitorResponseBody | 400                                         | application/json                            |
| apierrors.APIError                          | 4XX, 5XX                                    | \*/\*                                       |