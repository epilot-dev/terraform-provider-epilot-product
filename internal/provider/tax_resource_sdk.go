// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"encoding/json"
	"github.com/epilot-dev/terraform-provider-epilot-product/internal/sdk/pkg/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"time"
)

func (r *TaxResourceModel) ToCreateSDKType() *shared.TaxCreate {
	active := r.Active.ValueBool()
	description := r.Description.ValueString()
	rate := r.Rate.ValueString()
	region := shared.TaxCreateRegion(r.Region.ValueString())
	typeVar := shared.TaxCreateType(r.Type.ValueString())
	out := shared.TaxCreate{
		Active:      active,
		Description: description,
		Rate:        rate,
		Region:      region,
		Type:        typeVar,
	}
	return &out
}

func (r *TaxResourceModel) ToGetSDKType() *shared.TaxCreate {
	out := r.ToCreateSDKType()
	return out
}

func (r *TaxResourceModel) ToUpdateSDKType() *shared.TaxCreate {
	out := r.ToCreateSDKType()
	return out
}

func (r *TaxResourceModel) ToDeleteSDKType() *shared.TaxCreate {
	out := r.ToCreateSDKType()
	return out
}

func (r *TaxResourceModel) RefreshFromGetResponse(resp *shared.Tax) {
	if len(r.ACL) > len(resp.ACL) {
		r.ACL = r.ACL[:len(resp.ACL)]
	}
	for aclCount, aclItem := range resp.ACL {
		var acl1 EntityACL
		if aclItem.AdditionalProperties == nil {
			acl1.AdditionalProperties = types.StringNull()
		} else {
			additionalPropertiesResult, _ := json.Marshal(aclItem.AdditionalProperties)
			acl1.AdditionalProperties = types.StringValue(string(additionalPropertiesResult))
		}
		acl1.Delete = nil
		for _, v := range aclItem.Delete {
			acl1.Delete = append(acl1.Delete, types.StringValue(v))
		}
		acl1.Edit = nil
		for _, v := range aclItem.Edit {
			acl1.Edit = append(acl1.Edit, types.StringValue(v))
		}
		acl1.View = nil
		for _, v := range aclItem.View {
			acl1.View = append(acl1.View, types.StringValue(v))
		}
		if aclCount+1 > len(r.ACL) {
			r.ACL = append(r.ACL, acl1)
		} else {
			r.ACL[aclCount].AdditionalProperties = acl1.AdditionalProperties
			r.ACL[aclCount].Delete = acl1.Delete
			r.ACL[aclCount].Edit = acl1.Edit
			r.ACL[aclCount].View = acl1.View
		}
	}
	r.CreatedAt = types.StringValue(resp.CreatedAt.Format(time.RFC3339Nano))
	r.ID = types.StringValue(resp.ID)
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
	r.Description = types.StringValue(resp.Description)
	r.Rate = types.StringValue(resp.Rate)
	r.Region = types.StringValue(string(resp.Region))
	r.Type = types.StringValue(string(resp.Type))
}

func (r *TaxResourceModel) RefreshFromCreateResponse(resp *shared.Tax) {
	r.RefreshFromGetResponse(resp)
}

func (r *TaxResourceModel) RefreshFromUpdateResponse(resp *shared.Tax) {
	r.RefreshFromGetResponse(resp)
}
