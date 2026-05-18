## Go SDK Changes:
* `Swo.Entities.ListEntities()`: 
  *  `request.Request.Type` **Changed**
  *  `response.Entities[].HealthState.State` **Changed** (Breaking ⚠️)
* `Swo.Entities.GetEntityById()`:  `response.HealthState.State` **Changed** (Breaking ⚠️)
* `Swo.Dem.CreateUri()`: 
  *  `request.Request.AvailabilityCheckSettings.TestFrom` **Changed**
* `Swo.Dem.GetUri()`:  `response.AvailabilityCheckSettings.TestFrom` **Changed**
* `Swo.Dem.UpdateUri()`: 
  *  `request.Request.Dem.uri.AvailabilityCheckSettings.TestFrom` **Changed**
* `Swo.Dem.CreateWebsite()`: 
  * `request.Request.AvailabilityCheckSettings` **Changed**
    - `Authentication` **Added**
    - `TestFrom` **Changed**
* `Swo.Dem.GetWebsite()`: `response.AvailabilityCheckSettings` **Changed**
    - `Authentication` **Added**
    - `TestFrom` **Changed**
* `Swo.Dem.UpdateWebsite()`: 
  * `request.Request.Dem.website.AvailabilityCheckSettings` **Changed**
    - `Authentication` **Added**
    - `TestFrom` **Changed**
* `Swo.Metrics.ListMetrics()`:  `error.status[400]` **Added**
* `Swo.Metrics.ListMetricAttributes()`:  `error.status[400]` **Added**
* `Swo.Metrics.ListMetricAttributeValues()`:  `error.status[400]` **Added**
* `Swo.Metrics.ListMetricMeasurements()`:  `error.status[400]` **Added**
