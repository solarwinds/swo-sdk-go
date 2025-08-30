# CommonBadRequestErrorResponse


## Fields

| Field                                             | Type                                              | Required                                          | Description                                       | Example                                           |
| ------------------------------------------------- | ------------------------------------------------- | ------------------------------------------------- | ------------------------------------------------- | ------------------------------------------------- |
| `Message`                                         | *string*                                          | :heavy_check_mark:                                | Supporting description of the error.              | Error has occurred                                |
| `Target`                                          | **string*                                         | :heavy_minus_sign:                                | Indicates the invalid field.                      |                                                   |
| `Code`                                            | [*apierrors.Code](../../models/apierrors/code.md) | :heavy_minus_sign:                                | Uniquely identifies an error condition.           | InvalidRequest                                    |