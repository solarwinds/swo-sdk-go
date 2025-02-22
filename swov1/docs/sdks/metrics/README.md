# Metrics
(*Metrics*)

## Overview

### Available Operations

* [ListMetrics](#listmetrics) - List metrics
* [CreateCompositeMetric](#createcompositemetric) - Create composite metric
* [UpdateCompositeMetric](#updatecompositemetric) - Update composite metric
* [DeleteCompositeMetric](#deletecompositemetric) - Delete composite metric
* [GetMetricByName](#getmetricbyname) - Get metric info by name
* [ListMetricAttributes](#listmetricattributes) - List metric attribute names
* [ListMetricAttributeValues](#listmetricattributevalues) - List metric attribute values
* [ListMetricMeasurements](#listmetricmeasurements) - List metric measurement values, grouped by attributes, filtered by the filter. An empty list indicates no data points are available for the given parameters.

## ListMetrics

List metrics available within a time period. 

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
	"github.com/solarwinds/swo-sdk-go/swov1"
	"github.com/solarwinds/swo-sdk-go/swov1/models/components"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := swov1.New(
        swov1.WithSecurity(os.Getenv("SWO_API_TOKEN")),
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

| Error Type                                                 | Status Code                                                | Content Type                                               |
| ---------------------------------------------------------- | ---------------------------------------------------------- | ---------------------------------------------------------- |
| apierrors.CreateCompositeMetricResponseBody                | 400                                                        | application/json                                           |
| apierrors.CreateCompositeMetricMetricsResponseBody         | 403                                                        | application/json                                           |
| apierrors.CreateCompositeMetricMetricsResponseResponseBody | 500                                                        | application/json                                           |
| apierrors.APIError                                         | 4XX, 5XX                                                   | \*/\*                                                      |

## UpdateCompositeMetric

Update a composite metric given a metric name

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

    res, err := s.Metrics.UpdateCompositeMetric(ctx, operations.UpdateCompositeMetricRequest{
        Name: "<value>",
        UpdateCompositeMetric: components.UpdateCompositeMetric{
            DisplayName: "Disk IO rate",
            Description: "Disk bytes transferred per second",
            Formula: "rate(system.disk.io[5m])",
            Units: "bytes/s",
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
| `request`                                                                                          | [operations.UpdateCompositeMetricRequest](../../models/operations/updatecompositemetricrequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `opts`                                                                                             | [][operations.Option](../../models/operations/option.md)                                           | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |

### Response

**[*operations.UpdateCompositeMetricResponse](../../models/operations/updatecompositemetricresponse.md), error**

### Errors

| Error Type                                                    | Status Code                                                   | Content Type                                                  |
| ------------------------------------------------------------- | ------------------------------------------------------------- | ------------------------------------------------------------- |
| apierrors.UpdateCompositeMetricResponseBody                   | 401                                                           | application/json                                              |
| apierrors.UpdateCompositeMetricMetricsResponseBody            | 403                                                           | application/json                                              |
| apierrors.UpdateCompositeMetricMetricsResponseResponseBody    | 404                                                           | application/json                                              |
| apierrors.UpdateCompositeMetricMetricsResponse500ResponseBody | 500                                                           | application/json                                              |
| apierrors.APIError                                            | 4XX, 5XX                                                      | \*/\*                                                         |

## DeleteCompositeMetric

Delete a composite metric given a metric name

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

    res, err := s.Metrics.DeleteCompositeMetric(ctx, operations.DeleteCompositeMetricRequest{
        Name: "<value>",
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
| `request`                                                                                          | [operations.DeleteCompositeMetricRequest](../../models/operations/deletecompositemetricrequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `opts`                                                                                             | [][operations.Option](../../models/operations/option.md)                                           | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |

### Response

**[*operations.DeleteCompositeMetricResponse](../../models/operations/deletecompositemetricresponse.md), error**

### Errors

| Error Type                                         | Status Code                                        | Content Type                                       |
| -------------------------------------------------- | -------------------------------------------------- | -------------------------------------------------- |
| apierrors.DeleteCompositeMetricResponseBody        | 400                                                | application/json                                   |
| apierrors.DeleteCompositeMetricMetricsResponseBody | 403                                                | application/json                                   |
| apierrors.APIError                                 | 4XX, 5XX                                           | \*/\*                                              |

## GetMetricByName

Get info about a metric

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

| Error Type                            | Status Code                           | Content Type                          |
| ------------------------------------- | ------------------------------------- | ------------------------------------- |
| apierrors.GetMetricByNameResponseBody | 404                                   | application/json                      |
| apierrors.APIError                    | 4XX, 5XX                              | \*/\*                                 |

## ListMetricAttributes

List attribute names for a given metric.

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

| Error Type                                 | Status Code                                | Content Type                               |
| ------------------------------------------ | ------------------------------------------ | ------------------------------------------ |
| apierrors.ListMetricAttributesResponseBody | 404                                        | application/json                           |
| apierrors.APIError                         | 4XX, 5XX                                   | \*/\*                                      |

## ListMetricAttributeValues

List the values of a given metric's attribute.

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

| Error Type                                      | Status Code                                     | Content Type                                    |
| ----------------------------------------------- | ----------------------------------------------- | ----------------------------------------------- |
| apierrors.ListMetricAttributeValuesResponseBody | 404                                             | application/json                                |
| apierrors.APIError                              | 4XX, 5XX                                        | \*/\*                                           |

## ListMetricMeasurements

List metric measurement values, grouped by attributes, filtered by the filter. An empty list indicates no data points are available for the given parameters.

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

    res, err := s.Metrics.ListMetricMeasurements(ctx, operations.ListMetricMeasurementsRequest{
        Name: "<value>",
        SeriesType: components.MetricSeriesTypeScalar,
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

| Error Type                                   | Status Code                                  | Content Type                                 |
| -------------------------------------------- | -------------------------------------------- | -------------------------------------------- |
| apierrors.ListMetricMeasurementsResponseBody | 404                                          | application/json                             |
| apierrors.APIError                           | 4XX, 5XX                                     | \*/\*                                        |