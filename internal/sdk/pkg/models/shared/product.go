// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/epilot-dev/terraform-provider-epilot-product/internal/sdk/pkg/utils"
	"time"
)

type Feature struct {
	ID      *string  `json:"_id,omitempty"`
	Tags    []string `json:"_tags,omitempty"`
	Feature *string  `json:"feature,omitempty"`
}

func (o *Feature) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *Feature) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}

func (o *Feature) GetFeature() *string {
	if o == nil {
		return nil
	}
	return o.Feature
}

// ProductType - The type of Product:
//
// | type | description |
// |----| ----|
// | `product` | Represents a physical good |
// | `service` | Represents a service or virtual product |
type ProductType string

const (
	ProductTypeProduct ProductType = "product"
	ProductTypeService ProductType = "service"
)

func (e ProductType) ToPointer() *ProductType {
	return &e
}

func (e *ProductType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "product":
		fallthrough
	case "service":
		*e = ProductType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ProductType: %v", v)
	}
}

type Product struct {
	// Access control list (ACL) for an entity. Defines sharing access to external orgs or users.
	ACL       BaseEntityACL `json:"_acl"`
	CreatedAt time.Time     `json:"_created_at"`
	ID        string        `json:"_id"`
	// Organization Id the entity belongs to
	Org       string            `json:"_org"`
	Owners    []BaseEntityOwner `json:"_owners"`
	Schema    string            `json:"_schema"`
	Tags      []string          `json:"_tags"`
	Title     string            `json:"_title"`
	UpdatedAt time.Time         `json:"_updated_at"`
	Active    bool              `json:"active"`
	// The product code
	Code *string `json:"code,omitempty"`
	// A description of the product. Multi-line supported.
	Description *string   `json:"description,omitempty"`
	Feature     []Feature `json:"feature,omitempty"`
	// Not visible to customers, only in internal tables
	InternalName *string `json:"internal_name,omitempty"`
	// The description for the product
	Name         string        `json:"name"`
	PriceOptions *BaseRelation `json:"price_options,omitempty"`
	// The type of Product:
	//
	// | type | description |
	// |----| ----|
	// | `product` | Represents a physical good |
	// | `service` | Represents a service or virtual product |
	//
	Type *ProductType `default:"product" json:"type"`
}

func (p Product) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(p, "", false)
}

func (p *Product) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &p, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *Product) GetACL() BaseEntityACL {
	if o == nil {
		return BaseEntityACL{}
	}
	return o.ACL
}

func (o *Product) GetCreatedAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.CreatedAt
}

func (o *Product) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *Product) GetOrg() string {
	if o == nil {
		return ""
	}
	return o.Org
}

func (o *Product) GetOwners() []BaseEntityOwner {
	if o == nil {
		return []BaseEntityOwner{}
	}
	return o.Owners
}

func (o *Product) GetSchema() string {
	if o == nil {
		return ""
	}
	return o.Schema
}

func (o *Product) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}

func (o *Product) GetTitle() string {
	if o == nil {
		return ""
	}
	return o.Title
}

func (o *Product) GetUpdatedAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.UpdatedAt
}

func (o *Product) GetActive() bool {
	if o == nil {
		return false
	}
	return o.Active
}

func (o *Product) GetCode() *string {
	if o == nil {
		return nil
	}
	return o.Code
}

func (o *Product) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *Product) GetFeature() []Feature {
	if o == nil {
		return nil
	}
	return o.Feature
}

func (o *Product) GetInternalName() *string {
	if o == nil {
		return nil
	}
	return o.InternalName
}

func (o *Product) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *Product) GetPriceOptions() *BaseRelation {
	if o == nil {
		return nil
	}
	return o.PriceOptions
}

func (o *Product) GetType() *ProductType {
	if o == nil {
		return nil
	}
	return o.Type
}
