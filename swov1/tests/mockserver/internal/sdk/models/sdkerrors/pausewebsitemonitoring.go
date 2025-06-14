// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package sdkerrors

import (
	"encoding/json"
	"mockserver/internal/sdk/models/components"
)

// PauseWebsiteMonitoringResponseBody - The server cannot find the requested resource.
type PauseWebsiteMonitoringResponseBody struct {
	// Supporting description of the error
	Message  string                  `json:"message"`
	HTTPMeta components.HTTPMetadata `json:"-"`
}

var _ error = &PauseWebsiteMonitoringResponseBody{}

func (e *PauseWebsiteMonitoringResponseBody) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}
