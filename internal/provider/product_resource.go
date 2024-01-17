// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"github.com/epilot-dev/terraform-provider-epilot-product/internal/sdk"
	"github.com/epilot-dev/terraform-provider-epilot-product/internal/sdk/pkg/models/operations"
	"github.com/epilot-dev/terraform-provider-epilot-product/internal/validators"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &ProductResource{}
var _ resource.ResourceWithImportState = &ProductResource{}

func NewProductResource() resource.Resource {
	return &ProductResource{}
}

// ProductResource defines the resource implementation.
type ProductResource struct {
	client *sdk.SDK
}

// ProductResourceModel describes the resource data model.
type ProductResourceModel struct {
	ACL          BaseEntityACL     `tfsdk:"acl"`
	CreatedAt    types.String      `tfsdk:"created_at"`
	ID           types.String      `tfsdk:"id"`
	Org          types.String      `tfsdk:"org"`
	Owners       []BaseEntityOwner `tfsdk:"owners"`
	Schema       types.String      `tfsdk:"schema"`
	Tags         []types.String    `tfsdk:"tags"`
	Title        types.String      `tfsdk:"title"`
	UpdatedAt    types.String      `tfsdk:"updated_at"`
	Active       types.Bool        `tfsdk:"active"`
	Code         types.String      `tfsdk:"code"`
	Description  types.String      `tfsdk:"description"`
	Feature      []Feature         `tfsdk:"feature"`
	InternalName types.String      `tfsdk:"internal_name"`
	Name         types.String      `tfsdk:"name"`
	PriceOptions *BaseRelation     `tfsdk:"price_options"`
	Type         types.String      `tfsdk:"type"`
}

func (r *ProductResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_product"
}

func (r *ProductResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Product Resource",

		Attributes: map[string]schema.Attribute{
			"acl": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"additional_properties": schema.StringAttribute{
						Computed:    true,
						Description: `Parsed as JSON.`,
						Validators: []validator.String{
							validators.IsValidJSON(),
						},
					},
					"delete": schema.ListAttribute{
						Computed:    true,
						ElementType: types.StringType,
					},
					"edit": schema.ListAttribute{
						Computed:    true,
						ElementType: types.StringType,
					},
					"view": schema.ListAttribute{
						Computed:    true,
						ElementType: types.StringType,
					},
				},
				Description: `Access control list (ACL) for an entity. Defines sharing access to external orgs or users.`,
			},
			"created_at": schema.StringAttribute{
				Computed: true,
				Validators: []validator.String{
					validators.IsRFC3339(),
				},
			},
			"id": schema.StringAttribute{
				Computed:    true,
				Description: `The product id`,
			},
			"org": schema.StringAttribute{
				Computed:    true,
				Description: `Organization Id the entity belongs to`,
			},
			"owners": schema.ListNestedAttribute{
				Computed: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"org_id": schema.StringAttribute{
							Computed: true,
						},
						"user_id": schema.StringAttribute{
							Computed: true,
						},
					},
				},
			},
			"schema": schema.StringAttribute{
				Computed: true,
			},
			"tags": schema.ListAttribute{
				Computed:    true,
				ElementType: types.StringType,
			},
			"title": schema.StringAttribute{
				Computed: true,
			},
			"updated_at": schema.StringAttribute{
				Computed: true,
				Validators: []validator.String{
					validators.IsRFC3339(),
				},
			},
			"active": schema.BoolAttribute{
				Required: true,
			},
			"code": schema.StringAttribute{
				Computed:    true,
				Optional:    true,
				Description: `The product code`,
			},
			"description": schema.StringAttribute{
				Computed:    true,
				Optional:    true,
				Description: `A description of the product. Multi-line supported.`,
			},
			"feature": schema.ListNestedAttribute{
				Computed: true,
				Optional: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Computed: true,
							Optional: true,
						},
						"tags": schema.ListAttribute{
							Computed:    true,
							Optional:    true,
							ElementType: types.StringType,
						},
						"feature": schema.StringAttribute{
							Computed: true,
							Optional: true,
						},
					},
				},
			},
			"internal_name": schema.StringAttribute{
				Computed:    true,
				Optional:    true,
				Description: `Not visible to customers, only in internal tables`,
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: `The description for the product`,
			},
			"price_options": schema.SingleNestedAttribute{
				Computed: true,
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"dollar_relation": schema.ListNestedAttribute{
						Computed: true,
						Optional: true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"tags": schema.ListAttribute{
									Computed:    true,
									Optional:    true,
									ElementType: types.StringType,
								},
								"entity_id": schema.StringAttribute{
									Computed: true,
									Optional: true,
								},
							},
						},
					},
				},
			},
			"type": schema.StringAttribute{
				Computed: true,
				Optional: true,
				MarkdownDescription: `The type of Product:` + "\n" +
					`` + "\n" +
					`| type | description |` + "\n" +
					`|----| ----|` + "\n" +
					`| ` + "`" + `product` + "`" + ` | Represents a physical good |` + "\n" +
					`| ` + "`" + `service` + "`" + ` | Represents a service or virtual product |` + "\n" +
					`` + "\n" +
					`must be one of ["product", "service"]; Default: "product"`,
				Validators: []validator.String{
					stringvalidator.OneOf(
						"product",
						"service",
					),
				},
			},
		},
	}
}

func (r *ProductResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *ProductResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *ProductResourceModel
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

	request := *data.ToSharedProductCreate()
	res, err := r.client.Product.CreateProduct(ctx, request)
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
	if res.Product == nil {
		resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedProduct(res.Product)
	refreshPlan(ctx, plan, &data, resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ProductResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *ProductResourceModel
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

	var hydrate *bool
	productID := data.ID.ValueString()
	request := operations.GetProductRequest{
		Hydrate:   hydrate,
		ProductID: productID,
	}
	res, err := r.client.Product.GetProduct(ctx, request)
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
	if res.Product == nil {
		resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedProduct(res.Product)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ProductResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *ProductResourceModel
	var plan types.Object

	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	merge(ctx, req, resp, &data)
	if resp.Diagnostics.HasError() {
		return
	}

	productCreate := *data.ToSharedProductCreate()
	productID := data.ID.ValueString()
	request := operations.UpdateProductRequest{
		ProductCreate: productCreate,
		ProductID:     productID,
	}
	res, err := r.client.Product.UpdateProduct(ctx, request)
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
	if res.Product == nil {
		resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedProduct(res.Product)
	refreshPlan(ctx, plan, &data, resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ProductResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *ProductResourceModel
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

	productID := data.ID.ValueString()
	request := operations.DeleteProductRequest{
		ProductID: productID,
	}
	res, err := r.client.Product.DeleteProduct(ctx, request)
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

func (r *ProductResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), req.ID)...)
}
