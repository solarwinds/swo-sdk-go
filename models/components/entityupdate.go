// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type EntityUpdate struct {
	// Entity display name / alias. This value is equal to name unless it's explicitly overridden.
	DisplayName *string `json:"displayName,omitempty"`
	// Entity tags. Tag is a key-value pair, where there may be only single tag value for the same key.
	Tags map[string]string `json:"tags,omitempty"`
}

func (o *EntityUpdate) GetDisplayName() *string {
	if o == nil {
		return nil
	}
	return o.DisplayName
}

func (o *EntityUpdate) GetTags() map[string]string {
	if o == nil {
		return nil
	}
	return o.Tags
}
