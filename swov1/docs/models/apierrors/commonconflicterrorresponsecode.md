# CommonConflictErrorResponseCode

Uniquely identifies an error condition.

## Example Usage

```go
import (
	"github.com/solarwinds/swo-sdk-go/swov1/models/apierrors"
)

value := apierrors.CommonConflictErrorResponseCodeAccessForbidden
```


## Values

| Name                                                | Value                                               |
| --------------------------------------------------- | --------------------------------------------------- |
| `CommonConflictErrorResponseCodeAccessForbidden`    | AccessForbidden                                     |
| `CommonConflictErrorResponseCodeInternalError`      | InternalError                                       |
| `CommonConflictErrorResponseCodeInvalidRequest`     | InvalidRequest                                      |
| `CommonConflictErrorResponseCodeResourceConflict`   | ResourceConflict                                    |
| `CommonConflictErrorResponseCodeResourceNotFound`   | ResourceNotFound                                    |
| `CommonConflictErrorResponseCodeServiceUnavailable` | ServiceUnavailable                                  |
| `CommonConflictErrorResponseCodeUnauthorized`       | Unauthorized                                        |