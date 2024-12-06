// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/solarwinds/swo-sdk-go/models/components"
)

type ListMetricAttributeValuesRequest struct {
	// metric name
	Name string `pathParam:"style=simple,explode=false,name=name"`
	// attribute name
	AttributeName string `pathParam:"style=simple,explode=false,name=attributeName"`
	// Timestamp in ISO 8601 format in UTC timezone: yyyy-MM-ddTHH:mm:ssZ
	StartTime *string `queryParam:"style=form,explode=false,name=startTime"`
	// Timestamp in ISO 8601 format in UTC timezone: yyyy-MM-ddTHH:mm:ssZ
	EndTime *string `queryParam:"style=form,explode=false,name=endTime"`
	// Number of items in a response page. Default varies by API.
	PageSize *int `queryParam:"style=form,explode=false,name=pageSize"`
	// Token for the requested page
	SkipToken *string `queryParam:"style=form,explode=false,name=skipToken"`
}

func (o *ListMetricAttributeValuesRequest) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *ListMetricAttributeValuesRequest) GetAttributeName() string {
	if o == nil {
		return ""
	}
	return o.AttributeName
}

func (o *ListMetricAttributeValuesRequest) GetStartTime() *string {
	if o == nil {
		return nil
	}
	return o.StartTime
}

func (o *ListMetricAttributeValuesRequest) GetEndTime() *string {
	if o == nil {
		return nil
	}
	return o.EndTime
}

func (o *ListMetricAttributeValuesRequest) GetPageSize() *int {
	if o == nil {
		return nil
	}
	return o.PageSize
}

func (o *ListMetricAttributeValuesRequest) GetSkipToken() *string {
	if o == nil {
		return nil
	}
	return o.SkipToken
}

// ListMetricAttributeValuesResponseBody - The request has succeeded.
type ListMetricAttributeValuesResponseBody struct {
	Name     string                    `json:"name"`
	Values   []string                  `json:"values"`
	PageInfo components.CommonPageInfo `json:"pageInfo"`
}

func (o *ListMetricAttributeValuesResponseBody) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *ListMetricAttributeValuesResponseBody) GetValues() []string {
	if o == nil {
		return []string{}
	}
	return o.Values
}

func (o *ListMetricAttributeValuesResponseBody) GetPageInfo() components.CommonPageInfo {
	if o == nil {
		return components.CommonPageInfo{}
	}
	return o.PageInfo
}

type ListMetricAttributeValuesResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// The request has succeeded.
	Object *ListMetricAttributeValuesResponseBody

	Next func() (*ListMetricAttributeValuesResponse, error)
}

func (o *ListMetricAttributeValuesResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *ListMetricAttributeValuesResponse) GetObject() *ListMetricAttributeValuesResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}