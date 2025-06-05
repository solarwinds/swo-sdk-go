# UpdateWebsiteResponseBody

The server could not understand the request due to invalid syntax.


## Fields

| Field                                                              | Type                                                               | Required                                                           | Description                                                        | Example                                                            |
| ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ |
| `Message`                                                          | *string*                                                           | :heavy_check_mark:                                                 | Supporting description of the error                                | Bad request                                                        |
| `Target`                                                           | **string*                                                          | :heavy_minus_sign:                                                 | Indicates the invalid field                                        |                                                                    |
| `HTTPMeta`                                                         | [components.HTTPMetadata](../../models/components/httpmetadata.md) | :heavy_check_mark:                                                 | N/A                                                                |                                                                    |