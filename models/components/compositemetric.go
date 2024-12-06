// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type CompositeMetric struct {
	// Name of the composite metric
	Name string `json:"name"`
	// Display name of the composite metric. A short description of the metric.
	DisplayName string `json:"displayName"`
	// Description of the composite metric. A detailed description of the metric.
	Description string `json:"description"`
	// PromQL query to calculate the composite metric
	Formula string `json:"formula"`
	// Unit of the composite metric
	Units string `json:"units"`
}

func (o *CompositeMetric) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *CompositeMetric) GetDisplayName() string {
	if o == nil {
		return ""
	}
	return o.DisplayName
}

func (o *CompositeMetric) GetDescription() string {
	if o == nil {
		return ""
	}
	return o.Description
}

func (o *CompositeMetric) GetFormula() string {
	if o == nil {
		return ""
	}
	return o.Formula
}

func (o *CompositeMetric) GetUnits() string {
	if o == nil {
		return ""
	}
	return o.Units
}
