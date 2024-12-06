# Logs
(*Logs*)

## Overview

### Available Operations

* [SearchLogs](#searchlogs) - Search logs
* [ListLogArchives](#listlogarchives) - Retrieve location and metadata of log archives

## SearchLogs

Search logs within a time period

### Example Usage

```go
package main

import(
	"os"
	swosdkgo "github.com/solarwinds/swo-sdk-go"
	"github.com/solarwinds/swo-sdk-go/models/operations"
	"context"
	"log"
)

func main() {
    s := swosdkgo.New(
        swosdkgo.WithSecurity(os.Getenv("SOLARWINDS_BEARER_AUTH")),
    )

    ctx := context.Background()
    res, err := s.Logs.SearchLogs(ctx, operations.SearchLogsRequest{})
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
                for {
            // handle items
        
            res, err = res.Next()
        
            if err != nil {
                // handle error
            }
        
            if res == nil {
                break
            }
        }
        
    }
}
```

### Parameters

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |
| `request`                                                                    | [operations.SearchLogsRequest](../../models/operations/searchlogsrequest.md) | :heavy_check_mark:                                                           | The request object to use for the request.                                   |
| `opts`                                                                       | [][operations.Option](../../models/operations/option.md)                     | :heavy_minus_sign:                                                           | The options for this request.                                                |

### Response

**[*operations.SearchLogsResponse](../../models/operations/searchlogsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## ListLogArchives

Retrieves a list of log archives within a time period.

### Example Usage

```go
package main

import(
	"os"
	swosdkgo "github.com/solarwinds/swo-sdk-go"
	"github.com/solarwinds/swo-sdk-go/models/operations"
	"context"
	"log"
)

func main() {
    s := swosdkgo.New(
        swosdkgo.WithSecurity(os.Getenv("SOLARWINDS_BEARER_AUTH")),
    )

    ctx := context.Background()
    res, err := s.Logs.ListLogArchives(ctx, operations.ListLogArchivesRequest{
        StartTime: "<value>",
        EndTime: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
                for {
            // handle items
        
            res, err = res.Next()
        
            if err != nil {
                // handle error
            }
        
            if res == nil {
                break
            }
        }
        
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.ListLogArchivesRequest](../../models/operations/listlogarchivesrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `opts`                                                                                 | [][operations.Option](../../models/operations/option.md)                               | :heavy_minus_sign:                                                                     | The options for this request.                                                          |

### Response

**[*operations.ListLogArchivesResponse](../../models/operations/listlogarchivesresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |