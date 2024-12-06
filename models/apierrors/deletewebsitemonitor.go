// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package apierrors

import (
	"encoding/json"
	"github.com/solarwinds/swo-sdk-go/models/components"
)

type DeleteWebsiteMonitorDetails struct {
	Error string `json:"error"`
}

func (o *DeleteWebsiteMonitorDetails) GetError() string {
	if o == nil {
		return ""
	}
	return o.Error
}

type DeleteWebsiteMonitorInnerError struct {
	Code       string `json:"code"`
	InnerError string `json:"innerError"`
}

func (o *DeleteWebsiteMonitorInnerError) GetCode() string {
	if o == nil {
		return ""
	}
	return o.Code
}

func (o *DeleteWebsiteMonitorInnerError) GetInnerError() string {
	if o == nil {
		return ""
	}
	return o.InnerError
}

// DeleteWebsiteMonitorResponseBody - The server could not understand the request due to invalid syntax.
type DeleteWebsiteMonitorResponseBody struct {
	Code       string                          `json:"code"`
	Message    string                          `json:"message"`
	Target     *string                         `json:"target,omitempty"`
	Details    []DeleteWebsiteMonitorDetails   `json:"details,omitempty"`
	InnerError *DeleteWebsiteMonitorInnerError `json:"innerError,omitempty"`
	HTTPMeta   components.HTTPMetadata         `json:"-"`
}

var _ error = &DeleteWebsiteMonitorResponseBody{}

func (e *DeleteWebsiteMonitorResponseBody) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}
