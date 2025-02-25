# ListMetricAttributeValuesResponseBody

The server cannot find the requested resource.


## Fields

| Field                                                              | Type                                                               | Required                                                           | Description                                                        | Example                                                            |
| ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ |
| `Code`                                                             | *int64*                                                            | :heavy_check_mark:                                                 | HTTP status code as defined in RFC 2817                            | 404                                                                |
| `Message`                                                          | *string*                                                           | :heavy_check_mark:                                                 | Supporting description of the error                                | Entity does not exist                                              |
| `Target`                                                           | **string*                                                          | :heavy_minus_sign:                                                 | N/A                                                                |                                                                    |
| `HTTPMeta`                                                         | [components.HTTPMetadata](../../models/components/httpmetadata.md) | :heavy_check_mark:                                                 | N/A                                                                |                                                                    |