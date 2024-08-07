// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/epilot-dev/terraform-provider-epilot-product/internal/sdk/internal/utils"
)

// ProductCreateType - The type of Product:
//
// | type | description |
// |----| ----|
// | `product` | Represents a physical good |
// | `service` | Represents a service or virtual product |
type ProductCreateType string

const (
	ProductCreateTypeProduct ProductCreateType = "product"
	ProductCreateTypeService ProductCreateType = "service"
)

func (e ProductCreateType) ToPointer() *ProductCreateType {
	return &e
}
func (e *ProductCreateType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "product":
		fallthrough
	case "service":
		*e = ProductCreateType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ProductCreateType: %v", v)
	}
}

type ProductCreate struct {
	Active bool `json:"active"`
	// The product code
	Code *string `json:"code,omitempty"`
	// A description of the product. Multi-line supported.
	Description *string `json:"description,omitempty"`
	Feature     []any   `json:"feature,omitempty"`
	// Not visible to customers, only in internal tables
	InternalName *string `json:"internal_name,omitempty"`
	// The description for the product
	Name             string        `json:"name"`
	PriceOptions     *BaseRelation `json:"price_options,omitempty"`
	ProductDownloads any           `json:"product_downloads,omitempty"`
	ProductImages    any           `json:"product_images,omitempty"`
	// The type of Product:
	//
	// | type | description |
	// |----| ----|
	// | `product` | Represents a physical good |
	// | `service` | Represents a service or virtual product |
	//
	Type *ProductCreateType `default:"product" json:"type"`
}

func (p ProductCreate) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(p, "", false)
}

func (p *ProductCreate) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &p, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ProductCreate) GetActive() bool {
	if o == nil {
		return false
	}
	return o.Active
}

func (o *ProductCreate) GetCode() *string {
	if o == nil {
		return nil
	}
	return o.Code
}

func (o *ProductCreate) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *ProductCreate) GetFeature() []any {
	if o == nil {
		return nil
	}
	return o.Feature
}

func (o *ProductCreate) GetInternalName() *string {
	if o == nil {
		return nil
	}
	return o.InternalName
}

func (o *ProductCreate) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *ProductCreate) GetPriceOptions() *BaseRelation {
	if o == nil {
		return nil
	}
	return o.PriceOptions
}

func (o *ProductCreate) GetProductDownloads() any {
	if o == nil {
		return nil
	}
	return o.ProductDownloads
}

func (o *ProductCreate) GetProductImages() any {
	if o == nil {
		return nil
	}
	return o.ProductImages
}

func (o *ProductCreate) GetType() *ProductCreateType {
	if o == nil {
		return nil
	}
	return o.Type
}
