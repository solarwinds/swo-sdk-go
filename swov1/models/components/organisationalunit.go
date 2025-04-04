// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type OrganisationalUnit struct {
	// AWS Account ID/Org unit ID.
	ChildID string `json:"child_id"`
	// AWS Account Name/Org unit Name.
	ChildName string `json:"child_name"`
	// Parent Org unit ID.
	ParentID *string `json:"parent_id,omitempty"`
}

func (o *OrganisationalUnit) GetChildID() string {
	if o == nil {
		return ""
	}
	return o.ChildID
}

func (o *OrganisationalUnit) GetChildName() string {
	if o == nil {
		return ""
	}
	return o.ChildName
}

func (o *OrganisationalUnit) GetParentID() *string {
	if o == nil {
		return nil
	}
	return o.ParentID
}
