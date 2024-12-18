// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type Entity struct {
	// The ID of the entity
	ID string `json:"id"`
	// The type of the entity
	Type string `json:"type"`
	// The name of the entity
	Name *string `json:"name,omitempty"`
}

func (o *Entity) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *Entity) GetType() string {
	if o == nil {
		return ""
	}
	return o.Type
}

func (o *Entity) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}
