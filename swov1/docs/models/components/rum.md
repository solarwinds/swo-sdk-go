# Rum

    Use this field to configure real user monitoring (RUM) for the website.
    You are required to configure at least availability monitoring or real user monitoring to be able to create website.


## Fields

| Field                                                                 | Type                                                                  | Required                                                              | Description                                                           |
| --------------------------------------------------------------------- | --------------------------------------------------------------------- | --------------------------------------------------------------------- | --------------------------------------------------------------------- |
| `ApdexTimeInSeconds`                                                  | `*int`                                                                | :heavy_minus_sign:                                                    | Apdex time threshold in seconds for performance satisfaction scoring. |
| `Spa`                                                                 | `bool`                                                                | :heavy_check_mark:                                                    | Whether the website is a single-page application (SPA).               |