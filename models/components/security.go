// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type Security struct {
	BearerAuth *string `security:"scheme,type=http,subtype=bearer,name=Authorization,env=solarwinds_bearer_auth"`
}

func (o *Security) GetBearerAuth() *string {
	if o == nil {
		return nil
	}
	return o.BearerAuth
}
