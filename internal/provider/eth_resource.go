package provider

import (
	"context"
	"fmt"
	"terraform-provider-manta/manta"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ resource.Resource                = &ethResource{}
	_ resource.ResourceWithImportState = &ethResource{}
)

func NewethResource() resource.Resource {
	return &ethResource{}
}

type ethResource struct {
	client *manta.Wrapper
}

func (r *ethResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *ethResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_smd_interface"
}

type ethResourceModel struct {
	ID         types.String `tfsdk:"id"`
	Desc       types.String `tfsdk:"description"`
	MAC        types.String `tfsdk:"mac_address"`
	IPs        types.List   `tfsdk:"ip_addresses"`
	LastUpdate types.String `tfsdk:"last_update"`
	CompID     types.String `tfsdk:"component_id"`
	Type       types.String `tfsdk:"type"`
}

func (r *ethResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"description": schema.StringAttribute{
				Optional: true,
			},
			"mac_address": schema.StringAttribute{
				Required: true,
			},
			"ip_addresses": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
			},
			"last_update": schema.StringAttribute{
				Computed: true,
			},
			"component_id": schema.StringAttribute{
				Optional: true,
				Computed: true,
			},
			"type": schema.StringAttribute{
				Computed: true,
			},
		},
		Description: `Manage ethernet interfaces.`,
	}
}

func (r *ethResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan ethResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	ips := tfListToIps(ctx, plan.IPs)

	ethItem := manta.NodeInterface{
		IPs:    ips,
		CompID: plan.CompID.ValueString(),
		Type:   plan.Type.ValueString(),
		MAC:    plan.MAC.ValueString(),
	}

	err := r.client.AddEthernetInterface(ethItem)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating eth",
			"Could not create eth, unexpected error: "+err.Error(),
		)
		return
	}

	nodeEth, err := r.client.GetEthernetInterface(ethItem.MAC)

	plan.CompID = types.StringValue(nodeEth.CompID)
	plan.Type = types.StringValue(nodeEth.Type)
	plan.ID = types.StringValue(nodeEth.ID)
	plan.LastUpdate = types.StringValue(nodeEth.LastUpdate)

	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *ethResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
}

func (r *ethResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data ethResourceModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	ips := tfListToIps(ctx, data.IPs)
	updateReq := manta.NodeInterfacePatch{
		MAC: data.MAC.ValueString(),
		IPs: ips,
	}

	nodeEth, err := r.client.PatchEthernetInterface(updateReq)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error update eth",
			"Could not update eth, unexpected error: "+err.Error(),
		)
		return
	}

	data.CompID = types.StringValue(nodeEth.CompID)
	data.Type = types.StringValue(nodeEth.Type)
	data.ID = types.StringValue(nodeEth.ID)
	data.LastUpdate = types.StringValue(nodeEth.LastUpdate)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ethResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state ethResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	mac := state.MAC.ValueString()

	err := r.client.DeleteEthernetInterface(mac)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Deleting eth",
			"Could not delete eth, unexpected error: "+err.Error(),
		)
		return
	}
}

func (r *ethResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
}
