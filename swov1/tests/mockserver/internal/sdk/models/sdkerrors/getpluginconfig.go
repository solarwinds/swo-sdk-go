// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package sdkerrors

import (
	"encoding/json"
	"mockserver/internal/sdk/models/components"
)

// GetPluginConfigDboResponseBody - The server cannot find the requested resource.
type GetPluginConfigDboResponseBody struct {
	// Supporting description of the error
	Message  string                  `json:"message"`
	HTTPMeta components.HTTPMetadata `json:"-"`
}

var _ error = &GetPluginConfigDboResponseBody{}

func (e *GetPluginConfigDboResponseBody) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}

// GetPluginConfigResponseBody - The server could not understand the request due to invalid syntax.
type GetPluginConfigResponseBody struct {
	// Supporting description of the error
	Message string `json:"message"`
	// Indicates the invalid field
	Target   *string                 `json:"target,omitempty"`
	HTTPMeta components.HTTPMetadata `json:"-"`
}

var _ error = &GetPluginConfigResponseBody{}

func (e *GetPluginConfigResponseBody) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}
