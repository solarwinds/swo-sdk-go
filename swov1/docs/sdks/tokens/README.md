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

    res, err := s.Tokens.CreateToken(ctx, components.TokensCreateTokenRequest{
        Name: "<value>",
        Tags: components.Tags{
            Server: "<value>",
            TagWithoutValue: "<value>",
        },
        Type: components.TokensCreateTokenRequestTypeIngestion,
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.TokensCreateTokenResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [components.TokensCreateTokenRequest](../../models/components/tokenscreatetokenrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `opts`                                                                                     | [][operations.Option](../../models/operations/option.md)                                   | :heavy_minus_sign:                                                                         | The options for this request.                                                              |

### Response

**[*operations.CreateTokenResponse](../../models/operations/createtokenresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |