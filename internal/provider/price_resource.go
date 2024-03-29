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
var _ resource.Resource = &PriceResource{}
var _ resource.ResourceWithImportState = &PriceResource{}

func NewPriceResource() resource.Resource {
	return &PriceResource{}
}

// PriceResource defines the resource implementation.
type PriceResource struct {
	client *sdk.SDK
}

// PriceResourceModel describes the resource data model.
type PriceResourceModel struct {
	ID                     types.String     `tfsdk:"id"`
	Active                 types.Bool       `tfsdk:"active"`
	BillingDurationAmount  types.Number     `tfsdk:"billing_duration_amount"`
	BillingDurationUnit    types.String     `tfsdk:"billing_duration_unit"`
	Description            types.String     `tfsdk:"description"`
	IsCompositePrice       types.Bool       `tfsdk:"is_composite_price"`
	IsTaxInclusive         types.Bool       `tfsdk:"is_tax_inclusive"`
	LongDescription        types.String     `tfsdk:"long_description"`
	NoticeTimeAmount       types.Number     `tfsdk:"notice_time_amount"`
	NoticeTimeUnit         types.String     `tfsdk:"notice_time_unit"`
	PriceDisplayInJourneys types.String     `tfsdk:"price_display_in_journeys"`
	PricingModel           types.String     `tfsdk:"pricing_model"`
	RenewalDurationAmount  types.Number     `tfsdk:"renewal_duration_amount"`
	RenewalDurationUnit    types.String     `tfsdk:"renewal_duration_unit"`
	Tax                    types.String     `tfsdk:"tax"`
	TerminationTimeAmount  types.Number     `tfsdk:"termination_time_amount"`
	TerminationTimeUnit    types.String     `tfsdk:"termination_time_unit"`
	Tiers                  []PriceTier      `tfsdk:"tiers"`
	Type                   types.String     `tfsdk:"type"`
	Unit                   *PriceCreateUnit `tfsdk:"unit"`
	UnitAmount             types.Number     `tfsdk:"unit_amount"`
	UnitAmountCurrency     types.String     `tfsdk:"unit_amount_currency"`
	UnitAmountDecimal      types.String     `tfsdk:"unit_amount_decimal"`
	VariablePrice          types.Bool       `tfsdk:"variable_price"`
}

func (r *PriceResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_price"
}

