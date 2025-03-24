// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/epilot-dev/terraform-provider-epilot-product/internal/sdk/internal/utils"
)

type PriceCreateSchema string

const (
	PriceCreateSchemaPrice PriceCreateSchema = "price"
)

func (e PriceCreateSchema) ToPointer() *PriceCreateSchema {
	return &e
}
func (e *PriceCreateSchema) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "price":
		*e = PriceCreateSchema(v)
		return nil
	default:
		return fmt.Errorf("invalid value for PriceCreateSchema: %v", v)
	}
}

// PriceCreateBillingDurationUnit - The billing period duration unit
type PriceCreateBillingDurationUnit string

const (
	PriceCreateBillingDurationUnitWeeks  PriceCreateBillingDurationUnit = "weeks"
	PriceCreateBillingDurationUnitMonths PriceCreateBillingDurationUnit = "months"
	PriceCreateBillingDurationUnitYears  PriceCreateBillingDurationUnit = "years"
)

func (e PriceCreateBillingDurationUnit) ToPointer() *PriceCreateBillingDurationUnit {
	return &e
}
func (e *PriceCreateBillingDurationUnit) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "weeks":
		fallthrough
	case "months":
		fallthrough
	case "years":
		*e = PriceCreateBillingDurationUnit(v)
		return nil
	default:
		return fmt.Errorf("invalid value for PriceCreateBillingDurationUnit: %v", v)
	}
}

// PriceCreateNoticeTimeUnit - The notice period duration unit
type PriceCreateNoticeTimeUnit string

const (
	PriceCreateNoticeTimeUnitWeeks  PriceCreateNoticeTimeUnit = "weeks"
	PriceCreateNoticeTimeUnitMonths PriceCreateNoticeTimeUnit = "months"
	PriceCreateNoticeTimeUnitYears  PriceCreateNoticeTimeUnit = "years"
)

func (e PriceCreateNoticeTimeUnit) ToPointer() *PriceCreateNoticeTimeUnit {
	return &e
}
func (e *PriceCreateNoticeTimeUnit) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "weeks":
		fallthrough
	case "months":
		fallthrough
	case "years":
		*e = PriceCreateNoticeTimeUnit(v)
		return nil
	default:
		return fmt.Errorf("invalid value for PriceCreateNoticeTimeUnit: %v", v)
	}
}

// PriceCreatePriceComponents - A set of [price](/api/pricing#tag/simple_price_schema) components that define the composite price.
type PriceCreatePriceComponents struct {
	DollarRelation []PriceComponentRelation `json:"$relation,omitempty"`
}

func (o *PriceCreatePriceComponents) GetDollarRelation() []PriceComponentRelation {
	if o == nil {
		return nil
	}
	return o.DollarRelation
}

// PriceCreatePriceDisplayInJourneys - Defines the way the price amount is display in epilot journeys.
type PriceCreatePriceDisplayInJourneys string

const (
	PriceCreatePriceDisplayInJourneysShowPrice           PriceCreatePriceDisplayInJourneys = "show_price"
	PriceCreatePriceDisplayInJourneysShowAsStartingPrice PriceCreatePriceDisplayInJourneys = "show_as_starting_price"
	PriceCreatePriceDisplayInJourneysShowAsOnRequest     PriceCreatePriceDisplayInJourneys = "show_as_on_request"
)

func (e PriceCreatePriceDisplayInJourneys) ToPointer() *PriceCreatePriceDisplayInJourneys {
	return &e
}
func (e *PriceCreatePriceDisplayInJourneys) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "show_price":
		fallthrough
	case "show_as_starting_price":
		fallthrough
	case "show_as_on_request":
		*e = PriceCreatePriceDisplayInJourneys(v)
		return nil
	default:
		return fmt.Errorf("invalid value for PriceCreatePriceDisplayInJourneys: %v", v)
	}
}

