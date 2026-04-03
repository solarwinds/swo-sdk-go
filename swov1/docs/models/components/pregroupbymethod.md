# PreGroupByMethod

Aggregation method for secondary grouping, inside individual buckets. Has to be set together with `preGroupBy`.

## Example Usage

```go
import (
	"github.com/solarwinds/swo-sdk-go/swov1/models/components"
)

value := components.PreGroupByMethodAvg
```


## Values

| Name                    | Value                   |
| ----------------------- | ----------------------- |
| `PreGroupByMethodAvg`   | AVG                     |
| `PreGroupByMethodCount` | COUNT                   |
| `PreGroupByMethodMin`   | MIN                     |
| `PreGroupByMethodMax`   | MAX                     |
| `PreGroupByMethodSum`   | SUM                     |
| `PreGroupByMethodLast`  | LAST                    |