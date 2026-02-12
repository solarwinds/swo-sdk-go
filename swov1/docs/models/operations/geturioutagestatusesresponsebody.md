# GetURIOutageStatusesResponseBody

An array of outage statuses with pagination info


## Fields

| Field                                                                      | Type                                                                       | Required                                                                   | Description                                                                |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `Statuses`                                                                 | [][components.DemOutageStatus](../../models/components/demoutagestatus.md) | :heavy_check_mark:                                                         | Uri outage statuses                                                        |
| `PageInfo`                                                                 | [components.CommonPageInfo](../../models/components/commonpageinfo.md)     | :heavy_check_mark:                                                         | Pagination information                                                     |