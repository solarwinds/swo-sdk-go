# Metrics
(*Metrics*)

## Overview

### Available Operations

* [ListMetrics](#listmetrics) - List metrics
* [CreateCompositeMetric](#createcompositemetric) - Create composite metric
* [ListMultiMetricMeasurements](#listmultimetricmeasurements) - List measurements for a batch of metrics
* [UpdateCompositeMetric](#updatecompositemetric) - Update composite metric
* [DeleteCompositeMetric](#deletecompositemetric) - Delete composite metric
* [GetMetricByName](#getmetricbyname) - Get metric info by name
* [ListMetricAttributes](#listmetricattributes) - List metric attribute names
* [ListMetricAttributeValues](#listmetricattributevalues) - List metric attribute values
* [ListMetricMeasurements](#listmetricmeasurements) - List metric measurement values, grouped by attributes, filtered by the filter. An empty list indicates no data points are available for the given parameters.

## ListMetrics

List metrics available within a time period.

### Example Usage

<!-- UsageSnippet language="go" operationID="listMetrics" method="get" path="/v1/metrics" -->
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

Create a composite metric given a PromQL query.

### Example Usage

<!-- UsageSnippet language="go" operationID="createCompositeMetric" method="post" path="/v1/metrics" -->
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

    res, err := s.Metrics.CreateCompositeMetric(ctx, components.MetricsCompositeMetric{
        Name: "composite.custom.system.disk.io.rate",
        DisplayName: swov1.String("Disk IO rate"),
        Description: swov1.String("Disk bytes transferred per second"),
        Formula: "rate(system.disk.io[5m])",
        Units: swov1.String("bytes/s"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.MetricsCompositeMetric != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [components.MetricsCompositeMetric](../../models/components/metricscompositemetric.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `opts`                                                                                 | [][operations.Option](../../models/operations/option.md)                               | :heavy_minus_sign:                                                                     | The options for this request.                                                          |

### Response

**[*operations.CreateCompositeMetricResponse](../../models/operations/createcompositemetricresponse.md), error**

### Errors

| Error Type                                         | Status Code                                        | Content Type                                       |
| -------------------------------------------------- | -------------------------------------------------- | -------------------------------------------------- |
| apierrors.CreateCompositeMetricResponseBody        | 400                                                | application/json                                   |
| apierrors.CreateCompositeMetricMetricsResponseBody | 409                                                | application/json                                   |
| apierrors.APIError                                 | 4XX, 5XX                                           | \*/\*                                              |

## ListMultiMetricMeasurements

  List metric measurements, potentially filtered and aggregated. This operation mimics the
  capabilities found in GET `/v1/metrics/{name}/measurements`, extending it for an entire batch of
  metrics that can be requested at once, over the same time interval.

  It is legal to request the same metric multiple times, using different aggregations. For
  example, both the total for the interval (scalar) and a time series can be read simultaneously,
  using two different entries in the request payload. Entries in the response can be matched with
  the respective requests by metric name. When more than one exist for the same name, an `id`
  can be provided to disambiguate. This property will be echoed back unchanged and can be used
  for some or all entries, regardless of whether metrics repeat.

  By default, this endpoint will omit response entries that include no measurements. This is done
  only when the request contains enough information for the client to successfully match all
  responses back to their respective requests—i.e., when there is at most one entry for each
  combination of metric name and `id`. This provides a cleaner response, especially in case where
  multiple pages need to be traversed and data for most metric entries is exhausted early.
  Otherwise, empty entries will be included as well and the output will be fully positional. This
  positional mode can be forced with `forcePositional`.

  Pages can be navigated by following the links returned in `pageInfo`. Those requests must also
  be POST and must include the same payload that was initially sent. Behavior is undefined
  otherwise.

### Example Usage

<!-- UsageSnippet language="go" operationID="listMultiMetricMeasurements" method="post" path="/v1/metrics/measurements" -->
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

    res, err := s.Metrics.ListMultiMetricMeasurements(ctx, operations.ListMultiMetricMeasurementsRequest{
        RequestBody: operations.ListMultiMetricMeasurementsRequestBody{
            Metrics: []components.MetricsMeasurementsRequest{},
        },
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

| Parameter                                                                                                      | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                          | :heavy_check_mark:                                                                                             | The context to use for the request.                                                                            |
| `request`                                                                                                      | [operations.ListMultiMetricMeasurementsRequest](../../models/operations/listmultimetricmeasurementsrequest.md) | :heavy_check_mark:                                                                                             | The request object to use for the request.                                                                     |
| `opts`                                                                                                         | [][operations.Option](../../models/operations/option.md)                                                       | :heavy_minus_sign:                                                                                             | The options for this request.                                                                                  |

### Response

**[*operations.ListMultiMetricMeasurementsResponse](../../models/operations/listmultimetricmeasurementsresponse.md), error**

### Errors

| Error Type                                        | Status Code                                       | Content Type                                      |
| ------------------------------------------------- | ------------------------------------------------- | ------------------------------------------------- |
| apierrors.ListMultiMetricMeasurementsResponseBody | 400                                               | application/json                                  |
| apierrors.APIError                                | 4XX, 5XX                                          | \*/\*                                             |

## UpdateCompositeMetric

Update a composite metric given a metric name.

### Example Usage

<!-- UsageSnippet language="go" operationID="updateCompositeMetric" method="put" path="/v1/metrics/{name}" -->
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
        MetricsUpdateCompositeMetricRequest: components.MetricsUpdateCompositeMetricRequest{
            DisplayName: swov1.String("Disk IO rate"),
            Description: swov1.String("Disk bytes transferred per second"),
            Formula: "rate(system.disk.io[5m])",
            Units: swov1.String("bytes/s"),
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.MetricsCompositeMetric != nil {
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

| Error Type                                                 | Status Code                                                | Content Type                                               |
| ---------------------------------------------------------- | ---------------------------------------------------------- | ---------------------------------------------------------- |
| apierrors.UpdateCompositeMetricResponseBody                | 400                                                        | application/json                                           |
| apierrors.UpdateCompositeMetricMetricsResponseBody         | 403                                                        | application/json                                           |
| apierrors.UpdateCompositeMetricMetricsResponseResponseBody | 404                                                        | application/json                                           |
| apierrors.APIError                                         | 4XX, 5XX                                                   | \*/\*                                                      |

## DeleteCompositeMetric

Delete a composite metric given a metric name.

### Example Usage

<!-- UsageSnippet language="go" operationID="deleteCompositeMetric" method="delete" path="/v1/metrics/{name}" -->
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
| apierrors.DeleteCompositeMetricResponseBody        | 403                                                | application/json                                   |
| apierrors.DeleteCompositeMetricMetricsResponseBody | 404                                                | application/json                                   |
| apierrors.APIError                                 | 4XX, 5XX                                           | \*/\*                                              |

## GetMetricByName

Get information about a given metric.

### Example Usage

<!-- UsageSnippet language="go" operationID="getMetricByName" method="get" path="/v1/metrics/{name}" -->
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

<!-- UsageSnippet language="go" operationID="listMetricAttributes" method="get" path="/v1/metrics/{name}/attributes" -->
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

<!-- UsageSnippet language="go" operationID="listMetricAttributeValues" method="get" path="/v1/metrics/{name}/attributes/{attributeName}" -->
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

<!-- UsageSnippet language="go" operationID="listMetricMeasurements" method="get" path="/v1/metrics/{name}/measurements" -->
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
        SeriesType: components.MetricsMetricSeriesTypeTimeseries,
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