// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"encoding/json"
	tfTypes "github.com/epilot-dev/terraform-provider-epilot-product/internal/provider/types"
	"github.com/epilot-dev/terraform-provider-epilot-product/internal/sdk/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"math/big"
	"time"
)

func (r *PriceResourceModel) ToSharedPriceCreate() *shared.PriceCreate {
	additional := make(map[string]interface{})
	for additionalKey, additionalValue := range r.Additional {
		var additionalInst interface{}
		_ = json.Unmarshal([]byte(additionalValue.ValueString()), &additionalInst)
		additional[additionalKey] = additionalInst
	}
	var files *shared.BaseRelation
	if r.Files != nil {
		var dollarRelation []shared.DollarRelation = []shared.DollarRelation{}
		for _, dollarRelationItem := range r.Files.DollarRelation {
			var tags []string = []string{}
			for _, tagsItem := range dollarRelationItem.Tags {
				tags = append(tags, tagsItem.ValueString())
			}
			entityID := new(string)
			if !dollarRelationItem.EntityID.IsUnknown() && !dollarRelationItem.EntityID.IsNull() {
				*entityID = dollarRelationItem.EntityID.ValueString()
			} else {
				entityID = nil
			}
			dollarRelation = append(dollarRelation, shared.DollarRelation{
				Tags:     tags,
				EntityID: entityID,
			})
		}
		files = &shared.BaseRelation{
			DollarRelation: dollarRelation,
		}
	}
	var manifest []string = []string{}
	for _, manifestItem := range r.Manifest {
		manifest = append(manifest, manifestItem.ValueString())
	}
	var purpose []string = []string{}
	for _, purposeItem := range r.Purpose {
		purpose = append(purpose, purposeItem.ValueString())
	}
	schema := new(shared.PriceCreateSchema)
	if !r.Schema.IsUnknown() && !r.Schema.IsNull() {
		*schema = shared.PriceCreateSchema(r.Schema.ValueString())
	} else {
		schema = nil
	}
	var tags1 []string = []string{}
	for _, tagsItem1 := range r.Tags {
		tags1 = append(tags1, tagsItem1.ValueString())
	}
	var active bool
	active = r.Active.ValueBool()

	billingDurationAmount := new(float64)
	if !r.BillingDurationAmount.IsUnknown() && !r.BillingDurationAmount.IsNull() {
		*billingDurationAmount, _ = r.BillingDurationAmount.ValueBigFloat().Float64()
	} else {
		billingDurationAmount = nil
	}
	billingDurationUnit := new(shared.PriceCreateBillingDurationUnit)
	if !r.BillingDurationUnit.IsUnknown() && !r.BillingDurationUnit.IsNull() {
		*billingDurationUnit = shared.PriceCreateBillingDurationUnit(r.BillingDurationUnit.ValueString())
	} else {
		billingDurationUnit = nil
	}
	var description string
	description = r.Description.ValueString()

	isCompositePrice := new(bool)
	if !r.IsCompositePrice.IsUnknown() && !r.IsCompositePrice.IsNull() {
		*isCompositePrice = r.IsCompositePrice.ValueBool()
	} else {
		isCompositePrice = nil
	}
	isTaxInclusive := new(bool)
	if !r.IsTaxInclusive.IsUnknown() && !r.IsTaxInclusive.IsNull() {
		*isTaxInclusive = r.IsTaxInclusive.ValueBool()
	} else {
		isTaxInclusive = nil
	}
	longDescription := new(string)
	if !r.LongDescription.IsUnknown() && !r.LongDescription.IsNull() {
		*longDescription = r.LongDescription.ValueString()
	} else {
		longDescription = nil
	}
	noticeTimeAmount := new(float64)
	if !r.NoticeTimeAmount.IsUnknown() && !r.NoticeTimeAmount.IsNull() {
		*noticeTimeAmount, _ = r.NoticeTimeAmount.ValueBigFloat().Float64()
	} else {
		noticeTimeAmount = nil
	}
	noticeTimeUnit := new(shared.PriceCreateNoticeTimeUnit)
	if !r.NoticeTimeUnit.IsUnknown() && !r.NoticeTimeUnit.IsNull() {
		*noticeTimeUnit = shared.PriceCreateNoticeTimeUnit(r.NoticeTimeUnit.ValueString())
	} else {
		noticeTimeUnit = nil
	}
	var priceComponents *shared.PriceCreatePriceComponents
	if r.PriceComponents != nil {
		var dollarRelation1 []shared.PriceComponentRelation = []shared.PriceComponentRelation{}
		for _, dollarRelationItem1 := range r.PriceComponents.DollarRelation {
			var tags2 []string = []string{}
			for _, tagsItem2 := range dollarRelationItem1.Tags {
				tags2 = append(tags2, tagsItem2.ValueString())
			}
			entityId1 := new(string)
			if !dollarRelationItem1.EntityID.IsUnknown() && !dollarRelationItem1.EntityID.IsNull() {
				*entityId1 = dollarRelationItem1.EntityID.ValueString()
			} else {
				entityId1 = nil
			}
			dollarRelation1 = append(dollarRelation1, shared.PriceComponentRelation{
				Tags:     tags2,
				EntityID: entityId1,
			})
		}
		priceComponents = &shared.PriceCreatePriceComponents{
			DollarRelation: dollarRelation1,
		}
	}
	priceDisplayInJourneys := new(shared.PriceCreatePriceDisplayInJourneys)
	if !r.PriceDisplayInJourneys.IsUnknown() && !r.PriceDisplayInJourneys.IsNull() {
		*priceDisplayInJourneys = shared.PriceCreatePriceDisplayInJourneys(r.PriceDisplayInJourneys.ValueString())
	} else {
		priceDisplayInJourneys = nil
	}
	pricingModel := new(shared.PriceCreatePricingModel)
	if !r.PricingModel.IsUnknown() && !r.PricingModel.IsNull() {
		*pricingModel = shared.PriceCreatePricingModel(r.PricingModel.ValueString())
	} else {
		pricingModel = nil
	}
	renewalDurationAmount := new(float64)
	if !r.RenewalDurationAmount.IsUnknown() && !r.RenewalDurationAmount.IsNull() {
		*renewalDurationAmount, _ = r.RenewalDurationAmount.ValueBigFloat().Float64()
	} else {
		renewalDurationAmount = nil
	}
	renewalDurationUnit := new(shared.PriceCreateRenewalDurationUnit)
	if !r.RenewalDurationUnit.IsUnknown() && !r.RenewalDurationUnit.IsNull() {
		*renewalDurationUnit = shared.PriceCreateRenewalDurationUnit(r.RenewalDurationUnit.ValueString())
	} else {
		renewalDurationUnit = nil
	}
	var tax interface{}
	if !r.Tax.IsUnknown() && !r.Tax.IsNull() {
		_ = json.Unmarshal([]byte(r.Tax.ValueString()), &tax)
	}
	terminationTimeAmount := new(float64)
	if !r.TerminationTimeAmount.IsUnknown() && !r.TerminationTimeAmount.IsNull() {
		*terminationTimeAmount, _ = r.TerminationTimeAmount.ValueBigFloat().Float64()
	} else {
		terminationTimeAmount = nil
	}
	terminationTimeUnit := new(shared.PriceCreateTerminationTimeUnit)
	if !r.TerminationTimeUnit.IsUnknown() && !r.TerminationTimeUnit.IsNull() {
		*terminationTimeUnit = shared.PriceCreateTerminationTimeUnit(r.TerminationTimeUnit.ValueString())
	} else {
		terminationTimeUnit = nil
	}
	var tiers []shared.PriceTier = []shared.PriceTier{}
	for _, tiersItem := range r.Tiers {
		displayMode := new(shared.PriceTierDisplayMode)
		if !tiersItem.DisplayMode.IsUnknown() && !tiersItem.DisplayMode.IsNull() {
			*displayMode = shared.PriceTierDisplayMode(tiersItem.DisplayMode.ValueString())
		} else {
			displayMode = nil
		}
		flatFeeAmount := new(float64)
		if !tiersItem.FlatFeeAmount.IsUnknown() && !tiersItem.FlatFeeAmount.IsNull() {
			*flatFeeAmount, _ = tiersItem.FlatFeeAmount.ValueBigFloat().Float64()
		} else {
			flatFeeAmount = nil
		}
		flatFeeAmountDecimal := new(string)
		if !tiersItem.FlatFeeAmountDecimal.IsUnknown() && !tiersItem.FlatFeeAmountDecimal.IsNull() {
			*flatFeeAmountDecimal = tiersItem.FlatFeeAmountDecimal.ValueString()
		} else {
			flatFeeAmountDecimal = nil
		}
		unitAmount := new(float64)
		if !tiersItem.UnitAmount.IsUnknown() && !tiersItem.UnitAmount.IsNull() {
			*unitAmount, _ = tiersItem.UnitAmount.ValueBigFloat().Float64()
		} else {
			unitAmount = nil
		}
		unitAmountDecimal := new(string)
		if !tiersItem.UnitAmountDecimal.IsUnknown() && !tiersItem.UnitAmountDecimal.IsNull() {
			*unitAmountDecimal = tiersItem.UnitAmountDecimal.ValueString()
		} else {
			unitAmountDecimal = nil
		}
		upTo := new(float64)
		if !tiersItem.UpTo.IsUnknown() && !tiersItem.UpTo.IsNull() {
			*upTo, _ = tiersItem.UpTo.ValueBigFloat().Float64()
		} else {
			upTo = nil
		}
		tiers = append(tiers, shared.PriceTier{
			DisplayMode:          displayMode,
			FlatFeeAmount:        flatFeeAmount,
			FlatFeeAmountDecimal: flatFeeAmountDecimal,
			UnitAmount:           unitAmount,
			UnitAmountDecimal:    unitAmountDecimal,
			UpTo:                 upTo,
		})
	}
	typeVar := new(shared.PriceCreateType)
	if !r.Type.IsUnknown() && !r.Type.IsNull() {
		*typeVar = shared.PriceCreateType(r.Type.ValueString())
	} else {
		typeVar = nil
	}
	unit := new(string)
	if !r.Unit.IsUnknown() && !r.Unit.IsNull() {
		*unit = r.Unit.ValueString()
	} else {
		unit = nil
	}
	unitAmount1 := new(float64)
	if !r.UnitAmount.IsUnknown() && !r.UnitAmount.IsNull() {
		*unitAmount1, _ = r.UnitAmount.ValueBigFloat().Float64()
	} else {
		unitAmount1 = nil
	}
	unitAmountCurrency := new(string)
	if !r.UnitAmountCurrency.IsUnknown() && !r.UnitAmountCurrency.IsNull() {
		*unitAmountCurrency = r.UnitAmountCurrency.ValueString()
	} else {
		unitAmountCurrency = nil
	}
	unitAmountDecimal1 := new(string)
	if !r.UnitAmountDecimal.IsUnknown() && !r.UnitAmountDecimal.IsNull() {
		*unitAmountDecimal1 = r.UnitAmountDecimal.ValueString()
	} else {
		unitAmountDecimal1 = nil
	}
	variablePrice := new(bool)
	if !r.VariablePrice.IsUnknown() && !r.VariablePrice.IsNull() {
		*variablePrice = r.VariablePrice.ValueBool()
	} else {
		variablePrice = nil
	}
	out := shared.PriceCreate{
		Additional:             additional,
		Files:                  files,
		Manifest:               manifest,
		Purpose:                purpose,
		Schema:                 schema,
		Tags:                   tags1,
		Active:                 active,
		BillingDurationAmount:  billingDurationAmount,
		BillingDurationUnit:    billingDurationUnit,
		Description:            description,
		IsCompositePrice:       isCompositePrice,
		IsTaxInclusive:         isTaxInclusive,
		LongDescription:        longDescription,
		NoticeTimeAmount:       noticeTimeAmount,
		NoticeTimeUnit:         noticeTimeUnit,
		PriceComponents:        priceComponents,
		PriceDisplayInJourneys: priceDisplayInJourneys,
		PricingModel:           pricingModel,
		RenewalDurationAmount:  renewalDurationAmount,
		RenewalDurationUnit:    renewalDurationUnit,
		Tax:                    tax,
		TerminationTimeAmount:  terminationTimeAmount,
		TerminationTimeUnit:    terminationTimeUnit,
		Tiers:                  tiers,
		Type:                   typeVar,
		Unit:                   unit,
		UnitAmount:             unitAmount1,
		UnitAmountCurrency:     unitAmountCurrency,
		UnitAmountDecimal:      unitAmountDecimal1,
		VariablePrice:          variablePrice,
	}
	return &out
}

