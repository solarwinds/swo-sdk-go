# CreateURIResponseBody

The server could not understand the request due to invalid syntax.


## Fields

| Field                                                              | Type                                                               | Required                                                           | Description                                                        | Example                                                            |
| ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ |
| `Code`                                                             | *int64*                                                            | :heavy_check_mark:                                                 | HTTP status code as defined in RFC 2817                            | 400                                                                |
| `Message`                                                          | *string*                                                           | :heavy_check_mark:                                                 | Supporting description of the error                                | Bad request                                                        |
| `Target`                                                           | **string*                                                          | :heavy_minus_sign:                                                 | N/A                                                                |                                                                    |
| `HTTPMeta`                                                         | [components.HTTPMetadata](../../models/components/httpmetadata.md) | :heavy_check_mark:                                                 | N/A                                                                |                                                                    |