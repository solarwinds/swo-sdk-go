## Go SDK Changes:
* `Swo.CloudAccounts.ValidateMgmtAccountOnboarding()`: `request.Request` **Changed** (Breaking ⚠️)
    - `ManagementAccountId` **Added** (Breaking ⚠️)
    - `Request` **Removed** (Breaking ⚠️)
* `Swo.Dem.GetTransaction()`:  `response.TestDefinition.Commands[].Command.Enum(clickAt)` **Added** (Breaking ⚠️)
* `Swo.Entities.ListEntities()`:  `response.Entities[].Healthscore` **Removed** (Breaking ⚠️)
* `Swo.Entities.GetEntityById()`:  `response.Healthscore` **Removed** (Breaking ⚠️)
* `Swo.Metrics.ListMetricMeasurements()`: 
  *  `request.Request.BucketSizeInSeconds` **Removed** (Breaking ⚠️)
* `Swo.Dem.CreateTransaction()`: 
  *  `request.Request.TestDefinition.Commands[].Command.Enum(clickAt)` **Added**
* `Swo.Dem.UpdateTransaction()`: 
  *  `request.Request.Dem.transaction.TestDefinition.Commands[].Command.Enum(clickAt)` **Added**
