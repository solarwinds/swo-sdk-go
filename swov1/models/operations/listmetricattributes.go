// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/solarwinds/swo-sdk-go/swov1/internal/utils"
	"github.com/solarwinds/swo-sdk-go/swov1/models/components"
	"time"
)

type ListMetricAttributesRequest struct {
	// metric name
	Name string `pathParam:"style=simple,explode=false,name=name"`
	// Timestamp in ISO 8601 format in UTC timezone: yyyy-MM-ddTHH:mm:ssZ
	StartTime *time.Time `queryParam:"style=form,explode=false,name=startTime"`
	// Timestamp in ISO 8601 format in UTC timezone: yyyy-MM-ddTHH:mm:ssZ
	EndTime *time.Time `queryParam:"style=form,explode=false,name=endTime"`
	// Number of items in a response page. Default varies by API.
	PageSize *int `queryParam:"style=form,explode=false,name=pageSize"`
	// Token for the requested page
	SkipToken *string `queryParam:"style=form,explode=false,name=skipToken"`
}

func (l ListMetricAttributesRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(l, "", false)
}

func (l *ListMetricAttributesRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &l, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ListMetricAttributesRequest) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *ListMetricAttributesRequest) GetStartTime() *time.Time {
	if o == nil {
		return nil
	}
	return o.StartTime
}

func (o *ListMetricAttributesRequest) GetEndTime() *time.Time {
	if o == nil {
		return nil
	}
	return o.EndTime
}

func (o *ListMetricAttributesRequest) GetPageSize() *int {
	if o == nil {
		return nil
	}
	return o.PageSize
}

func (o *ListMetricAttributesRequest) GetSkipToken() *string {
	if o == nil {
		return nil
	}
	return o.SkipToken
}

// ListMetricAttributesResponseBody - The request has succeeded.
type ListMetricAttributesResponseBody struct {
	Names    []string                  `json:"names"`
	PageInfo components.CommonPageInfo `json:"pageInfo"`
}

func (o *ListMetricAttributesResponseBody) GetNames() []string {
	if o == nil {
		return []string{}
	}
	return o.Names
}

func (o *ListMetricAttributesResponseBody) GetPageInfo() components.CommonPageInfo {
	if o == nil {
		return components.CommonPageInfo{}
	}
	return o.PageInfo
}

type ListMetricAttributesResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// The request has succeeded.
	Object *ListMetricAttributesResponseBody

	Next func() (*ListMetricAttributesResponse, error)
}

func (o *ListMetricAttributesResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *ListMetricAttributesResponse) GetObject() *ListMetricAttributesResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
