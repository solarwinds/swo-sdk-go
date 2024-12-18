// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"encoding/json"
	"fmt"
)

type MetricsAggregationMethods string

const (
	MetricsAggregationMethodsCount MetricsAggregationMethods = "COUNT"
	MetricsAggregationMethodsMin   MetricsAggregationMethods = "MIN"
	MetricsAggregationMethodsMax   MetricsAggregationMethods = "MAX"
	MetricsAggregationMethodsAvg   MetricsAggregationMethods = "AVG"
	MetricsAggregationMethodsSum   MetricsAggregationMethods = "SUM"
	MetricsAggregationMethodsLast  MetricsAggregationMethods = "LAST"
)

func (e MetricsAggregationMethods) ToPointer() *MetricsAggregationMethods {
	return &e
}
func (e *MetricsAggregationMethods) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "COUNT":
		fallthrough
	case "MIN":
		fallthrough
	case "MAX":
		fallthrough
	case "AVG":
		fallthrough
	case "SUM":
		fallthrough
	case "LAST":
		*e = MetricsAggregationMethods(v)
		return nil
	default:
		return fmt.Errorf("invalid value for MetricsAggregationMethods: %v", v)
	}
}
