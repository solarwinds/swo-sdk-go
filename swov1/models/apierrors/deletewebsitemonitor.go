// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package apierrors

import (
	"encoding/json"
	"github.com/solarwinds/swo-sdk-go/swov1/models/components"
)

// DeleteWebsiteMonitorResponseBody - The server could not understand the request due to invalid syntax.
type DeleteWebsiteMonitorResponseBody struct {
	// HTTP status code as defined in RFC 2817
	Code int64 `json:"code"`
	// Supporting description of the error
	Message  string                  `json:"message"`
	Target   *string                 `json:"target,omitempty"`
	HTTPMeta components.HTTPMetadata `json:"-"`
}

var _ error = &DeleteWebsiteMonitorResponseBody{}

func (e *DeleteWebsiteMonitorResponseBody) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}
