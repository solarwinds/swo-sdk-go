// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/solarwinds/swo-sdk-go/swov1/models/components"
)

type CreateOrgStructureResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
}

func (o *CreateOrgStructureResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}
