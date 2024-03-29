// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"github.com/epilot-dev/terraform-provider-epilot-product/internal/sdk/pkg/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *TaxResourceModel) ToSharedTaxCreate() *shared.TaxCreate {
	active := r.Active.ValueBool()
	description := new(string)
	if !r.Description.IsUnknown() && !r.Description.IsNull() {
		*description = r.Description.ValueString()
	} else {
		description = nil
	}
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

func (r *TaxResourceModel) RefreshFromSharedTax(resp *shared.Tax) {
	r.ID = types.StringValue(resp.ID)
	r.Active = types.BoolValue(resp.Active)
	r.Description = types.StringPointerValue(resp.Description)
	r.Rate = types.StringValue(resp.Rate)
	r.Region = types.StringValue(string(resp.Region))
	r.Type = types.StringValue(string(resp.Type))
}
