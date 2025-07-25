// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"encoding/json"
	"fmt"
)

// IntegrationType - Integration Type would either be STREAMING or POLLING
type IntegrationType string

const (
	IntegrationTypeStreaming IntegrationType = "STREAMING"
	IntegrationTypePolling   IntegrationType = "POLLING"
)

func (e IntegrationType) ToPointer() *IntegrationType {
	return &e
}
func (e *IntegrationType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "STREAMING":
		fallthrough
	case "POLLING":
		*e = IntegrationType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for IntegrationType: %v", v)
	}
}

type CloudAccountsAwsUpdateIntegrationResponse struct {
	// String array of AWS Regions in which the AWS Account is integrated.
	SelectedRegions []string `json:"selectedRegions"`
	// New External ID generated that is to be integrated as part of AWS Account's trust relationship.
	ExternalID string `json:"externalId"`
	// Integration ID generated by SolarWinds Observability that is unique for each Account integrated.
	IntegrationID string `json:"integrationId"`
	// Integration Type would either be STREAMING or POLLING
	IntegrationType IntegrationType `json:"integrationType"`
	// Flag to indicate if the AWS Account is newly integrated to SolarWinds Observability under the Control Tower Management Account Integration.
	IsNewAccount bool `json:"isNewAccount"`
}

func (o *CloudAccountsAwsUpdateIntegrationResponse) GetSelectedRegions() []string {
	if o == nil {
		return []string{}
	}
	return o.SelectedRegions
}

func (o *CloudAccountsAwsUpdateIntegrationResponse) GetExternalID() string {
	if o == nil {
		return ""
	}
	return o.ExternalID
}

func (o *CloudAccountsAwsUpdateIntegrationResponse) GetIntegrationID() string {
	if o == nil {
		return ""
	}
	return o.IntegrationID
}

func (o *CloudAccountsAwsUpdateIntegrationResponse) GetIntegrationType() IntegrationType {
	if o == nil {
		return IntegrationType("")
	}
	return o.IntegrationType
}

func (o *CloudAccountsAwsUpdateIntegrationResponse) GetIsNewAccount() bool {
	if o == nil {
		return false
	}
	return o.IsNewAccount
}
