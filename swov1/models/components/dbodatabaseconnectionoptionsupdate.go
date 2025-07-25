// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"github.com/solarwinds/swo-sdk-go/swov1/internal/utils"
)

type DboDatabaseConnectionOptionsUpdate struct {
	// Database server host
	Host *string `json:"host,omitempty"`
	// Database server port
	Port *string `default:"" json:"port"`
	// Database schema name
	Dbname *string `default:"" json:"dbname"`
	// Encrypted credentials for connecting to the database server when using basic authentication (username, password)
	// can be generated using this command:
	// ./dbo-headless-installer -swoparams=<SwoParamasJsonFile> --encrypt-creds --user=<USERNAME> --password=<PASSWORD>
	// Use the dbo-headless-installer binary located at:
	// https://agent-binaries.cloud.solarwinds.com/?prefix=dbo-headless-installer/latest/
	EncryptedCredentials *string `default:"" json:"encryptedCredentials"`
	// Username for connecting to database server needed only for auth methods other than basic auth
	User *string `default:"" json:"user"`
	// Enable ssl when agent connects to database server
	SslEnabled *bool `default:"false" json:"sslEnabled"`
	// SSL connection options, when sslEnabled is true
	SslOptions *DboDatabaseSslOptionsUpdate `json:"sslOptions,omitempty"`
	// Cloud region in case of database managed by cloud provider, required for IAM authentication
	CloudRegion *string `default:"" json:"cloudRegion"`
	// binding for packet sniffing for sniffer captureMethod (on-host), example: 0.0.0.0:6379,[::]:6379
	Bindings *string `default:"" json:"bindings"`
}

func (d DboDatabaseConnectionOptionsUpdate) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *DboDatabaseConnectionOptionsUpdate) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *DboDatabaseConnectionOptionsUpdate) GetHost() *string {
	if o == nil {
		return nil
	}
	return o.Host
}

func (o *DboDatabaseConnectionOptionsUpdate) GetPort() *string {
	if o == nil {
		return nil
	}
	return o.Port
}

func (o *DboDatabaseConnectionOptionsUpdate) GetDbname() *string {
	if o == nil {
		return nil
	}
	return o.Dbname
}

func (o *DboDatabaseConnectionOptionsUpdate) GetEncryptedCredentials() *string {
	if o == nil {
		return nil
	}
	return o.EncryptedCredentials
}

func (o *DboDatabaseConnectionOptionsUpdate) GetUser() *string {
	if o == nil {
		return nil
	}
	return o.User
}

func (o *DboDatabaseConnectionOptionsUpdate) GetSslEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.SslEnabled
}

func (o *DboDatabaseConnectionOptionsUpdate) GetSslOptions() *DboDatabaseSslOptionsUpdate {
	if o == nil {
		return nil
	}
	return o.SslOptions
}

func (o *DboDatabaseConnectionOptionsUpdate) GetCloudRegion() *string {
	if o == nil {
		return nil
	}
	return o.CloudRegion
}

func (o *DboDatabaseConnectionOptionsUpdate) GetBindings() *string {
	if o == nil {
		return nil
	}
	return o.Bindings
}
