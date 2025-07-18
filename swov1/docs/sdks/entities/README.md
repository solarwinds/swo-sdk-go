# Entities
(*Entities*)

## Overview

### Available Operations

* [ListEntities](#listentities) - Get a list of entities by type. A returned empty list indicates no entities matched the given parameters.
* [GetEntityByID](#getentitybyid) - Get an entity by ID
* [UpdateEntityByID](#updateentitybyid) - Update an entity by ID

## ListEntities

Get a list of entities by type. A returned empty list indicates no entities matched the given parameters.

### Example Usage

```go
package main

import(
	"context"
	"os"
	"github.com/solarwinds/swo-sdk-go/swov1"
	"github.com/solarwinds/swo-sdk-go/swov1/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := swov1.New(
        swov1.WithSecurity(os.Getenv("SWO_API_TOKEN")),
    )

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

| Error Type                                 | Status Code                                | Content Type                               |
| ------------------------------------------ | ------------------------------------------ | ------------------------------------------ |
| apierrors.ListEntitiesResponseBody         | 400                                        | application/json                           |
| apierrors.ListEntitiesEntitiesResponseBody | 401                                        | application/json                           |
| apierrors.APIError                         | 4XX, 5XX                                   | \*/\*                                      |

## GetEntityByID

Get an entity by ID

### Example Usage

```go
package main

import(
	"context"
	"os"
	"github.com/solarwinds/swo-sdk-go/swov1"
	"github.com/solarwinds/swo-sdk-go/swov1/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := swov1.New(
        swov1.WithSecurity(os.Getenv("SWO_API_TOKEN")),
    )

    res, err := s.Entities.GetEntityByID(ctx, operations.GetEntityByIDRequest{
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.EntitiesEntity != nil {
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

| Error Type                                          | Status Code                                         | Content Type                                        |
| --------------------------------------------------- | --------------------------------------------------- | --------------------------------------------------- |
| apierrors.GetEntityByIDResponseBody                 | 400                                                 | application/json                                    |
| apierrors.GetEntityByIDEntitiesResponseBody         | 401                                                 | application/json                                    |
| apierrors.GetEntityByIDEntitiesResponseResponseBody | 404                                                 | application/json                                    |
| apierrors.APIError                                  | 4XX, 5XX                                            | \*/\*                                               |

## UpdateEntityByID

Update an entity by ID

### Example Usage

```go
package main

import(
	"context"
	"os"
	"github.com/solarwinds/swo-sdk-go/swov1"
	"github.com/solarwinds/swo-sdk-go/swov1/models/components"
	"github.com/solarwinds/swo-sdk-go/swov1/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := swov1.New(
        swov1.WithSecurity(os.Getenv("SWO_API_TOKEN")),
    )

    res, err := s.Entities.UpdateEntityByID(ctx, operations.UpdateEntityByIDRequest{
        ID: "<id>",
        EntitiesEntity: components.EntitiesEntityInput{
            DisplayName: swov1.String("SyslogTest"),
            Tags: map[string]*string{
                "gg.tk.token": swov1.String("test"),
                "kfi.tk.token": swov1.String("qa-test"),
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.UpdateEntityByIDRequest](../../models/operations/updateentitybyidrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `opts`                                                                                   | [][operations.Option](../../models/operations/option.md)                                 | :heavy_minus_sign:                                                                       | The options for this request.                                                            |

### Response

**[*operations.UpdateEntityByIDResponse](../../models/operations/updateentitybyidresponse.md), error**

### Errors

| Error Type                                             | Status Code                                            | Content Type                                           |
| ------------------------------------------------------ | ------------------------------------------------------ | ------------------------------------------------------ |
| apierrors.UpdateEntityByIDResponseBody                 | 400                                                    | application/json                                       |
| apierrors.UpdateEntityByIDEntitiesResponseBody         | 401                                                    | application/json                                       |
| apierrors.UpdateEntityByIDEntitiesResponseResponseBody | 404                                                    | application/json                                       |
| apierrors.APIError                                     | 4XX, 5XX                                               | \*/\*                                                  |