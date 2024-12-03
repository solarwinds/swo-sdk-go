# CreateWebsiteMonitorResponseBody

The server could not understand the request due to invalid syntax.


## Fields

| Field                                                              | Type                                                               | Required                                                           | Description                                                        |
| ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ |
| `Code`                                                             | *string*                                                           | :heavy_check_mark:                                                 | N/A                                                                |
| `Message`                                                          | *string*                                                           | :heavy_check_mark:                                                 | N/A                                                                |
| `Target`                                                           | **string*                                                          | :heavy_minus_sign:                                                 | N/A                                                                |
| `Details`                                                          | [][apierrors.Details](../../models/apierrors/details.md)           | :heavy_minus_sign:                                                 | N/A                                                                |
| `InnerError`                                                       | [*apierrors.InnerError](../../models/apierrors/innererror.md)      | :heavy_minus_sign:                                                 | N/A                                                                |
| `HTTPMeta`                                                         | [components.HTTPMetadata](../../models/components/httpmetadata.md) | :heavy_check_mark:                                                 | N/A                                                                |