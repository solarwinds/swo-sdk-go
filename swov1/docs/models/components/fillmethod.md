# FillMethod

Method for filling missing data points in the range.

## Example Usage

```go
import (
	"github.com/solarwinds/swo-sdk-go/swov1/models/components"
)

value := components.FillMethodNone
```


## Values

| Name                      | Value                     |
| ------------------------- | ------------------------- |
| `FillMethodNone`          | NONE                      |
| `FillMethodZeroFill`      | ZERO_FILL                 |
| `FillMethodNullFill`      | NULL_FILL                 |
| `FillMethodLastValueFill` | LAST_VALUE_FILL           |