func (r *PriceResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Price Resource",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: `The price id`,
			},
			"active": schema.BoolAttribute{
				Required:    true,
				Description: `Whether the price can be used for new purchases.`,
			},
			"billing_duration_amount": schema.NumberAttribute{
				Computed:    true,
				Optional:    true,
				Description: `The billing period duration`,
			},
			"billing_duration_unit": schema.StringAttribute{
				Computed:    true,
				Optional:    true,
				Description: `The billing period duration unit. must be one of ["weeks", "months", "years"]`,
				Validators: []validator.String{
					stringvalidator.OneOf(
						"weeks",
						"months",
						"years",
					),
				},
			},
			"description": schema.StringAttribute{
				Required:    true,
				Description: `A brief description of the price.`,
			},
			"is_composite_price": schema.BoolAttribute{
				Computed:    true,
				Optional:    true,
				Description: `The flag for prices that contain price components.`,
			},
			"is_tax_inclusive": schema.BoolAttribute{
				Computed:    true,
				Optional:    true,
				Description: `Specifies whether the price is considered ` + "`" + `inclusive` + "`" + ` of taxes or not. Default: false`,
			},
			"long_description": schema.StringAttribute{
				Computed:    true,
				Optional:    true,
				Description: `A detailed description of the price. This is shown on the order document and order table. Multi-line supported.`,
			},
			"notice_time_amount": schema.NumberAttribute{
				Computed:    true,
				Optional:    true,
				Description: `The notice period duration`,
			},
			"notice_time_unit": schema.StringAttribute{
				Computed:    true,
				Optional:    true,
				Description: `The notice period duration unit. must be one of ["weeks", "months", "years"]`,
				Validators: []validator.String{
					stringvalidator.OneOf(
						"weeks",
						"months",
						"years",
					),
				},
			},
			"price_display_in_journeys": schema.StringAttribute{
				Computed:    true,
				Optional:    true,
				Description: `Defines the way the price amount is display in epilot journeys. must be one of ["show_price", "show_as_starting_price", "show_as_on_request"]`,
				Validators: []validator.String{
					stringvalidator.OneOf(
						"show_price",
						"show_as_starting_price",
						"show_as_on_request",
					),
				},
			},
			"pricing_model": schema.StringAttribute{
				Computed: true,
				Optional: true,
				MarkdownDescription: `Describes how to compute the price per period. Either ` + "`" + `per_unit` + "`" + `, ` + "`" + `tiered_graduated` + "`" + ` or ` + "`" + `tiered_volume` + "`" + `.` + "\n" +
					`- ` + "`" + `per_unit` + "`" + ` indicates that the fixed amount (specified in unit_amount or unit_amount_decimal) will be charged per unit in quantity` + "\n" +
					`- ` + "`" + `tiered_graduated` + "`" + ` indicates that the unit pricing will be computed using tiers attribute. The customer pays the price per unit in every range their purchase rises through.` + "\n" +
					`- ` + "`" + `tiered_volume` + "`" + ` indicates that the unit pricing will be computed using tiers attribute. The customer pays the same unit price for all purchased units.` + "\n" +
					`- ` + "`" + `tiered_flatfee` + "`" + ` While similar to tiered_volume, tiered flat fee charges for the same price (flat) for the entire range instead using the unit price to multiply the quantity.` + "\n" +
					`` + "\n" +
					`must be one of ["per_unit", "tiered_volume", "tiered_graduated", "tiered_flatfee"]; Default: "per_unit"`,
				Validators: []validator.String{
					stringvalidator.OneOf(
						"per_unit",
						"tiered_volume",
						"tiered_graduated",
						"tiered_flatfee",
					),
				},
			},
			"renewal_duration_amount": schema.NumberAttribute{
				Computed:    true,
				Optional:    true,
				Description: `The renewal period duration`,
			},
			"renewal_duration_unit": schema.StringAttribute{
				Computed:    true,
				Optional:    true,
				Description: `The renewal period duration unit. must be one of ["weeks", "months", "years"]`,
				Validators: []validator.String{
					stringvalidator.OneOf(
						"weeks",
						"months",
						"years",
					),
				},
			},
			"tax": schema.StringAttribute{
				Computed:    true,
				Optional:    true,
				Description: `Parsed as JSON.`,
				Validators: []validator.String{
					validators.IsValidJSON(),
				},
			},
			"termination_time_amount": schema.NumberAttribute{
				Computed:    true,
				Optional:    true,
				Description: `The termination period duration`,
			},
			"termination_time_unit": schema.StringAttribute{
				Computed:    true,
				Optional:    true,
				Description: `The termination period duration unit. must be one of ["weeks", "months", "years"]`,
				Validators: []validator.String{
					stringvalidator.OneOf(
						"weeks",
						"months",
						"years",
					),
				},
			},
			"tiers": schema.ListNestedAttribute{
				Computed: true,
				Optional: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"display_mode": schema.StringAttribute{
							Computed:    true,
							Optional:    true,
							Description: `must be one of ["hidden", "on_request"]`,
							Validators: []validator.String{
								stringvalidator.OneOf(
									"hidden",
									"on_request",
								),
							},
						},
						"flat_fee_amount": schema.NumberAttribute{
							Computed: true,
							Optional: true,
						},
						"flat_fee_amount_decimal": schema.StringAttribute{
							Computed: true,
							Optional: true,
						},
						"unit_amount": schema.NumberAttribute{
							Computed: true,
							Optional: true,
						},
						"unit_amount_decimal": schema.StringAttribute{
							Computed: true,
							Optional: true,
						},
						"up_to": schema.NumberAttribute{
							Computed: true,
							Optional: true,
						},
					},
				},
				MarkdownDescription: `Defines an array of tiers. Each tier has an upper bound, an unit amount and a flat fee.` + "\n" +
					``,
			},
			"type": schema.StringAttribute{
				Computed:    true,
				Optional:    true,
				Description: `One of ` + "`" + `one_time` + "`" + ` or ` + "`" + `recurring` + "`" + ` depending on whether the price is for a one-time purchase or a recurring (subscription) purchase. must be one of ["one_time", "recurring"]; Default: "one_time"`,
				Validators: []validator.String{
					stringvalidator.OneOf(
						"one_time",
						"recurring",
					),
				},
			},
			"unit": schema.SingleNestedAttribute{
				Computed: true,
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"str": schema.StringAttribute{
						Computed: true,
						Optional: true,
					},
					"one": schema.StringAttribute{
						Computed:    true,
						Optional:    true,
						Description: `built-in units. must be one of ["kw", "kwh", "m", "m2", "l", "cubic-meter", "cubic-meter-h", "ls", "a", "kva", "w", "wp", "kwp"]`,
						Validators: []validator.String{
							stringvalidator.OneOf(
								"kw",
								"kwh",
								"m",
								"m2",
								"l",
								"cubic-meter",
								"cubic-meter-h",
								"ls",
								"a",
								"kva",
								"w",
								"wp",
								"kwp",
							),
						},
					},
				},
				Description: `The unit of measurement used for display purposes and possibly for calculations when the price is variable.`,
				Validators: []validator.Object{
					validators.ExactlyOneChild(),
				},
			},
			"unit_amount": schema.NumberAttribute{
				Computed:    true,
				Optional:    true,
				Description: `The unit amount in cents to be charged, represented as a whole integer if possible.`,
			},
			"unit_amount_currency": schema.StringAttribute{
				Computed:    true,
				Optional:    true,
				Description: `Three-letter ISO currency code, in lowercase.`,
			},
			"unit_amount_decimal": schema.StringAttribute{
				Computed:    true,
				Optional:    true,
				Description: `The unit amount in cents to be charged, represented as a decimal string with at most 12 decimal places.`,
			},
			"variable_price": schema.BoolAttribute{
				Computed:    true,
				Optional:    true,
				Description: `The flag for prices that can be influenced by external variables such as user input. Default: false`,
			},
		},
	}
}

