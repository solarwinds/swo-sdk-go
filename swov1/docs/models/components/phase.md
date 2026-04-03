# Phase

Indicates at which phase of the failure occurred

## Example Usage

```go
import (
	"github.com/solarwinds/swo-sdk-go/swov1/models/components"
)

value := components.PhaseResolve
```


## Values

| Name              | Value             |
| ----------------- | ----------------- |
| `PhaseResolve`    | resolve           |
| `PhaseConnection` | connection        |
| `PhaseRequest`    | request           |
| `PhaseResponse`   | response          |
| `PhaseUnknown`    | unknown           |