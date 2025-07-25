// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"encoding/json"
	"fmt"
)

type DboDatabaseAuthMethod string

const (
	DboDatabaseAuthMethodBasic             DboDatabaseAuthMethod = "basic"
	DboDatabaseAuthMethodAwsiam            DboDatabaseAuthMethod = "awsiam"
	DboDatabaseAuthMethodWindows           DboDatabaseAuthMethod = "windows"
	DboDatabaseAuthMethodEntraclientsecret DboDatabaseAuthMethod = "entraclientsecret"
	DboDatabaseAuthMethodEntraclientcert   DboDatabaseAuthMethod = "entraclientcert"
	DboDatabaseAuthMethodCert              DboDatabaseAuthMethod = "cert"
)

func (e DboDatabaseAuthMethod) ToPointer() *DboDatabaseAuthMethod {
	return &e
}
func (e *DboDatabaseAuthMethod) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "basic":
		fallthrough
	case "awsiam":
		fallthrough
	case "windows":
		fallthrough
	case "entraclientsecret":
		fallthrough
	case "entraclientcert":
		fallthrough
	case "cert":
		*e = DboDatabaseAuthMethod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DboDatabaseAuthMethod: %v", v)
	}
}
