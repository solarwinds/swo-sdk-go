// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type ActivateAwsIntegrationRequest struct {
	// AWS Control Tower Management Account ID.
	ManagementAccountID string `json:"managementAccountId"`
	// AWS Accounts ID to be integrated.
	AccountID string `json:"accountId"`
	// True to enable the integration, false to disable.
	Enable bool `json:"enable"`
}

func (o *ActivateAwsIntegrationRequest) GetManagementAccountID() string {
	if o == nil {
		return ""
	}
	return o.ManagementAccountID
}

func (o *ActivateAwsIntegrationRequest) GetAccountID() string {
	if o == nil {
		return ""
	}
	return o.AccountID
}

func (o *ActivateAwsIntegrationRequest) GetEnable() bool {
	if o == nil {
		return false
	}
	return o.Enable
}
