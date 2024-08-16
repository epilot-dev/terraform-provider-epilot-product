// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/epilot-dev/terraform-provider-epilot-product/internal/sdk/internal/utils"
)

type ProductPatchSchema string

const (
	ProductPatchSchemaProduct ProductPatchSchema = "product"
)

func (e ProductPatchSchema) ToPointer() *ProductPatchSchema {
	return &e
}
func (e *ProductPatchSchema) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "product":
		*e = ProductPatchSchema(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ProductPatchSchema: %v", v)
	}
}

// ProductPatchType - The type of Product:
//
// | type | description |
// |----| ----|
// | `product` | Represents a physical good |
// | `service` | Represents a service or virtual product |
type ProductPatchType string

const (
	ProductPatchTypeProduct ProductPatchType = "product"
	ProductPatchTypeService ProductPatchType = "service"
)

func (e ProductPatchType) ToPointer() *ProductPatchType {
	return &e
}
func (e *ProductPatchType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "product":
		fallthrough
	case "service":
		*e = ProductPatchType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ProductPatchType: %v", v)
	}
}

type ProductPatch struct {
	// Additional fields that are not part of the schema
	Additional        map[string]any      `json:"__additional,omitempty"`
	AvailabilityFiles *BaseRelation       `json:"_availability_files,omitempty"`
	Files             *BaseRelation       `json:"_files,omitempty"`
	Purpose           []string            `json:"_purpose,omitempty"`
	Schema            *ProductPatchSchema `json:"_schema,omitempty"`
	Tags              []string            `json:"_tags,omitempty"`
	Active            *bool               `json:"active,omitempty"`
	// The product code
	Code *string `json:"code,omitempty"`
	// A description of the product. Multi-line supported.
	Description *string `json:"description,omitempty"`
	Feature     []any   `json:"feature,omitempty"`
	// Not visible to customers, only in internal tables
	InternalName *string `json:"internal_name,omitempty"`
	// The description for the product
	Name             *string       `json:"name,omitempty"`
	PriceOptions     *BaseRelation `json:"price_options,omitempty"`
	ProductDownloads *BaseRelation `json:"product_downloads,omitempty"`
	ProductImages    *BaseRelation `json:"product_images,omitempty"`
	// The type of Product:
	//
	// | type | description |
	// |----| ----|
	// | `product` | Represents a physical good |
	// | `service` | Represents a service or virtual product |
	//
	Type *ProductPatchType `default:"product" json:"type"`
}

func (p ProductPatch) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(p, "", false)
}

func (p *ProductPatch) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &p, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ProductPatch) GetAdditional() map[string]any {
	if o == nil {
		return nil
	}
	return o.Additional
}

func (o *ProductPatch) GetAvailabilityFiles() *BaseRelation {
	if o == nil {
		return nil
	}
	return o.AvailabilityFiles
}

func (o *ProductPatch) GetFiles() *BaseRelation {
	if o == nil {
		return nil
	}
	return o.Files
}

func (o *ProductPatch) GetPurpose() []string {
	if o == nil {
		return nil
	}
	return o.Purpose
}

func (o *ProductPatch) GetSchema() *ProductPatchSchema {
	if o == nil {
		return nil
	}
	return o.Schema
}

func (o *ProductPatch) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}

func (o *ProductPatch) GetActive() *bool {
	if o == nil {
		return nil
	}
	return o.Active
}

func (o *ProductPatch) GetCode() *string {
	if o == nil {
		return nil
	}
	return o.Code
}

func (o *ProductPatch) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *ProductPatch) GetFeature() []any {
	if o == nil {
		return nil
	}
	return o.Feature
}

func (o *ProductPatch) GetInternalName() *string {
	if o == nil {
		return nil
	}
	return o.InternalName
}

func (o *ProductPatch) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *ProductPatch) GetPriceOptions() *BaseRelation {
	if o == nil {
		return nil
	}
	return o.PriceOptions
}

func (o *ProductPatch) GetProductDownloads() *BaseRelation {
	if o == nil {
		return nil
	}
	return o.ProductDownloads
}

func (o *ProductPatch) GetProductImages() *BaseRelation {
	if o == nil {
		return nil
	}
	return o.ProductImages
}

func (o *ProductPatch) GetType() *ProductPatchType {
	if o == nil {
		return nil
	}
	return o.Type
}
