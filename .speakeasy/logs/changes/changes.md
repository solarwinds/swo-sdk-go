## Go SDK Changes:
* `Swo.Entities.UpdateEntityById()`: `response` **Changed** (Breaking ⚠️)
    - `Status[200]` **Added** (Breaking ⚠️)
    - `Status[202]` **Removed** (Breaking ⚠️)
* `Swo.Metrics.ListMetricMeasurements()`:  `response.Groupings[].Measurements[].Value` **Changed** (Breaking ⚠️)
* `Swo.Metrics.ListMultiMetricMeasurements()`:  `response.Metrics[].Groupings[].Measurements[].Value` **Changed** (Breaking ⚠️)
* `Swo.Dbo.GetPublicKey()`: `error` **Changed** (Breaking ⚠️)
    - `Status[400]` **Removed** (Breaking ⚠️)
    - `Status[404]` **Removed** (Breaking ⚠️)
* `Swo.Dbo.GetPluginConfig()`:  `error.status[403]` **Added**
* `Swo.Dbo.UpdateDatabase()`:  `error.status[403]` **Added**
* `Swo.Dbo.ObserveDatabase()`:  `error.status[403]` **Added**
* `Swo.Dbo.GetPlugins()`:  `error.status[403]` **Added**
* `Swo.Dbo.PluginOperation()`:  `error.status[403]` **Added**
* `Swo.Dbo.UnobserveDatabase()`:  `error.status[403]` **Added**
* `Swo.Dbo.DeleteDatabase()`:  `error.status[403]` **Added**
* `Swo.Dbo.SetConfig()`:  `error.status[403]` **Added**
* `Swo.Dbo.GetConfig()`:  `error.status[403]` **Added**
