## Go SDK Changes:
* `Swo.Dem.DeleteTransaction()`: `response` **Changed** (Breaking ⚠️)
    - `Status[200]` **Removed** (Breaking ⚠️)
    - `Status[204]` **Added** (Breaking ⚠️)
* `Swo.Dbo.UnobserveDatabase()`: **Added**
* `Swo.Dem.GetUriOutageStatuses()`: **Added**
* `Swo.Dem.GetUriTestResults()`: **Added**
* `Swo.Dem.GetWebsiteOutageStatuses()`: **Added**
* `Swo.Dem.GetWebsiteTestResults()`: **Added**
* `Swo.Entities.ListEntities()`:  `response.Entities[].HealthState` **Added**
* `Swo.Entities.GetEntityById()`:  `response.HealthState` **Added**
