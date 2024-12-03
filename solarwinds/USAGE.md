<!-- Start SDK Example Usage [usage] -->
```go
package main

import (
	"context"
	"github.com/solarwinds/swo-sdk-go/solarwinds"
	"github.com/solarwinds/swo-sdk-go/solarwinds/models/components"
	"log"
	"os"
)

func main() {
	s := solarwinds.New(
		solarwinds.WithSecurity(os.Getenv("SWO_BEARER_AUTH")),
	)

	ctx := context.Background()
	res, err := s.Changeevents.CreateChangeEvent(ctx, components.ChangeEvent{
		ID:        solarwinds.Int64(1731676626),
		Name:      "app-deploys",
		Title:     "deployed v45",
		Timestamp: solarwinds.Int64(1731676626),
		Source:    solarwinds.String("foo3.example.com"),
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