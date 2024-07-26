// Code generated by Speakeasy (https://speakeasyapi.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/epilot-dev/terraform-provider-epilot-product/internal/sdk/internal/utils"
)

// PricePatchBillingDurationUnit - The billing period duration unit
type PricePatchBillingDurationUnit string

const (
	PricePatchBillingDurationUnitWeeks  PricePatchBillingDurationUnit = "weeks"
	PricePatchBillingDurationUnitMonths PricePatchBillingDurationUnit = "months"
	PricePatchBillingDurationUnitYears  PricePatchBillingDurationUnit = "years"
)

func (e PricePatchBillingDurationUnit) ToPointer() *PricePatchBillingDurationUnit {
	return &e
}
func (e *PricePatchBillingDurationUnit) UnmarshalJSON(data []byte) error {
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
		*e = PricePatchBillingDurationUnit(v)
		return nil
	default:
		return fmt.Errorf("invalid value for PricePatchBillingDurationUnit: %v", v)
	}
}

// PricePatchNoticeTimeUnit - The notice period duration unit
type PricePatchNoticeTimeUnit string

const (
	PricePatchNoticeTimeUnitWeeks  PricePatchNoticeTimeUnit = "weeks"
	PricePatchNoticeTimeUnitMonths PricePatchNoticeTimeUnit = "months"
	PricePatchNoticeTimeUnitYears  PricePatchNoticeTimeUnit = "years"
)

func (e PricePatchNoticeTimeUnit) ToPointer() *PricePatchNoticeTimeUnit {
	return &e
}
func (e *PricePatchNoticeTimeUnit) UnmarshalJSON(data []byte) error {
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
		*e = PricePatchNoticeTimeUnit(v)
		return nil
	default:
		return fmt.Errorf("invalid value for PricePatchNoticeTimeUnit: %v", v)
	}
}

// PricePatchPriceComponents - A set of [price](/api/pricing#tag/simple_price_schema) components that define the composite price.
type PricePatchPriceComponents struct {
	DollarRelation []PriceComponentRelation `json:"$relation,omitempty"`
}

func (o *PricePatchPriceComponents) GetDollarRelation() []PriceComponentRelation {
	if o == nil {
		return nil
	}
	return o.DollarRelation
}

// PricePatchPriceDisplayInJourneys - Defines the way the price amount is display in epilot journeys.
type PricePatchPriceDisplayInJourneys string

const (
	PricePatchPriceDisplayInJourneysShowPrice           PricePatchPriceDisplayInJourneys = "show_price"
	PricePatchPriceDisplayInJourneysShowAsStartingPrice PricePatchPriceDisplayInJourneys = "show_as_starting_price"
	PricePatchPriceDisplayInJourneysShowAsOnRequest     PricePatchPriceDisplayInJourneys = "show_as_on_request"
)

func (e PricePatchPriceDisplayInJourneys) ToPointer() *PricePatchPriceDisplayInJourneys {
	return &e
}
func (e *PricePatchPriceDisplayInJourneys) UnmarshalJSON(data []byte) error {
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
		*e = PricePatchPriceDisplayInJourneys(v)
		return nil
	default:
		return fmt.Errorf("invalid value for PricePatchPriceDisplayInJourneys: %v", v)
	}
}

// PricePatchPricingModel - Describes how to compute the price per period. Either `per_unit`, `tiered_graduated` or `tiered_volume`.
// - `per_unit` indicates that the fixed amount (specified in unit_amount or unit_amount_decimal) will be charged per unit in quantity
// - `tiered_graduated` indicates that the unit pricing will be computed using tiers attribute. The customer pays the price per unit in every range their purchase rises through.
// - `tiered_volume` indicates that the unit pricing will be computed using tiers attribute. The customer pays the same unit price for all purchased units.
// - `tiered_flatfee` While similar to tiered_volume, tiered flat fee charges for the same price (flat) for the entire range instead using the unit price to multiply the quantity.
type PricePatchPricingModel string

const (
	PricePatchPricingModelPerUnit         PricePatchPricingModel = "per_unit"
	PricePatchPricingModelTieredVolume    PricePatchPricingModel = "tiered_volume"
	PricePatchPricingModelTieredGraduated PricePatchPricingModel = "tiered_graduated"
	PricePatchPricingModelTieredFlatfee   PricePatchPricingModel = "tiered_flatfee"
)

