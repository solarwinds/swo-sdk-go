# DashboardsDashboardWidgetInput

Widget input for create and update operations.


## Fields

| Field                            | Type                             | Required                         | Description                      |
| -------------------------------- | -------------------------------- | -------------------------------- | -------------------------------- |
| `ID`                             | `string`                         | :heavy_check_mark:               | Widget identifier.               |
| `Type`                           | `string`                         | :heavy_check_mark:               | Widget type.                     |
| `Title`                          | `*string`                        | :heavy_minus_sign:               | Widget display title.            |
| `Properties`                     | map[string]`any`                 | :heavy_minus_sign:               | Widget configuration properties. |