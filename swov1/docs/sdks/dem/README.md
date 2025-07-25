# Dem
(*Dem*)

## Overview

### Available Operations

* [ListProbes](#listprobes) - Get a list of existing synthetic probes
* [GetDemSettings](#getdemsettings) - Get DEM settings
* [SetDemSettings](#setdemsettings) - Set DEM settings
* [CreateURI](#createuri) - Create URI monitoring configuration
* [GetURI](#geturi) - Get URI monitoring configuration
* [UpdateURI](#updateuri) - Update URI monitoring configuration
* [DeleteURI](#deleteuri) - Delete URI
* [PauseURIMonitoring](#pauseurimonitoring) - Pause monitoring of the URI
* [UnpauseURIMonitoring](#unpauseurimonitoring) - Unpause monitoring of the URI
* [CreateWebsite](#createwebsite) - Create website monitoring configuration
* [GetWebsite](#getwebsite) - Get website monitoring configuration
* [UpdateWebsite](#updatewebsite) - Update website monitoring configuration
* [DeleteWebsite](#deletewebsite) - Delete website
* [PauseWebsiteMonitoring](#pausewebsitemonitoring) - Pause monitoring of a website
* [UnpauseWebsiteMonitoring](#unpausewebsitemonitoring) - Unpause monitoring of a website

## ListProbes

Get a list of existing synthetic probes

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

| Error Type                       | Status Code                      | Content Type                     |
| -------------------------------- | -------------------------------- | -------------------------------- |
| apierrors.ListProbesResponseBody | 500                              | application/json                 |
| apierrors.APIError               | 4XX, 5XX                         | \*/\*                            |

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
    if res.DemOutageConfiguration != nil {
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

    res, err := s.Dem.SetDemSettings(ctx, components.DemOutageConfiguration{
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

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [components.DemOutageConfiguration](../../models/components/demoutageconfiguration.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `opts`                                                                                 | [][operations.Option](../../models/operations/option.md)                               | :heavy_minus_sign:                                                                     | The options for this request.                                                          |

### Response

**[*operations.SetDemSettingsResponse](../../models/operations/setdemsettingsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## CreateURI

Create URI monitoring configuration

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

    res, err := s.Dem.CreateURI(ctx, components.DemURI{
        Name: "solarwinds.com",
        IPOrDomain: "solarwinds.com",
        AvailabilityCheckSettings: components.DemURIAvailabilityCheckSettingsInput{
            PlatformOptions: &components.PlatformOptions{
                ProbePlatforms: []components.DemProbePlatform{
                    components.DemProbePlatformAws,
                },
                TestFromAll: swov1.Bool(true),
            },
            TestFrom: components.DemTestFrom{
                Type: components.TypeRegion,
                Values: []string{
                    "NA",
                },
            },
            TestIntervalInSeconds: 14400,
            OutageConfiguration: nil,
            Ping: &components.Ping{
                Enabled: true,
            },
            TCP: &components.TCP{
                Enabled: true,
                Port: 443,
                StringToSend: swov1.String("GET / HTTP/1.1\r\n" +
                "Host: solarwinds.com\r\n" +
                "Connection: close\r\n" +
                "\r\n" +
                ""),
                StringToExpect: swov1.String("HTTP/1.1 200 OK"),
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

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `request`                                                | [components.DemURI](../../models/components/demuri.md)   | :heavy_check_mark:                                       | The request object to use for the request.               |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.CreateURIResponse](../../models/operations/createuriresponse.md), error**

### Errors

| Error Type                      | Status Code                     | Content Type                    |
| ------------------------------- | ------------------------------- | ------------------------------- |
| apierrors.CreateURIResponseBody | 400                             | application/json                |
| apierrors.APIError              | 4XX, 5XX                        | \*/\*                           |

## GetURI

Get URI monitoring configuration

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

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| apierrors.GetURIResponseBody | 404                          | application/json             |
| apierrors.APIError           | 4XX, 5XX                     | \*/\*                        |

## UpdateURI

Update URI monitoring configuration

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

    res, err := s.Dem.UpdateURI(ctx, operations.UpdateURIRequest{
        EntityID: "<id>",
        DemURI: components.DemURI{
            Name: "solarwinds.com",
            IPOrDomain: "solarwinds.com",
            AvailabilityCheckSettings: components.DemURIAvailabilityCheckSettingsInput{
                PlatformOptions: &components.PlatformOptions{
                    ProbePlatforms: []components.DemProbePlatform{
                        components.DemProbePlatformAws,
                    },
                    TestFromAll: swov1.Bool(true),
                },
                TestFrom: components.DemTestFrom{
                    Type: components.TypeRegion,
                    Values: []string{
                        "NA",
                    },
                },
                TestIntervalInSeconds: 14400,
                OutageConfiguration: &components.OutageConfiguration{
                    FailingTestLocations: components.DemURIAvailabilityCheckSettingsInputFailingTestLocationsAll,
                    ConsecutiveForDown: 2,
                },
                Ping: &components.Ping{
                    Enabled: true,
                },
                TCP: &components.TCP{
                    Enabled: true,
                    Port: 443,
                    StringToSend: swov1.String("GET / HTTP/1.1\r\n" +
                    "Host: solarwinds.com\r\n" +
                    "Connection: close\r\n" +
                    "\r\n" +
                    ""),
                    StringToExpect: swov1.String("HTTP/1.1 200 OK"),
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

| Error Type                         | Status Code                        | Content Type                       |
| ---------------------------------- | ---------------------------------- | ---------------------------------- |
| apierrors.UpdateURIResponseBody    | 400                                | application/json                   |
| apierrors.UpdateURIDemResponseBody | 404                                | application/json                   |
| apierrors.APIError                 | 4XX, 5XX                           | \*/\*                              |

## DeleteURI

Delete URI

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

    res, err := s.Dem.DeleteURI(ctx, operations.DeleteURIRequest{
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

| Parameter                                                                  | Type                                                                       | Required                                                                   | Description                                                                |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `ctx`                                                                      | [context.Context](https://pkg.go.dev/context#Context)                      | :heavy_check_mark:                                                         | The context to use for the request.                                        |
| `request`                                                                  | [operations.DeleteURIRequest](../../models/operations/deleteurirequest.md) | :heavy_check_mark:                                                         | The request object to use for the request.                                 |
| `opts`                                                                     | [][operations.Option](../../models/operations/option.md)                   | :heavy_minus_sign:                                                         | The options for this request.                                              |

### Response

**[*operations.DeleteURIResponse](../../models/operations/deleteuriresponse.md), error**

### Errors

| Error Type                      | Status Code                     | Content Type                    |
| ------------------------------- | ------------------------------- | ------------------------------- |
| apierrors.DeleteURIResponseBody | 404                             | application/json                |
| apierrors.APIError              | 4XX, 5XX                        | \*/\*                           |

## PauseURIMonitoring

Pause monitoring of the URI

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

| Error Type                               | Status Code                              | Content Type                             |
| ---------------------------------------- | ---------------------------------------- | ---------------------------------------- |
| apierrors.PauseURIMonitoringResponseBody | 404                                      | application/json                         |
| apierrors.APIError                       | 4XX, 5XX                                 | \*/\*                                    |

## UnpauseURIMonitoring

Unpause monitoring of the URI

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

| Error Type                                 | Status Code                                | Content Type                               |
| ------------------------------------------ | ------------------------------------------ | ------------------------------------------ |
| apierrors.UnpauseURIMonitoringResponseBody | 404                                        | application/json                           |
| apierrors.APIError                         | 4XX, 5XX                                   | \*/\*                                      |

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

    res, err := s.Dem.CreateWebsite(ctx, components.DemWebsite{
        Name: "solarwinds.com",
        URL: "https://www.solarwinds.com",
        AvailabilityCheckSettings: &components.AvailabilityCheckSettings{
            PlatformOptions: &components.DemWebsitePlatformOptions{
                ProbePlatforms: []components.DemProbePlatform{
                    components.DemProbePlatformAws,
                },
                TestFromAll: swov1.Bool(true),
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
                Enabled: swov1.Bool(true),
                DaysPriorToExpiration: swov1.Int(7),
                IgnoreIntermediateCertificates: swov1.Bool(true),
            },
            CustomHeaders: []components.DemCustomHeaders{
                components.DemCustomHeaders{
                    Name: "string",
                    Value: "string",
                },
            },
            AllowInsecureRenegotiation: swov1.Bool(true),
            PostData: swov1.String("{\"example\": \"value\"}"),
        },
        Tags: []components.CommonTag{
            components.CommonTag{
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

| Error Type                       | Status Code                      | Content Type                     |
| -------------------------------- | -------------------------------- | -------------------------------- |
| apierrors.GetWebsiteResponseBody | 404                              | application/json                 |
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
        DemWebsite: components.DemWebsite{
            Name: "solarwinds.com",
            URL: "https://www.solarwinds.com",
            AvailabilityCheckSettings: &components.AvailabilityCheckSettings{
                PlatformOptions: &components.DemWebsitePlatformOptions{
                    ProbePlatforms: []components.DemProbePlatform{
                        components.DemProbePlatformAws,
                    },
                    TestFromAll: swov1.Bool(true),
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
                    Enabled: swov1.Bool(true),
                    DaysPriorToExpiration: swov1.Int(7),
                    IgnoreIntermediateCertificates: swov1.Bool(true),
                },
                CustomHeaders: []components.DemCustomHeaders{
                    components.DemCustomHeaders{
                        Name: "string",
                        Value: "string",
                    },
                },
                AllowInsecureRenegotiation: swov1.Bool(true),
                PostData: swov1.String("{\"example\": \"value\"}"),
            },
            Tags: []components.CommonTag{
                components.CommonTag{
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

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| apierrors.UpdateWebsiteResponseBody    | 400                                    | application/json                       |
| apierrors.UpdateWebsiteDemResponseBody | 404                                    | application/json                       |
| apierrors.APIError                     | 4XX, 5XX                               | \*/\*                                  |

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
    if res.CommonEntityID != nil {
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
| apierrors.DeleteWebsiteResponseBody | 404                                 | application/json                    |
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

| Error Type                                   | Status Code                                  | Content Type                                 |
| -------------------------------------------- | -------------------------------------------- | -------------------------------------------- |
| apierrors.PauseWebsiteMonitoringResponseBody | 404                                          | application/json                             |
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

| Error Type                                     | Status Code                                    | Content Type                                   |
| ---------------------------------------------- | ---------------------------------------------- | ---------------------------------------------- |
| apierrors.UnpauseWebsiteMonitoringResponseBody | 404                                            | application/json                               |
| apierrors.APIError                             | 4XX, 5XX                                       | \*/\*                                          |