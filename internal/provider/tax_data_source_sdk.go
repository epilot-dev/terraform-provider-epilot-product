// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"encoding/json"
	tfTypes "github.com/epilot-dev/terraform-provider-epilot-product/internal/provider/types"
	"github.com/epilot-dev/terraform-provider-epilot-product/internal/sdk/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"time"
)

func (r *TaxDataSourceModel) RefreshFromSharedTax(resp *shared.Tax) {
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
		r.Description = types.StringPointerValue(resp.Description)
		r.Rate = types.StringValue(resp.Rate)
		r.Region = types.StringValue(resp.Region)
		r.Type = types.StringValue(string(resp.Type))
	}
}
