// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type ProbePlatformOptions struct {
	// Cloud platforms hosting synthetic probes.
	ProbePlatforms []ProbePlatform `json:"probePlatforms"`
	TestFromAll    *bool           `json:"testFromAll,omitempty"`
}

func (o *ProbePlatformOptions) GetProbePlatforms() []ProbePlatform {
	if o == nil {
		return []ProbePlatform{}
	}
	return o.ProbePlatforms
}

func (o *ProbePlatformOptions) GetTestFromAll() *bool {
	if o == nil {
		return nil
	}
	return o.TestFromAll
}