## Go SDK Changes:
* `Swo.Integrations.ListIntegrations()`: **Added**
* `Swo.Integrations.ListIntegrationsByType()`: **Added**
* `Swo.Integrations.GetIntegration()`: **Added**
* `Swo.Integrations.DeleteIntegration()`: **Added**
* `Swo.Dbo.ObserveDatabase()`: `request.Request` **Changed**
    - `DbConnOptions.AdditionalOptions` **Added**
    - `DbType.Enum(oracle)` **Added**
* `Swo.Dbo.UpdateDatabase()`: 
  *  `request.Request.Dbo.updateDatabaseRequest.DbConnOptions.AdditionalOptions` **Added**
* `Swo.Dem.CreateTransaction()`: 
  *  `request.Request.TestDefinition.TestFrom` **Changed**
* `Swo.Dem.GetTransaction()`:  `response.TestDefinition.TestFrom` **Changed**
* `Swo.Dem.UpdateTransaction()`: 
  *  `request.Request.Dem.transaction.TestDefinition.TestFrom` **Changed**
* `Swo.Dem.PauseTransactionMonitoring()`:  `error.status[400]` **Added**
* `Swo.Dem.UnpauseTransactionMonitoring()`:  `error.status[400]` **Added**
* `Swo.Entities.ListEntities()`: 
  *  `request.Request.Search` **Added**
* `Swo.Tokens.CreateToken()`:  `error.status[400]` **Added**
