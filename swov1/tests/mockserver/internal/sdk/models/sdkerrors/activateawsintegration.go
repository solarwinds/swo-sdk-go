// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package sdkerrors

import (
	"encoding/json"
	"mockserver/internal/sdk/models/components"
)

// ActivateAwsIntegrationInternalServerError - Server error
type ActivateAwsIntegrationInternalServerError struct {
	// Uniquely identifies an error condition.
	Code *components.CommonDefaultErrorCode `json:"code,omitempty"`
	// Supporting description of the error
	Message string `json:"message"`
	// Indicates the invalid field
	Target   *string                 `json:"target,omitempty"`
	HTTPMeta components.HTTPMetadata `json:"-"`
}

var _ error = &ActivateAwsIntegrationInternalServerError{}

func (e *ActivateAwsIntegrationInternalServerError) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}

// ActivateAwsIntegrationNotFoundError - The server cannot find the requested resource.
type ActivateAwsIntegrationNotFoundError struct {
	// Uniquely identifies an error condition.
	Code *components.CommonDefaultErrorCode `json:"code,omitempty"`
	// Supporting description of the error
	Message string `json:"message"`
	// Indicates the invalid field
	Target   *string                 `json:"target,omitempty"`
	HTTPMeta components.HTTPMetadata `json:"-"`
}

var _ error = &ActivateAwsIntegrationNotFoundError{}

func (e *ActivateAwsIntegrationNotFoundError) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}

// ActivateAwsIntegrationUnauthorizedError - Access is unauthorized.
type ActivateAwsIntegrationUnauthorizedError struct {
	// Uniquely identifies an error condition.
	Code *components.CommonDefaultErrorCode `json:"code,omitempty"`
	// Supporting description of the error
	Message string `json:"message"`
	// Indicates the invalid field
	Target   *string                 `json:"target,omitempty"`
	HTTPMeta components.HTTPMetadata `json:"-"`
}

var _ error = &ActivateAwsIntegrationUnauthorizedError{}

func (e *ActivateAwsIntegrationUnauthorizedError) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}

// ActivateAwsIntegrationBadRequestError - The server could not understand the request due to invalid syntax.
type ActivateAwsIntegrationBadRequestError struct {
	// Uniquely identifies an error condition.
	Code *components.CommonDefaultErrorCode `json:"code,omitempty"`
	// Supporting description of the error
	Message string `json:"message"`
	// Indicates the invalid field
	Target   *string                 `json:"target,omitempty"`
	HTTPMeta components.HTTPMetadata `json:"-"`
}

var _ error = &ActivateAwsIntegrationBadRequestError{}

func (e *ActivateAwsIntegrationBadRequestError) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}
