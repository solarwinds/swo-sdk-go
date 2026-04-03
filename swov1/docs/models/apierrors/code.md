# Code

Uniquely identifies an error condition.

## Example Usage

```go
import (
	"github.com/solarwinds/swo-sdk-go/swov1/models/apierrors"
)

value := apierrors.CodeAccessForbidden
```


## Values

| Name                     | Value                    |
| ------------------------ | ------------------------ |
| `CodeAccessForbidden`    | AccessForbidden          |
| `CodeInternalError`      | InternalError            |
| `CodeInvalidRequest`     | InvalidRequest           |
| `CodeResourceConflict`   | ResourceConflict         |
| `CodeResourceNotFound`   | ResourceNotFound         |
| `CodeServiceUnavailable` | ServiceUnavailable       |
| `CodeUnauthorized`       | Unauthorized             |