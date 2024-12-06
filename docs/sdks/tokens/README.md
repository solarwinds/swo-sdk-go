# Tokens
(*Tokens*)

## Overview

### Available Operations

* [CreateToken](#createtoken) - Create ingestion token

## CreateToken

Create ingestion token

### Example Usage

```go
package main

import(
	"os"
	swosdkgo "github.com/solarwinds/swo-sdk-go"
	"github.com/solarwinds/swo-sdk-go/models/components"
	"context"
	"log"
)

func main() {
    s := swosdkgo.New(
        swosdkgo.WithSecurity(os.Getenv("SOLARWINDS_BEARER_AUTH")),
    )

    ctx := context.Background()
    res, err := s.Tokens.CreateToken(ctx, components.CreateTokenRequest{
        Name: "<value>",
        Tags: components.Tags{
            Server: "<value>",
            TagWithoutValue: "<value>",
        },
        Type: components.TypeIngestion,
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

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [components.CreateTokenRequest](../../models/components/createtokenrequest.md) | :heavy_check_mark:                                                             | The request object to use for the request.                                     |
| `opts`                                                                         | [][operations.Option](../../models/operations/option.md)                       | :heavy_minus_sign:                                                             | The options for this request.                                                  |

### Response

**[*operations.CreateTokenResponse](../../models/operations/createtokenresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |