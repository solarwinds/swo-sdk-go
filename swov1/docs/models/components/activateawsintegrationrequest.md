# ActivateAwsIntegrationRequest


## Fields

| Field                                             | Type                                              | Required                                          | Description                                       |
| ------------------------------------------------- | ------------------------------------------------- | ------------------------------------------------- | ------------------------------------------------- |
| `ManagementAccountID`                             | *string*                                          | :heavy_check_mark:                                | AWS Control Tower Management Account ID.          |
| `AccountID`                                       | *string*                                          | :heavy_check_mark:                                | AWS Accounts ID to be integrated.                 |
| `Enable`                                          | *bool*                                            | :heavy_check_mark:                                | True to enable the integration, false to disable. |