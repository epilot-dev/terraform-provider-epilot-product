// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"encoding/json"
	"github.com/epilot-dev/terraform-provider-epilot-product/internal/sdk/pkg/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"math/big"
	"time"
)

func (r *PriceDataSourceModel) RefreshFromGetResponse(resp *shared.Price) {
	if resp.ACL.AdditionalProperties == nil {
		r.ACL.AdditionalProperties = types.StringNull()
	} else {
		additionalPropertiesResult, _ := json.Marshal(resp.ACL.AdditionalProperties)
		r.ACL.AdditionalProperties = types.StringValue(string(additionalPropertiesResult))
	}
	r.ACL.Delete = nil
	for _, v := range resp.ACL.Delete {
		r.ACL.Delete = append(r.ACL.Delete, types.StringValue(v))
	}
	r.ACL.Edit = nil
	for _, v := range resp.ACL.Edit {
		r.ACL.Edit = append(r.ACL.Edit, types.StringValue(v))
	}
	r.ACL.View = nil
	for _, v := range resp.ACL.View {
		r.ACL.View = append(r.ACL.View, types.StringValue(v))
	}
	r.CreatedAt = types.StringValue(resp.CreatedAt.Format(time.RFC3339Nano))
	r.Org = types.StringValue(resp.Org)
	if len(r.Owners) > len(resp.Owners) {
		r.Owners = r.Owners[:len(resp.Owners)]
	}
	for ownersCount, ownersItem := range resp.Owners {
		var owners1 EntityOwner
		owners1.OrgID = types.StringValue(ownersItem.OrgID)
		if ownersItem.UserID != nil {
			owners1.UserID = types.StringValue(*ownersItem.UserID)
		} else {
			owners1.UserID = types.StringNull()
		}
		if ownersCount+1 > len(r.Owners) {
			r.Owners = append(r.Owners, owners1)
		} else {
			r.Owners[ownersCount].OrgID = owners1.OrgID
			r.Owners[ownersCount].UserID = owners1.UserID
		}
	}
	r.Schema = types.StringValue(resp.Schema)
	r.Tags = nil
	for _, v := range resp.Tags {
		r.Tags = append(r.Tags, types.StringValue(v))
	}
	r.Title = types.StringValue(resp.Title)
	r.UpdatedAt = types.StringValue(resp.UpdatedAt.Format(time.RFC3339Nano))
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
	r.ID = types.StringValue(resp.ID)
	if resp.IsCompositePrice != nil {
		r.IsCompositePrice = types.BoolValue(*resp.IsCompositePrice)
	} else {
		r.IsCompositePrice = types.BoolNull()
	}
	if resp.IsTaxInclusive != nil {
		r.IsTaxInclusive = types.BoolValue(*resp.IsTaxInclusive)
	} else {
		r.IsTaxInclusive = types.BoolNull()
	}
	if resp.LongDescription != nil {
		r.LongDescription = types.StringValue(*resp.LongDescription)
	} else {
		r.LongDescription = types.StringNull()
	}
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
		r.Tax = nil
	} else {
		r.Tax = &BaseRelation{}
		if len(r.Tax.DollarRelation) > len(resp.Tax.DollarRelation) {
			r.Tax.DollarRelation = r.Tax.DollarRelation[:len(resp.Tax.DollarRelation)]
		}
		for dollarRelationCount, dollarRelationItem := range resp.Tax.DollarRelation {
			var dollarRelation1 DollarRelation
			dollarRelation1.Tags = nil
			for _, v := range dollarRelationItem.Tags {
				dollarRelation1.Tags = append(dollarRelation1.Tags, types.StringValue(v))
			}
			if dollarRelationItem.EntityID != nil {
				dollarRelation1.EntityID = types.StringValue(*dollarRelationItem.EntityID)
			} else {
				dollarRelation1.EntityID = types.StringNull()
			}
			if dollarRelationCount+1 > len(r.Tax.DollarRelation) {
				r.Tax.DollarRelation = append(r.Tax.DollarRelation, dollarRelation1)
			} else {
				r.Tax.DollarRelation[dollarRelationCount].Tags = dollarRelation1.Tags
				r.Tax.DollarRelation[dollarRelationCount].EntityID = dollarRelation1.EntityID
			}
		}
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
	if len(r.Tiers) > len(resp.Tiers) {
		r.Tiers = r.Tiers[:len(resp.Tiers)]
	}
	for tiersCount, tiersItem := range resp.Tiers {
		var tiers1 PriceTier
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
		if tiersItem.FlatFeeAmountDecimal != nil {
			tiers1.FlatFeeAmountDecimal = types.StringValue(*tiersItem.FlatFeeAmountDecimal)
		} else {
			tiers1.FlatFeeAmountDecimal = types.StringNull()
		}
		if tiersItem.UnitAmount != nil {
			tiers1.UnitAmount = types.NumberValue(big.NewFloat(float64(*tiersItem.UnitAmount)))
		} else {
			tiers1.UnitAmount = types.NumberNull()
		}
		if tiersItem.UnitAmountDecimal != nil {
			tiers1.UnitAmountDecimal = types.StringValue(*tiersItem.UnitAmountDecimal)
		} else {
			tiers1.UnitAmountDecimal = types.StringNull()
		}
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
	if resp.Unit == nil {
		r.Unit = nil
	} else {
		r.Unit = &PriceCreateUnit{}
		if resp.Unit.One != nil {
			if resp.Unit.One != nil {
				r.Unit.One = types.StringValue(string(*resp.Unit.One))
			} else {
				r.Unit.One = types.StringNull()
			}
		}
		if resp.Unit.Str != nil {
			if resp.Unit.Str != nil {
				r.Unit.Str = types.StringValue(*resp.Unit.Str)
			} else {
				r.Unit.Str = types.StringNull()
			}
		}
	}
	if resp.UnitAmount != nil {
		r.UnitAmount = types.NumberValue(big.NewFloat(float64(*resp.UnitAmount)))
	} else {
		r.UnitAmount = types.NumberNull()
	}
	if resp.UnitAmountCurrency != nil {
		r.UnitAmountCurrency = types.StringValue(*resp.UnitAmountCurrency)
	} else {
		r.UnitAmountCurrency = types.StringNull()
	}
	if resp.UnitAmountDecimal != nil {
		r.UnitAmountDecimal = types.StringValue(*resp.UnitAmountDecimal)
	} else {
		r.UnitAmountDecimal = types.StringNull()
	}
	if resp.VariablePrice != nil {
		r.VariablePrice = types.BoolValue(*resp.VariablePrice)
	} else {
		r.VariablePrice = types.BoolNull()
	}
}
