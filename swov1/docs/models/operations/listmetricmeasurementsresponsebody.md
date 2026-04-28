# ListMetricMeasurementsResponseBody

The request has succeeded.


## Fields

| Field                                                                      | Type                                                                       | Required                                                                   | Description                                                                |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `Groupings`                                                                | [][components.MetricsGrouping](../../models/components/metricsgrouping.md) | :heavy_check_mark:                                                         | Measurement data grouped by attributes.                                    |
| `BucketSizeInSeconds`                                                      | `int`                                                                      | :heavy_check_mark:                                                         | Bucket size used for computing time series points, in seconds.             |
| `PageInfo`                                                                 | [components.CommonPageInfo](../../models/components/commonpageinfo.md)     | :heavy_check_mark:                                                         | Pagination information.                                                    |