// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/solarwinds/swo-sdk-go/swov1/models/components"
)

type GetEntityByIDRequest struct {
	// The entity's ID
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *GetEntityByIDRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type GetEntityByIDResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// The request has succeeded.
	EntitiesEntity *components.EntitiesEntity
}

func (o *GetEntityByIDResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetEntityByIDResponse) GetEntitiesEntity() *components.EntitiesEntity {
	if o == nil {
		return nil
	}
	return o.EntitiesEntity
}