// PriceCreatePricingModel - Describes how to compute the price per period. Either `per_unit`, `tiered_graduated` or `tiered_volume`.
// - `per_unit` indicates that the fixed amount (specified in unit_amount or unit_amount_decimal) will be charged per unit in quantity
// - `tiered_graduated` indicates that the unit pricing will be computed using tiers attribute. The customer pays the price per unit in every range their purchase rises through.
// - `tiered_volume` indicates that the unit pricing will be computed using tiers attribute. The customer pays the same unit price for all purchased units.
// - `tiered_flatfee` While similar to tiered_volume, tiered flat fee charges for the same price (flat) for the entire range instead using the unit price to multiply the quantity.
type PriceCreatePricingModel string

const (
	PriceCreatePricingModelPerUnit         PriceCreatePricingModel = "per_unit"
	PriceCreatePricingModelTieredVolume    PriceCreatePricingModel = "tiered_volume"
	PriceCreatePricingModelTieredGraduated PriceCreatePricingModel = "tiered_graduated"
	PriceCreatePricingModelTieredFlatfee   PriceCreatePricingModel = "tiered_flatfee"
)

func (e PriceCreatePricingModel) ToPointer() *PriceCreatePricingModel {
	return &e
}
func (e *PriceCreatePricingModel) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "per_unit":
		fallthrough
	case "tiered_volume":
		fallthrough
	case "tiered_graduated":
		fallthrough
	case "tiered_flatfee":
		*e = PriceCreatePricingModel(v)
		return nil
	default:
		return fmt.Errorf("invalid value for PriceCreatePricingModel: %v", v)
	}
}

// PriceCreateRenewalDurationUnit - The renewal period duration unit
type PriceCreateRenewalDurationUnit string

const (
	PriceCreateRenewalDurationUnitWeeks  PriceCreateRenewalDurationUnit = "weeks"
	PriceCreateRenewalDurationUnitMonths PriceCreateRenewalDurationUnit = "months"
	PriceCreateRenewalDurationUnitYears  PriceCreateRenewalDurationUnit = "years"
)

func (e PriceCreateRenewalDurationUnit) ToPointer() *PriceCreateRenewalDurationUnit {
	return &e
}
func (e *PriceCreateRenewalDurationUnit) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "weeks":
		fallthrough
	case "months":
		fallthrough
	case "years":
		*e = PriceCreateRenewalDurationUnit(v)
		return nil
	default:
		return fmt.Errorf("invalid value for PriceCreateRenewalDurationUnit: %v", v)
	}
}

// PriceCreateTerminationTimeUnit - The termination period duration unit
type PriceCreateTerminationTimeUnit string

const (
	PriceCreateTerminationTimeUnitWeeks  PriceCreateTerminationTimeUnit = "weeks"
	PriceCreateTerminationTimeUnitMonths PriceCreateTerminationTimeUnit = "months"
	PriceCreateTerminationTimeUnitYears  PriceCreateTerminationTimeUnit = "years"
)

func (e PriceCreateTerminationTimeUnit) ToPointer() *PriceCreateTerminationTimeUnit {
	return &e
}
func (e *PriceCreateTerminationTimeUnit) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "weeks":
		fallthrough
	case "months":
		fallthrough
	case "years":
		*e = PriceCreateTerminationTimeUnit(v)
		return nil
	default:
		return fmt.Errorf("invalid value for PriceCreateTerminationTimeUnit: %v", v)
	}
}

// PriceCreateType - One of `one_time` or `recurring` depending on whether the price is for a one-time purchase or a recurring (subscription) purchase.
type PriceCreateType string

const (
	PriceCreateTypeOneTime   PriceCreateType = "one_time"
	PriceCreateTypeRecurring PriceCreateType = "recurring"
)

func (e PriceCreateType) ToPointer() *PriceCreateType {
	return &e
}
func (e *PriceCreateType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "one_time":
		fallthrough
	case "recurring":
		*e = PriceCreateType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for PriceCreateType: %v", v)
	}
}

