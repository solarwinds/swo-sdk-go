// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"mockserver/internal/sdk/models/components"
)

type UnpauseWebsiteMonitoringRequest struct {
	EntityID string `pathParam:"style=simple,explode=false,name=entityId"`
}

func (o *UnpauseWebsiteMonitoringRequest) GetEntityID() string {
	if o == nil {
		return ""
	}
	return o.EntityID
}

type UnpauseWebsiteMonitoringResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// The request has succeeded.
	CommonEntityID *components.CommonEntityID
}

func (o *UnpauseWebsiteMonitoringResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *UnpauseWebsiteMonitoringResponse) GetCommonEntityID() *components.CommonEntityID {
	if o == nil {
		return nil
	}
	return o.CommonEntityID
}
