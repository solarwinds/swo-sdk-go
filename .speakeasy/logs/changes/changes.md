## Go SDK Changes:
* `Swo.Dem.PauseTransactionMonitoring()`:  `error.status[400]` **Removed** (Breaking ⚠️)
* `Swo.Dem.UnpauseTransactionMonitoring()`:  `error.status[400]` **Removed** (Breaking ⚠️)
* `Swo.Dem.CreateTransaction()`: 
  * `request.Request.TestDefinition` **Changed**
    - `AdvancedDataDisabled` **Added**
    - `Commands[].SkipOnFailure` **Added**
    - `TimeoutInSeconds` **Added**
* `Swo.Dem.GetTransaction()`: `response.TestDefinition` **Changed**
    - `AdvancedDataDisabled` **Added**
    - `Commands[].SkipOnFailure` **Added**
    - `TimeoutInSeconds` **Added**
* `Swo.Dem.UpdateTransaction()`: 
  * `request.Request.Dem.transaction.TestDefinition` **Changed**
    - `AdvancedDataDisabled` **Added**
    - `Commands[].SkipOnFailure` **Added**
    - `TimeoutInSeconds` **Added**
