<!-- Start SDK Example Usage [usage] -->
```go
package main

import (
	"context"
	swosdkgo "github.com/solarwinds/swo-sdk-go"
	"github.com/solarwinds/swo-sdk-go/models/components"
	"log"
	"os"
)

func main() {
	s := swosdkgo.New(
		swosdkgo.WithSecurity(os.Getenv("SOLARWINDS_BEARER_AUTH")),
	)

	ctx := context.Background()
	res, err := s.Changeevents.CreateChangeEvent(ctx, components.ChangeEvent{
		ID:        swosdkgo.Int64(1731676626),
		Name:      "app-deploys",
		Title:     "deployed v45",
		Timestamp: swosdkgo.Int64(1731676626),
		Source:    swosdkgo.String("foo3.example.com"),
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