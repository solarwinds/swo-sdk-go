// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/solarwinds/swo-sdk-go/swov1/models/components"
)

type GetPublicKeyResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// The request has succeeded.
	DboDatabaseCredentialsPublicKeyResponse *components.DboDatabaseCredentialsPublicKeyResponse
}

func (o *GetPublicKeyResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetPublicKeyResponse) GetDboDatabaseCredentialsPublicKeyResponse() *components.DboDatabaseCredentialsPublicKeyResponse {
	if o == nil {
		return nil
	}
	return o.DboDatabaseCredentialsPublicKeyResponse
}
