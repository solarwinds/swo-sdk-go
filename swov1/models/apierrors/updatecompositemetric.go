// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package apierrors

import (
	"encoding/json"
	"fmt"
	"github.com/solarwinds/swo-sdk-go/swov1/models/components"
)

type UpdateCompositeMetricMetricsResponseCode string

const (
	UpdateCompositeMetricMetricsResponseCodeInternalServerError UpdateCompositeMetricMetricsResponseCode = "InternalServerError"
)

func (e UpdateCompositeMetricMetricsResponseCode) ToPointer() *UpdateCompositeMetricMetricsResponseCode {
	return &e
}
func (e *UpdateCompositeMetricMetricsResponseCode) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "InternalServerError":
		*e = UpdateCompositeMetricMetricsResponseCode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for UpdateCompositeMetricMetricsResponseCode: %v", v)
	}
}

// UpdateCompositeMetricMetricsResponse500ResponseBody - Server error
type UpdateCompositeMetricMetricsResponse500ResponseBody struct {
	Code     UpdateCompositeMetricMetricsResponseCode `json:"code"`
	Message  string                                   `json:"message"`
	HTTPMeta components.HTTPMetadata                  `json:"-"`
}

var _ error = &UpdateCompositeMetricMetricsResponse500ResponseBody{}

func (e *UpdateCompositeMetricMetricsResponse500ResponseBody) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}

type UpdateCompositeMetricMetricsCode string

const (
	UpdateCompositeMetricMetricsCodeNotFound UpdateCompositeMetricMetricsCode = "NotFound"
)

func (e UpdateCompositeMetricMetricsCode) ToPointer() *UpdateCompositeMetricMetricsCode {
	return &e
}
func (e *UpdateCompositeMetricMetricsCode) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "NotFound":
		*e = UpdateCompositeMetricMetricsCode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for UpdateCompositeMetricMetricsCode: %v", v)
	}
}

// UpdateCompositeMetricMetricsResponseResponseBody - The server cannot find the requested resource.
type UpdateCompositeMetricMetricsResponseResponseBody struct {
	Code     UpdateCompositeMetricMetricsCode `json:"code"`
	Message  string                           `json:"message"`
	HTTPMeta components.HTTPMetadata          `json:"-"`
}

var _ error = &UpdateCompositeMetricMetricsResponseResponseBody{}

func (e *UpdateCompositeMetricMetricsResponseResponseBody) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}

// UpdateCompositeMetricMetricsResponseBody - Access is forbidden.
type UpdateCompositeMetricMetricsResponseBody struct {
	// HTTP status code as defined in RFC 2817
	Code int64 `json:"code"`
	// Supporting description of the error
	Message  string                  `json:"message"`
	Target   *string                 `json:"target,omitempty"`
	HTTPMeta components.HTTPMetadata `json:"-"`
}

var _ error = &UpdateCompositeMetricMetricsResponseBody{}

func (e *UpdateCompositeMetricMetricsResponseBody) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}

type UpdateCompositeMetricCode string

const (
	UpdateCompositeMetricCodeUnauthorized UpdateCompositeMetricCode = "Unauthorized"
)

func (e UpdateCompositeMetricCode) ToPointer() *UpdateCompositeMetricCode {
	return &e
}
func (e *UpdateCompositeMetricCode) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Unauthorized":
		*e = UpdateCompositeMetricCode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for UpdateCompositeMetricCode: %v", v)
	}
}

// UpdateCompositeMetricResponseBody - Access is unauthorized.
type UpdateCompositeMetricResponseBody struct {
	Code     UpdateCompositeMetricCode `json:"code"`
	Message  string                    `json:"message"`
	HTTPMeta components.HTTPMetadata   `json:"-"`
}

var _ error = &UpdateCompositeMetricResponseBody{}

func (e *UpdateCompositeMetricResponseBody) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}