func (r *PriceResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *PriceResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *PriceResourceModel
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

	request := *data.ToSharedPriceCreate()
	res, err := r.client.Price.CreatePrice(ctx, request)
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
	if res.Price == nil {
		resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedPrice(res.Price)
	refreshPlan(ctx, plan, &data, resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *PriceResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *PriceResourceModel
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
	priceID := data.ID.ValueString()
	request := operations.GetPriceRequest{
		Hydrate: hydrate,
		PriceID: priceID,
	}
	res, err := r.client.Price.GetPrice(ctx, request)
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
	if res.Price == nil {
		resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedPrice(res.Price)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *PriceResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *PriceResourceModel
	var plan types.Object

	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	merge(ctx, req, resp, &data)
	if resp.Diagnostics.HasError() {
		return
	}

	priceCreate := *data.ToSharedPriceCreate()
	priceID := data.ID.ValueString()
	request := operations.UpdatePriceRequest{
		PriceCreate: priceCreate,
		PriceID:     priceID,
	}
	res, err := r.client.Price.UpdatePrice(ctx, request)
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
	if res.Price == nil {
		resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedPrice(res.Price)
	refreshPlan(ctx, plan, &data, resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *PriceResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *PriceResourceModel
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

	priceID := data.ID.ValueString()
	request := operations.DeletePriceRequest{
		PriceID: priceID,
	}
	res, err := r.client.Price.DeletePrice(ctx, request)
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

func (r *PriceResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), req.ID)...)
}
