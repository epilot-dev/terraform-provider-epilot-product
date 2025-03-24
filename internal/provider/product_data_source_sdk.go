// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"encoding/json"
	tfTypes "github.com/epilot-dev/terraform-provider-epilot-product/internal/provider/types"
	"github.com/epilot-dev/terraform-provider-epilot-product/internal/sdk/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"time"
)

func (r *ProductDataSourceModel) RefreshFromSharedProduct(resp *shared.Product) {
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
		if resp.AvailabilityFiles == nil {
			r.AvailabilityFiles = nil
		} else {
			r.AvailabilityFiles = &tfTypes.BaseRelation{}
			r.AvailabilityFiles.DollarRelation = []tfTypes.DollarRelation{}
			if len(r.AvailabilityFiles.DollarRelation) > len(resp.AvailabilityFiles.DollarRelation) {
				r.AvailabilityFiles.DollarRelation = r.AvailabilityFiles.DollarRelation[:len(resp.AvailabilityFiles.DollarRelation)]
			}
			for dollarRelationCount, dollarRelationItem := range resp.AvailabilityFiles.DollarRelation {
				var dollarRelation1 tfTypes.DollarRelation
				if dollarRelationItem.Tags != nil {
					dollarRelation1.Tags = make([]types.String, 0, len(dollarRelationItem.Tags))
					for _, v := range dollarRelationItem.Tags {
						dollarRelation1.Tags = append(dollarRelation1.Tags, types.StringValue(v))
					}
				}
				dollarRelation1.EntityID = types.StringPointerValue(dollarRelationItem.EntityID)
				if dollarRelationCount+1 > len(r.AvailabilityFiles.DollarRelation) {
					r.AvailabilityFiles.DollarRelation = append(r.AvailabilityFiles.DollarRelation, dollarRelation1)
				} else {
					r.AvailabilityFiles.DollarRelation[dollarRelationCount].Tags = dollarRelation1.Tags
					r.AvailabilityFiles.DollarRelation[dollarRelationCount].EntityID = dollarRelation1.EntityID
				}
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
			for dollarRelationCount1, dollarRelationItem1 := range resp.Files.DollarRelation {
				var dollarRelation3 tfTypes.DollarRelation
				if dollarRelationItem1.Tags != nil {
					dollarRelation3.Tags = make([]types.String, 0, len(dollarRelationItem1.Tags))
					for _, v := range dollarRelationItem1.Tags {
						dollarRelation3.Tags = append(dollarRelation3.Tags, types.StringValue(v))
					}
				}
				dollarRelation3.EntityID = types.StringPointerValue(dollarRelationItem1.EntityID)
				if dollarRelationCount1+1 > len(r.Files.DollarRelation) {
					r.Files.DollarRelation = append(r.Files.DollarRelation, dollarRelation3)
				} else {
					r.Files.DollarRelation[dollarRelationCount1].Tags = dollarRelation3.Tags
					r.Files.DollarRelation[dollarRelationCount1].EntityID = dollarRelation3.EntityID
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
		r.Categories = make([]types.String, 0, len(resp.Categories))
		for _, v := range resp.Categories {
			r.Categories = append(r.Categories, types.StringValue(v))
		}
		r.Code = types.StringPointerValue(resp.Code)
		r.Description = types.StringPointerValue(resp.Description)
		r.Feature = nil
		for _, featureItem := range resp.Feature {
			var feature1 types.String
			feature1Result, _ := json.Marshal(featureItem)
			feature1 = types.StringValue(string(feature1Result))
			r.Feature = append(r.Feature, feature1)
		}
		r.InternalName = types.StringPointerValue(resp.InternalName)
		r.Name = types.StringValue(resp.Name)
		if resp.PriceOptions == nil {
			r.PriceOptions = nil
		} else {
			r.PriceOptions = &tfTypes.BaseRelation{}
			r.PriceOptions.DollarRelation = []tfTypes.DollarRelation{}
			if len(r.PriceOptions.DollarRelation) > len(resp.PriceOptions.DollarRelation) {
				r.PriceOptions.DollarRelation = r.PriceOptions.DollarRelation[:len(resp.PriceOptions.DollarRelation)]
			}
			for dollarRelationCount2, dollarRelationItem2 := range resp.PriceOptions.DollarRelation {
				var dollarRelation5 tfTypes.DollarRelation
				if dollarRelationItem2.Tags != nil {
					dollarRelation5.Tags = make([]types.String, 0, len(dollarRelationItem2.Tags))
					for _, v := range dollarRelationItem2.Tags {
						dollarRelation5.Tags = append(dollarRelation5.Tags, types.StringValue(v))
					}
				}
				dollarRelation5.EntityID = types.StringPointerValue(dollarRelationItem2.EntityID)
				if dollarRelationCount2+1 > len(r.PriceOptions.DollarRelation) {
					r.PriceOptions.DollarRelation = append(r.PriceOptions.DollarRelation, dollarRelation5)
				} else {
					r.PriceOptions.DollarRelation[dollarRelationCount2].Tags = dollarRelation5.Tags
					r.PriceOptions.DollarRelation[dollarRelationCount2].EntityID = dollarRelation5.EntityID
				}
			}
		}
		if resp.ProductDownloads == nil {
			r.ProductDownloads = nil
		} else {
			r.ProductDownloads = &tfTypes.BaseRelation{}
			r.ProductDownloads.DollarRelation = []tfTypes.DollarRelation{}
			if len(r.ProductDownloads.DollarRelation) > len(resp.ProductDownloads.DollarRelation) {
				r.ProductDownloads.DollarRelation = r.ProductDownloads.DollarRelation[:len(resp.ProductDownloads.DollarRelation)]
			}
			for dollarRelationCount3, dollarRelationItem3 := range resp.ProductDownloads.DollarRelation {
				var dollarRelation7 tfTypes.DollarRelation
				if dollarRelationItem3.Tags != nil {
					dollarRelation7.Tags = make([]types.String, 0, len(dollarRelationItem3.Tags))
					for _, v := range dollarRelationItem3.Tags {
						dollarRelation7.Tags = append(dollarRelation7.Tags, types.StringValue(v))
					}
				}
				dollarRelation7.EntityID = types.StringPointerValue(dollarRelationItem3.EntityID)
				if dollarRelationCount3+1 > len(r.ProductDownloads.DollarRelation) {
					r.ProductDownloads.DollarRelation = append(r.ProductDownloads.DollarRelation, dollarRelation7)
				} else {
					r.ProductDownloads.DollarRelation[dollarRelationCount3].Tags = dollarRelation7.Tags
					r.ProductDownloads.DollarRelation[dollarRelationCount3].EntityID = dollarRelation7.EntityID
				}
			}
		}
		if resp.ProductImages == nil {
			r.ProductImages = nil
		} else {
			r.ProductImages = &tfTypes.BaseRelation{}
			r.ProductImages.DollarRelation = []tfTypes.DollarRelation{}
			if len(r.ProductImages.DollarRelation) > len(resp.ProductImages.DollarRelation) {
				r.ProductImages.DollarRelation = r.ProductImages.DollarRelation[:len(resp.ProductImages.DollarRelation)]
			}
			for dollarRelationCount4, dollarRelationItem4 := range resp.ProductImages.DollarRelation {
				var dollarRelation9 tfTypes.DollarRelation
				if dollarRelationItem4.Tags != nil {
					dollarRelation9.Tags = make([]types.String, 0, len(dollarRelationItem4.Tags))
					for _, v := range dollarRelationItem4.Tags {
						dollarRelation9.Tags = append(dollarRelation9.Tags, types.StringValue(v))
					}
				}
				dollarRelation9.EntityID = types.StringPointerValue(dollarRelationItem4.EntityID)
				if dollarRelationCount4+1 > len(r.ProductImages.DollarRelation) {
					r.ProductImages.DollarRelation = append(r.ProductImages.DollarRelation, dollarRelation9)
				} else {
					r.ProductImages.DollarRelation[dollarRelationCount4].Tags = dollarRelation9.Tags
					r.ProductImages.DollarRelation[dollarRelationCount4].EntityID = dollarRelation9.EntityID
				}
			}
		}
		if resp.Type != nil {
			r.Type = types.StringValue(string(*resp.Type))
		} else {
			r.Type = types.StringNull()
		}
	}
}
