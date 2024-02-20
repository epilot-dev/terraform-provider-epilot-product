// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"encoding/json"
	"github.com/epilot-dev/terraform-provider-epilot-product/internal/sdk/pkg/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"time"
)

func (r *ProductDataSourceModel) RefreshFromSharedProduct(resp *shared.Product) {
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
		var owners1 BaseEntityOwner
		owners1.OrgID = types.StringValue(ownersItem.OrgID)
		owners1.UserID = types.StringPointerValue(ownersItem.UserID)
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
	r.Code = types.StringPointerValue(resp.Code)
	r.Description = types.StringPointerValue(resp.Description)
	r.Feature = nil
	for _, featureItem := range resp.Feature {
		var feature1 types.String
		feature1Result, _ := json.Marshal(featureItem)
		feature1 = types.StringValue(string(feature1Result))
		r.Feature = append(r.Feature, feature1)
	}
	r.ID = types.StringValue(resp.ID)
	r.InternalName = types.StringPointerValue(resp.InternalName)
	r.Name = types.StringValue(resp.Name)
	if resp.PriceOptions == nil {
		r.PriceOptions = nil
	} else {
		r.PriceOptions = &BaseRelation{}
		if len(r.PriceOptions.DollarRelation) > len(resp.PriceOptions.DollarRelation) {
			r.PriceOptions.DollarRelation = r.PriceOptions.DollarRelation[:len(resp.PriceOptions.DollarRelation)]
		}
		for dollarRelationCount, dollarRelationItem := range resp.PriceOptions.DollarRelation {
			var dollarRelation1 DollarRelation
			dollarRelation1.Tags = nil
			for _, v := range dollarRelationItem.Tags {
				dollarRelation1.Tags = append(dollarRelation1.Tags, types.StringValue(v))
			}
			dollarRelation1.EntityID = types.StringPointerValue(dollarRelationItem.EntityID)
			if dollarRelationCount+1 > len(r.PriceOptions.DollarRelation) {
				r.PriceOptions.DollarRelation = append(r.PriceOptions.DollarRelation, dollarRelation1)
			} else {
				r.PriceOptions.DollarRelation[dollarRelationCount].Tags = dollarRelation1.Tags
				r.PriceOptions.DollarRelation[dollarRelationCount].EntityID = dollarRelation1.EntityID
			}
		}
	}
	if resp.ProductImages == nil {
		r.ProductImages = types.StringNull()
	} else {
		productImagesResult, _ := json.Marshal(resp.ProductImages)
		r.ProductImages = types.StringValue(string(productImagesResult))
	}
	if resp.Type != nil {
		r.Type = types.StringValue(string(*resp.Type))
	} else {
		r.Type = types.StringNull()
	}
}
