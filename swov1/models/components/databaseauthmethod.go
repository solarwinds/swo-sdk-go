// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"encoding/json"
	"fmt"
)

type DatabaseAuthMethod string

const (
	DatabaseAuthMethodBasic             DatabaseAuthMethod = "basic"
	DatabaseAuthMethodAwsiam            DatabaseAuthMethod = "awsiam"
	DatabaseAuthMethodWindows           DatabaseAuthMethod = "windows"
	DatabaseAuthMethodEntraclientsecret DatabaseAuthMethod = "entraclientsecret"
	DatabaseAuthMethodEntraclientcert   DatabaseAuthMethod = "entraclientcert"
	DatabaseAuthMethodCert              DatabaseAuthMethod = "cert"
)

func (e DatabaseAuthMethod) ToPointer() *DatabaseAuthMethod {
	return &e
}
func (e *DatabaseAuthMethod) UnmarshalJSON(data []byte) error {
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
		*e = DatabaseAuthMethod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DatabaseAuthMethod: %v", v)
	}
}
