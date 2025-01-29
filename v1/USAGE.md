<!-- Start SDK Example Usage [usage] -->
```go
package main

import (
	"context"
	"github.com/solarwinds/swo-sdk-go/v1"
	"github.com/solarwinds/swo-sdk-go/v1/models/components"
	"log"
	"os"
)

func main() {
	ctx := context.Background()

	s := v1.New(
		v1.WithSecurity(os.Getenv("SWO_API_TOKEN")),
	)

	res, err := s.Changeevents.CreateChangeEvent(ctx, components.ChangeEvent{
		ID:        v1.Int64(1731676626),
		Name:      "app-deploys",
		Title:     "deployed v45",
		Timestamp: v1.Int64(1731676626),
		Source:    v1.String("foo3.example.com"),
		Tags: map[string]string{
			"app":         "foo",
			"environment": "production",
		},
		Links: []components.CommonLink{
			components.CommonLink{
				Rel:  "self",
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
<!-- End SDK Example Usage [usage] -->