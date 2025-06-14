// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package sdkerrors

import (
	"encoding/json"
	"mockserver/internal/sdk/models/components"
)

// UpdateWebsiteDemResponseBody - The server cannot find the requested resource.
type UpdateWebsiteDemResponseBody struct {
	// Supporting description of the error
	Message  string                  `json:"message"`
	HTTPMeta components.HTTPMetadata `json:"-"`
}

var _ error = &UpdateWebsiteDemResponseBody{}

func (e *UpdateWebsiteDemResponseBody) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}

// UpdateWebsiteResponseBody - The server could not understand the request due to invalid syntax.
type UpdateWebsiteResponseBody struct {
	// Supporting description of the error
	Message string `json:"message"`
	// Indicates the invalid field
	Target   *string                 `json:"target,omitempty"`
	HTTPMeta components.HTTPMetadata `json:"-"`
}

var _ error = &UpdateWebsiteResponseBody{}

func (e *UpdateWebsiteResponseBody) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}
