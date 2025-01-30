# CommonMetricInfo


## Fields

| Field                                                      | Type                                                       | Required                                                   | Description                                                | Example                                                    |
| ---------------------------------------------------------- | ---------------------------------------------------------- | ---------------------------------------------------------- | ---------------------------------------------------------- | ---------------------------------------------------------- |
| `Name`                                                     | *string*                                                   | :heavy_check_mark:                                         | Name of the metric                                         | composite.k8s.pod.container.status.restarts.increase       |
| `Units`                                                    | *string*                                                   | :heavy_check_mark:                                         | Units of the metric                                        | count                                                      |
| `Formula`                                                  | **string*                                                  | :heavy_minus_sign:                                         | Formula for the metric                                     | increase(k8s.kube_pod_container_status_restarts_total[5m]) |
| `LastReportedTime`                                         | *string*                                                   | :heavy_check_mark:                                         | Last reported time of the metric                           | 2021-01-01T00:00:00Z                                       |