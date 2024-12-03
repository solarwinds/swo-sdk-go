// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/solarwinds/swo-sdk-go/solarwinds/models/components"
)

type ListLogArchivesRequest struct {
	// List archives beginning at this time. Timestamp in ISO 8601 format in UTC timezone: yyyy-MM-ddTHH:mm:ssZ. Must be within the the last year.
	StartTime string `queryParam:"style=form,explode=false,name=startTime"`
	// List archives ending at this time. Timestamp in ISO 8601 format in UTC timezone: yyyy-MM-ddTHH:mm:ssZ. Must be within the the last year.
	EndTime string `queryParam:"style=form,explode=false,name=endTime"`
	// Number of items in a response page. Default varies by API.
	PageSize *int `queryParam:"style=form,explode=false,name=pageSize"`
	// Token for the requested page
	SkipToken *string `queryParam:"style=form,explode=false,name=skipToken"`
}

func (o *ListLogArchivesRequest) GetStartTime() string {
	if o == nil {
		return ""
	}
	return o.StartTime
}

func (o *ListLogArchivesRequest) GetEndTime() string {
	if o == nil {
		return ""
	}
	return o.EndTime
}

func (o *ListLogArchivesRequest) GetPageSize() *int {
	if o == nil {
		return nil
	}
	return o.PageSize
}

func (o *ListLogArchivesRequest) GetSkipToken() *string {
	if o == nil {
		return nil
	}
	return o.SkipToken
}

// ListLogArchivesResponseBody - The request has succeeded.
type ListLogArchivesResponseBody struct {
	LogArchives []components.LogsArchive  `json:"logArchives"`
	PageInfo    components.CommonPageInfo `json:"pageInfo"`
}

func (o *ListLogArchivesResponseBody) GetLogArchives() []components.LogsArchive {
	if o == nil {
		return []components.LogsArchive{}
	}
	return o.LogArchives
}

func (o *ListLogArchivesResponseBody) GetPageInfo() components.CommonPageInfo {
	if o == nil {
		return components.CommonPageInfo{}
	}
	return o.PageInfo
}

type ListLogArchivesResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// The request has succeeded.
	Object *ListLogArchivesResponseBody

	Next func() (*ListLogArchivesResponse, error)
}

func (o *ListLogArchivesResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *ListLogArchivesResponse) GetObject() *ListLogArchivesResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}