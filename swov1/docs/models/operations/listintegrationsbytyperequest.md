# ListIntegrationsByTypeRequest


## Fields

| Field                                                      | Type                                                       | Required                                                   | Description                                                |
| ---------------------------------------------------------- | ---------------------------------------------------------- | ---------------------------------------------------------- | ---------------------------------------------------------- |
| `Type`                                                     | `string`                                                   | :heavy_check_mark:                                         | Integration type, e.g. nginx or mysql.                     |
| `PageSize`                                                 | `*int`                                                     | :heavy_minus_sign:                                         | Number of items in a response page. Default varies by API. |
| `SkipToken`                                                | `*string`                                                  | :heavy_minus_sign:                                         | Token for the requested page.                              |