func (e PricePatchPricingModel) ToPointer() *PricePatchPricingModel {
	return &e
}
func (e *PricePatchPricingModel) UnmarshalJSON(data []byte) error {
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
		*e = PricePatchPricingModel(v)
		return nil
	default:
		return fmt.Errorf("invalid value for PricePatchPricingModel: %v", v)
	}
}

// PricePatchRenewalDurationUnit - The renewal period duration unit
type PricePatchRenewalDurationUnit string

const (
	PricePatchRenewalDurationUnitWeeks  PricePatchRenewalDurationUnit = "weeks"
	PricePatchRenewalDurationUnitMonths PricePatchRenewalDurationUnit = "months"
	PricePatchRenewalDurationUnitYears  PricePatchRenewalDurationUnit = "years"
)

func (e PricePatchRenewalDurationUnit) ToPointer() *PricePatchRenewalDurationUnit {
	return &e
}
func (e *PricePatchRenewalDurationUnit) UnmarshalJSON(data []byte) error {
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
		*e = PricePatchRenewalDurationUnit(v)
		return nil
	default:
		return fmt.Errorf("invalid value for PricePatchRenewalDurationUnit: %v", v)
	}
}

// PricePatchTerminationTimeUnit - The termination period duration unit
type PricePatchTerminationTimeUnit string

const (
	PricePatchTerminationTimeUnitWeeks  PricePatchTerminationTimeUnit = "weeks"
	PricePatchTerminationTimeUnitMonths PricePatchTerminationTimeUnit = "months"
	PricePatchTerminationTimeUnitYears  PricePatchTerminationTimeUnit = "years"
)

func (e PricePatchTerminationTimeUnit) ToPointer() *PricePatchTerminationTimeUnit {
	return &e
}
func (e *PricePatchTerminationTimeUnit) UnmarshalJSON(data []byte) error {
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
		*e = PricePatchTerminationTimeUnit(v)
		return nil
	default:
		return fmt.Errorf("invalid value for PricePatchTerminationTimeUnit: %v", v)
	}
}

// PricePatchType - One of `one_time` or `recurring` depending on whether the price is for a one-time purchase or a recurring (subscription) purchase.
type PricePatchType string

const (
	PricePatchTypeOneTime   PricePatchType = "one_time"
	PricePatchTypeRecurring PricePatchType = "recurring"
)

func (e PricePatchType) ToPointer() *PricePatchType {
	return &e
}
func (e *PricePatchType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "one_time":
		fallthrough
	case "recurring":
		*e = PricePatchType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for PricePatchType: %v", v)
	}
}

type PricePatch struct {
	// Whether the price can be used for new purchases.
	Active *bool `json:"active,omitempty"`
	// The billing period duration
	BillingDurationAmount *float64 `json:"billing_duration_amount,omitempty"`
	// The billing period duration unit
	BillingDurationUnit *PricePatchBillingDurationUnit `json:"billing_duration_unit,omitempty"`
	// A brief description of the price.
	Description *string `json:"description,omitempty"`
	// The flag for prices that contain price components.
	IsCompositePrice *bool `json:"is_composite_price,omitempty"`
	// Specifies whether the price is considered `inclusive` of taxes or not.
	IsTaxInclusive *bool `default:"false" json:"is_tax_inclusive"`
	// A detailed description of the price. This is shown on the order document and order table. Multi-line supported.
	LongDescription *string `json:"long_description,omitempty"`
	// The notice period duration
	NoticeTimeAmount *float64 `json:"notice_time_amount,omitempty"`
	// The notice period duration unit
	NoticeTimeUnit *PricePatchNoticeTimeUnit `json:"notice_time_unit,omitempty"`
	// A set of [price](/api/pricing#tag/simple_price_schema) components that define the composite price.
	PriceComponents *PricePatchPriceComponents `json:"price_components,omitempty"`
	// Defines the way the price amount is display in epilot journeys.
	PriceDisplayInJourneys *PricePatchPriceDisplayInJourneys `json:"price_display_in_journeys,omitempty"`
	// Describes how to compute the price per period. Either `per_unit`, `tiered_graduated` or `tiered_volume`.
	// - `per_unit` indicates that the fixed amount (specified in unit_amount or unit_amount_decimal) will be charged per unit in quantity
	// - `tiered_graduated` indicates that the unit pricing will be computed using tiers attribute. The customer pays the price per unit in every range their purchase rises through.
	// - `tiered_volume` indicates that the unit pricing will be computed using tiers attribute. The customer pays the same unit price for all purchased units.
	// - `tiered_flatfee` While similar to tiered_volume, tiered flat fee charges for the same price (flat) for the entire range instead using the unit price to multiply the quantity.
	//
	PricingModel *PricePatchPricingModel `default:"per_unit" json:"pricing_model"`
	// The renewal period duration
	RenewalDurationAmount *float64 `json:"renewal_duration_amount,omitempty"`
	// The renewal period duration unit
	RenewalDurationUnit *PricePatchRenewalDurationUnit `json:"renewal_duration_unit,omitempty"`
	Tax                 any                            `json:"tax,omitempty"`
	// The termination period duration
	TerminationTimeAmount *float64 `json:"termination_time_amount,omitempty"`
	// The termination period duration unit
	TerminationTimeUnit *PricePatchTerminationTimeUnit `json:"termination_time_unit,omitempty"`
	// Defines an array of tiers. Each tier has an upper bound, an unit amount and a flat fee.
	//
	Tiers []PriceTier `json:"tiers,omitempty"`
	// One of `one_time` or `recurring` depending on whether the price is for a one-time purchase or a recurring (subscription) purchase.
	Type *PricePatchType `default:"one_time" json:"type"`
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

func (p PricePatch) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(p, "", false)
}

