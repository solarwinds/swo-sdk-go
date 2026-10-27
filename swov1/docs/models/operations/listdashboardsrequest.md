# ListDashboardsRequest


## Fields

| Field                                                        | Type                                                         | Required                                                     | Description                                                  |
| ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ |
| `Name`                                                       | `*string`                                                    | :heavy_minus_sign:                                           | Filter by dashboard name.                                    |
| `IsOwnedByCurrentUser`                                       | `*bool`                                                      | :heavy_minus_sign:                                           | When true, return only dashboards owned by the current user. |
| `PageSize`                                                   | `*int`                                                       | :heavy_minus_sign:                                           | Number of items in a response page. Default varies by API.   |
| `SkipToken`                                                  | `*string`                                                    | :heavy_minus_sign:                                           | Token for the requested page.                                |