// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"encoding/json"
	"fmt"
	"mockserver/internal/sdk/utils"
)

// DatabaseSslOptionsUpdateSslMode - SSL mode such as require, verify-ca, verify-full as applicable
type DatabaseSslOptionsUpdateSslMode string

const (
	DatabaseSslOptionsUpdateSslModeRequire    DatabaseSslOptionsUpdateSslMode = "require"
	DatabaseSslOptionsUpdateSslModeVerifyCa   DatabaseSslOptionsUpdateSslMode = "verify-ca"
	DatabaseSslOptionsUpdateSslModeVerifyFull DatabaseSslOptionsUpdateSslMode = "verify-full"
)

func (e DatabaseSslOptionsUpdateSslMode) ToPointer() *DatabaseSslOptionsUpdateSslMode {
	return &e
}
func (e *DatabaseSslOptionsUpdateSslMode) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "require":
		fallthrough
	case "verify-ca":
		fallthrough
	case "verify-full":
		*e = DatabaseSslOptionsUpdateSslMode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DatabaseSslOptionsUpdateSslMode: %v", v)
	}
}

type DatabaseSslOptionsUpdate struct {
	// SSL mode such as require, verify-ca, verify-full as applicable
	SslMode *DatabaseSslOptionsUpdateSslMode `default:"require" json:"sslMode"`
	// CA file path
	SslCAPath *string `default:"" json:"sslCAPath"`
	// SSL key file path
	SslKeyPath *string `default:"" json:"sslKeyPath"`
	// SSL cert file path
	SslCertPath *string `default:"" json:"sslCertPath"`
}

func (d DatabaseSslOptionsUpdate) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *DatabaseSslOptionsUpdate) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *DatabaseSslOptionsUpdate) GetSslMode() *DatabaseSslOptionsUpdateSslMode {
	if o == nil {
		return nil
	}
	return o.SslMode
}

func (o *DatabaseSslOptionsUpdate) GetSslCAPath() *string {
	if o == nil {
		return nil
	}
	return o.SslCAPath
}

func (o *DatabaseSslOptionsUpdate) GetSslKeyPath() *string {
	if o == nil {
		return nil
	}
	return o.SslKeyPath
}

func (o *DatabaseSslOptionsUpdate) GetSslCertPath() *string {
	if o == nil {
		return nil
	}
	return o.SslCertPath
}
