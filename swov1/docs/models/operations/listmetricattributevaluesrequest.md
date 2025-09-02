# ListMetricAttributeValuesRequest


## Fields

| Field                                                              | Type                                                               | Required                                                           | Description                                                        |
| ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ |
| `Name`                                                             | *string*                                                           | :heavy_check_mark:                                                 | Metric name.                                                       |
| `AttributeName`                                                    | *string*                                                           | :heavy_check_mark:                                                 | attribute name                                                     |
| `StartTime`                                                        | [*time.Time](https://pkg.go.dev/time#Time)                         | :heavy_minus_sign:                                                 | Timestamp in ISO 8601 format in UTC timezone: yyyy-MM-ddTHH:mm:ssZ |
| `EndTime`                                                          | [*time.Time](https://pkg.go.dev/time#Time)                         | :heavy_minus_sign:                                                 | Timestamp in ISO 8601 format in UTC timezone: yyyy-MM-ddTHH:mm:ssZ |
| `PageSize`                                                         | **int*                                                             | :heavy_minus_sign:                                                 | Number of items in a response page. Default varies by API.         |
| `SkipToken`                                                        | **string*                                                          | :heavy_minus_sign:                                                 | Token for the requested page.                                      |