// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"mockserver/internal/sdk/models/components"
)

type CreateWebsiteResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// The request has succeeded and a new resource has been created as a result.
	CommonEntityID *components.CommonEntityID
}

func (o *CreateWebsiteResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *CreateWebsiteResponse) GetCommonEntityID() *components.CommonEntityID {
	if o == nil {
		return nil
	}
	return o.CommonEntityID
}
