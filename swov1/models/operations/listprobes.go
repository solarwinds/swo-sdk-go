// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/solarwinds/swo-sdk-go/swov1/models/components"
)

// ListProbesResponseBody - The request has succeeded.
type ListProbesResponseBody struct {
	// Synthetic probes used to perform availability tests.
	Probes []components.Probe `json:"probes"`
}

func (o *ListProbesResponseBody) GetProbes() []components.Probe {
	if o == nil {
		return []components.Probe{}
	}
	return o.Probes
}

type ListProbesResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// The request has succeeded.
	Object *ListProbesResponseBody
}

func (o *ListProbesResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *ListProbesResponse) GetObject() *ListProbesResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
