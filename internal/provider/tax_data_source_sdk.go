// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"github.com/epilot-dev/terraform-provider-epilot-product/internal/sdk/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *TaxDataSourceModel) RefreshFromSharedTax(resp *shared.Tax) {
	if resp != nil {
		r.ID = types.StringValue(resp.ID)
		r.Active = types.BoolValue(resp.Active)
		r.Description = types.StringPointerValue(resp.Description)
		r.Rate = types.StringValue(resp.Rate)
		r.Region = types.StringValue(string(resp.Region))
		r.Type = types.StringValue(string(resp.Type))
	}
}
