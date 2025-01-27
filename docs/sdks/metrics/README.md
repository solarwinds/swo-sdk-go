# Metrics
(*Metrics*)

## Overview

### Available Operations

* [ListMetrics](#listmetrics) - List metrics
* [CreateCompositeMetric](#createcompositemetric) - Create composite metric
* [GetMetricByName](#getmetricbyname) - Get metric info by name
* [ListMetricAttributes](#listmetricattributes) - List metric attribute names
* [ListMetricAttributeValues](#listmetricattributevalues) - List metric attribute values
* [ListMetricMeasurements](#listmetricmeasurements) - List metric measurement values, grouped by attributes, filtered by the filter

## ListMetrics

List metrics seen within a time period

### Example Usage

```go
package main

import(
	"context"
	"os"
	swosdkgo "github.com/solarwinds/swo-sdk-go"
	"github.com/solarwinds/swo-sdk-go/models/operations"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := swosdkgo.New(
        swosdkgo.WithSecurity(os.Getenv("SWO_API_TOKEN")),
    )

    res, err := s.Metrics.ListMetrics(ctx, operations.ListMetricsRequest{})
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

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [operations.ListMetricsRequest](../../models/operations/listmetricsrequest.md) | :heavy_check_mark:                                                             | The request object to use for the request.                                     |
| `opts`                                                                         | [][operations.Option](../../models/operations/option.md)                       | :heavy_minus_sign:                                                             | The options for this request.                                                  |

### Response

**[*operations.ListMetricsResponse](../../models/operations/listmetricsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## CreateCompositeMetric

Create a composite metric given a PromQL query

### Example Usage

```go
package main

import(
	"context"
	"os"
	swosdkgo "github.com/solarwinds/swo-sdk-go"
	"github.com/solarwinds/swo-sdk-go/models/components"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := swosdkgo.New(
        swosdkgo.WithSecurity(os.Getenv("SWO_API_TOKEN")),
    )

    res, err := s.Metrics.CreateCompositeMetric(ctx, components.CompositeMetric{
        Name: "composite.custom.system.disk.io.rate",
        DisplayName: "Disk IO rate",
        Description: "Disk bytes transferred per second",
        Formula: "rate(system.disk.io[5m])",
        Units: "bytes/s",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CompositeMetric != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                | Type                                                                     | Required                                                                 | Description                                                              |
| ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ |
| `ctx`                                                                    | [context.Context](https://pkg.go.dev/context#Context)                    | :heavy_check_mark:                                                       | The context to use for the request.                                      |
| `request`                                                                | [components.CompositeMetric](../../models/components/compositemetric.md) | :heavy_check_mark:                                                       | The request object to use for the request.                               |
| `opts`                                                                   | [][operations.Option](../../models/operations/option.md)                 | :heavy_minus_sign:                                                       | The options for this request.                                            |

### Response

**[*operations.CreateCompositeMetricResponse](../../models/operations/createcompositemetricresponse.md), error**

### Errors

| Error Type                                         | Status Code                                        | Content Type                                       |
| -------------------------------------------------- | -------------------------------------------------- | -------------------------------------------------- |
| apierrors.CreateCompositeMetricResponseBody        | 400                                                | application/json                                   |
| apierrors.CreateCompositeMetricMetricsResponseBody | 403                                                | application/json                                   |
| apierrors.APIError                                 | 4XX, 5XX                                           | \*/\*                                              |

## GetMetricByName

Get info about a metric

### Example Usage

```go
package main

import(
	"context"
	"os"
	swosdkgo "github.com/solarwinds/swo-sdk-go"
	"github.com/solarwinds/swo-sdk-go/models/operations"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := swosdkgo.New(
        swosdkgo.WithSecurity(os.Getenv("SWO_API_TOKEN")),
    )

    res, err := s.Metrics.GetMetricByName(ctx, operations.GetMetricByNameRequest{
        Name: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CommonMetricInfo != nil {
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

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.GetMetricByNameRequest](../../models/operations/getmetricbynamerequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `opts`                                                                                 | [][operations.Option](../../models/operations/option.md)                               | :heavy_minus_sign:                                                                     | The options for this request.                                                          |

### Response

**[*operations.GetMetricByNameResponse](../../models/operations/getmetricbynameresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## ListMetricAttributes

List all attribute names defined for the given metric

### Example Usage

```go
package main

import(
	"context"
	"os"
	swosdkgo "github.com/solarwinds/swo-sdk-go"
	"github.com/solarwinds/swo-sdk-go/models/operations"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := swosdkgo.New(
        swosdkgo.WithSecurity(os.Getenv("SWO_API_TOKEN")),
    )

    res, err := s.Metrics.ListMetricAttributes(ctx, operations.ListMetricAttributesRequest{
        Name: "<value>",
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
| `request`                                                                                        | [operations.ListMetricAttributesRequest](../../models/operations/listmetricattributesrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `opts`                                                                                           | [][operations.Option](../../models/operations/option.md)                                         | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |

### Response

**[*operations.ListMetricAttributesResponse](../../models/operations/listmetricattributesresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## ListMetricAttributeValues

List values of a metric's attribute

### Example Usage

```go
package main

import(
	"context"
	"os"
	swosdkgo "github.com/solarwinds/swo-sdk-go"
	"github.com/solarwinds/swo-sdk-go/models/operations"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := swosdkgo.New(
        swosdkgo.WithSecurity(os.Getenv("SWO_API_TOKEN")),
    )

    res, err := s.Metrics.ListMetricAttributeValues(ctx, operations.ListMetricAttributeValuesRequest{
        Name: "<value>",
        AttributeName: "<value>",
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

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |
| `request`                                                                                                  | [operations.ListMetricAttributeValuesRequest](../../models/operations/listmetricattributevaluesrequest.md) | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |
| `opts`                                                                                                     | [][operations.Option](../../models/operations/option.md)                                                   | :heavy_minus_sign:                                                                                         | The options for this request.                                                                              |

### Response

**[*operations.ListMetricAttributeValuesResponse](../../models/operations/listmetricattributevaluesresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## ListMetricMeasurements

List metric measurement values, grouped by attributes, filtered by the filter

### Example Usage

```go
package main

import(
	"context"
	"os"
	swosdkgo "github.com/solarwinds/swo-sdk-go"
	"github.com/solarwinds/swo-sdk-go/models/operations"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := swosdkgo.New(
        swosdkgo.WithSecurity(os.Getenv("SWO_API_TOKEN")),
    )

    res, err := s.Metrics.ListMetricMeasurements(ctx, operations.ListMetricMeasurementsRequest{
        Name: "<value>",
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

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.ListMetricMeasurementsRequest](../../models/operations/listmetricmeasurementsrequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `opts`                                                                                               | [][operations.Option](../../models/operations/option.md)                                             | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |

### Response

**[*operations.ListMetricMeasurementsResponse](../../models/operations/listmetricmeasurementsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |