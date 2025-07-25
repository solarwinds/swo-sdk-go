// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/solarwinds/swo-sdk-go/swov1/models/components"
)

type UpdateAwsIntegrationResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// The request has succeeded.
	CloudAccountsAwsUpdateIntegrationResponse *components.CloudAccountsAwsUpdateIntegrationResponse
}

func (o *UpdateAwsIntegrationResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *UpdateAwsIntegrationResponse) GetCloudAccountsAwsUpdateIntegrationResponse() *components.CloudAccountsAwsUpdateIntegrationResponse {
	if o == nil {
		return nil
	}
	return o.CloudAccountsAwsUpdateIntegrationResponse
}
