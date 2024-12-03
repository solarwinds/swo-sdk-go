# Entities
(*Entities*)

## Overview

### Available Operations

* [ListEntities](#listentities) - Get a list of entities
* [GetEntityByID](#getentitybyid) - Get an entity by ID

## ListEntities

Get a list of entities

### Example Usage

```go
package main

import(
	"os"
	"github.com/solarwinds/swo-sdk-go/solarwinds"
	"github.com/solarwinds/swo-sdk-go/solarwinds/models/operations"
	"context"
	"log"
)

func main() {
    s := solarwinds.New(
        solarwinds.WithSecurity(os.Getenv("SWO_BEARER_AUTH")),
    )

    ctx := context.Background()
    res, err := s.Entities.ListEntities(ctx, operations.ListEntitiesRequest{
        Type: "<value>",
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

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [operations.ListEntitiesRequest](../../models/operations/listentitiesrequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |
| `opts`                                                                           | [][operations.Option](../../models/operations/option.md)                         | :heavy_minus_sign:                                                               | The options for this request.                                                    |

### Response

**[*operations.ListEntitiesResponse](../../models/operations/listentitiesresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## GetEntityByID

Get an entity by ID

### Example Usage

```go
package main

import(
	"os"
	"github.com/solarwinds/swo-sdk-go/solarwinds"
	"github.com/solarwinds/swo-sdk-go/solarwinds/models/operations"
	"context"
	"log"
)

func main() {
    s := solarwinds.New(
        solarwinds.WithSecurity(os.Getenv("SWO_BEARER_AUTH")),
    )

    ctx := context.Background()
    res, err := s.Entities.GetEntityByID(ctx, operations.GetEntityByIDRequest{
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Entity != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.GetEntityByIDRequest](../../models/operations/getentitybyidrequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |
| `opts`                                                                             | [][operations.Option](../../models/operations/option.md)                           | :heavy_minus_sign:                                                                 | The options for this request.                                                      |

### Response

**[*operations.GetEntityByIDResponse](../../models/operations/getentitybyidresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |