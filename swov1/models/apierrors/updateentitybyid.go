// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package apierrors

import (
	"encoding/json"
	"github.com/solarwinds/swo-sdk-go/swov1/models/components"
)

// UpdateEntityByIDEntitiesResponseResponseBody - The server cannot find the requested resource.
type UpdateEntityByIDEntitiesResponseResponseBody struct {
	// HTTP status code as defined in RFC 2817
	Code int64 `json:"code"`
	// Supporting description of the error
	Message  string                  `json:"message"`
	Target   *string                 `json:"target,omitempty"`
	HTTPMeta components.HTTPMetadata `json:"-"`
}

var _ error = &UpdateEntityByIDEntitiesResponseResponseBody{}

func (e *UpdateEntityByIDEntitiesResponseResponseBody) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}

// UpdateEntityByIDEntitiesResponseBody - Access is unauthorized.
type UpdateEntityByIDEntitiesResponseBody struct {
	// HTTP status code as defined in RFC 2817
	Code int64 `json:"code"`
	// Supporting description of the error
	Message  string                  `json:"message"`
	Target   *string                 `json:"target,omitempty"`
	HTTPMeta components.HTTPMetadata `json:"-"`
}

var _ error = &UpdateEntityByIDEntitiesResponseBody{}

func (e *UpdateEntityByIDEntitiesResponseBody) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}

// UpdateEntityByIDResponseBody - The server could not understand the request due to invalid syntax.
type UpdateEntityByIDResponseBody struct {
	// HTTP status code as defined in RFC 2817
	Code int64 `json:"code"`
	// Supporting description of the error
	Message  string                  `json:"message"`
	Target   *string                 `json:"target,omitempty"`
	HTTPMeta components.HTTPMetadata `json:"-"`
}

var _ error = &UpdateEntityByIDResponseBody{}

func (e *UpdateEntityByIDResponseBody) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}
