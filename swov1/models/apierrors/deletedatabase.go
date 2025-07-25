// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package apierrors

import (
	"encoding/json"
	"github.com/solarwinds/swo-sdk-go/swov1/models/components"
)

// DeleteDatabaseDboResponseBody - The server cannot find the requested resource.
type DeleteDatabaseDboResponseBody struct {
	// Uniquely identifies an error condition.
	Code *components.CommonDefaultErrorCode `json:"code,omitempty"`
	// Supporting description of the error
	Message string `json:"message"`
	// Indicates the invalid field
	Target   *string                 `json:"target,omitempty"`
	HTTPMeta components.HTTPMetadata `json:"-"`
}

var _ error = &DeleteDatabaseDboResponseBody{}

func (e *DeleteDatabaseDboResponseBody) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}

// DeleteDatabaseResponseBody - The server could not understand the request due to invalid syntax.
type DeleteDatabaseResponseBody struct {
	// Uniquely identifies an error condition.
	Code *components.CommonDefaultErrorCode `json:"code,omitempty"`
	// Supporting description of the error
	Message string `json:"message"`
	// Indicates the invalid field
	Target   *string                 `json:"target,omitempty"`
	HTTPMeta components.HTTPMetadata `json:"-"`
}

var _ error = &DeleteDatabaseResponseBody{}

func (e *DeleteDatabaseResponseBody) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}
