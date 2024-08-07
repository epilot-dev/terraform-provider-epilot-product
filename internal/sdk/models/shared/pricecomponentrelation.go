// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type PriceComponentRelation struct {
	// An arbitrary set of tags attached to the composite price - component relation
	Tags []string `json:"_tags,omitempty"`
	// The id of the price component
	EntityID *string `json:"entity_id,omitempty"`
}

func (o *PriceComponentRelation) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}

func (o *PriceComponentRelation) GetEntityID() *string {
	if o == nil {
		return nil
	}
	return o.EntityID
}
