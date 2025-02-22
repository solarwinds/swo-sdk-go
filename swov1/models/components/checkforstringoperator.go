// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"encoding/json"
	"fmt"
)

type CheckForStringOperator string

const (
	CheckForStringOperatorContains       CheckForStringOperator = "CONTAINS"
	CheckForStringOperatorDoesNotContain CheckForStringOperator = "DOES_NOT_CONTAIN"
)

func (e CheckForStringOperator) ToPointer() *CheckForStringOperator {
	return &e
}
func (e *CheckForStringOperator) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "CONTAINS":
		fallthrough
	case "DOES_NOT_CONTAIN":
		*e = CheckForStringOperator(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CheckForStringOperator: %v", v)
	}
}