type PriceCreate struct {
	// Additional fields that are not part of the schema
	Additional map[string]any `json:"__additional,omitempty"`
	Files      *BaseRelation  `json:"_files,omitempty"`
	// Manifest ID used to create/update the entity
	Manifest []string           `json:"_manifest,omitempty"`
	Purpose  []string           `json:"_purpose,omitempty"`
	Schema   *PriceCreateSchema `json:"_schema,omitempty"`
	Tags     []string           `json:"_tags,omitempty"`
	// Whether the price can be used for new purchases.
	Active bool `json:"active"`
	// The billing period duration
	BillingDurationAmount *float64 `json:"billing_duration_amount,omitempty"`
	// The billing period duration unit
	BillingDurationUnit *PriceCreateBillingDurationUnit `json:"billing_duration_unit,omitempty"`
	// A brief description of the price.
	Description string `json:"description"`
	// The flag for prices that contain price components.
	IsCompositePrice *bool `json:"is_composite_price,omitempty"`
	// Specifies whether the price is considered `inclusive` of taxes or not.
	IsTaxInclusive *bool `default:"false" json:"is_tax_inclusive"`
	// A detailed description of the price. This is shown on the order document and order table. Multi-line supported.
	LongDescription *string `json:"long_description,omitempty"`
	// The notice period duration
	NoticeTimeAmount *float64 `json:"notice_time_amount,omitempty"`
	// The notice period duration unit
	NoticeTimeUnit *PriceCreateNoticeTimeUnit `json:"notice_time_unit,omitempty"`
	// A set of [price](/api/pricing#tag/simple_price_schema) components that define the composite price.
	PriceComponents *PriceCreatePriceComponents `json:"price_components,omitempty"`
	// Defines the way the price amount is display in epilot journeys.
	PriceDisplayInJourneys *PriceCreatePriceDisplayInJourneys `json:"price_display_in_journeys,omitempty"`
	// Describes how to compute the price per period. Either `per_unit`, `tiered_graduated` or `tiered_volume`.
	// - `per_unit` indicates that the fixed amount (specified in unit_amount or unit_amount_decimal) will be charged per unit in quantity
	// - `tiered_graduated` indicates that the unit pricing will be computed using tiers attribute. The customer pays the price per unit in every range their purchase rises through.
	// - `tiered_volume` indicates that the unit pricing will be computed using tiers attribute. The customer pays the same unit price for all purchased units.
	// - `tiered_flatfee` While similar to tiered_volume, tiered flat fee charges for the same price (flat) for the entire range instead using the unit price to multiply the quantity.
	//
	PricingModel *PriceCreatePricingModel `default:"per_unit" json:"pricing_model"`
	// The renewal period duration
	RenewalDurationAmount *float64 `json:"renewal_duration_amount,omitempty"`
	// The renewal period duration unit
	RenewalDurationUnit *PriceCreateRenewalDurationUnit `json:"renewal_duration_unit,omitempty"`
	Tax                 any                             `json:"tax,omitempty"`
	// The termination period duration
	TerminationTimeAmount *float64 `json:"termination_time_amount,omitempty"`
	// The termination period duration unit
	TerminationTimeUnit *PriceCreateTerminationTimeUnit `json:"termination_time_unit,omitempty"`
	// Defines an array of tiers. Each tier has an upper bound, an unit amount and a flat fee.
	//
	Tiers []PriceTier `json:"tiers,omitempty"`
	// One of `one_time` or `recurring` depending on whether the price is for a one-time purchase or a recurring (subscription) purchase.
	Type *PriceCreateType `default:"one_time" json:"type"`
	// The unit of measurement used for display purposes and possibly for calculations when the price is variable.
	Unit *string `json:"unit,omitempty"`
	// The unit amount in cents to be charged, represented as a whole integer if possible.
	UnitAmount *float64 `json:"unit_amount,omitempty"`
	// Three-letter ISO currency code, in lowercase.
	UnitAmountCurrency *string `json:"unit_amount_currency,omitempty"`
	// The unit amount in cents to be charged, represented as a decimal string with at most 12 decimal places.
	UnitAmountDecimal *string `json:"unit_amount_decimal,omitempty"`
	// The flag for prices that can be influenced by external variables such as user input.
	VariablePrice *bool `default:"false" json:"variable_price"`
}

func (p PriceCreate) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(p, "", false)
}

func (p *PriceCreate) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &p, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *PriceCreate) GetAdditional() map[string]any {
	if o == nil {
		return nil
	}
	return o.Additional
}

