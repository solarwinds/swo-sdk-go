// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package sdkerrors

import (
	"encoding/json"
	"mockserver/internal/sdk/models/components"
)

// UpdateAwsIntegrationCloudAccountsResponseResponseBody - Server error
type UpdateAwsIntegrationCloudAccountsResponseResponseBody struct {
	// Supporting description of the error
	Message  string                  `json:"message"`
	HTTPMeta components.HTTPMetadata `json:"-"`
}

var _ error = &UpdateAwsIntegrationCloudAccountsResponseResponseBody{}

func (e *UpdateAwsIntegrationCloudAccountsResponseResponseBody) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}

// UpdateAwsIntegrationCloudAccountsResponseBody - Access is unauthorized.
type UpdateAwsIntegrationCloudAccountsResponseBody struct {
	// Supporting description of the error
	Message  string                  `json:"message"`
	HTTPMeta components.HTTPMetadata `json:"-"`
}

var _ error = &UpdateAwsIntegrationCloudAccountsResponseBody{}

func (e *UpdateAwsIntegrationCloudAccountsResponseBody) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}

// UpdateAwsIntegrationResponseBody - The server could not understand the request due to invalid syntax.
type UpdateAwsIntegrationResponseBody struct {
	// Supporting description of the error
	Message string `json:"message"`
	// Indicates the invalid field
	Target   *string                 `json:"target,omitempty"`
	HTTPMeta components.HTTPMetadata `json:"-"`
}

var _ error = &UpdateAwsIntegrationResponseBody{}

func (e *UpdateAwsIntegrationResponseBody) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}
