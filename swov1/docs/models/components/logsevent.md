# LogsEvent


## Fields

| Field                                                | Type                                                 | Required                                             | Description                                          | Example                                              |
| ---------------------------------------------------- | ---------------------------------------------------- | ---------------------------------------------------- | ---------------------------------------------------- | ---------------------------------------------------- |
| `ID`                                                 | `string`                                             | :heavy_check_mark:                                   | Unique identifier of the log event.                  | 1793698955374546944                                  |
| `Time`                                               | `string`                                             | :heavy_check_mark:                                   | Timestamp of the log event.                          | 2024-01-01T00:00:00Z                                 |
| `Message`                                            | `string`                                             | :heavy_check_mark:                                   | Log message content.                                 | This is a log message                                |
| `Hostname`                                           | `string`                                             | :heavy_check_mark:                                   | Hostname of the server that generated the log event. | webserver.example.com                                |
| `Severity`                                           | `string`                                             | :heavy_check_mark:                                   | Severity level of the log event.                     | INFO                                                 |
| `Program`                                            | `string`                                             | :heavy_check_mark:                                   | Name of the program that generated the log event.    | httpd                                                |