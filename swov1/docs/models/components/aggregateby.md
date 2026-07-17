# AggregateBy

Aggregation method used to group measurements. Required field for all non-composite metrics. Ignored for composites defined in the schema, whose definitions are assumed to contain an aggregation function already. It is applicable, though, to custom metrics that are either pre-defined or user-defined.

## Example Usage

```go
import (
	"github.com/solarwinds/swo-sdk-go/swov1/models/components"
)

value := components.AggregateByAvg
```


## Values

| Name               | Value              |
| ------------------ | ------------------ |
| `AggregateByAvg`   | AVG                |
| `AggregateByCount` | COUNT              |
| `AggregateByMin`   | MIN                |
| `AggregateByMax`   | MAX                |
| `AggregateBySum`   | SUM                |
| `AggregateByLast`  | LAST               |