func (r *PriceResourceModel) RefreshFromSharedPrice(resp *shared.Price) {
	if resp != nil {
		if resp.Additional != nil {
			r.Additional = make(map[string]types.String, len(resp.Additional))
			for key, value := range resp.Additional {
				result, _ := json.Marshal(value)
				r.Additional[key] = types.StringValue(string(result))
			}
		}
		if resp.ACL == nil {
			r.ACL = nil
		} else {
			r.ACL = &tfTypes.BaseEntityACL{}
			r.ACL.Delete = make([]types.String, 0, len(resp.ACL.Delete))
			for _, v := range resp.ACL.Delete {
				r.ACL.Delete = append(r.ACL.Delete, types.StringValue(v))
			}
			r.ACL.Edit = make([]types.String, 0, len(resp.ACL.Edit))
			for _, v := range resp.ACL.Edit {
				r.ACL.Edit = append(r.ACL.Edit, types.StringValue(v))
			}
			r.ACL.View = make([]types.String, 0, len(resp.ACL.View))
			for _, v := range resp.ACL.View {
				r.ACL.View = append(r.ACL.View, types.StringValue(v))
			}
		}
		if resp.CreatedAt != nil {
			r.CreatedAt = types.StringValue(resp.CreatedAt.Format(time.RFC3339Nano))
		} else {
			r.CreatedAt = types.StringNull()
		}
		if resp.Files == nil {
			r.Files = nil
		} else {
			r.Files = &tfTypes.BaseRelation{}
			r.Files.DollarRelation = []tfTypes.DollarRelation{}
			if len(r.Files.DollarRelation) > len(resp.Files.DollarRelation) {
				r.Files.DollarRelation = r.Files.DollarRelation[:len(resp.Files.DollarRelation)]
			}
			for dollarRelationCount, dollarRelationItem := range resp.Files.DollarRelation {
				var dollarRelation1 tfTypes.DollarRelation
				if dollarRelationItem.Tags != nil {
					dollarRelation1.Tags = make([]types.String, 0, len(dollarRelationItem.Tags))
					for _, v := range dollarRelationItem.Tags {
						dollarRelation1.Tags = append(dollarRelation1.Tags, types.StringValue(v))
					}
				}
				dollarRelation1.EntityID = types.StringPointerValue(dollarRelationItem.EntityID)
				if dollarRelationCount+1 > len(r.Files.DollarRelation) {
					r.Files.DollarRelation = append(r.Files.DollarRelation, dollarRelation1)
				} else {
					r.Files.DollarRelation[dollarRelationCount].Tags = dollarRelation1.Tags
					r.Files.DollarRelation[dollarRelationCount].EntityID = dollarRelation1.EntityID
				}
			}
		}
		r.ID = types.StringPointerValue(resp.ID)
		r.Manifest = make([]types.String, 0, len(resp.Manifest))
		for _, v := range resp.Manifest {
			r.Manifest = append(r.Manifest, types.StringValue(v))
		}
		r.Org = types.StringValue(resp.Org)
		r.Owners = []tfTypes.BaseEntityOwner{}
		if len(r.Owners) > len(resp.Owners) {
			r.Owners = r.Owners[:len(resp.Owners)]
		}
		for ownersCount, ownersItem := range resp.Owners {
			var owners1 tfTypes.BaseEntityOwner
			owners1.OrgID = types.StringValue(ownersItem.OrgID)
			owners1.UserID = types.StringPointerValue(ownersItem.UserID)
			if ownersCount+1 > len(r.Owners) {
				r.Owners = append(r.Owners, owners1)
			} else {
				r.Owners[ownersCount].OrgID = owners1.OrgID
				r.Owners[ownersCount].UserID = owners1.UserID
			}
		}
		if resp.Purpose != nil {
			r.Purpose = make([]types.String, 0, len(resp.Purpose))
			for _, v := range resp.Purpose {
				r.Purpose = append(r.Purpose, types.StringValue(v))
			}
		}
		r.Schema = types.StringValue(string(resp.Schema))
		if resp.Tags != nil {
			r.Tags = make([]types.String, 0, len(resp.Tags))
			for _, v := range resp.Tags {
				r.Tags = append(r.Tags, types.StringValue(v))
			}
		}
		r.Title = types.StringPointerValue(resp.Title)
		if resp.UpdatedAt != nil {
			r.UpdatedAt = types.StringValue(resp.UpdatedAt.Format(time.RFC3339Nano))
		} else {
			r.UpdatedAt = types.StringNull()
		}
		r.Active = types.BoolValue(resp.Active)
		if resp.BillingDurationAmount != nil {
			r.BillingDurationAmount = types.NumberValue(big.NewFloat(float64(*resp.BillingDurationAmount)))
		} else {
			r.BillingDurationAmount = types.NumberNull()
		}
		if resp.BillingDurationUnit != nil {
			r.BillingDurationUnit = types.StringValue(string(*resp.BillingDurationUnit))
		} else {
			r.BillingDurationUnit = types.StringNull()
		}
		r.Description = types.StringValue(resp.Description)
		r.IsCompositePrice = types.BoolPointerValue(resp.IsCompositePrice)
		r.IsTaxInclusive = types.BoolPointerValue(resp.IsTaxInclusive)
		r.LongDescription = types.StringPointerValue(resp.LongDescription)
		if resp.NoticeTimeAmount != nil {
			r.NoticeTimeAmount = types.NumberValue(big.NewFloat(float64(*resp.NoticeTimeAmount)))
		} else {
			r.NoticeTimeAmount = types.NumberNull()
		}
		if resp.NoticeTimeUnit != nil {
			r.NoticeTimeUnit = types.StringValue(string(*resp.NoticeTimeUnit))
		} else {
			r.NoticeTimeUnit = types.StringNull()
		}
		if resp.PriceComponents == nil {
			r.PriceComponents = nil
		} else {
			r.PriceComponents = &tfTypes.PriceCreatePriceComponents{}
			r.PriceComponents.DollarRelation = []tfTypes.PriceComponentRelation{}
			if len(r.PriceComponents.DollarRelation) > len(resp.PriceComponents.DollarRelation) {
				r.PriceComponents.DollarRelation = r.PriceComponents.DollarRelation[:len(resp.PriceComponents.DollarRelation)]
			}
			for dollarRelationCount1, dollarRelationItem1 := range resp.PriceComponents.DollarRelation {
				var dollarRelation3 tfTypes.PriceComponentRelation
				dollarRelation3.Tags = make([]types.String, 0, len(dollarRelationItem1.Tags))
				for _, v := range dollarRelationItem1.Tags {
					dollarRelation3.Tags = append(dollarRelation3.Tags, types.StringValue(v))
				}
				dollarRelation3.EntityID = types.StringPointerValue(dollarRelationItem1.EntityID)
				if dollarRelationCount1+1 > len(r.PriceComponents.DollarRelation) {
					r.PriceComponents.DollarRelation = append(r.PriceComponents.DollarRelation, dollarRelation3)
				} else {
					r.PriceComponents.DollarRelation[dollarRelationCount1].Tags = dollarRelation3.Tags
					r.PriceComponents.DollarRelation[dollarRelationCount1].EntityID = dollarRelation3.EntityID
				}
			}
		}
		if resp.PriceDisplayInJourneys != nil {
			r.PriceDisplayInJourneys = types.StringValue(string(*resp.PriceDisplayInJourneys))
		} else {
			r.PriceDisplayInJourneys = types.StringNull()
		}
		if resp.PricingModel != nil {
			r.PricingModel = types.StringValue(string(*resp.PricingModel))
		} else {
			r.PricingModel = types.StringNull()
		}
		if resp.RenewalDurationAmount != nil {
			r.RenewalDurationAmount = types.NumberValue(big.NewFloat(float64(*resp.RenewalDurationAmount)))
		} else {
			r.RenewalDurationAmount = types.NumberNull()
		}
		if resp.RenewalDurationUnit != nil {
			r.RenewalDurationUnit = types.StringValue(string(*resp.RenewalDurationUnit))
		} else {
			r.RenewalDurationUnit = types.StringNull()
		}
		if resp.Tax == nil {
			r.Tax = types.StringNull()
		} else {
			taxResult, _ := json.Marshal(resp.Tax)
			r.Tax = types.StringValue(string(taxResult))
		}
		if resp.TerminationTimeAmount != nil {
			r.TerminationTimeAmount = types.NumberValue(big.NewFloat(float64(*resp.TerminationTimeAmount)))
		} else {
			r.TerminationTimeAmount = types.NumberNull()
		}
		if resp.TerminationTimeUnit != nil {
			r.TerminationTimeUnit = types.StringValue(string(*resp.TerminationTimeUnit))
		} else {
			r.TerminationTimeUnit = types.StringNull()
		}
		r.Tiers = []tfTypes.PriceTier{}
		if len(r.Tiers) > len(resp.Tiers) {
			r.Tiers = r.Tiers[:len(resp.Tiers)]
		}
		for tiersCount, tiersItem := range resp.Tiers {
			var tiers1 tfTypes.PriceTier
			if tiersItem.DisplayMode != nil {
				tiers1.DisplayMode = types.StringValue(string(*tiersItem.DisplayMode))
			} else {
				tiers1.DisplayMode = types.StringNull()
			}
			if tiersItem.FlatFeeAmount != nil {
				tiers1.FlatFeeAmount = types.NumberValue(big.NewFloat(float64(*tiersItem.FlatFeeAmount)))
			} else {
				tiers1.FlatFeeAmount = types.NumberNull()
			}
			tiers1.FlatFeeAmountDecimal = types.StringPointerValue(tiersItem.FlatFeeAmountDecimal)
			if tiersItem.UnitAmount != nil {
				tiers1.UnitAmount = types.NumberValue(big.NewFloat(float64(*tiersItem.UnitAmount)))
			} else {
				tiers1.UnitAmount = types.NumberNull()
			}
			tiers1.UnitAmountDecimal = types.StringPointerValue(tiersItem.UnitAmountDecimal)
			if tiersItem.UpTo != nil {
				tiers1.UpTo = types.NumberValue(big.NewFloat(float64(*tiersItem.UpTo)))
			} else {
				tiers1.UpTo = types.NumberNull()
			}
			if tiersCount+1 > len(r.Tiers) {
				r.Tiers = append(r.Tiers, tiers1)
			} else {
				r.Tiers[tiersCount].DisplayMode = tiers1.DisplayMode
				r.Tiers[tiersCount].FlatFeeAmount = tiers1.FlatFeeAmount
				r.Tiers[tiersCount].FlatFeeAmountDecimal = tiers1.FlatFeeAmountDecimal
				r.Tiers[tiersCount].UnitAmount = tiers1.UnitAmount
				r.Tiers[tiersCount].UnitAmountDecimal = tiers1.UnitAmountDecimal
				r.Tiers[tiersCount].UpTo = tiers1.UpTo
			}
		}
		if resp.Type != nil {
			r.Type = types.StringValue(string(*resp.Type))
		} else {
			r.Type = types.StringNull()
		}
		r.Unit = types.StringPointerValue(resp.Unit)
		if resp.UnitAmount != nil {
			r.UnitAmount = types.NumberValue(big.NewFloat(float64(*resp.UnitAmount)))
		} else {
			r.UnitAmount = types.NumberNull()
		}
		r.UnitAmountCurrency = types.StringPointerValue(resp.UnitAmountCurrency)
		r.UnitAmountDecimal = types.StringPointerValue(resp.UnitAmountDecimal)
		r.VariablePrice = types.BoolPointerValue(resp.VariablePrice)
	}
}

