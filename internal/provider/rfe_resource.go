package provider

import (
	"context"
	"fmt"
	"terraform-provider-manta/manta"
	"time"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ resource.Resource = &rfeResource{}
)

// NewrfeResource is a helper function to simplify the provider implementation.
func NewrfeResource() resource.Resource {
	return &rfeResource{}
}

// rfeResource is the resource implementation.
type rfeResource struct {
	client *manta.Wrapper
}

// Configure adds the provider configured client to the resource.
func (r *rfeResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Add a nil check when handling ProviderData because Terraform
	// sets that data after it calls the ConfigureProvider RPC.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*manta.Wrapper)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			fmt.Sprintf("Expected *manta.Wrapper, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.client = client
}

// Metadata returns the resource type name.
func (r *rfeResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_rfe"
}

type rfeResourceModel struct {
	ID                 types.String `tfsdk:"id"`
	Type               types.String `tfsdk:"type"`
	Hostname           types.String `tfsdk:"hostname"`
	Domain             types.String `tfsdk:"domain"`
	FQDN               types.String `tfsdk:"fqdn"`
	User               types.String `tfsdk:"user"`
	Password           types.String `tfsdk:"password"`
	RediscoverOnUpdate types.Bool   `tfsdk:"rediscoveronupdate"`
	Enabled            types.Bool   `tfsdk:"enabled"`
	LastUpdated        types.String `tfsdk:"last_updated"`
}

// Schema defines the schema for the resource.
func (r *rfeResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Required: true,
			},
			"type": schema.StringAttribute{
				Computed: true,
			},
			"hostname": schema.StringAttribute{
				Optional: true,
				Computed: true,
			},
			"domain": schema.StringAttribute{
				Optional: true,
				Computed: true,
			},
			"fqdn": schema.StringAttribute{
				Optional: true,
				Computed: true,
			},
			"user": schema.StringAttribute{
				Optional: true,
			},
			"password": schema.StringAttribute{
				Optional: true,
			},
			"rediscoveronupdate": schema.BoolAttribute{
				Optional: true,
				Computed: true,
			},
			"enabled": schema.BoolAttribute{
				Optional: true,
				Computed: true,
			},
			"last_updated": schema.StringAttribute{
				Computed: true,
			},
		},
		Description: "Rfe represents a single rfe.",
	}
}

// Create creates the resource and sets the initial Terraform state.
func (r *rfeResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "---- Create function ----")
	// Retrieve values from plan
	var plan rfeResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Generate API request body from plan
	var rfeItem manta.RfeItem

	rfeItem = manta.RfeItem{
		ID:                 string(plan.ID.ValueString()),
		Hostname:           string(plan.Hostname.ValueString()),
		Domain:             string(plan.Domain.ValueString()),
		FQDN:               string(plan.FQDN.ValueString()),
		User:               string(plan.User.ValueString()),
		Password:           string(plan.Password.ValueString()),
		RediscoverOnUpdate: bool(plan.RediscoverOnUpdate.ValueBool()),
		Enabled:            bool(plan.Enabled.ValueBool()),
	}

	// Create new rfe
	rfeCreated, err := r.client.AddRfe(rfeItem)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating rfe",
			"Could not create rfe, unexpected error: "+err.Error(),
		)
		return
	}

	plan.Type = types.StringValue(rfeCreated.Type)
	plan.Hostname = types.StringValue(rfeCreated.Hostname)
	plan.Domain = types.StringValue(rfeCreated.Domain)
	plan.FQDN = types.StringValue(rfeCreated.FQDN)
	plan.Enabled = types.BoolValue(rfeCreated.Enabled)
	plan.RediscoverOnUpdate = types.BoolValue(rfeCreated.RediscoverOnUpdate)
	plan.LastUpdated = types.StringValue(time.Now().Format(time.RFC850))

	// Set state to fully populated data
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "---- End Create function ----")
}

// Read refreshes the Terraform state with the latest data.
func (r *rfeResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	tflog.Debug(ctx, "---- Read function ----")

	// Get current state
	var state rfeResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Get refreshed order value from HashiCups
	rfes, err := r.client.GetRfe()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Reading HashiCups Order",
			"Could not read HashiCups order ID "+state.ID.ValueString()+": "+err.Error(),
		)
		return
	}

	for _, rfe := range rfes {
		tflog.Debug(ctx, "rfe id: "+rfe.ID)
	}

	// Set refreshed state
	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Update updates the resource and sets the updated Terraform state on success.
func (r *rfeResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
}

// Delete deletes the resource and removes the Terraform state on success.
func (r *rfeResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	tflog.Debug(ctx, "---- Delete function ----")

	// Retrieve values from state
	var state rfeResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Delete existing order
	out, err := r.client.DeleteRfe(state.ID.ValueString())
	tflog.Debug(ctx, out)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Deleting HashiCups Order",
			"Could not delete order, unexpected error: "+err.Error(),
		)
		return
	}
}
