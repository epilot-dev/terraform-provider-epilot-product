// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	speakeasy_boolplanmodifier "github.com/epilot-dev/terraform-provider-epilot-product/internal/planmodifiers/boolplanmodifier"
	speakeasy_listplanmodifier "github.com/epilot-dev/terraform-provider-epilot-product/internal/planmodifiers/listplanmodifier"
	speakeasy_mapplanmodifier "github.com/epilot-dev/terraform-provider-epilot-product/internal/planmodifiers/mapplanmodifier"
	speakeasy_numberplanmodifier "github.com/epilot-dev/terraform-provider-epilot-product/internal/planmodifiers/numberplanmodifier"
	speakeasy_objectplanmodifier "github.com/epilot-dev/terraform-provider-epilot-product/internal/planmodifiers/objectplanmodifier"
	speakeasy_stringplanmodifier "github.com/epilot-dev/terraform-provider-epilot-product/internal/planmodifiers/stringplanmodifier"
	tfTypes "github.com/epilot-dev/terraform-provider-epilot-product/internal/provider/types"
	"github.com/epilot-dev/terraform-provider-epilot-product/internal/sdk"
	"github.com/epilot-dev/terraform-provider-epilot-product/internal/sdk/models/operations"
	"github.com/epilot-dev/terraform-provider-epilot-product/internal/validators"
	speakeasy_objectvalidators "github.com/epilot-dev/terraform-provider-epilot-product/internal/validators/objectvalidators"
	speakeasy_stringvalidators "github.com/epilot-dev/terraform-provider-epilot-product/internal/validators/stringvalidators"
	"github.com/hashicorp/terraform-plugin-framework-validators/mapvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &CouponResource{}
var _ resource.ResourceWithImportState = &CouponResource{}

func NewCouponResource() resource.Resource {
	return &CouponResource{}
}

// CouponResource defines the resource implementation.
type CouponResource struct {
	client *sdk.SDK
}

// CouponResourceModel describes the resource data model.
type CouponResourceModel struct {
	ACL                *tfTypes.BaseEntityACL    `tfsdk:"acl"`
	Active             types.Bool                `tfsdk:"active"`
	Additional         map[string]types.String   `tfsdk:"additional"`
	CashbackPeriod     types.String              `tfsdk:"cashback_period"`
	Category           types.String              `tfsdk:"category"`
	CreatedAt          types.String              `tfsdk:"created_at"`
	Description        types.String              `tfsdk:"description"`
	Files              *tfTypes.BaseRelation     `tfsdk:"files"`
	FixedValue         types.Number              `tfsdk:"fixed_value"`
	FixedValueCurrency types.String              `tfsdk:"fixed_value_currency"`
	FixedValueDecimal  types.String              `tfsdk:"fixed_value_decimal"`
	ID                 types.String              `tfsdk:"id"`
	Manifest           []types.String            `tfsdk:"manifest"`
	Name               types.String              `tfsdk:"name"`
	Org                types.String              `tfsdk:"org"`
	Owners             []tfTypes.BaseEntityOwner `tfsdk:"owners"`
	PercentageValue    types.String              `tfsdk:"percentage_value"`
	Prices             *tfTypes.BaseRelation     `tfsdk:"prices"`
	PromoCodeUsage     types.String              `tfsdk:"promo_code_usage"`
	PromoCodes         []tfTypes.PromoCode       `tfsdk:"promo_codes"`
	Purpose            []types.String            `tfsdk:"purpose"`
	RequiresPromoCode  types.Bool                `tfsdk:"requires_promo_code"`
	Schema             types.String              `tfsdk:"schema"`
	Tags               []types.String            `tfsdk:"tags"`
	Title              types.String              `tfsdk:"title"`
	Type               types.String              `tfsdk:"type"`
	UpdatedAt          types.String              `tfsdk:"updated_at"`
}

func (r *CouponResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_coupon"
}

