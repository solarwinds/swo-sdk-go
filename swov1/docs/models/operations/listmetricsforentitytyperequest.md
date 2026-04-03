# ListMetricsForEntityTypeRequest


## Fields

| Field                                                              | Type                                                               | Required                                                           | Description                                                        |
| ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ |
| `Type`                                                             | `string`                                                           | :heavy_check_mark:                                                 | Entity type to list metrics for.                                   |
| `StartTime`                                                        | [*time.Time](https://pkg.go.dev/time#Time)                         | :heavy_minus_sign:                                                 | Timestamp in ISO 8601 format in UTC timezone: yyyy-MM-ddTHH:mm:ssZ |
| `EndTime`                                                          | [*time.Time](https://pkg.go.dev/time#Time)                         | :heavy_minus_sign:                                                 | Timestamp in ISO 8601 format in UTC timezone: yyyy-MM-ddTHH:mm:ssZ |