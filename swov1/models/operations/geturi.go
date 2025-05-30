// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"github.com/solarwinds/swo-sdk-go/swov1/internal/utils"
	"github.com/solarwinds/swo-sdk-go/swov1/models/components"
	"time"
)

type GetURIRequest struct {
	EntityID string `pathParam:"style=simple,explode=false,name=entityId"`
}

func (o *GetURIRequest) GetEntityID() string {
	if o == nil {
		return ""
	}
	return o.EntityID
}

type Status string

const (
	StatusUp          Status = "up"
	StatusDown        Status = "down"
	StatusPaused      Status = "paused"
	StatusMaintenance Status = "maintenance"
	StatusUnknown     Status = "unknown"
)

func (e Status) ToPointer() *Status {
	return &e
}
func (e *Status) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "up":
		fallthrough
	case "down":
		fallthrough
	case "paused":
		fallthrough
	case "maintenance":
		fallthrough
	case "unknown":
		*e = Status(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Status: %v", v)
	}
}

// PlatformOptions - Configure cloud platforms of the synthetic availability test probes. If omitted or set to null, no particular cloud platform will be enforced.
type PlatformOptions struct {
	// Cloud platforms hosting synthetic probes.
	ProbePlatforms []components.ProbePlatform `json:"probePlatforms"`
	//   Use this field to configure whether availability tests should be performed from all selected
	//   platforms or one randomly selected platform. It has no effect if you provided only one platform
	//   in the `probePlatforms` field.
	//
	//   If set to true, a separate test is made from each of the selected platforms.
	//
	//   If set to false, only one of the selected platforms is chosen every time.
	//
	//   If omitted, the previous setting will stay in effect. If there is no previous setting, the value
	//   will default to false.
	TestFromAll *bool `json:"testFromAll,omitempty"`
}

func (o *PlatformOptions) GetProbePlatforms() []components.ProbePlatform {
	if o == nil {
		return []components.ProbePlatform{}
	}
	return o.ProbePlatforms
}

func (o *PlatformOptions) GetTestFromAll() *bool {
	if o == nil {
		return nil
	}
	return o.TestFromAll
}

// FailingTestLocations - How many locations must report a failure for an entity to be considered down.
type FailingTestLocations string

const (
	FailingTestLocationsAll FailingTestLocations = "all"
	FailingTestLocationsAny FailingTestLocations = "any"
)

func (e FailingTestLocations) ToPointer() *FailingTestLocations {
	return &e
}
func (e *FailingTestLocations) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "all":
		fallthrough
	case "any":
		*e = FailingTestLocations(v)
		return nil
	default:
		return fmt.Errorf("invalid value for FailingTestLocations: %v", v)
	}
}

// OutageConfiguration -   Default conditions when the entity is considered down.
//
//	If omitted or set to null, organization configuration will be used for this entity.
type OutageConfiguration struct {
	// How many locations must report a failure for an entity to be considered down.
	FailingTestLocations FailingTestLocations `json:"failingTestLocations"`
	// Number of consecutive failing tests for an entity to be considered down.
	ConsecutiveForDown int `json:"consecutiveForDown"`
}

func (o *OutageConfiguration) GetFailingTestLocations() FailingTestLocations {
	if o == nil {
		return FailingTestLocations("")
	}
	return o.FailingTestLocations
}

func (o *OutageConfiguration) GetConsecutiveForDown() int {
	if o == nil {
		return 0
	}
	return o.ConsecutiveForDown
}

// Ping -   Use this field to configure ping tests for the URI. If omitted or set to null, ping tests will be disabled.
//
//	One test type (ping or TCP) must be enabled for a URI.
type Ping struct {
	// Use this field to configure ping tests for the URI. If omitted or set to false, ping tests will be disabled.
	// One test type (ping or TCP) must be enabled for a URI.
	Enabled bool `json:"enabled"`
}

func (o *Ping) GetEnabled() bool {
	if o == nil {
		return false
	}
	return o.Enabled
}

// TCP -   Use this field to configure TCP tests for the URI. If omitted or set to null, TCP tests will be disabled.
//
//	One test type (ping or TCP) must be enabled for a URI.
type TCP struct {
	// Use this field to configure TCP tests for the URI. If omitted or set to false, TCP tests will be disabled.
	// One test type (ping or TCP) must be enabled for a URI.
	Enabled bool `json:"enabled"`
	// Port number to be used in TCP tests.
	Port int `json:"port"`
	// Use this field to specify a string to send in the body of a TCP request.
	StringToSend *string `json:"stringToSend,omitempty"`
	// Use this field to specify a string to search for in the body of a TCP response.
	StringToExpect *string `json:"stringToExpect,omitempty"`
}

func (o *TCP) GetEnabled() bool {
	if o == nil {
		return false
	}
	return o.Enabled
}

func (o *TCP) GetPort() int {
	if o == nil {
		return 0
	}
	return o.Port
}

func (o *TCP) GetStringToSend() *string {
	if o == nil {
		return nil
	}
	return o.StringToSend
}

func (o *TCP) GetStringToExpect() *string {
	if o == nil {
		return nil
	}
	return o.StringToExpect
}

// Protocol used to test availability of the URI.
type Protocol string

const (
	ProtocolPing Protocol = "PING"
	ProtocolTCP  Protocol = "TCP"
)

func (e Protocol) ToPointer() *Protocol {
	return &e
}
func (e *Protocol) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "PING":
		fallthrough
	case "TCP":
		*e = Protocol(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Protocol: %v", v)
	}
}

