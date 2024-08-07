// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"github.com/epilot-dev/terraform-provider-epilot-product/internal/sdk"
	"github.com/epilot-dev/terraform-provider-epilot-product/internal/sdk/models/operations"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
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
	Active      types.Bool   `tfsdk:"active"`
	Description types.String `tfsdk:"description"`
	ID          types.String `tfsdk:"id"`
	Rate        types.String `tfsdk:"rate"`
	Region      types.String `tfsdk:"region"`
	Type        types.String `tfsdk:"type"`
}

func (r *TaxResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_tax"
}

func (r *TaxResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Tax Resource",
		Attributes: map[string]schema.Attribute{
			"active": schema.BoolAttribute{
				Required: true,
			},
			"description": schema.StringAttribute{
				Computed: true,
				Optional: true,
			},
			"id": schema.StringAttribute{
				Computed: true,
			},
			"rate": schema.StringAttribute{
				Required: true,
			},
			"region": schema.StringAttribute{
				Required:    true,
				Description: `must be one of ["DE", "AT", "CH"]`,
				Validators: []validator.String{
					stringvalidator.OneOf(
						"DE",
						"AT",
						"CH",
					),
				},
			},
			"type": schema.StringAttribute{
				Required:    true,
				Description: `must be one of ["VAT", "Custom"]`,
				Validators: []validator.String{
					stringvalidator.OneOf(
						"VAT",
						"Custom",
					),
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

	// read.tax.hydrateread.tax.hydrate impedance mismatch: "boolean" != "class"trace=["Tax#create.req"]
	// {"Name":"Tax","Type":"class","Original":{"Name":"TaxCreate","OriginalName":"TaxCreate","ContextStack":[{"Type":"refType","Identifier":"Schemas","Used":false,"MustUse":false},{"Type":"refName","Identifier":"TaxCreate","Used":false,"MustUse":false},{"Type":"component","Identifier":"true","Used":false,"MustUse":false}],"Type":"class","ItemType":null,"Fields":[{"Name":"active","OriginalName":"active","Type":{"Name":"","OriginalName":"","ContextStack":[],"Type":"boolean","ItemType":null,"Fields":[],"Validations":{"MinItems":null,"MinLength":null,"Minimum":null,"MaxItems":null,"MaxLength":null,"Maximum":null,"Pattern":null,"UniqueItems":null},"AssociatedTypes":[],"Enum":null,"Scope":"","IsComponent":false,"Truncated":false,"Comments":null,"Input":false,"Output":false,"Extensions":{},"Examples":[],"Format":"","Discriminator":null,"ComplexAny":false,"OutputLocation":"","ResolvedModel":"","EventStreamEnvelope":false,"ResponseEnvelope":false},"Comments":null,"Annotations":[{"Ignore":false,"FieldName":"active"}],"Nullable":false,"Optional":false,"SerializationMethod":null,"ErrorMessage":false,"Const":null,"Default":null,"IsAdditionalProperties":false,"ParameterIndex":null},{"Name":"description","OriginalName":"description","Type":{"Name":"","OriginalName":"","ContextStack":[],"Type":"string","ItemType":null,"Fields":[],"Validations":{"MinItems":null,"MinLength":null,"Minimum":null,"MaxItems":null,"MaxLength":null,"Maximum":null,"Pattern":null,"UniqueItems":null},"AssociatedTypes":[],"Enum":null,"Scope":"","IsComponent":false,"Truncated":false,"Comments":null,"Input":false,"Output":false,"Extensions":{},"Examples":[],"Format":"","Discriminator":null,"ComplexAny":false,"OutputLocation":"","ResolvedModel":"","EventStreamEnvelope":false,"ResponseEnvelope":false},"Comments":null,"Annotations":[{"Ignore":false,"FieldName":"description"}],"Nullable":false,"Optional":true,"SerializationMethod":null,"ErrorMessage":false,"Const":null,"Default":null,"IsAdditionalProperties":false,"ParameterIndex":null},{"Name":"rate","OriginalName":"rate","Type":{"Name":"","OriginalName":"","ContextStack":[],"Type":"string","ItemType":null,"Fields":[],"Validations":{"MinItems":null,"MinLength":null,"Minimum":null,"MaxItems":null,"MaxLength":null,"Maximum":null,"Pattern":null,"UniqueItems":null},"AssociatedTypes":[],"Enum":null,"Scope":"","IsComponent":false,"Truncated":false,"Comments":null,"Input":false,"Output":false,"Extensions":{},"Examples":[],"Format":"","Discriminator":null,"ComplexAny":false,"OutputLocation":"","ResolvedModel":"","EventStreamEnvelope":false,"ResponseEnvelope":false},"Comments":null,"Annotations":[{"Ignore":false,"FieldName":"rate"}],"Nullable":false,"Optional":false,"SerializationMethod":null,"ErrorMessage":false,"Const":null,"Default":null,"IsAdditionalProperties":false,"ParameterIndex":null},{"Name":"region","OriginalName":"region","Type":{"Name":"TaxCreate_region","OriginalName":"region","ContextStack":[{"Type":"refType","Identifier":"Schemas","Used":false,"MustUse":false},{"Type":"refName","Identifier":"TaxCreate","Used":true,"MustUse":false}],"Type":"enum","ItemType":null,"Fields":[],"Validations":{"MinItems":null,"MinLength":null,"Minimum":null,"MaxItems":null,"MaxLength":null,"Maximum":null,"Pattern":null,"UniqueItems":null},"AssociatedTypes":[],"Enum":{"Type":{"Name":"","OriginalName":"","ContextStack":[],"Type":"string","ItemType":null,"Fields":[],"Validations":{"MinItems":null,"MinLength":null,"Minimum":null,"MaxItems":null,"MaxLength":null,"Maximum":null,"Pattern":null,"UniqueItems":null},"AssociatedTypes":[],"Enum":null,"Scope":"","IsComponent":false,"Truncated":false,"Comments":null,"Input":false,"Output":false,"Extensions":{},"Examples":[],"Format":"","Discriminator":null,"ComplexAny":false,"OutputLocation":"","ResolvedModel":"","EventStreamEnvelope":false,"ResponseEnvelope":false},"Values":["DE","AT","CH"],"Names":[],"Open":false,"Format":""},"Scope":"shared","IsComponent":false,"Truncated":false,"Comments":null,"Input":false,"Output":false,"Extensions":{},"Examples":[],"Format":"","Discriminator":null,"ComplexAny":false,"OutputLocation":"models/shared","ResolvedModel":"TaxCreate","EventStreamEnvelope":false,"ResponseEnvelope":false},"Comments":null,"Annotations":[{"Ignore":false,"FieldName":"region"}],"Nullable":false,"Optional":false,"SerializationMethod":null,"ErrorMessage":false,"Const":null,"Default":null,"IsAdditionalProperties":false,"ParameterIndex":null},{"Name":"type","OriginalName":"type","Type":{"Name":"TaxCreate_type","OriginalName":"type","ContextStack":[{"Type":"refType","Identifier":"Schemas","Used":false,"MustUse":false},{"Type":"refName","Identifier":"TaxCreate","Used":true,"MustUse":false}],"Type":"enum","ItemType":null,"Fields":[],"Validations":{"MinItems":null,"MinLength":null,"Minimum":null,"MaxItems":null,"MaxLength":null,"Maximum":null,"Pattern":null,"UniqueItems":null},"AssociatedTypes":[],"Enum":{"Type":{"Name":"","OriginalName":"","ContextStack":[],"Type":"string","ItemType":null,"Fields":[],"Validations":{"MinItems":null,"MinLength":null,"Minimum":null,"MaxItems":null,"MaxLength":null,"Maximum":null,"Pattern":null,"UniqueItems":null},"AssociatedTypes":[],"Enum":null,"Scope":"","IsComponent":false,"Truncated":false,"Comments":null,"Input":false,"Output":false,"Extensions":{},"Examples":[],"Format":"","Discriminator":null,"ComplexAny":false,"OutputLocation":"","ResolvedModel":"","EventStreamEnvelope":false,"ResponseEnvelope":false},"Values":["VAT","Custom"],"Names":[],"Open":false,"Format":""},"Scope":"shared","IsComponent":false,"Truncated":false,"Comments":null,"Input":false,"Output":false,"Extensions":{},"Examples":[],"Format":"","Discriminator":null,"ComplexAny":false,"OutputLocation":"models/shared","ResolvedModel":"TaxCreate","EventStreamEnvelope":false,"ResponseEnvelope":false},"Comments":null,"Annotations":[{"Ignore":false,"FieldName":"type"}],"Nullable":false,"Optional":false,"SerializationMethod":null,"ErrorMessage":false,"Const":null,"Default":null,"IsAdditionalProperties":false,"ParameterIndex":null}],"Validations":{"MinItems":null,"MinLength":null,"Minimum":null,"MaxItems":null,"MaxLength":null,"Maximum":null,"Pattern":null,"UniqueItems":null},"AssociatedTypes":[],"Enum":null,"Scope":"shared","IsComponent":true,"Truncated":false,"Comments":null,"Input":false,"Output":false,"Extensions":{"x-speakeasy-entity":"Tax"},"Examples":[],"Format":"","Discriminator":null,"ComplexAny":false,"OutputLocation":"models/shared","ResolvedModel":"TaxCreate","EventStreamEnvelope":false,"ResponseEnvelope":false},"ContextStack":[{"Type":"refType","Identifier":"Schemas","Used":false,"MustUse":false},{"Type":"refName","Identifier":"TaxCreate","Used":false,"MustUse":false},{"Type":"component","Identifier":"true","Used":false,"MustUse":false}],"OutputLocation":"models/shared","Format":"","CircularReference":null,"ResolvedModel":"TaxCreate","Enum":null,"ItemType":null,"Scope":"shared","AssociatedTypes":[],"Discriminator":null,"ResponseEnvelope":false,"OriginalName":"TaxCreate","Extensions":{"x-speakeasy-trace":{"Tax#create.req":true},"x-speakeasy-param-computed":true,"x-untouched":true,"x-speakeasy-in-get-request":true,"x-speakeasy-in-get":true,"x-speakeasy-root":true,"x-speakeasy-entity":"Tax"},"Examples":[],"Fields":[{"Optional":false,"IsAdditionalProperties":false,"Name":"_id","Type":{"ResolvedModel":"","Type":"string","Extensions":{"x-speakeasy-trace":{"Tax#create.resp.id":true},"x-speakeasy-param-computed":true,"x-untouched":true,"x-speakeasy-param-readonly":true,"x-speakeasy-in-get":true},"CircularReference":null,"Fields":[],"EventStreamEnvelope":false,"ResponseEnvelope":false,"Comments":null,"Examples":[{"Value":{"Kind":8,"Style":0,"Tag":"!!str","Value":"123e4567-e89b-12d3-a456-426614174000","Anchor":"","Alias":null,"Content":[],"HeadComment":"","LineComment":"","FootComment":"","Line":309,"Column":16},"Description":"","IsTest":false}],"Name":"","IsComponent":false,"OutputLocation":"","AssociatedTypes":[],"ItemType":null,"Output":false,"Scope":"","Enum":null,"Discriminator":null,"Truncated":false,"Original":{"Name":"","OriginalName":"","ContextStack":[],"Type":"string","ItemType":null,"Fields":[],"Validations":{"MinItems":null,"MinLength":null,"Minimum":null,"MaxItems":null,"MaxLength":null,"Maximum":null,"Pattern":null,"UniqueItems":null},"AssociatedTypes":[],"Enum":null,"Scope":"","IsComponent":false,"Truncated":false,"Comments":null,"Input":false,"Output":false,"Extensions":{},"Examples":[{"Value":{"Kind":8,"Style":0,"Tag":"!!str","Value":"123e4567-e89b-12d3-a456-426614174000","Anchor":"","Alias":null,"Content":[],"HeadComment":"","LineComment":"","FootComment":"","Line":309,"Column":16},"Description":"","IsTest":false}],"Format":"uuid","Discriminator":null,"ComplexAny":false,"OutputLocation":"","ResolvedModel":"","EventStreamEnvelope":false,"ResponseEnvelope":false},"ContextStack":[],"Validations":{"Pattern":null,"UniqueItems":null,"MinItems":null,"MinLength":null,"Minimum":null,"MaxItems":null,"MaxLength":null,"Maximum":null},"OriginalName":"","Format":"uuid","Input":false},"Nullable":false,"Default":null,"Const":null,"Comments":null,"Annotations":[{"Ignore":false,"FieldName":"_id"}],"ErrorMessage":false,"OriginalName":"_id"},{"Nullable":false,"Annotations":[{"Ignore":false,"FieldName":"active"}],"Name":"active","Type":{"Examples":[],"Fields":[],"ResolvedModel":"","Comments":null,"OutputLocation":"","AssociatedTypes":[],"Discriminator":null,"Truncated":false,"OriginalName":"","Name":"","Enum":null,"ResponseEnvelope":false,"ContextStack":[],"Extensions":{"x-speakeasy-trace":{"Tax#create.req.active":true},"x-speakeasy-param-computed":true,"x-untouched":true,"x-speakeasy-in-get":true},"CircularReference":null,"Scope":"","EventStreamEnvelope":false,"Type":"boolean","Input":false,"Output":false,"Validations":{"MaxLength":null,"Maximum":null,"Pattern":null,"UniqueItems":null,"MinItems":null,"MinLength":null,"Minimum":null,"MaxItems":null},"Original":{"Name":"","OriginalName":"","ContextStack":[],"Type":"boolean","ItemType":null,"Fields":[],"Validations":{"MinItems":null,"MinLength":null,"Minimum":null,"MaxItems":null,"MaxLength":null,"Maximum":null,"Pattern":null,"UniqueItems":null},"AssociatedTypes":[],"Enum":null,"Scope":"","IsComponent":false,"Truncated":false,"Comments":null,"Input":false,"Output":false,"Extensions":{},"Examples":[],"Format":"","Discriminator":null,"ComplexAny":false,"OutputLocation":"","ResolvedModel":"","EventStreamEnvelope":false,"ResponseEnvelope":false},"Format":"","IsComponent":false,"ItemType":null},"Optional":false,"Default":null,"Const":null,"Comments":null,"IsAdditionalProperties":false,"ErrorMessage":false,"OriginalName":"active"},{"Nullable":false,"Default":null,"Const":null,"OriginalName":"description","Name":"description","Comments":null,"Annotations":[{"Ignore":false,"FieldName":"description"}],"IsAdditionalProperties":false,"ErrorMessage":false,"Type":{"Name":"","Discriminator":null,"Format":"","IsComponent":false,"ResolvedModel":"","Type":"string","CircularReference":null,"Validations":{"UniqueItems":null,"MinItems":null,"MinLength":null,"Minimum":null,"MaxItems":null,"MaxLength":null,"Maximum":null,"Pattern":null},"ResponseEnvelope":false,"Original":{"Name":"","OriginalName":"","ContextStack":[],"Type":"string","ItemType":null,"Fields":[],"Validations":{"MinItems":null,"MinLength":null,"Minimum":null,"MaxItems":null,"MaxLength":null,"Maximum":null,"Pattern":null,"UniqueItems":null},"AssociatedTypes":[],"Enum":null,"Scope":"","IsComponent":false,"Truncated":false,"Comments":null,"Input":false,"Output":false,"Extensions":{},"Examples":[],"Format":"","Discriminator":null,"ComplexAny":false,"OutputLocation":"","ResolvedModel":"","EventStreamEnvelope":false,"ResponseEnvelope":false},"OriginalName":"","Enum":null,"Output":false,"Truncated":false,"AssociatedTypes":[],"Fields":[],"ItemType":null,"Scope":"","EventStreamEnvelope":false,"Comments":null,"Input":false,"OutputLocation":"","ContextStack":[],"Extensions":{"x-speakeasy-in-get":true,"x-speakeasy-trace":{"Tax#create.req.description":true},"x-speakeasy-param-computed":true,"x-untouched":true},"Examples":[]},"Optional":true},{"Default":null,"Const":null,"Annotations":[{"Ignore":false,"FieldName":"rate"}],"IsAdditionalProperties":false,"ErrorMessage":false,"Type":{"Input":false,"AssociatedTypes":[],"Comments":null,"Discriminator":null,"Fields":[],"IsComponent":false,"ItemType":null,"CircularReference":null,"ContextStack":[],"Type":"string","Extensions":{"x-untouched":true,"x-speakeasy-in-get":true,"x-speakeasy-trace":{"Tax#create.req.rate":true},"x-speakeasy-param-computed":true},"OutputLocation":"","Name":"","Examples":[],"Format":"","ResolvedModel":"","Scope":"","EventStreamEnvelope":false,"Truncated":false,"Enum":null,"Validations":{"Minimum":null,"MaxItems":null,"MaxLength":null,"Maximum":null,"Pattern":null,"UniqueItems":null,"MinItems":null,"MinLength":null},"Original":{"Name":"","OriginalName":"","ContextStack":[],"Type":"string","ItemType":null,"Fields":[],"Validations":{"MinItems":null,"MinLength":null,"Minimum":null,"MaxItems":null,"MaxLength":null,"Maximum":null,"Pattern":null,"UniqueItems":null},"AssociatedTypes":[],"Enum":null,"Scope":"","IsComponent":false,"Truncated":false,"Comments":null,"Input":false,"Output":false,"Extensions":{},"Examples":[],"Format":"","Discriminator":null,"ComplexAny":false,"OutputLocation":"","ResolvedModel":"","EventStreamEnvelope":false,"ResponseEnvelope":false},"OriginalName":"","ResponseEnvelope":false,"Output":false},"Name":"rate","Nullable":false,"OriginalName":"rate","Optional":false,"Comments":null},{"IsAdditionalProperties":false,"OriginalName":"region","Name":"region","Nullable":false,"Default":null,"Const":null,"Comments":null,"Type":{"Name":"TaxCreate_region","Validations":{"MaxLength":null,"Maximum":null,"Pattern":null,"UniqueItems":null,"MinItems":null,"MinLength":null,"Minimum":null,"MaxItems":null},"CircularReference":null,"Original":{"Name":"TaxCreate_region","OriginalName":"region","ContextStack":[{"Type":"refType","Identifier":"Schemas","Used":false,"MustUse":false},{"Type":"refName","Identifier":"TaxCreate","Used":true,"MustUse":false}],"Type":"enum","ItemType":null,"Fields":[],"Validations":{"MinItems":null,"MinLength":null,"Minimum":null,"MaxItems":null,"MaxLength":null,"Maximum":null,"Pattern":null,"UniqueItems":null},"AssociatedTypes":[],"Enum":{"Type":{"Name":"","OriginalName":"","ContextStack":[],"Type":"string","ItemType":null,"Fields":[],"Validations":{"MinItems":null,"MinLength":null,"Minimum":null,"MaxItems":null,"MaxLength":null,"Maximum":null,"Pattern":null,"UniqueItems":null},"AssociatedTypes":[],"Enum":null,"Scope":"","IsComponent":false,"Truncated":false,"Comments":null,"Input":false,"Output":false,"Extensions":{},"Examples":[],"Format":"","Discriminator":null,"ComplexAny":false,"OutputLocation":"","ResolvedModel":"","EventStreamEnvelope":false,"ResponseEnvelope":false},"Values":["DE","AT","CH"],"Names":[],"Open":false,"Format":""},"Scope":"shared","IsComponent":false,"Truncated":false,"Comments":null,"Input":false,"Output":false,"Extensions":{},"Examples":[],"Format":"","Discriminator":null,"ComplexAny":false,"OutputLocation":"models/shared","ResolvedModel":"TaxCreate","EventStreamEnvelope":false,"ResponseEnvelope":false},"OutputLocation":"models/shared","Scope":"shared","ContextStack":[{"Type":"refType","Identifier":"Schemas","Used":false,"MustUse":false},{"Type":"refName","Identifier":"TaxCreate","Used":true,"MustUse":false}],"Truncated":false,"Fields":[],"Enum":{"Open":false,"Names":[],"Values":["DE","AT","CH"],"Type":{"ResolvedModel":"","EventStreamEnvelope":false,"Fields":[],"Input":false,"OriginalName":"","Extensions":{},"Discriminator":null,"Enum":null,"Comments":null,"Validations":{"MinLength":null,"Minimum":null,"MaxItems":null,"MaxLength":null,"Maximum":null,"Pattern":null,"UniqueItems":null,"MinItems":null},"Type":"string","IsComponent":false,"Original":{"Name":"","OriginalName":"","ContextStack":[],"Type":"string","ItemType":null,"Fields":[],"Validations":{"MinItems":null,"MinLength":null,"Minimum":null,"MaxItems":null,"MaxLength":null,"Maximum":null,"Pattern":null,"UniqueItems":null},"AssociatedTypes":[],"Enum":null,"Scope":"","IsComponent":false,"Truncated":false,"Comments":null,"Input":false,"Output":false,"Extensions":{},"Examples":[],"Format":"","Discriminator":null,"ComplexAny":false,"OutputLocation":"","ResolvedModel":"","EventStreamEnvelope":false,"ResponseEnvelope":false},"ContextStack":[],"Examples":[],"Output":false,"ResponseEnvelope":false,"Name":"","Scope":"","CircularReference":null,"Format":"","ItemType":null,"Truncated":false,"OutputLocation":"","AssociatedTypes":[]}},"AssociatedTypes":[],"Comments":null,"IsComponent":false,"ResponseEnvelope":false,"Extensions":{"x-speakeasy-trace":{"Tax#create.req.region":true},"x-speakeasy-param-computed":true,"x-untouched":true,"x-speakeasy-in-get":true},"Examples":[],"Format":"","EventStreamEnvelope":false,"Type":"enum","OriginalName":"region","Output":false,"ItemType":null,"ResolvedModel":"TaxCreate","Discriminator":null,"Input":false},"Optional":false,"Annotations":[{"Ignore":false,"FieldName":"region"}],"ErrorMessage":false},{"Const":null,"Annotations":[{"ParamType":"pathParam","Name":"taxId","Serialization":"","Style":"simple","Explode":false,"FieldType":{"Type":"string","Validations":{"MaxItems":null,"MaxLength":null,"Maximum":null,"Pattern":null,"UniqueItems":null,"MinItems":null,"MinLength":null,"Minimum":null},"AssociatedTypes":[],"Enum":null,"Scope":"","Format":"uuid","EventStreamEnvelope":false,"Output":false,"Extensions":{"x-speakeasy-match":"id"},"Examples":[{"IsTest":false,"Value":{"Content":[],"HeadComment":"","LineComment":"","Kind":8,"Style":0,"Tag":"!!str","Value":"123e4567-e89b-12d3-a456-426614174000","Alias":null,"FootComment":"","Anchor":"","Line":309,"Column":16},"Description":""}],"Name":"","ContextStack":[],"Fields":[],"Comments":{"ExtendedComments":{},"Deprecated":false,"DeprecationReplacement":"","DeprecationMessage":"","Summary":"","Description":"The tax id","ExternalDocs":null},"OutputLocation":"","ResponseEnvelope":false,"OriginalName":"","ItemType":null,"IsComponent":false,"Truncated":false,"Input":false,"Discriminator":null,"ComplexAny":false,"ResolvedModel":""},"Hidden":false}],"OriginalName":"taxId","Name":"taxId","Nullable":false,"Default":null,"ErrorMessage":false,"Type":{"ResponseEnvelope":false,"Original":{"Name":"","OriginalName":"","ContextStack":[],"Type":"string","ItemType":null,"Fields":[],"Validations":{"MinItems":null,"MinLength":null,"Minimum":null,"MaxItems":null,"MaxLength":null,"Maximum":null,"Pattern":null,"UniqueItems":null},"AssociatedTypes":[],"Enum":null,"Scope":"","IsComponent":false,"Truncated":false,"Comments":{"Summary":"","Description":"The tax id","ExternalDocs":null,"ExtendedComments":{},"Deprecated":false,"DeprecationReplacement":"","DeprecationMessage":""},"Input":false,"Output":false,"Extensions":{"x-speakeasy-match":"id"},"Examples":[{"Value":{"Kind":8,"Style":0,"Tag":"!!str","Value":"123e4567-e89b-12d3-a456-426614174000","Anchor":"","Alias":null,"Content":[],"HeadComment":"","LineComment":"","FootComment":"","Line":309,"Column":16},"Description":"","IsTest":false}],"Format":"uuid","Discriminator":null,"ComplexAny":false,"OutputLocation":"","ResolvedModel":"","EventStreamEnvelope":false,"ResponseEnvelope":false},"OutputLocation":"","Extensions":{"x-speakeasy-match":"id","x-speakeasy-trace":{"Tax#get.req.tax_id":true}},"Discriminator":null,"IsComponent":false,"CircularReference":null,"ItemType":null,"Input":false,"Output":false,"Format":"uuid","AssociatedTypes":[],"EventStreamEnvelope":false,"Fields":[],"Enum":null,"ContextStack":[],"Type":"string","Scope":"","Validations":{"UniqueItems":null,"MinItems":null,"MinLength":null,"Minimum":null,"MaxItems":null,"MaxLength":null,"Maximum":null,"Pattern":null},"ResolvedModel":"","Name":"","Examples":[{"Value":{"Kind":8,"Style":0,"Tag":"!!str","Value":"123e4567-e89b-12d3-a456-426614174000","Anchor":"","Alias":null,"Content":[],"HeadComment":"","LineComment":"","FootComment":"","Line":309,"Column":16},"Description":"","IsTest":false}],"Truncated":false,"OriginalName":"","Comments":{"DeprecationMessage":"","Summary":"","Description":"The tax id","ExternalDocs":null,"ExtendedComments":{},"Deprecated":false,"DeprecationReplacement":""}},"Optional":false,"Comments":{"Description":"The tax id","ExternalDocs":null,"ExtendedComments":{},"Deprecated":false,"DeprecationReplacement":"","DeprecationMessage":"","Summary":""},"IsAdditionalProperties":false},{"Nullable":false,"Const":null,"Comments":null,"Annotations":[{"Ignore":false,"FieldName":"type"}],"IsAdditionalProperties":false,"OriginalName":"type","Type":{"Scope":"shared","ContextStack":[{"Type":"refType","Identifier":"Schemas","Used":false,"MustUse":false},{"Type":"refName","Identifier":"TaxCreate","Used":true,"MustUse":false}],"Input":false,"OutputLocation":"models/shared","Type":"enum","Extensions":{"x-speakeasy-trace":{"Tax#create.req.type":true},"x-speakeasy-param-computed":true,"x-untouched":true,"x-speakeasy-in-get":true},"AssociatedTypes":[],"Examples":[],"Enum":{"Type":{"OriginalName":"","Type":"string","Discriminator":null,"Fields":[],"IsComponent":false,"Scope":"","Truncated":false,"ContextStack":[],"Comments":null,"Extensions":{},"AssociatedTypes":[],"ResponseEnvelope":false,"Original":{"Name":"","OriginalName":"","ContextStack":[],"Type":"string","ItemType":null,"Fields":[],"Validations":{"MinItems":null,"MinLength":null,"Minimum":null,"MaxItems":null,"MaxLength":null,"Maximum":null,"Pattern":null,"UniqueItems":null},"AssociatedTypes":[],"Enum":null,"Scope":"","IsComponent":false,"Truncated":false,"Comments":null,"Input":false,"Output":false,"Extensions":{},"Examples":[],"Format":"","Discriminator":null,"ComplexAny":false,"OutputLocation":"","ResolvedModel":"","EventStreamEnvelope":false,"ResponseEnvelope":false},"Name":"","ResolvedModel":"","Input":false,"CircularReference":null,"Examples":[],"EventStreamEnvelope":false,"Enum":null,"Validations":{"MinItems":null,"MinLength":null,"Minimum":null,"MaxItems":null,"MaxLength":null,"Maximum":null,"Pattern":null,"UniqueItems":null},"Format":"","ItemType":null,"Output":false,"OutputLocation":""},"Open":false,"Names":[],"Values":["VAT","Custom"]},"Comments":null,"Name":"TaxCreate_type","ResolvedModel":"TaxCreate","Discriminator":null,"ItemType":null,"Output":false,"Format":"","IsComponent":false,"EventStreamEnvelope":false,"ResponseEnvelope":false,"Original":{"Name":"TaxCreate_type","OriginalName":"type","ContextStack":[{"Type":"refType","Identifier":"Schemas","Used":false,"MustUse":false},{"Type":"refName","Identifier":"TaxCreate","Used":true,"MustUse":false}],"Type":"enum","ItemType":null,"Fields":[],"Validations":{"MinItems":null,"MinLength":null,"Minimum":null,"MaxItems":null,"MaxLength":null,"Maximum":null,"Pattern":null,"UniqueItems":null},"AssociatedTypes":[],"Enum":{"Type":{"Name":"","OriginalName":"","ContextStack":[],"Type":"string","ItemType":null,"Fields":[],"Validations":{"MinItems":null,"MinLength":null,"Minimum":null,"MaxItems":null,"MaxLength":null,"Maximum":null,"Pattern":null,"UniqueItems":null},"AssociatedTypes":[],"Enum":null,"Scope":"","IsComponent":false,"Truncated":false,"Comments":null,"Input":false,"Output":false,"Extensions":{},"Examples":[],"Format":"","Discriminator":null,"ComplexAny":false,"OutputLocation":"","ResolvedModel":"","EventStreamEnvelope":false,"ResponseEnvelope":false},"Values":["VAT","Custom"],"Names":[],"Open":false,"Format":""},"Scope":"shared","IsComponent":false,"Truncated":false,"Comments":null,"Input":false,"Output":false,"Extensions":{},"Examples":[],"Format":"","Discriminator":null,"ComplexAny":false,"OutputLocation":"models/shared","ResolvedModel":"TaxCreate","EventStreamEnvelope":false,"ResponseEnvelope":false},"OriginalName":"type","Fields":[],"Truncated":false,"Validations":{"Minimum":null,"MaxItems":null,"MaxLength":null,"Maximum":null,"Pattern":null,"UniqueItems":null,"MinItems":null,"MinLength":null},"CircularReference":null},"Name":"type","Optional":false,"Default":null,"ErrorMessage":false}],"Input":false,"Output":false,"EventStreamEnvelope":false,"Truncated":false,"Validations":{"MaxItems":null,"MaxLength":null,"Maximum":null,"Pattern":null,"UniqueItems":null,"MinItems":null,"MinLength":null,"Minimum":null},"Comments":null,"IsComponent":true}
	// {"ContextStack":[],"CircularReference":null,"OriginalName":"","Discriminator":null,"Format":"","Input":false,"IsComponent":false,"Truncated":false,"Output":false,"OutputLocation":"","Extensions":{},"AssociatedTypes":[],"Comments":{"ExternalDocs":null,"ExtendedComments":{},"Deprecated":false,"DeprecationReplacement":"","DeprecationMessage":"","Summary":"","Description":"Hydrates entities in relations when passed true"},"Scope":"","Validations":{"Pattern":null,"UniqueItems":null,"MinItems":null,"MinLength":null,"Minimum":null,"MaxItems":null,"MaxLength":null,"Maximum":null},"EventStreamEnvelope":false,"ResponseEnvelope":false,"Enum":null,"Original":{"Name":"","OriginalName":"","ContextStack":[],"Type":"boolean","ItemType":null,"Fields":[],"Validations":{"MinItems":null,"MinLength":null,"Minimum":null,"MaxItems":null,"MaxLength":null,"Maximum":null,"Pattern":null,"UniqueItems":null},"AssociatedTypes":[],"Enum":null,"Scope":"","IsComponent":false,"Truncated":false,"Comments":{"Summary":"","Description":"Hydrates entities in relations when passed true","ExternalDocs":null,"ExtendedComments":{},"Deprecated":false,"DeprecationReplacement":"","DeprecationMessage":""},"Input":false,"Output":false,"Extensions":{},"Examples":[],"Format":"","Discriminator":null,"ComplexAny":false,"OutputLocation":"","ResolvedModel":"","EventStreamEnvelope":false,"ResponseEnvelope":false},"Name":"","Type":"boolean","Fields":[],"ItemType":null,"Examples":[],"ResolvedModel":""}
	var hydrate *bool
	var taxID string
	taxID = data.ID.ValueString()

	request := operations.GetTaxRequest{
		Hydrate: hydrate,
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

	taxCreate := *data.ToSharedTaxCreate()
	var taxID string
	taxID = data.ID.ValueString()

	request := operations.UpdateTaxRequest{
		TaxCreate: taxCreate,
		TaxID:     taxID,
	}
	res, err := r.client.Tax.UpdateTax(ctx, request)
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
