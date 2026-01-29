# GetWebsiteOutageStatusesResponseBody

An array of outage statues with pagination info


## Fields

| Field                                                                      | Type                                                                       | Required                                                                   | Description                                                                |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `Statuses`                                                                 | [][components.DemOutageStatus](../../models/components/demoutagestatus.md) | :heavy_check_mark:                                                         | Website outage statuses                                                    |
| `PageInfo`                                                                 | [components.CommonPageInfo](../../models/components/commonpageinfo.md)     | :heavy_check_mark:                                                         | Pagination information                                                     |