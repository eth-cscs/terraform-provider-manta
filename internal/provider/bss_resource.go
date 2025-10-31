package provider

import (
	"context"
	"fmt"
	"terraform-provider-manta/manta"

	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ resource.Resource                = &bssResource{}
	_ resource.ResourceWithImportState = &bssResource{}
)

func NewbssResource() resource.Resource {
	return &bssResource{}
}

type bssResource struct {
	client *manta.Wrapper
}

func (r *bssResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *bssResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_bss"
}

type bssResourceModel struct {
	Hosts  types.List   `tfsdk:"hosts"`
	Macs   types.List   `tfsdk:"macs"`
	Nids   types.List   `tfsdk:"nids"`
	Params types.String `tfsdk:"params"`
	Kernel types.String `tfsdk:"kernel"`
	Initrd types.String `tfsdk:"initrd"`
}

func (r *bssResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"hosts": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
			},
			"macs": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
			},
			"nids": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
			},
			"params": schema.StringAttribute{
				Optional: true,
			},
			"kernel": schema.StringAttribute{
				Optional: true,
			},
			"initrd": schema.StringAttribute{
				Optional: true,
			},
		},
	}
}

func (r *bssResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan bssResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	hosts := tfListToStringArray(ctx, plan.Hosts)
	macs := tfListToStringArray(ctx, plan.Macs)
	nids := tfListToStringArray(ctx, plan.Nids)

	bssItem := manta.BssParams{
		Hosts:  hosts,
		Macs:   macs,
		Nids:   nids,
		Params: string(plan.Params.ValueString()),
		Initrd: string(plan.Initrd.ValueString()),
		Kernel: string(plan.Kernel.ValueString()),
	}

	_, err := r.client.AddBss(bssItem)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating bss",
			"Could not create bss, unexpected error: "+err.Error(),
		)
		return
	}

	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *bssResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
}

func (r *bssResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan bssResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	hosts := tfListToStringArray(ctx, plan.Hosts)
	macs := tfListToStringArray(ctx, plan.Macs)
	nids := tfListToStringArray(ctx, plan.Nids)

	bssItem := manta.BssParams{
		Hosts:  hosts,
		Macs:   macs,
		Nids:   nids,
		Params: string(plan.Params.ValueString()),
		Initrd: string(plan.Initrd.ValueString()),
		Kernel: string(plan.Kernel.ValueString()),
	}

	_, err := r.client.UpdateBss(bssItem)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating bss",
			"Could not create bss, unexpected error: "+err.Error(),
		)
		return
	}

	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *bssResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state bssResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	hosts := tfListToStringArray(ctx, state.Hosts)
	macs := tfListToStringArray(ctx, state.Macs)
	nids := tfListToStringArray(ctx, state.Nids)

	bssItem := manta.BssParams{
		Hosts:  hosts,
		Macs:   macs,
		Nids:   nids,
		Params: string(state.Params.ValueString()),
		Initrd: string(state.Initrd.ValueString()),
		Kernel: string(state.Kernel.ValueString()),
	}

	out, err := r.client.DeleteBss(bssItem)
	tflog.Debug(ctx, out)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Deleting BSS",
			"Could not delete bss, unexpected error: "+err.Error(),
		)
		return
	}
}

func (r *bssResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
}
