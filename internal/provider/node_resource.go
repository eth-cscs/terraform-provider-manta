package provider

import (
	"context"
	"fmt"
	"terraform-provider-manta/manta"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ resource.Resource = &nodeResource{}
)

// NewnodeResource is a helper function to simplify the provider implementation.
func NewnodeResource() resource.Resource {
	return &nodeResource{}
}

// nodeResource is the resource implementation.
type nodeResource struct {
	client *manta.Wrapper
}

// Configure adds the provider configured client to the resource.
func (r *nodeResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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
func (r *nodeResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_node"
}

type nodeResourceModel struct {
	ID      types.String `tfsdk:"id"`
	Type    types.String `tfsdk:"type"`
	State   types.String `tfsdk:"state"`
	Flag    types.String `tfsdk:"flag"`
	Enabled types.Bool   `tfsdk:"enabled"`
	Role    types.String `tfsdk:"role"`
	NID     types.Int64  `tfsdk:"nid"`
	NetType types.String `tfsdk:"nettype"`
	Arch    types.String `tfsdk:"arch"`
	Class   types.String `tfsdk:"class"`
}

// Schema defines the schema for the resource.
func (r *nodeResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Required: true,
			},
			"type": schema.StringAttribute{
				Computed: true,
			},
			"state": schema.StringAttribute{
				Optional: true,
				Computed: true,
			},
			"flag": schema.StringAttribute{
				Computed: true,
			},
			"enabled": schema.BoolAttribute{
				Computed: true,
			},
			"role": schema.StringAttribute{
				Computed: true,
			},
			"nid": schema.Int64Attribute{
				Computed: true,
			},
			"nettype": schema.StringAttribute{
				Computed: true,
			},
			"arch": schema.StringAttribute{
				Computed: true,
			},
			"class": schema.StringAttribute{
				Computed: true,
			},
		},
		Description: "Node represents a single node.",
	}
}

// Create creates the resource and sets the initial Terraform state.
func (r *nodeResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "---- Create function ----")

	var plan nodeResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	//var node NodeItem

	state := string(plan.State.ValueString())

	if state != "" {
		_, err := r.client.PowerNodeId(string(plan.ID.ValueString()), state)
		if err != nil {
			resp.Diagnostics.AddError(
				"Error cannot swith the power status of the node: ", err.Error(),
			)
			return
		}
	}

	// Create new rfe
	node, err := r.client.GetNodeId(string(plan.ID.ValueString()))
	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating rfe",
			"Could not create rfe, unexpected error: "+err.Error(),
		)
		return
	}

	plan.Type = types.StringValue(node.Type)
	plan.State = types.StringValue(node.State)
	plan.Flag = types.StringValue(node.Flag)
	plan.Enabled = types.BoolValue(node.Enabled)
	plan.Role = types.StringValue(node.Role)
	plan.NID = types.Int64Value(int64(node.NID))
	plan.NetType = types.StringValue(node.NetType)
	plan.Arch = types.StringValue(node.Arch)
	plan.Class = types.StringValue(node.Class)

	// Set state to fully populated data
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "---- End Create function ----")
}

// Read refreshes the Terraform state with the latest data.
func (r *nodeResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
}

// Update updates the resource and sets the updated Terraform state on success.
func (r *nodeResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
}

// Delete deletes the resource and removes the Terraform state on success.
func (r *nodeResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
}
