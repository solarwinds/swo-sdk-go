// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package apierrors

import (
	"encoding/json"
	"github.com/solarwinds/swo-sdk-go/swov1/models/components"
)

// CreateCompositeMetricMetricsResponseResponseBody - Server error
type CreateCompositeMetricMetricsResponseResponseBody struct {
	// HTTP status code as defined in RFC 2817
	Code int64 `json:"code"`
	// Supporting description of the error
	Message  string                  `json:"message"`
	Target   *string                 `json:"target,omitempty"`
	HTTPMeta components.HTTPMetadata `json:"-"`
}

var _ error = &CreateCompositeMetricMetricsResponseResponseBody{}

func (e *CreateCompositeMetricMetricsResponseResponseBody) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}

// CreateCompositeMetricMetricsResponseBody - Access is forbidden.
type CreateCompositeMetricMetricsResponseBody struct {
	// HTTP status code as defined in RFC 2817
	Code int64 `json:"code"`
	// Supporting description of the error
	Message  string                  `json:"message"`
	Target   *string                 `json:"target,omitempty"`
	HTTPMeta components.HTTPMetadata `json:"-"`
}

var _ error = &CreateCompositeMetricMetricsResponseBody{}

func (e *CreateCompositeMetricMetricsResponseBody) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}

// CreateCompositeMetricResponseBody - The server could not understand the request due to invalid syntax.
type CreateCompositeMetricResponseBody struct {
	// HTTP status code as defined in RFC 2817
	Code int64 `json:"code"`
	// Supporting description of the error
	Message  string                  `json:"message"`
	Target   *string                 `json:"target,omitempty"`
	HTTPMeta components.HTTPMetadata `json:"-"`
}

var _ error = &CreateCompositeMetricResponseBody{}

func (e *CreateCompositeMetricResponseBody) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}
