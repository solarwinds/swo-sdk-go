// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"mockserver/internal/sdk/models/components"
	"mockserver/internal/sdk/utils"
	"time"
)

type ListMetricsRequest struct {
	// metric name
	Name *string `queryParam:"style=form,explode=false,name=name"`
	// Timestamp in ISO 8601 format in UTC timezone: yyyy-MM-ddTHH:mm:ssZ
	StartTime *time.Time `queryParam:"style=form,explode=false,name=startTime"`
	// Timestamp in ISO 8601 format in UTC timezone: yyyy-MM-ddTHH:mm:ssZ
	EndTime *time.Time `queryParam:"style=form,explode=false,name=endTime"`
	// Number of items in a response page. Default varies by API.
	PageSize *int `queryParam:"style=form,explode=false,name=pageSize"`
	// Token for the requested page.
	SkipToken *string `queryParam:"style=form,explode=false,name=skipToken"`
}

func (l ListMetricsRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(l, "", false)
}

func (l *ListMetricsRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &l, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ListMetricsRequest) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *ListMetricsRequest) GetStartTime() *time.Time {
	if o == nil {
		return nil
	}
	return o.StartTime
}

func (o *ListMetricsRequest) GetEndTime() *time.Time {
	if o == nil {
		return nil
	}
	return o.EndTime
}

func (o *ListMetricsRequest) GetPageSize() *int {
	if o == nil {
		return nil
	}
	return o.PageSize
}

func (o *ListMetricsRequest) GetSkipToken() *string {
	if o == nil {
		return nil
	}
	return o.SkipToken
}

// ListMetricsResponseBody - The request has succeeded.
type ListMetricsResponseBody struct {
	MetricsInfo []components.CommonMetricInfo `json:"metricsInfo"`
	PageInfo    components.CommonPageInfo     `json:"pageInfo"`
}

func (o *ListMetricsResponseBody) GetMetricsInfo() []components.CommonMetricInfo {
	if o == nil {
		return []components.CommonMetricInfo{}
	}
	return o.MetricsInfo
}

func (o *ListMetricsResponseBody) GetPageInfo() components.CommonPageInfo {
	if o == nil {
		return components.CommonPageInfo{}
	}
	return o.PageInfo
}

type ListMetricsResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// The request has succeeded.
	Object *ListMetricsResponseBody

	Next func() (*ListMetricsResponse, error)
}

func (o *ListMetricsResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *ListMetricsResponse) GetObject() *ListMetricsResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
