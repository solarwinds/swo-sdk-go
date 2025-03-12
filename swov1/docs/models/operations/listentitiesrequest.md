# ListEntitiesRequest


## Fields

| Field                                                      | Type                                                       | Required                                                   | Description                                                |
| ---------------------------------------------------------- | ---------------------------------------------------------- | ---------------------------------------------------------- | ---------------------------------------------------------- |
| `Type`                                                     | *string*                                                   | :heavy_check_mark:                                         | The entity type to search for                              |
| `Name`                                                     | **string*                                                  | :heavy_minus_sign:                                         | The entity name to search for                              |
| `PageSize`                                                 | **int*                                                     | :heavy_minus_sign:                                         | Number of items in a response page. Default varies by API. |
| `SkipToken`                                                | **string*                                                  | :heavy_minus_sign:                                         | Token for the requested page.                              |