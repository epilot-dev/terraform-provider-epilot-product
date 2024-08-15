// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/epilot-dev/terraform-provider-epilot-product/internal/sdk/internal/utils"
	"time"
)

type Schema string

const (
	SchemaPrice Schema = "price"
)

func (e Schema) ToPointer() *Schema {
	return &e
}
func (e *Schema) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "price":
		*e = Schema(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Schema: %v", v)
	}
}

// BillingDurationUnit - The billing period duration unit
type BillingDurationUnit string

const (
	BillingDurationUnitWeeks  BillingDurationUnit = "weeks"
	BillingDurationUnitMonths BillingDurationUnit = "months"
	BillingDurationUnitYears  BillingDurationUnit = "years"
)

func (e BillingDurationUnit) ToPointer() *BillingDurationUnit {
	return &e
}
func (e *BillingDurationUnit) UnmarshalJSON(data []byte) error {
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
		*e = BillingDurationUnit(v)
		return nil
	default:
		return fmt.Errorf("invalid value for BillingDurationUnit: %v", v)
	}
}

// NoticeTimeUnit - The notice period duration unit
type NoticeTimeUnit string

const (
	NoticeTimeUnitWeeks  NoticeTimeUnit = "weeks"
	NoticeTimeUnitMonths NoticeTimeUnit = "months"
	NoticeTimeUnitYears  NoticeTimeUnit = "years"
)

func (e NoticeTimeUnit) ToPointer() *NoticeTimeUnit {
	return &e
}
func (e *NoticeTimeUnit) UnmarshalJSON(data []byte) error {
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
		*e = NoticeTimeUnit(v)
		return nil
	default:
		return fmt.Errorf("invalid value for NoticeTimeUnit: %v", v)
	}
}

// PriceComponents - A set of [price](/api/pricing#tag/simple_price_schema) components that define the composite price.
type PriceComponents struct {
	DollarRelation []PriceComponentRelation `json:"$relation,omitempty"`
}

func (o *PriceComponents) GetDollarRelation() []PriceComponentRelation {
	if o == nil {
		return nil
	}
	return o.DollarRelation
}

// PriceDisplayInJourneys - Defines the way the price amount is display in epilot journeys.
type PriceDisplayInJourneys string

const (
	PriceDisplayInJourneysShowPrice           PriceDisplayInJourneys = "show_price"
	PriceDisplayInJourneysShowAsStartingPrice PriceDisplayInJourneys = "show_as_starting_price"
	PriceDisplayInJourneysShowAsOnRequest     PriceDisplayInJourneys = "show_as_on_request"
)

func (e PriceDisplayInJourneys) ToPointer() *PriceDisplayInJourneys {
	return &e
}
func (e *PriceDisplayInJourneys) UnmarshalJSON(data []byte) error {
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
		*e = PriceDisplayInJourneys(v)
		return nil
	default:
		return fmt.Errorf("invalid value for PriceDisplayInJourneys: %v", v)
	}
}

// PricingModel - Describes how to compute the price per period. Either `per_unit`, `tiered_graduated` or `tiered_volume`.
// - `per_unit` indicates that the fixed amount (specified in unit_amount or unit_amount_decimal) will be charged per unit in quantity
// - `tiered_graduated` indicates that the unit pricing will be computed using tiers attribute. The customer pays the price per unit in every range their purchase rises through.
// - `tiered_volume` indicates that the unit pricing will be computed using tiers attribute. The customer pays the same unit price for all purchased units.
// - `tiered_flatfee` While similar to tiered_volume, tiered flat fee charges for the same price (flat) for the entire range instead using the unit price to multiply the quantity.
type PricingModel string

const (
	PricingModelPerUnit         PricingModel = "per_unit"
	PricingModelTieredVolume    PricingModel = "tiered_volume"
	PricingModelTieredGraduated PricingModel = "tiered_graduated"
	PricingModelTieredFlatfee   PricingModel = "tiered_flatfee"
)

func (e PricingModel) ToPointer() *PricingModel {
	return &e
}
func (e *PricingModel) UnmarshalJSON(data []byte) error {
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
		*e = PricingModel(v)
		return nil
	default:
		return fmt.Errorf("invalid value for PricingModel: %v", v)
	}
}

// RenewalDurationUnit - The renewal period duration unit
type RenewalDurationUnit string

const (
	RenewalDurationUnitWeeks  RenewalDurationUnit = "weeks"
	RenewalDurationUnitMonths RenewalDurationUnit = "months"
	RenewalDurationUnitYears  RenewalDurationUnit = "years"
)

