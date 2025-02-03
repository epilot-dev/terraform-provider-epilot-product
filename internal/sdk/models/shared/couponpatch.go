// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type CouponPatchSchema string

const (
	CouponPatchSchemaCoupon CouponPatchSchema = "coupon"
)

func (e CouponPatchSchema) ToPointer() *CouponPatchSchema {
	return &e
}
func (e *CouponPatchSchema) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "coupon":
		*e = CouponPatchSchema(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CouponPatchSchema: %v", v)
	}
}

// CouponPatchCashbackPeriod - The cashback period, for now it's limited to either 0 months or 12 months
type CouponPatchCashbackPeriod string

const (
	CouponPatchCashbackPeriodZero   CouponPatchCashbackPeriod = "0"
	CouponPatchCashbackPeriodTwelve CouponPatchCashbackPeriod = "12"
)

func (e CouponPatchCashbackPeriod) ToPointer() *CouponPatchCashbackPeriod {
	return &e
}
func (e *CouponPatchCashbackPeriod) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "0":
		fallthrough
	case "12":
		*e = CouponPatchCashbackPeriod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CouponPatchCashbackPeriod: %v", v)
	}
}

type CouponPatchCategory string

const (
	CouponPatchCategoryDiscount CouponPatchCategory = "discount"
	CouponPatchCategoryCashback CouponPatchCategory = "cashback"
)

func (e CouponPatchCategory) ToPointer() *CouponPatchCategory {
	return &e
}
func (e *CouponPatchCategory) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "discount":
		fallthrough
	case "cashback":
		*e = CouponPatchCategory(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CouponPatchCategory: %v", v)
	}
}

type CouponPatchType string

const (
	CouponPatchTypeFixed      CouponPatchType = "fixed"
	CouponPatchTypePercentage CouponPatchType = "percentage"
)

func (e CouponPatchType) ToPointer() *CouponPatchType {
	return &e
}
func (e *CouponPatchType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "fixed":
		fallthrough
	case "percentage":
		*e = CouponPatchType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CouponPatchType: %v", v)
	}
}

type CouponPatch struct {
	// Additional fields that are not part of the schema
	Additional map[string]any `json:"__additional,omitempty"`
	Files      *BaseRelation  `json:"_files,omitempty"`
	// Manifest ID used to create/update the entity
	Manifest []string           `json:"_manifest,omitempty"`
	Schema   *CouponPatchSchema `json:"_schema,omitempty"`
	Tags     []string           `json:"_tags,omitempty"`
	Active   *bool              `json:"active,omitempty"`
	// The cashback period, for now it's limited to either 0 months or 12 months
	CashbackPeriod *CouponPatchCashbackPeriod `json:"cashback_period,omitempty"`
	Category       *CouponPatchCategory       `json:"category,omitempty"`
	Description    *string                    `json:"description,omitempty"`
	// Use if type is set to fixed. The fixed amount in cents to be discounted, represented as a whole integer.
	FixedValue *float64 `json:"fixed_value,omitempty"`
	// Use if type is set to fixed. Three-letter ISO currency code, in lowercase.
	FixedValueCurrency *string `json:"fixed_value_currency,omitempty"`
	// Use if type is set to fixed. The unit amount in cents to be discounted, represented as a decimal string with at most 12 decimal places.
	FixedValueDecimal *string `json:"fixed_value_decimal,omitempty"`
	Name              *string `json:"name,omitempty"`
	// Use if type is set to percentage. The percentage to be discounted, represented as a whole integer.
	PercentageValue *string       `json:"percentage_value,omitempty"`
	Prices          *BaseRelation `json:"prices,omitempty"`
	// Map of ids of promo codes with their usage count
	PromoCodeUsage any         `json:"promo_code_usage,omitempty"`
	PromoCodes     []PromoCode `json:"promo_codes,omitempty"`
	// Whether the coupon requires a promo code to be applied
	RequiresPromoCode *bool            `json:"requires_promo_code,omitempty"`
	Type              *CouponPatchType `json:"type,omitempty"`
}

func (o *CouponPatch) GetAdditional() map[string]any {
	if o == nil {
		return nil
	}
	return o.Additional
}

func (o *CouponPatch) GetFiles() *BaseRelation {
	if o == nil {
		return nil
	}
	return o.Files
}

func (o *CouponPatch) GetManifest() []string {
	if o == nil {
		return nil
	}
	return o.Manifest
}

func (o *CouponPatch) GetSchema() *CouponPatchSchema {
	if o == nil {
		return nil
	}
	return o.Schema
}

func (o *CouponPatch) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}

func (o *CouponPatch) GetActive() *bool {
	if o == nil {
		return nil
	}
	return o.Active
}

func (o *CouponPatch) GetCashbackPeriod() *CouponPatchCashbackPeriod {
	if o == nil {
		return nil
	}
	return o.CashbackPeriod
}

func (o *CouponPatch) GetCategory() *CouponPatchCategory {
	if o == nil {
		return nil
	}
	return o.Category
}

func (o *CouponPatch) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *CouponPatch) GetFixedValue() *float64 {
	if o == nil {
		return nil
	}
	return o.FixedValue
}

func (o *CouponPatch) GetFixedValueCurrency() *string {
	if o == nil {
		return nil
	}
	return o.FixedValueCurrency
}

func (o *CouponPatch) GetFixedValueDecimal() *string {
	if o == nil {
		return nil
	}
	return o.FixedValueDecimal
}

func (o *CouponPatch) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *CouponPatch) GetPercentageValue() *string {
	if o == nil {
		return nil
	}
	return o.PercentageValue
}

func (o *CouponPatch) GetPrices() *BaseRelation {
	if o == nil {
		return nil
	}
	return o.Prices
}

func (o *CouponPatch) GetPromoCodeUsage() any {
	if o == nil {
		return nil
	}
	return o.PromoCodeUsage
}

func (o *CouponPatch) GetPromoCodes() []PromoCode {
	if o == nil {
		return nil
	}
	return o.PromoCodes
}

func (o *CouponPatch) GetRequiresPromoCode() *bool {
	if o == nil {
		return nil
	}
	return o.RequiresPromoCode
}

func (o *CouponPatch) GetType() *CouponPatchType {
	if o == nil {
		return nil
	}
	return o.Type
}
