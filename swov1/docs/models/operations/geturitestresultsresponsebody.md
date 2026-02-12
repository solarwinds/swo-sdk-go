# GetURITestResultsResponseBody

An array of test results with pagination info


## Fields

| Field                                                                  | Type                                                                   | Required                                                               | Description                                                            |
| ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- |
| `Results`                                                              | [][components.DemTestResult](../../models/components/demtestresult.md) | :heavy_check_mark:                                                     | Uri test results                                                       |
| `PageInfo`                                                             | [components.CommonPageInfo](../../models/components/commonpageinfo.md) | :heavy_check_mark:                                                     | Pagination information                                                 |