func (p *PricePatch) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &p, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *PricePatch) GetActive() *bool {
	if o == nil {
		return nil
	}
	return o.Active
}

func (o *PricePatch) GetBillingDurationAmount() *float64 {
	if o == nil {
		return nil
	}
	return o.BillingDurationAmount
}

func (o *PricePatch) GetBillingDurationUnit() *PricePatchBillingDurationUnit {
	if o == nil {
		return nil
	}
	return o.BillingDurationUnit
}

func (o *PricePatch) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *PricePatch) GetIsCompositePrice() *bool {
	if o == nil {
		return nil
	}
	return o.IsCompositePrice
}

func (o *PricePatch) GetIsTaxInclusive() *bool {
	if o == nil {
		return nil
	}
	return o.IsTaxInclusive
}

func (o *PricePatch) GetLongDescription() *string {
	if o == nil {
		return nil
	}
	return o.LongDescription
}

func (o *PricePatch) GetNoticeTimeAmount() *float64 {
	if o == nil {
		return nil
	}
	return o.NoticeTimeAmount
}

func (o *PricePatch) GetNoticeTimeUnit() *PricePatchNoticeTimeUnit {
	if o == nil {
		return nil
	}
	return o.NoticeTimeUnit
}

func (o *PricePatch) GetPriceComponents() *PricePatchPriceComponents {
	if o == nil {
		return nil
	}
	return o.PriceComponents
}

func (o *PricePatch) GetPriceDisplayInJourneys() *PricePatchPriceDisplayInJourneys {
	if o == nil {
		return nil
	}
	return o.PriceDisplayInJourneys
}

func (o *PricePatch) GetPricingModel() *PricePatchPricingModel {
	if o == nil {
		return nil
	}
	return o.PricingModel
}

func (o *PricePatch) GetRenewalDurationAmount() *float64 {
	if o == nil {
		return nil
	}
	return o.RenewalDurationAmount
}

func (o *PricePatch) GetRenewalDurationUnit() *PricePatchRenewalDurationUnit {
	if o == nil {
		return nil
	}
	return o.RenewalDurationUnit
}

func (o *PricePatch) GetTax() any {
	if o == nil {
		return nil
	}
	return o.Tax
}

func (o *PricePatch) GetTerminationTimeAmount() *float64 {
	if o == nil {
		return nil
	}
	return o.TerminationTimeAmount
}

func (o *PricePatch) GetTerminationTimeUnit() *PricePatchTerminationTimeUnit {
	if o == nil {
		return nil
	}
	return o.TerminationTimeUnit
}

func (o *PricePatch) GetTiers() []PriceTier {
	if o == nil {
		return nil
	}
	return o.Tiers
}

func (o *PricePatch) GetType() *PricePatchType {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *PricePatch) GetUnit() *string {
	if o == nil {
		return nil
	}
	return o.Unit
}

func (o *PricePatch) GetUnitAmount() *float64 {
	if o == nil {
		return nil
	}
	return o.UnitAmount
}

func (o *PricePatch) GetUnitAmountCurrency() *string {
	if o == nil {
		return nil
	}
	return o.UnitAmountCurrency
}

func (o *PricePatch) GetUnitAmountDecimal() *string {
	if o == nil {
		return nil
	}
	return o.UnitAmountDecimal
}

func (o *PricePatch) GetVariablePrice() *bool {
	if o == nil {
		return nil
	}
	return o.VariablePrice
}