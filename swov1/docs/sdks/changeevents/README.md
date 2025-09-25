# Changeevents
(*Changeevents*)

## Overview

### Available Operations

* [CreateChangeEvent](#createchangeevent) - Create an event

## CreateChangeEvent

Create an event

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

    res, err := s.Changeevents.CreateChangeEvent(ctx, components.ChangeEvent{
        ID: swov1.Int64(1731676626),
        Name: "app-deploys",
        Title: "deployed v45",
        Timestamp: swov1.Int64(1731676626),
        Source: swov1.String("foo3.example.com"),
        Tags: map[string]string{
            "app": "foo",
            "environment": "production",
        },
        Links: []components.CommonLink{
            components.CommonLink{
                Rel: "self",
                Href: "https://example.com",
            },
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

| Parameter                                                        | Type                                                             | Required                                                         | Description                                                      |
| ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- |
| `ctx`                                                            | [context.Context](https://pkg.go.dev/context#Context)            | :heavy_check_mark:                                               | The context to use for the request.                              |
| `request`                                                        | [components.ChangeEvent](../../models/components/changeevent.md) | :heavy_check_mark:                                               | The request object to use for the request.                       |
| `opts`                                                           | [][operations.Option](../../models/operations/option.md)         | :heavy_minus_sign:                                               | The options for this request.                                    |

### Response

**[*operations.CreateChangeEventResponse](../../models/operations/createchangeeventresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |