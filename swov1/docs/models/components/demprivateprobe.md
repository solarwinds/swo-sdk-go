# DemPrivateProbe


## Fields

| Field                                                    | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ID`                                                     | `string`                                                 | :heavy_check_mark:                                       | Unique identifier of the private probe.                  | probe-1                                                  |
| `Name`                                                   | `string`                                                 | :heavy_check_mark:                                       | Human-readable name of the private probe.                | Washington                                               |
| `Active`                                                 | `bool`                                                   | :heavy_check_mark:                                       | Indicates whether the private probe is currently active. | true                                                     |
| `Version`                                                | `string`                                                 | :heavy_check_mark:                                       | Version of the private probe configuration.              | 1.0.3                                                    |
| `LastSeen`                                               | [time.Time](https://pkg.go.dev/time#Time)                | :heavy_check_mark:                                       | Timestamp of the last time the private probe was seen.   | 2025-01-15T14:31:19.735Z                                 |