func (e RenewalDurationUnit) ToPointer() *RenewalDurationUnit {
	return &e
}
func (e *RenewalDurationUnit) UnmarshalJSON(data []byte) error {
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
		*e = RenewalDurationUnit(v)
		return nil
	default:
		return fmt.Errorf("invalid value for RenewalDurationUnit: %v", v)
	}
}

// TerminationTimeUnit - The termination period duration unit
type TerminationTimeUnit string

const (
	TerminationTimeUnitWeeks  TerminationTimeUnit = "weeks"
	TerminationTimeUnitMonths TerminationTimeUnit = "months"
	TerminationTimeUnitYears  TerminationTimeUnit = "years"
)

func (e TerminationTimeUnit) ToPointer() *TerminationTimeUnit {
	return &e
}
func (e *TerminationTimeUnit) UnmarshalJSON(data []byte) error {
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
		*e = TerminationTimeUnit(v)
		return nil
	default:
		return fmt.Errorf("invalid value for TerminationTimeUnit: %v", v)
	}
}

// Type - One of `one_time` or `recurring` depending on whether the price is for a one-time purchase or a recurring (subscription) purchase.
type Type string

const (
	TypeOneTime   Type = "one_time"
	TypeRecurring Type = "recurring"
)

func (e Type) ToPointer() *Type {
	return &e
}
func (e *Type) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "one_time":
		fallthrough
	case "recurring":
		*e = Type(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Type: %v", v)
	}
}

type Price struct {
	// Additional fields that are not part of the schema
	Additional map[string]any `json:"__additional,omitempty"`
	// Access control list (ACL) for an entity. Defines sharing access to external orgs or users.
	ACL       *BaseEntityACL `json:"_acl,omitempty"`
	CreatedAt *time.Time     `json:"_created_at,omitempty"`
	Files     *BaseRelation  `json:"_files,omitempty"`
	ID        *string        `json:"_id,omitempty"`
	// Organization Id the entity belongs to
	Org       string            `json:"_org"`
	Owners    []BaseEntityOwner `json:"_owners,omitempty"`
	Schema    Schema            `json:"_schema"`
	Tags      []string          `json:"_tags,omitempty"`
	Title     *string           `json:"_title,omitempty"`
	UpdatedAt *time.Time        `json:"_updated_at,omitempty"`
	// Whether the price can be used for new purchases.
	Active bool `json:"active"`
	// The billing period duration
	BillingDurationAmount *float64 `json:"billing_duration_amount,omitempty"`
	// The billing period duration unit
	BillingDurationUnit *BillingDurationUnit `json:"billing_duration_unit,omitempty"`
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
	NoticeTimeUnit *NoticeTimeUnit `json:"notice_time_unit,omitempty"`
	// A set of [price](/api/pricing#tag/simple_price_schema) components that define the composite price.
	PriceComponents *PriceComponents `json:"price_components,omitempty"`
	// Defines the way the price amount is display in epilot journeys.
	PriceDisplayInJourneys *PriceDisplayInJourneys `json:"price_display_in_journeys,omitempty"`
	// Describes how to compute the price per period. Either `per_unit`, `tiered_graduated` or `tiered_volume`.
	// - `per_unit` indicates that the fixed amount (specified in unit_amount or unit_amount_decimal) will be charged per unit in quantity
	// - `tiered_graduated` indicates that the unit pricing will be computed using tiers attribute. The customer pays the price per unit in every range their purchase rises through.
	// - `tiered_volume` indicates that the unit pricing will be computed using tiers attribute. The customer pays the same unit price for all purchased units.
	// - `tiered_flatfee` While similar to tiered_volume, tiered flat fee charges for the same price (flat) for the entire range instead using the unit price to multiply the quantity.
	//
	PricingModel *PricingModel `default:"per_unit" json:"pricing_model"`
	// The renewal period duration
	RenewalDurationAmount *float64 `json:"renewal_duration_amount,omitempty"`
	// The renewal period duration unit
	RenewalDurationUnit *RenewalDurationUnit `json:"renewal_duration_unit,omitempty"`
	Tax                 any                  `json:"tax,omitempty"`
	// The termination period duration
	TerminationTimeAmount *float64 `json:"termination_time_amount,omitempty"`
	// The termination period duration unit
	TerminationTimeUnit *TerminationTimeUnit `json:"termination_time_unit,omitempty"`
	// Defines an array of tiers. Each tier has an upper bound, an unit amount and a flat fee.
	//
	Tiers []PriceTier `json:"tiers,omitempty"`
	// One of `one_time` or `recurring` depending on whether the price is for a one-time purchase or a recurring (subscription) purchase.
	Type *Type `default:"one_time" json:"type"`
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

func (p Price) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(p, "", false)
}

