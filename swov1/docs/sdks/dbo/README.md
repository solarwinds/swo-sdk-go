# Dbo
(*Dbo*)

## Overview

### Available Operations

* [ObserveDatabase](#observedatabase) - Add database observability to a database
* [GetPublicKey](#getpublickey) - Get public key for encrypting database credentials locally
* [UpdateDatabase](#updatedatabase) - Update an observed database
* [DeleteDatabase](#deletedatabase) - Delete an observed database
* [GetPluginConfig](#getpluginconfig) - Get configuration of plugins observing a database
* [GetPlugins](#getplugins) - Get status of plugins observing a database
* [PluginOperation](#pluginoperation) - Apply an operation on a database observability plugin

## ObserveDatabase

Add database observability to a database

### Example Usage

<!-- UsageSnippet language="go" operationID="observeDatabase" method="post" path="/v1/dbo/databases" -->
```go
package main

import(
	"context"
	"os"
	"github.com/solarwinds/swo-sdk-go/swov1"
	"github.com/solarwinds/swo-sdk-go/swov1/models/components"
	"log"
)

func main() {
    ctx := context.Background()

    s := swov1.New(
        swov1.WithSecurity(os.Getenv("SWO_API_TOKEN")),
    )

    res, err := s.Dbo.ObserveDatabase(ctx, components.DboObserveDatabaseRequest{
        Name: "<value>",
        AgentID: "<id>",
        DbType: components.DboDatabaseTypeMongo,
        AuthMethod: components.DboDatabaseAuthMethodEntraclientsecret,
        DbConnOptions: components.DboDatabaseConnectionOptions{
            Host: "mixed-scrap.com",
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [components.DboObserveDatabaseRequest](../../models/components/dboobservedatabaserequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../models/operations/option.md)                                     | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.ObserveDatabaseResponse](../../models/operations/observedatabaseresponse.md), error**

### Errors

| Error Type                                | Status Code                               | Content Type                              |
| ----------------------------------------- | ----------------------------------------- | ----------------------------------------- |
| apierrors.CommonBadRequestErrorResponse   | 400                                       | application/json                          |
| apierrors.CommonUnauthorizedErrorResponse | 401                                       | application/json                          |
| apierrors.CommonInternalErrorResponse     | 500                                       | application/json                          |
| apierrors.APIError                        | 4XX, 5XX                                  | \*/\*                                     |

## GetPublicKey

Get public key for encrypting database credentials locally

### Example Usage

<!-- UsageSnippet language="go" operationID="getPublicKey" method="get" path="/v1/dbo/databases/credentials/public-key" -->
```go
package main

import(
	"context"
	"os"
	"github.com/solarwinds/swo-sdk-go/swov1"
	"log"
)

func main() {
    ctx := context.Background()

    s := swov1.New(
        swov1.WithSecurity(os.Getenv("SWO_API_TOKEN")),
    )

    res, err := s.Dbo.GetPublicKey(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.DboDatabaseCredentialsPublicKeyResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetPublicKeyResponse](../../models/operations/getpublickeyresponse.md), error**

### Errors

| Error Type                                | Status Code                               | Content Type                              |
| ----------------------------------------- | ----------------------------------------- | ----------------------------------------- |
| apierrors.CommonBadRequestErrorResponse   | 400                                       | application/json                          |
| apierrors.CommonUnauthorizedErrorResponse | 401                                       | application/json                          |
| apierrors.CommonNotFoundErrorResponse     | 404                                       | application/json                          |
| apierrors.CommonInternalErrorResponse     | 500                                       | application/json                          |
| apierrors.APIError                        | 4XX, 5XX                                  | \*/\*                                     |

## UpdateDatabase

Update an observed database

### Example Usage

<!-- UsageSnippet language="go" operationID="updateDatabase" method="patch" path="/v1/dbo/databases/{entityId}" -->
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

    res, err := s.Dbo.UpdateDatabase(ctx, operations.UpdateDatabaseRequest{
        EntityID: "<id>",
        DboUpdateDatabaseRequest: components.DboUpdateDatabaseRequest{},
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

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.UpdateDatabaseRequest](../../models/operations/updatedatabaserequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |
| `opts`                                                                               | [][operations.Option](../../models/operations/option.md)                             | :heavy_minus_sign:                                                                   | The options for this request.                                                        |

### Response

**[*operations.UpdateDatabaseResponse](../../models/operations/updatedatabaseresponse.md), error**

### Errors

| Error Type                                | Status Code                               | Content Type                              |
| ----------------------------------------- | ----------------------------------------- | ----------------------------------------- |
| apierrors.CommonBadRequestErrorResponse   | 400                                       | application/json                          |
| apierrors.CommonUnauthorizedErrorResponse | 401                                       | application/json                          |
| apierrors.CommonNotFoundErrorResponse     | 404                                       | application/json                          |
| apierrors.CommonInternalErrorResponse     | 500                                       | application/json                          |
| apierrors.APIError                        | 4XX, 5XX                                  | \*/\*                                     |

## DeleteDatabase

Delete an observed database

### Example Usage

<!-- UsageSnippet language="go" operationID="deleteDatabase" method="delete" path="/v1/dbo/databases/{entityId}" -->
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

    res, err := s.Dbo.DeleteDatabase(ctx, operations.DeleteDatabaseRequest{
        EntityID: "<id>",
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

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.DeleteDatabaseRequest](../../models/operations/deletedatabaserequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |
| `opts`                                                                               | [][operations.Option](../../models/operations/option.md)                             | :heavy_minus_sign:                                                                   | The options for this request.                                                        |

### Response

**[*operations.DeleteDatabaseResponse](../../models/operations/deletedatabaseresponse.md), error**

### Errors

| Error Type                                | Status Code                               | Content Type                              |
| ----------------------------------------- | ----------------------------------------- | ----------------------------------------- |
| apierrors.CommonBadRequestErrorResponse   | 400                                       | application/json                          |
| apierrors.CommonUnauthorizedErrorResponse | 401                                       | application/json                          |
| apierrors.CommonNotFoundErrorResponse     | 404                                       | application/json                          |
| apierrors.CommonInternalErrorResponse     | 500                                       | application/json                          |
| apierrors.APIError                        | 4XX, 5XX                                  | \*/\*                                     |

## GetPluginConfig

Get configuration of plugins observing a database

### Example Usage

<!-- UsageSnippet language="go" operationID="getPluginConfig" method="get" path="/v1/dbo/databases/{entityId}/pluginConfig" -->
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

    res, err := s.Dbo.GetPluginConfig(ctx, operations.GetPluginConfigRequest{
        EntityID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DboDatabasePluginConfigResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.GetPluginConfigRequest](../../models/operations/getpluginconfigrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `opts`                                                                                 | [][operations.Option](../../models/operations/option.md)                               | :heavy_minus_sign:                                                                     | The options for this request.                                                          |

### Response

**[*operations.GetPluginConfigResponse](../../models/operations/getpluginconfigresponse.md), error**

### Errors

| Error Type                                | Status Code                               | Content Type                              |
| ----------------------------------------- | ----------------------------------------- | ----------------------------------------- |
| apierrors.CommonBadRequestErrorResponse   | 400                                       | application/json                          |
| apierrors.CommonUnauthorizedErrorResponse | 401                                       | application/json                          |
| apierrors.CommonNotFoundErrorResponse     | 404                                       | application/json                          |
| apierrors.CommonInternalErrorResponse     | 500                                       | application/json                          |
| apierrors.APIError                        | 4XX, 5XX                                  | \*/\*                                     |

## GetPlugins

Get status of plugins observing a database

### Example Usage

<!-- UsageSnippet language="go" operationID="getPlugins" method="get" path="/v1/dbo/databases/{entityId}/plugins" -->
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

    res, err := s.Dbo.GetPlugins(ctx, operations.GetPluginsRequest{
        EntityID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DboDatabasePluginStatusResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |
| `request`                                                                    | [operations.GetPluginsRequest](../../models/operations/getpluginsrequest.md) | :heavy_check_mark:                                                           | The request object to use for the request.                                   |
| `opts`                                                                       | [][operations.Option](../../models/operations/option.md)                     | :heavy_minus_sign:                                                           | The options for this request.                                                |

### Response

**[*operations.GetPluginsResponse](../../models/operations/getpluginsresponse.md), error**

### Errors

| Error Type                                | Status Code                               | Content Type                              |
| ----------------------------------------- | ----------------------------------------- | ----------------------------------------- |
| apierrors.CommonBadRequestErrorResponse   | 400                                       | application/json                          |
| apierrors.CommonUnauthorizedErrorResponse | 401                                       | application/json                          |
| apierrors.CommonNotFoundErrorResponse     | 404                                       | application/json                          |
| apierrors.CommonInternalErrorResponse     | 500                                       | application/json                          |
| apierrors.APIError                        | 4XX, 5XX                                  | \*/\*                                     |

## PluginOperation

Apply an operation on a database observability plugin

### Example Usage

<!-- UsageSnippet language="go" operationID="pluginOperation" method="post" path="/v1/dbo/databases/{entityId}/plugins/operation/{operation}" -->
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

    res, err := s.Dbo.PluginOperation(ctx, operations.PluginOperationRequest{
        EntityID: "<id>",
        Operation: "<value>",
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

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.PluginOperationRequest](../../models/operations/pluginoperationrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `opts`                                                                                 | [][operations.Option](../../models/operations/option.md)                               | :heavy_minus_sign:                                                                     | The options for this request.                                                          |

### Response

**[*operations.PluginOperationResponse](../../models/operations/pluginoperationresponse.md), error**

### Errors

| Error Type                                | Status Code                               | Content Type                              |
| ----------------------------------------- | ----------------------------------------- | ----------------------------------------- |
| apierrors.CommonBadRequestErrorResponse   | 400                                       | application/json                          |
| apierrors.CommonUnauthorizedErrorResponse | 401                                       | application/json                          |
| apierrors.CommonNotFoundErrorResponse     | 404                                       | application/json                          |
| apierrors.CommonInternalErrorResponse     | 500                                       | application/json                          |
| apierrors.APIError                        | 4XX, 5XX                                  | \*/\*                                     |