func (r *PriceResourceModel) ToSharedPricePatch() *shared.PricePatch {
	additional := make(map[string]interface{})
	for additionalKey, additionalValue := range r.Additional {
		var additionalInst interface{}
		_ = json.Unmarshal([]byte(additionalValue.ValueString()), &additionalInst)
		additional[additionalKey] = additionalInst
	}
	var files *shared.BaseRelation
	if r.Files != nil {
		var dollarRelation []shared.DollarRelation = []shared.DollarRelation{}
		for _, dollarRelationItem := range r.Files.DollarRelation {
			var tags []string = []string{}
			for _, tagsItem := range dollarRelationItem.Tags {
				tags = append(tags, tagsItem.ValueString())
			}
			entityID := new(string)
			if !dollarRelationItem.EntityID.IsUnknown() && !dollarRelationItem.EntityID.IsNull() {
				*entityID = dollarRelationItem.EntityID.ValueString()
			} else {
				entityID = nil
			}
			dollarRelation = append(dollarRelation, shared.DollarRelation{
				Tags:     tags,
				EntityID: entityID,
			})
		}
		files = &shared.BaseRelation{
			DollarRelation: dollarRelation,
		}
	}
	var manifest []string = []string{}
	for _, manifestItem := range r.Manifest {
		manifest = append(manifest, manifestItem.ValueString())
	}
	var purpose []string = []string{}
	for _, purposeItem := range r.Purpose {
		purpose = append(purpose, purposeItem.ValueString())
	}
	schema := new(shared.PricePatchSchema)
	if !r.Schema.IsUnknown() && !r.Schema.IsNull() {
		*schema = shared.PricePatchSchema(r.Schema.ValueString())
	} else {
		schema = nil
	}
	var tags1 []string = []string{}
	for _, tagsItem1 := range r.Tags {
		tags1 = append(tags1, tagsItem1.ValueString())
	}
	active := new(bool)
	if !r.Active.IsUnknown() && !r.Active.IsNull() {
		*active = r.Active.ValueBool()
	} else {
		active = nil
	}
	billingDurationAmount := new(float64)
	if !r.BillingDurationAmount.IsUnknown() && !r.BillingDurationAmount.IsNull() {
		*billingDurationAmount, _ = r.BillingDurationAmount.ValueBigFloat().Float64()
	} else {
		billingDurationAmount = nil
	}
	billingDurationUnit := new(shared.PricePatchBillingDurationUnit)
	if !r.BillingDurationUnit.IsUnknown() && !r.BillingDurationUnit.IsNull() {
		*billingDurationUnit = shared.PricePatchBillingDurationUnit(r.BillingDurationUnit.ValueString())
	} else {
		billingDurationUnit = nil
	}
	description := new(string)
	if !r.Description.IsUnknown() && !r.Description.IsNull() {
		*description = r.Description.ValueString()
	} else {
		description = nil
	}
	isCompositePrice := new(bool)
	if !r.IsCompositePrice.IsUnknown() && !r.IsCompositePrice.IsNull() {
		*isCompositePrice = r.IsCompositePrice.ValueBool()
	} else {
		isCompositePrice = nil
	}
	isTaxInclusive := new(bool)
	if !r.IsTaxInclusive.IsUnknown() && !r.IsTaxInclusive.IsNull() {
		*isTaxInclusive = r.IsTaxInclusive.ValueBool()
	} else {
		isTaxInclusive = nil
	}
	longDescription := new(string)
	if !r.LongDescription.IsUnknown() && !r.LongDescription.IsNull() {
		*longDescription = r.LongDescription.ValueString()
	} else {
		longDescription = nil
	}
	noticeTimeAmount := new(float64)
	if !r.NoticeTimeAmount.IsUnknown() && !r.NoticeTimeAmount.IsNull() {
		*noticeTimeAmount, _ = r.NoticeTimeAmount.ValueBigFloat().Float64()
	} else {
		noticeTimeAmount = nil
	}
	noticeTimeUnit := new(shared.PricePatchNoticeTimeUnit)
	if !r.NoticeTimeUnit.IsUnknown() && !r.NoticeTimeUnit.IsNull() {
		*noticeTimeUnit = shared.PricePatchNoticeTimeUnit(r.NoticeTimeUnit.ValueString())
	} else {
		noticeTimeUnit = nil
	}
	var priceComponents *shared.PricePatchPriceComponents
	if r.PriceComponents != nil {
		var dollarRelation1 []shared.PriceComponentRelation = []shared.PriceComponentRelation{}
		for _, dollarRelationItem1 := range r.PriceComponents.DollarRelation {
			var tags2 []string = []string{}
			for _, tagsItem2 := range dollarRelationItem1.Tags {
				tags2 = append(tags2, tagsItem2.ValueString())
			}
			entityId1 := new(string)
			if !dollarRelationItem1.EntityID.IsUnknown() && !dollarRelationItem1.EntityID.IsNull() {
				*entityId1 = dollarRelationItem1.EntityID.ValueString()
			} else {
				entityId1 = nil
			}
			dollarRelation1 = append(dollarRelation1, shared.PriceComponentRelation{
				Tags:     tags2,
				EntityID: entityId1,
			})
		}
		priceComponents = &shared.PricePatchPriceComponents{
			DollarRelation: dollarRelation1,
		}
	}
	priceDisplayInJourneys := new(shared.PricePatchPriceDisplayInJourneys)
	if !r.PriceDisplayInJourneys.IsUnknown() && !r.PriceDisplayInJourneys.IsNull() {
		*priceDisplayInJourneys = shared.PricePatchPriceDisplayInJourneys(r.PriceDisplayInJourneys.ValueString())
	} else {
		priceDisplayInJourneys = nil
	}
	pricingModel := new(shared.PricePatchPricingModel)
	if !r.PricingModel.IsUnknown() && !r.PricingModel.IsNull() {
		*pricingModel = shared.PricePatchPricingModel(r.PricingModel.ValueString())
	} else {
		pricingModel = nil
	}
	renewalDurationAmount := new(float64)
	if !r.RenewalDurationAmount.IsUnknown() && !r.RenewalDurationAmount.IsNull() {
		*renewalDurationAmount, _ = r.RenewalDurationAmount.ValueBigFloat().Float64()
	} else {
		renewalDurationAmount = nil
	}
	renewalDurationUnit := new(shared.PricePatchRenewalDurationUnit)
	if !r.RenewalDurationUnit.IsUnknown() && !r.RenewalDurationUnit.IsNull() {
		*renewalDurationUnit = shared.PricePatchRenewalDurationUnit(r.RenewalDurationUnit.ValueString())
	} else {
		renewalDurationUnit = nil
	}
	var tax interface{}
	if !r.Tax.IsUnknown() && !r.Tax.IsNull() {
		_ = json.Unmarshal([]byte(r.Tax.ValueString()), &tax)
	}
	terminationTimeAmount := new(float64)
	if !r.TerminationTimeAmount.IsUnknown() && !r.TerminationTimeAmount.IsNull() {
		*terminationTimeAmount, _ = r.TerminationTimeAmount.ValueBigFloat().Float64()
	} else {
		terminationTimeAmount = nil
	}
	terminationTimeUnit := new(shared.PricePatchTerminationTimeUnit)
	if !r.TerminationTimeUnit.IsUnknown() && !r.TerminationTimeUnit.IsNull() {
		*terminationTimeUnit = shared.PricePatchTerminationTimeUnit(r.TerminationTimeUnit.ValueString())
	} else {
		terminationTimeUnit = nil
	}
	var tiers []shared.PriceTier = []shared.PriceTier{}
	for _, tiersItem := range r.Tiers {
		displayMode := new(shared.PriceTierDisplayMode)
		if !tiersItem.DisplayMode.IsUnknown() && !tiersItem.DisplayMode.IsNull() {
			*displayMode = shared.PriceTierDisplayMode(tiersItem.DisplayMode.ValueString())
		} else {
			displayMode = nil
		}
		flatFeeAmount := new(float64)
		if !tiersItem.FlatFeeAmount.IsUnknown() && !tiersItem.FlatFeeAmount.IsNull() {
			*flatFeeAmount, _ = tiersItem.FlatFeeAmount.ValueBigFloat().Float64()
		} else {
			flatFeeAmount = nil
		}
		flatFeeAmountDecimal := new(string)
		if !tiersItem.FlatFeeAmountDecimal.IsUnknown() && !tiersItem.FlatFeeAmountDecimal.IsNull() {
			*flatFeeAmountDecimal = tiersItem.FlatFeeAmountDecimal.ValueString()
		} else {
			flatFeeAmountDecimal = nil
		}
		unitAmount := new(float64)
		if !tiersItem.UnitAmount.IsUnknown() && !tiersItem.UnitAmount.IsNull() {
			*unitAmount, _ = tiersItem.UnitAmount.ValueBigFloat().Float64()
		} else {
			unitAmount = nil
		}
		unitAmountDecimal := new(string)
		if !tiersItem.UnitAmountDecimal.IsUnknown() && !tiersItem.UnitAmountDecimal.IsNull() {
			*unitAmountDecimal = tiersItem.UnitAmountDecimal.ValueString()
		} else {
			unitAmountDecimal = nil
		}
		upTo := new(float64)
		if !tiersItem.UpTo.IsUnknown() && !tiersItem.UpTo.IsNull() {
			*upTo, _ = tiersItem.UpTo.ValueBigFloat().Float64()
		} else {
			upTo = nil
		}
		tiers = append(tiers, shared.PriceTier{
			DisplayMode:          displayMode,
			FlatFeeAmount:        flatFeeAmount,
			FlatFeeAmountDecimal: flatFeeAmountDecimal,
			UnitAmount:           unitAmount,
			UnitAmountDecimal:    unitAmountDecimal,
			UpTo:                 upTo,
		})
	}
	typeVar := new(shared.PricePatchType)
	if !r.Type.IsUnknown() && !r.Type.IsNull() {
		*typeVar = shared.PricePatchType(r.Type.ValueString())
	} else {
		typeVar = nil
	}
	unit := new(string)
	if !r.Unit.IsUnknown() && !r.Unit.IsNull() {
		*unit = r.Unit.ValueString()
	} else {
		unit = nil
	}
	unitAmount1 := new(float64)
	if !r.UnitAmount.IsUnknown() && !r.UnitAmount.IsNull() {
		*unitAmount1, _ = r.UnitAmount.ValueBigFloat().Float64()
	} else {
		unitAmount1 = nil
	}
	unitAmountCurrency := new(string)
	if !r.UnitAmountCurrency.IsUnknown() && !r.UnitAmountCurrency.IsNull() {
		*unitAmountCurrency = r.UnitAmountCurrency.ValueString()
	} else {
		unitAmountCurrency = nil
	}
	unitAmountDecimal1 := new(string)
	if !r.UnitAmountDecimal.IsUnknown() && !r.UnitAmountDecimal.IsNull() {
		*unitAmountDecimal1 = r.UnitAmountDecimal.ValueString()
	} else {
		unitAmountDecimal1 = nil
	}
	variablePrice := new(bool)
	if !r.VariablePrice.IsUnknown() && !r.VariablePrice.IsNull() {
		*variablePrice = r.VariablePrice.ValueBool()
	} else {
		variablePrice = nil
	}
	out := shared.PricePatch{
		Additional:             additional,
		Files:                  files,
		Manifest:               manifest,
		Purpose:                purpose,
		Schema:                 schema,
		Tags:                   tags1,
		Active:                 active,
		BillingDurationAmount:  billingDurationAmount,
		BillingDurationUnit:    billingDurationUnit,
		Description:            description,
		IsCompositePrice:       isCompositePrice,
		IsTaxInclusive:         isTaxInclusive,
		LongDescription:        longDescription,
		NoticeTimeAmount:       noticeTimeAmount,
		NoticeTimeUnit:         noticeTimeUnit,
		PriceComponents:        priceComponents,
		PriceDisplayInJourneys: priceDisplayInJourneys,
		PricingModel:           pricingModel,
		RenewalDurationAmount:  renewalDurationAmount,
		RenewalDurationUnit:    renewalDurationUnit,
		Tax:                    tax,
		TerminationTimeAmount:  terminationTimeAmount,
		TerminationTimeUnit:    terminationTimeUnit,
		Tiers:                  tiers,
		Type:                   typeVar,
		Unit:                   unit,
		UnitAmount:             unitAmount1,
		UnitAmountCurrency:     unitAmountCurrency,
		UnitAmountDecimal:      unitAmountDecimal1,
		VariablePrice:          variablePrice,
	}
	return &out
}