func (p *Price) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &p, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *Price) GetAdditional() map[string]any {
	if o == nil {
		return nil
	}
	return o.Additional
}

func (o *Price) GetACL() *BaseEntityACL {
	if o == nil {
		return nil
	}
	return o.ACL
}

func (o *Price) GetCreatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *Price) GetFiles() *BaseRelation {
	if o == nil {
		return nil
	}
	return o.Files
}

func (o *Price) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *Price) GetOrg() string {
	if o == nil {
		return ""
	}
	return o.Org
}

func (o *Price) GetOwners() []BaseEntityOwner {
	if o == nil {
		return nil
	}
	return o.Owners
}

func (o *Price) GetSchema() Schema {
	if o == nil {
		return Schema("")
	}
	return o.Schema
}

func (o *Price) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}

func (o *Price) GetTitle() *string {
	if o == nil {
		return nil
	}
	return o.Title
}

func (o *Price) GetUpdatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}

func (o *Price) GetActive() bool {
	if o == nil {
		return false
	}
	return o.Active
}

func (o *Price) GetBillingDurationAmount() *float64 {
	if o == nil {
		return nil
	}
	return o.BillingDurationAmount
}

func (o *Price) GetBillingDurationUnit() *BillingDurationUnit {
	if o == nil {
		return nil
	}
	return o.BillingDurationUnit
}

func (o *Price) GetDescription() string {
	if o == nil {
		return ""
	}
	return o.Description
}

func (o *Price) GetIsCompositePrice() *bool {
	if o == nil {
		return nil
	}
	return o.IsCompositePrice
}

func (o *Price) GetIsTaxInclusive() *bool {
	if o == nil {
		return nil
	}
	return o.IsTaxInclusive
}

func (o *Price) GetLongDescription() *string {
	if o == nil {
		return nil
	}
	return o.LongDescription
}

func (o *Price) GetNoticeTimeAmount() *float64 {
	if o == nil {
		return nil
	}
	return o.NoticeTimeAmount
}

func (o *Price) GetNoticeTimeUnit() *NoticeTimeUnit {
	if o == nil {
		return nil
	}
	return o.NoticeTimeUnit
}

func (o *Price) GetPriceComponents() *PriceComponents {
	if o == nil {
		return nil
	}
	return o.PriceComponents
}

func (o *Price) GetPriceDisplayInJourneys() *PriceDisplayInJourneys {
	if o == nil {
		return nil
	}
	return o.PriceDisplayInJourneys
}

func (o *Price) GetPricingModel() *PricingModel {
	if o == nil {
		return nil
	}
	return o.PricingModel
}

func (o *Price) GetRenewalDurationAmount() *float64 {
	if o == nil {
		return nil
	}
	return o.RenewalDurationAmount
}

func (o *Price) GetRenewalDurationUnit() *RenewalDurationUnit {
	if o == nil {
		return nil
	}
	return o.RenewalDurationUnit
}

func (o *Price) GetTax() any {
	if o == nil {
		return nil
	}
	return o.Tax
}

func (o *Price) GetTerminationTimeAmount() *float64 {
	if o == nil {
		return nil
	}
	return o.TerminationTimeAmount
}

func (o *Price) GetTerminationTimeUnit() *TerminationTimeUnit {
	if o == nil {
		return nil
	}
	return o.TerminationTimeUnit
}

func (o *Price) GetTiers() []PriceTier {
	if o == nil {
		return nil
	}
	return o.Tiers
}

func (o *Price) GetType() *Type {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *Price) GetUnit() *string {
	if o == nil {
		return nil
	}
	return o.Unit
}

func (o *Price) GetUnitAmount() *float64 {
	if o == nil {
		return nil
	}
	return o.UnitAmount
}

func (o *Price) GetUnitAmountCurrency() *string {
	if o == nil {
		return nil
	}
	return o.UnitAmountCurrency
}

func (o *Price) GetUnitAmountDecimal() *string {
	if o == nil {
		return nil
	}
	return o.UnitAmountDecimal
}

func (o *Price) GetVariablePrice() *bool {
	if o == nil {
		return nil
	}
	return o.VariablePrice
}