func (r *CouponResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Coupon Resource",
		Attributes: map[string]schema.Attribute{
			"acl": schema.SingleNestedAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.Object{
					speakeasy_objectplanmodifier.SuppressDiff(speakeasy_objectplanmodifier.ExplicitSuppress),
				},
				Attributes: map[string]schema.Attribute{
					"delete": schema.ListAttribute{
						Computed: true,
						PlanModifiers: []planmodifier.List{
							speakeasy_listplanmodifier.SuppressDiff(speakeasy_listplanmodifier.ExplicitSuppress),
						},
						ElementType: types.StringType,
					},
					"edit": schema.ListAttribute{
						Computed: true,
						PlanModifiers: []planmodifier.List{
							speakeasy_listplanmodifier.SuppressDiff(speakeasy_listplanmodifier.ExplicitSuppress),
						},
						ElementType: types.StringType,
					},
					"view": schema.ListAttribute{
						Computed: true,
						PlanModifiers: []planmodifier.List{
							speakeasy_listplanmodifier.SuppressDiff(speakeasy_listplanmodifier.ExplicitSuppress),
						},
						ElementType: types.StringType,
					},
				},
				Description: `Access control list (ACL) for an entity. Defines sharing access to external orgs or users.`,
			},
			"active": schema.BoolAttribute{
				Required: true,
				PlanModifiers: []planmodifier.Bool{
					speakeasy_boolplanmodifier.SuppressDiff(speakeasy_boolplanmodifier.ExplicitSuppress),
				},
			},
			"additional": schema.MapAttribute{
				Computed: true,
				Optional: true,
				PlanModifiers: []planmodifier.Map{
					speakeasy_mapplanmodifier.SuppressDiff(speakeasy_mapplanmodifier.ExplicitSuppress),
				},
				ElementType: types.StringType,
				Description: `Additional fields that are not part of the schema`,
				Validators: []validator.Map{
					mapvalidator.ValueStringsAre(validators.IsValidJSON()),
				},
			},
			"cashback_period": schema.StringAttribute{
				Computed: true,
				Optional: true,
				PlanModifiers: []planmodifier.String{
					speakeasy_stringplanmodifier.SuppressDiff(speakeasy_stringplanmodifier.ExplicitSuppress),
				},
				Description: `The cashback period, for now it's limited to either 0 months or 12 months. must be one of ["0", "12"]`,
				Validators: []validator.String{
					stringvalidator.OneOf("0", "12"),
				},
			},
			"category": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					speakeasy_stringplanmodifier.SuppressDiff(speakeasy_stringplanmodifier.ExplicitSuppress),
				},
				Description: `must be one of ["discount", "cashback"]`,
				Validators: []validator.String{
					stringvalidator.OneOf(
						"discount",
						"cashback",
					),
				},
			},
			"created_at": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					speakeasy_stringplanmodifier.SuppressDiff(speakeasy_stringplanmodifier.ExplicitSuppress),
				},
				Validators: []validator.String{
					validators.IsRFC3339(),
				},
			},
			"description": schema.StringAttribute{
				Computed: true,
				Optional: true,
				PlanModifiers: []planmodifier.String{
					speakeasy_stringplanmodifier.SuppressDiff(speakeasy_stringplanmodifier.ExplicitSuppress),
				},
			},
			"files": schema.SingleNestedAttribute{
				Computed: true,
				Optional: true,
				PlanModifiers: []planmodifier.Object{
					speakeasy_objectplanmodifier.SuppressDiff(speakeasy_objectplanmodifier.ExplicitSuppress),
				},
				Attributes: map[string]schema.Attribute{
					"dollar_relation": schema.ListNestedAttribute{
						Computed: true,
						Optional: true,
						PlanModifiers: []planmodifier.List{
							speakeasy_listplanmodifier.SuppressDiff(speakeasy_listplanmodifier.ExplicitSuppress),
						},
						NestedObject: schema.NestedAttributeObject{
							Validators: []validator.Object{
								speakeasy_objectvalidators.NotNull(),
							},
							PlanModifiers: []planmodifier.Object{
								speakeasy_objectplanmodifier.SuppressDiff(speakeasy_objectplanmodifier.ExplicitSuppress),
							},
							Attributes: map[string]schema.Attribute{
								"entity_id": schema.StringAttribute{
									Computed: true,
									Optional: true,
									PlanModifiers: []planmodifier.String{
										speakeasy_stringplanmodifier.SuppressDiff(speakeasy_stringplanmodifier.ExplicitSuppress),
									},
								},
								"tags": schema.ListAttribute{
									Computed: true,
									Optional: true,
									PlanModifiers: []planmodifier.List{
										speakeasy_listplanmodifier.SuppressDiff(speakeasy_listplanmodifier.ExplicitSuppress),
									},
									ElementType: types.StringType,
								},
							},
						},
					},
				},
			},
			"fixed_value": schema.NumberAttribute{
				Computed: true,
				Optional: true,
				PlanModifiers: []planmodifier.Number{
					speakeasy_numberplanmodifier.SuppressDiff(speakeasy_numberplanmodifier.ExplicitSuppress),
				},
				Description: `Use if type is set to fixed. The fixed amount in cents to be discounted, represented as a whole integer.`,
			},
			"fixed_value_currency": schema.StringAttribute{
				Computed: true,
				Optional: true,
				PlanModifiers: []planmodifier.String{
					speakeasy_stringplanmodifier.SuppressDiff(speakeasy_stringplanmodifier.ExplicitSuppress),
				},
				Description: `Use if type is set to fixed. Three-letter ISO currency code, in lowercase.`,
			},
			"fixed_value_decimal": schema.StringAttribute{
				Computed: true,
				Optional: true,
				PlanModifiers: []planmodifier.String{
					speakeasy_stringplanmodifier.SuppressDiff(speakeasy_stringplanmodifier.ExplicitSuppress),
				},
				Description: `Use if type is set to fixed. The unit amount in cents to be discounted, represented as a decimal string with at most 12 decimal places.`,
			},
			"id": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					speakeasy_stringplanmodifier.SuppressDiff(speakeasy_stringplanmodifier.ExplicitSuppress),
				},
			},
			"manifest": schema.ListAttribute{
				Computed: true,
				Optional: true,
				PlanModifiers: []planmodifier.List{
					speakeasy_listplanmodifier.SuppressDiff(speakeasy_listplanmodifier.ExplicitSuppress),
				},
				ElementType: types.StringType,
				Description: `Manifest ID used to create/update the entity`,
			},
			"name": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					speakeasy_stringplanmodifier.SuppressDiff(speakeasy_stringplanmodifier.ExplicitSuppress),
				},
			},
			"org": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					speakeasy_stringplanmodifier.SuppressDiff(speakeasy_stringplanmodifier.ExplicitSuppress),
				},
				Description: `Organization Id the entity belongs to`,
			},
			"owners": schema.ListNestedAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.List{
					speakeasy_listplanmodifier.SuppressDiff(speakeasy_listplanmodifier.ExplicitSuppress),
				},
				NestedObject: schema.NestedAttributeObject{
					PlanModifiers: []planmodifier.Object{
						speakeasy_objectplanmodifier.SuppressDiff(speakeasy_objectplanmodifier.ExplicitSuppress),
					},
					Attributes: map[string]schema.Attribute{
						"org_id": schema.StringAttribute{
							Computed: true,
							PlanModifiers: []planmodifier.String{
								speakeasy_stringplanmodifier.SuppressDiff(speakeasy_stringplanmodifier.ExplicitSuppress),
							},
						},
						"user_id": schema.StringAttribute{
							Computed: true,
							PlanModifiers: []planmodifier.String{
								speakeasy_stringplanmodifier.SuppressDiff(speakeasy_stringplanmodifier.ExplicitSuppress),
							},
						},
					},
				},
			},
			"percentage_value": schema.StringAttribute{
				Computed: true,
				Optional: true,
				PlanModifiers: []planmodifier.String{
					speakeasy_stringplanmodifier.SuppressDiff(speakeasy_stringplanmodifier.ExplicitSuppress),
				},
				Description: `Use if type is set to percentage. The percentage to be discounted, represented as a whole integer.`,
			},
			"prices": schema.SingleNestedAttribute{
				Computed: true,
				Optional: true,
				PlanModifiers: []planmodifier.Object{
					speakeasy_objectplanmodifier.SuppressDiff(speakeasy_objectplanmodifier.ExplicitSuppress),
				},
				Attributes: map[string]schema.Attribute{
					"dollar_relation": schema.ListNestedAttribute{
						Computed: true,
						Optional: true,
						PlanModifiers: []planmodifier.List{
							speakeasy_listplanmodifier.SuppressDiff(speakeasy_listplanmodifier.ExplicitSuppress),
						},
						NestedObject: schema.NestedAttributeObject{
							Validators: []validator.Object{
								speakeasy_objectvalidators.NotNull(),
							},
							PlanModifiers: []planmodifier.Object{
								speakeasy_objectplanmodifier.SuppressDiff(speakeasy_objectplanmodifier.ExplicitSuppress),
							},
							Attributes: map[string]schema.Attribute{
								"entity_id": schema.StringAttribute{
									Computed: true,
									Optional: true,
									PlanModifiers: []planmodifier.String{
										speakeasy_stringplanmodifier.SuppressDiff(speakeasy_stringplanmodifier.ExplicitSuppress),
									},
								},
								"tags": schema.ListAttribute{
									Computed: true,
									Optional: true,
									PlanModifiers: []planmodifier.List{
										speakeasy_listplanmodifier.SuppressDiff(speakeasy_listplanmodifier.ExplicitSuppress),
									},
									ElementType: types.StringType,
								},
							},
						},
					},
				},
			},
			"promo_code_usage": schema.StringAttribute{
				Computed: true,
				Optional: true,
				PlanModifiers: []planmodifier.String{
					speakeasy_stringplanmodifier.SuppressDiff(speakeasy_stringplanmodifier.ExplicitSuppress),
				},
				Description: `Map of ids of promo codes with their usage count. Parsed as JSON.`,
				Validators: []validator.String{
					validators.IsValidJSON(),
				},
			},
			"promo_codes": schema.ListNestedAttribute{
				Computed: true,
				Optional: true,
				PlanModifiers: []planmodifier.List{
					speakeasy_listplanmodifier.SuppressDiff(speakeasy_listplanmodifier.ExplicitSuppress),
				},
				NestedObject: schema.NestedAttributeObject{
					Validators: []validator.Object{
						speakeasy_objectvalidators.NotNull(),
					},
					PlanModifiers: []planmodifier.Object{
						speakeasy_objectplanmodifier.SuppressDiff(speakeasy_objectplanmodifier.ExplicitSuppress),
					},
					Attributes: map[string]schema.Attribute{
						"code": schema.StringAttribute{
							Computed: true,
							Optional: true,
							PlanModifiers: []planmodifier.String{
								speakeasy_stringplanmodifier.SuppressDiff(speakeasy_stringplanmodifier.ExplicitSuppress),
							},
							Description: `The code of the promo code. Not Null`,
							Validators: []validator.String{
								speakeasy_stringvalidators.NotNull(),
							},
						},
						"has_usage_limit": schema.BoolAttribute{
							Computed: true,
							Optional: true,
							PlanModifiers: []planmodifier.Bool{
								speakeasy_boolplanmodifier.SuppressDiff(speakeasy_boolplanmodifier.ExplicitSuppress),
							},
							Description: `Whether the promo code has a usage limit`,
						},
						"id": schema.StringAttribute{
							Computed: true,
							Optional: true,
							PlanModifiers: []planmodifier.String{
								speakeasy_stringplanmodifier.SuppressDiff(speakeasy_stringplanmodifier.ExplicitSuppress),
							},
							Description: `The id of the promo code. Not Null`,
							Validators: []validator.String{
								speakeasy_stringvalidators.NotNull(),
							},
						},
						"usage_limit": schema.NumberAttribute{
							Computed: true,
							Optional: true,
							PlanModifiers: []planmodifier.Number{
								speakeasy_numberplanmodifier.SuppressDiff(speakeasy_numberplanmodifier.ExplicitSuppress),
							},
							Description: `The usage limit of the promo code`,
						},
					},
				},
			},
			"purpose": schema.ListAttribute{
				Computed: true,
				Optional: true,
				PlanModifiers: []planmodifier.List{
					speakeasy_listplanmodifier.SuppressDiff(speakeasy_listplanmodifier.ExplicitSuppress),
				},
				ElementType: types.StringType,
			},
			"requires_promo_code": schema.BoolAttribute{
				Computed: true,
				Optional: true,
				PlanModifiers: []planmodifier.Bool{
					speakeasy_boolplanmodifier.SuppressDiff(speakeasy_boolplanmodifier.ExplicitSuppress),
				},
				Description: `Whether the coupon requires a promo code to be applied`,
			},
			"schema": schema.StringAttribute{
				Computed: true,
				Optional: true,
				PlanModifiers: []planmodifier.String{
					speakeasy_stringplanmodifier.SuppressDiff(speakeasy_stringplanmodifier.ExplicitSuppress),
				},
				Description: `must be "coupon"`,
				Validators: []validator.String{
					stringvalidator.OneOf("coupon"),
				},
			},
			"tags": schema.ListAttribute{
				Computed: true,
				Optional: true,
				PlanModifiers: []planmodifier.List{
					speakeasy_listplanmodifier.SuppressDiff(speakeasy_listplanmodifier.ExplicitSuppress),
				},
				ElementType: types.StringType,
			},
			"title": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					speakeasy_stringplanmodifier.SuppressDiff(speakeasy_stringplanmodifier.ExplicitSuppress),
				},
			},
			"type": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					speakeasy_stringplanmodifier.SuppressDiff(speakeasy_stringplanmodifier.ExplicitSuppress),
				},
				Description: `must be one of ["fixed", "percentage"]`,
				Validators: []validator.String{
					stringvalidator.OneOf(
						"fixed",
						"percentage",
					),
				},
			},
			"updated_at": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					speakeasy_stringplanmodifier.SuppressDiff(speakeasy_stringplanmodifier.ExplicitSuppress),
				},
				Validators: []validator.String{
					validators.IsRFC3339(),
				},
			},
		},
	}
}