// AvailabilityCheckSettings - Availability tests configuration for the URI.
type AvailabilityCheckSettings struct {
	// Configure cloud platforms of the synthetic availability test probes. If omitted or set to null, no particular cloud platform will be enforced.
	PlatformOptions *PlatformOptions `json:"platformOptions,omitempty"`
	//   Configure locations of the synthetic availability test probes.
	//   Acceptable values depend on the selected type and actual values of existing probes.
	TestFrom components.TestFrom `json:"testFrom"`
	// Configure how often availability tests should be performed. Provide a number of seconds that is one of 60, 300, 600, 900, 1800, 3600, 7200, 14400.
	TestIntervalInSeconds float64 `json:"testIntervalInSeconds"`
	//   Default conditions when the entity is considered down.
	//   If omitted or set to null, organization configuration will be used for this entity.
	OutageConfiguration *OutageConfiguration `json:"outageConfiguration,omitempty"`
	//   Use this field to configure ping tests for the URI. If omitted or set to null, ping tests will be disabled.
	//   One test type (ping or TCP) must be enabled for a URI.
	Ping *Ping `json:"ping,omitempty"`
	//   Use this field to configure TCP tests for the URI. If omitted or set to null, TCP tests will be disabled.
	//   One test type (ping or TCP) must be enabled for a URI.
	TCP *TCP `json:"tcp,omitempty"`
	// Protocol used to test availability of the URI.
	Protocol Protocol `json:"protocol"`
}

func (o *AvailabilityCheckSettings) GetPlatformOptions() *PlatformOptions {
	if o == nil {
		return nil
	}
	return o.PlatformOptions
}

func (o *AvailabilityCheckSettings) GetTestFrom() components.TestFrom {
	if o == nil {
		return components.TestFrom{}
	}
	return o.TestFrom
}

func (o *AvailabilityCheckSettings) GetTestIntervalInSeconds() float64 {
	if o == nil {
		return 0.0
	}
	return o.TestIntervalInSeconds
}

func (o *AvailabilityCheckSettings) GetOutageConfiguration() *OutageConfiguration {
	if o == nil {
		return nil
	}
	return o.OutageConfiguration
}

func (o *AvailabilityCheckSettings) GetPing() *Ping {
	if o == nil {
		return nil
	}
	return o.Ping
}

func (o *AvailabilityCheckSettings) GetTCP() *TCP {
	if o == nil {
		return nil
	}
	return o.TCP
}

func (o *AvailabilityCheckSettings) GetProtocol() Protocol {
	if o == nil {
		return Protocol("")
	}
	return o.Protocol
}

// GetURIResponseBody - The request has succeeded.
type GetURIResponseBody struct {
	ID     string `json:"id"`
	Type   string `json:"type"`
	Status Status `json:"status"`
	//   Name of the URI, which must be unique within the organization.
	//   The name must also not contain any control characters, any white space other than space (U+0020), or any consecutive, leading or trailing spaces.
	Name string `json:"name"`
	// IP/domain address of the URI.
	IPOrDomain string `json:"ipOrDomain"`
	// Availability tests configuration for the URI.
	AvailabilityCheckSettings AvailabilityCheckSettings `json:"availabilityCheckSettings"`
	// Entity tags. Tag is a key-value pair, where there may be only single tag value for the same key.
	Tags []components.Tag `json:"tags,omitempty"`
	// Time when the last outage started.
	LastOutageStartTime *time.Time `json:"lastOutageStartTime,omitempty"`
	// Time when the last outage ended.
	LastOutageEndTime *time.Time `json:"lastOutageEndTime,omitempty"`
	// Time when the last test was performed.
	LastTestTime *time.Time `json:"lastTestTime,omitempty"`
	// Last time when a synthetic test failed.
	LastErrorTime *time.Time `json:"lastErrorTime,omitempty"`
	// Response time from the last synthetic check in milliseconds.
	LastResponseTime *int `json:"lastResponseTime,omitempty"`
}

func (g GetURIResponseBody) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(g, "", false)
}

func (g *GetURIResponseBody) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &g, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *GetURIResponseBody) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *GetURIResponseBody) GetType() string {
	if o == nil {
		return ""
	}
	return o.Type
}

func (o *GetURIResponseBody) GetStatus() Status {
	if o == nil {
		return Status("")
	}
	return o.Status
}

func (o *GetURIResponseBody) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *GetURIResponseBody) GetIPOrDomain() string {
	if o == nil {
		return ""
	}
	return o.IPOrDomain
}

func (o *GetURIResponseBody) GetAvailabilityCheckSettings() AvailabilityCheckSettings {
	if o == nil {
		return AvailabilityCheckSettings{}
	}
	return o.AvailabilityCheckSettings
}

func (o *GetURIResponseBody) GetTags() []components.Tag {
	if o == nil {
		return nil
	}
	return o.Tags
}

func (o *GetURIResponseBody) GetLastOutageStartTime() *time.Time {
	if o == nil {
		return nil
	}
	return o.LastOutageStartTime
}

func (o *GetURIResponseBody) GetLastOutageEndTime() *time.Time {
	if o == nil {
		return nil
	}
	return o.LastOutageEndTime
}

func (o *GetURIResponseBody) GetLastTestTime() *time.Time {
	if o == nil {
		return nil
	}
	return o.LastTestTime
}

func (o *GetURIResponseBody) GetLastErrorTime() *time.Time {
	if o == nil {
		return nil
	}
	return o.LastErrorTime
}

func (o *GetURIResponseBody) GetLastResponseTime() *int {
	if o == nil {
		return nil
	}
	return o.LastResponseTime
}

type GetURIResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// The request has succeeded.
	Object *GetURIResponseBody
}

func (o *GetURIResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetURIResponse) GetObject() *GetURIResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