func (o *PriceCreate) GetFiles() *BaseRelation {
	if o == nil {
		return nil
	}
	return o.Files
}

func (o *PriceCreate) GetManifest() []string {
	if o == nil {
		return nil
	}
	return o.Manifest
}

func (o *PriceCreate) GetPurpose() []string {
	if o == nil {
		return nil
	}
	return o.Purpose
}

func (o *PriceCreate) GetSchema() *PriceCreateSchema {
	if o == nil {
		return nil
	}
	return o.Schema
}

func (o *PriceCreate) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}

func (o *PriceCreate) GetActive() bool {
	if o == nil {
		return false
	}
	return o.Active
}

func (o *PriceCreate) GetBillingDurationAmount() *float64 {
	if o == nil {
		return nil
	}
	return o.BillingDurationAmount
}

func (o *PriceCreate) GetBillingDurationUnit() *PriceCreateBillingDurationUnit {
	if o == nil {
		return nil
	}
	return o.BillingDurationUnit
}

func (o *PriceCreate) GetDescription() string {
	if o == nil {
		return ""
	}
	return o.Description
}

func (o *PriceCreate) GetIsCompositePrice() *bool {
	if o == nil {
		return nil
	}
	return o.IsCompositePrice
}

func (o *PriceCreate) GetIsTaxInclusive() *bool {
	if o == nil {
		return nil
	}
	return o.IsTaxInclusive
}

func (o *PriceCreate) GetLongDescription() *string {
	if o == nil {
		return nil
	}
	return o.LongDescription
}

func (o *PriceCreate) GetNoticeTimeAmount() *float64 {
	if o == nil {
		return nil
	}
	return o.NoticeTimeAmount
}

func (o *PriceCreate) GetNoticeTimeUnit() *PriceCreateNoticeTimeUnit {
	if o == nil {
		return nil
	}
	return o.NoticeTimeUnit
}

func (o *PriceCreate) GetPriceComponents() *PriceCreatePriceComponents {
	if o == nil {
		return nil
	}
	return o.PriceComponents
}

func (o *PriceCreate) GetPriceDisplayInJourneys() *PriceCreatePriceDisplayInJourneys {
	if o == nil {
		return nil
	}
	return o.PriceDisplayInJourneys
}

func (o *PriceCreate) GetPricingModel() *PriceCreatePricingModel {
	if o == nil {
		return nil
	}
	return o.PricingModel
}

func (o *PriceCreate) GetRenewalDurationAmount() *float64 {
	if o == nil {
		return nil
	}
	return o.RenewalDurationAmount
}

func (o *PriceCreate) GetRenewalDurationUnit() *PriceCreateRenewalDurationUnit {
	if o == nil {
		return nil
	}
	return o.RenewalDurationUnit
}

func (o *PriceCreate) GetTax() any {
	if o == nil {
		return nil
	}
	return o.Tax
}

func (o *PriceCreate) GetTerminationTimeAmount() *float64 {
	if o == nil {
		return nil
	}
	return o.TerminationTimeAmount
}

func (o *PriceCreate) GetTerminationTimeUnit() *PriceCreateTerminationTimeUnit {
	if o == nil {
		return nil
	}
	return o.TerminationTimeUnit
}

func (o *PriceCreate) GetTiers() []PriceTier {
	if o == nil {
		return nil
	}
	return o.Tiers
}

func (o *PriceCreate) GetType() *PriceCreateType {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *PriceCreate) GetUnit() *string {
	if o == nil {
		return nil
	}
	return o.Unit
}

func (o *PriceCreate) GetUnitAmount() *float64 {
	if o == nil {
		return nil
	}
	return o.UnitAmount
}

func (o *PriceCreate) GetUnitAmountCurrency() *string {
	if o == nil {
		return nil
	}
	return o.UnitAmountCurrency
}

func (o *PriceCreate) GetUnitAmountDecimal() *string {
	if o == nil {
		return nil
	}
	return o.UnitAmountDecimal
}

func (o *PriceCreate) GetVariablePrice() *bool {
	if o == nil {
		return nil
	}
	return o.VariablePrice
}
