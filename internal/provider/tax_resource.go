// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	speakeasy_boolplanmodifier "github.com/epilot-dev/terraform-provider-epilot-product/internal/planmodifiers/boolplanmodifier"
	speakeasy_listplanmodifier "github.com/epilot-dev/terraform-provider-epilot-product/internal/planmodifiers/listplanmodifier"
	speakeasy_mapplanmodifier "github.com/epilot-dev/terraform-provider-epilot-product/internal/planmodifiers/mapplanmodifier"
	speakeasy_objectplanmodifier "github.com/epilot-dev/terraform-provider-epilot-product/internal/planmodifiers/objectplanmodifier"
	speakeasy_stringplanmodifier "github.com/epilot-dev/terraform-provider-epilot-product/internal/planmodifiers/stringplanmodifier"
	tfTypes "github.com/epilot-dev/terraform-provider-epilot-product/internal/provider/types"
	"github.com/epilot-dev/terraform-provider-epilot-product/internal/sdk"
	"github.com/epilot-dev/terraform-provider-epilot-product/internal/sdk/models/operations"
	"github.com/epilot-dev/terraform-provider-epilot-product/internal/validators"
	speakeasy_objectvalidators "github.com/epilot-dev/terraform-provider-epilot-product/internal/validators/objectvalidators"
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
var _ resource.Resource = &TaxResource{}
var _ resource.ResourceWithImportState = &TaxResource{}

func NewTaxResource() resource.Resource {
	return &TaxResource{}
}

// TaxResource defines the resource implementation.
type TaxResource struct {
	client *sdk.SDK
}

// TaxResourceModel describes the resource data model.
type TaxResourceModel struct {
	ACL         *tfTypes.BaseEntityACL    `tfsdk:"acl"`
	Active      types.Bool                `tfsdk:"active"`
	Additional  map[string]types.String   `tfsdk:"additional"`
	CreatedAt   types.String              `tfsdk:"created_at"`
	Description types.String              `tfsdk:"description"`
	Files       *tfTypes.BaseRelation     `tfsdk:"files"`
	ID          types.String              `tfsdk:"id"`
	Manifest    []types.String            `tfsdk:"manifest"`
	Org         types.String              `tfsdk:"org"`
	Owners      []tfTypes.BaseEntityOwner `tfsdk:"owners"`
	Purpose     []types.String            `tfsdk:"purpose"`
	Rate        types.String              `tfsdk:"rate"`
	Region      types.String              `tfsdk:"region"`
	Schema      types.String              `tfsdk:"schema"`
	Tags        []types.String            `tfsdk:"tags"`
	Title       types.String              `tfsdk:"title"`
	Type        types.String              `tfsdk:"type"`
	UpdatedAt   types.String              `tfsdk:"updated_at"`
}

func (r *TaxResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_tax"
}

func (r *TaxResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Tax Resource",
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
			"purpose": schema.ListAttribute{
				Computed: true,
				Optional: true,
				PlanModifiers: []planmodifier.List{
					speakeasy_listplanmodifier.SuppressDiff(speakeasy_listplanmodifier.ExplicitSuppress),
				},
				ElementType: types.StringType,
			},
			"rate": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					speakeasy_stringplanmodifier.SuppressDiff(speakeasy_stringplanmodifier.ExplicitSuppress),
				},
			},
			"region": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					speakeasy_stringplanmodifier.SuppressDiff(speakeasy_stringplanmodifier.ExplicitSuppress),
				},
			},
			"schema": schema.StringAttribute{
				Computed: true,
				Optional: true,
				PlanModifiers: []planmodifier.String{
					speakeasy_stringplanmodifier.SuppressDiff(speakeasy_stringplanmodifier.ExplicitSuppress),
				},
				Description: `must be "tax"`,
				Validators: []validator.String{
					stringvalidator.OneOf("tax"),
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
				Description: `must be one of ["VAT", "Custom"]`,
				Validators: []validator.String{
					stringvalidator.OneOf(
						"VAT",
						"Custom",
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

func (r *TaxResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *TaxResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *TaxResourceModel
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

	request := *data.ToSharedTaxCreate()
	res, err := r.client.Tax.CreateTax(ctx, request)
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
	if !(res.Tax != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedTax(res.Tax)
	refreshPlan(ctx, plan, &data, resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *TaxResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *TaxResourceModel
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

	// read.tax.hydrateread.tax.hydrate impedance mismatch: boolean != classtrace=["Tax#create.req"]
	var hydrate *bool
	// read.tax.strictread.tax.strict impedance mismatch: boolean != classtrace=["Tax#create.req"]
	var strict *bool
	var taxID string
	taxID = data.ID.ValueString()

	request := operations.GetTaxRequest{
		Hydrate: hydrate,
		Strict:  strict,
		TaxID:   taxID,
	}
	res, err := r.client.Tax.GetTax(ctx, request)
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
	if !(res.Tax != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedTax(res.Tax)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *TaxResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *TaxResourceModel
	var plan types.Object

	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	merge(ctx, req, resp, &data)
	if resp.Diagnostics.HasError() {
		return
	}

	taxPatch := *data.ToSharedTaxPatch()
	var taxID string
	taxID = data.ID.ValueString()

	request := operations.PatchTaxRequest{
		TaxPatch: taxPatch,
		TaxID:    taxID,
	}
	res, err := r.client.Tax.PatchTax(ctx, request)
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
	if !(res.Tax != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedTax(res.Tax)
	refreshPlan(ctx, plan, &data, resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *TaxResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *TaxResourceModel
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

	var taxID string
	taxID = data.ID.ValueString()

	request := operations.DeleteTaxRequest{
		TaxID: taxID,
	}
	res, err := r.client.Tax.DeleteTax(ctx, request)
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

func (r *TaxResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), req.ID)...)
}