func (r *CouponResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*sdk.SDK)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *sdk.SDK, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.client = client
}

func (r *CouponResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *CouponResourceModel
	var plan types.Object

	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(plan.As(ctx, &data, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})...)

	if resp.Diagnostics.HasError() {
		return
	}

	request := *data.ToSharedCouponCreate()
	res, err := r.client.Coupon.CreateCoupon(ctx, request)
	if err != nil {
		resp.Diagnostics.AddError("failure to invoke API", err.Error())
		if res != nil && res.RawResponse != nil {
			resp.Diagnostics.AddError("unexpected http request/response", debugResponse(res.RawResponse))
		}
		return
	}
	if res == nil {
		resp.Diagnostics.AddError("unexpected response from API", fmt.Sprintf("%v", res))
		return
	}
	if res.StatusCode != 201 {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res.StatusCode), debugResponse(res.RawResponse))
		return
	}
	if !(res.Coupon != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedCoupon(res.Coupon)
	refreshPlan(ctx, plan, &data, resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *CouponResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *CouponResourceModel
	var item types.Object

	resp.Diagnostics.Append(req.State.Get(ctx, &item)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(item.As(ctx, &data, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})...)

	if resp.Diagnostics.HasError() {
		return
	}

	var couponID string
	couponID = data.ID.ValueString()

	// read.coupon.hydrateread.coupon.hydrate impedance mismatch: boolean != classtrace=["Coupon#create.req"]
	var hydrate *bool
	// read.coupon.strictread.coupon.strict impedance mismatch: boolean != classtrace=["Coupon#create.req"]
	var strict *bool
	request := operations.GetCouponRequest{
		CouponID: couponID,
		Hydrate:  hydrate,
		Strict:   strict,
	}
	res, err := r.client.Coupon.GetCoupon(ctx, request)
	if err != nil {
		resp.Diagnostics.AddError("failure to invoke API", err.Error())
		if res != nil && res.RawResponse != nil {
			resp.Diagnostics.AddError("unexpected http request/response", debugResponse(res.RawResponse))
		}
		return
	}
	if res == nil {
		resp.Diagnostics.AddError("unexpected response from API", fmt.Sprintf("%v", res))
		return
	}
	if res.StatusCode == 404 {
		resp.State.RemoveResource(ctx)
		return
	}
	if res.StatusCode != 200 {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res.StatusCode), debugResponse(res.RawResponse))
		return
	}
	if !(res.Coupon != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedCoupon(res.Coupon)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *CouponResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *CouponResourceModel
	var plan types.Object

	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	merge(ctx, req, resp, &data)
	if resp.Diagnostics.HasError() {
		return
	}

	couponPatch := *data.ToSharedCouponPatch()
	var couponID string
	couponID = data.ID.ValueString()

	request := operations.PatchCouponRequest{
		CouponPatch: couponPatch,
		CouponID:    couponID,
	}
	res, err := r.client.Coupon.PatchCoupon(ctx, request)
	if err != nil {
		resp.Diagnostics.AddError("failure to invoke API", err.Error())
		if res != nil && res.RawResponse != nil {
			resp.Diagnostics.AddError("unexpected http request/response", debugResponse(res.RawResponse))
		}
		return
	}
	if res == nil {
		resp.Diagnostics.AddError("unexpected response from API", fmt.Sprintf("%v", res))
		return
	}
	if res.StatusCode != 200 {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res.StatusCode), debugResponse(res.RawResponse))
		return
	}
	if !(res.Coupon != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedCoupon(res.Coupon)
	refreshPlan(ctx, plan, &data, resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *CouponResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *CouponResourceModel
	var item types.Object

	resp.Diagnostics.Append(req.State.Get(ctx, &item)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(item.As(ctx, &data, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})...)

	if resp.Diagnostics.HasError() {
		return
	}

	var couponID string
	couponID = data.ID.ValueString()

	request := operations.DeleteCouponRequest{
		CouponID: couponID,
	}
	res, err := r.client.Coupon.DeleteCoupon(ctx, request)
	if err != nil {
		resp.Diagnostics.AddError("failure to invoke API", err.Error())
		if res != nil && res.RawResponse != nil {
			resp.Diagnostics.AddError("unexpected http request/response", debugResponse(res.RawResponse))
		}
		return
	}
	if res == nil {
		resp.Diagnostics.AddError("unexpected response from API", fmt.Sprintf("%v", res))
		return
	}
	if res.StatusCode != 200 {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res.StatusCode), debugResponse(res.RawResponse))
		return
	}

}

func (r *CouponResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), req.ID)...)
}
