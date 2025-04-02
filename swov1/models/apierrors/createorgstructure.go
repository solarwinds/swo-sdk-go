// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package apierrors

import (
	"encoding/json"
	"github.com/solarwinds/swo-sdk-go/swov1/models/components"
)

// CreateOrgStructureCloudAccountsResponse500ResponseBody - Server error
type CreateOrgStructureCloudAccountsResponse500ResponseBody struct {
	// HTTP status code as defined in RFC 2817
	Code int64 `json:"code"`
	// Supporting description of the error
	Message  string                  `json:"message"`
	Target   *string                 `json:"target,omitempty"`
	HTTPMeta components.HTTPMetadata `json:"-"`
}

var _ error = &CreateOrgStructureCloudAccountsResponse500ResponseBody{}

func (e *CreateOrgStructureCloudAccountsResponse500ResponseBody) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}

// CreateOrgStructureCloudAccountsResponseResponseBody - The server cannot find the requested resource.
type CreateOrgStructureCloudAccountsResponseResponseBody struct {
	// HTTP status code as defined in RFC 2817
	Code int64 `json:"code"`
	// Supporting description of the error
	Message  string                  `json:"message"`
	Target   *string                 `json:"target,omitempty"`
	HTTPMeta components.HTTPMetadata `json:"-"`
}

var _ error = &CreateOrgStructureCloudAccountsResponseResponseBody{}

func (e *CreateOrgStructureCloudAccountsResponseResponseBody) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}

// CreateOrgStructureCloudAccountsResponseBody - Access is unauthorized.
type CreateOrgStructureCloudAccountsResponseBody struct {
	// HTTP status code as defined in RFC 2817
	Code int64 `json:"code"`
	// Supporting description of the error
	Message  string                  `json:"message"`
	Target   *string                 `json:"target,omitempty"`
	HTTPMeta components.HTTPMetadata `json:"-"`
}

var _ error = &CreateOrgStructureCloudAccountsResponseBody{}

func (e *CreateOrgStructureCloudAccountsResponseBody) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}

// CreateOrgStructureResponseBody - The server could not understand the request due to invalid syntax.
type CreateOrgStructureResponseBody struct {
	// HTTP status code as defined in RFC 2817
	Code int64 `json:"code"`
	// Supporting description of the error
	Message  string                  `json:"message"`
	Target   *string                 `json:"target,omitempty"`
	HTTPMeta components.HTTPMetadata `json:"-"`
}

var _ error = &CreateOrgStructureResponseBody{}

func (e *CreateOrgStructureResponseBody) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}
