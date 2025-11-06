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
	_ resource.Resource                = &cloudinitResource{}
	_ resource.ResourceWithImportState = &cloudinitResource{}
)

func NewcloudinitResource() resource.Resource {
	return &cloudinitResource{}
}

type cloudinitResource struct {
	client *manta.Wrapper
}

func (r *cloudinitResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *cloudinitResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_cloudinit_defaults"
}

type cloudinitResourceModel struct {
	BaseUrl    types.String `tfsdk:"base_url"`
	PublicKeys types.List   `tfsdk:"public_keys"`
}

func (r *cloudinitResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"base_url": schema.StringAttribute{
				MarkdownDescription: "URL to cloud-init",
				Required:            true,
			},
			"public_keys": schema.ListAttribute{
				MarkdownDescription: "Public keys add in `~/.ssh/authorized_keys`",
				Required:            true,
				ElementType:         types.StringType,
			},
		},
		Description: `Set default meta-data for cluster in cloud-init.`,
	}
}

func (r *cloudinitResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan cloudinitResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	publicKeys := tfListToStringArray(ctx, plan.PublicKeys)

	cloudinit := manta.ClusterDefaults{
		BaseUrl:    string(plan.BaseUrl.ValueString()),
		PublicKeys: publicKeys,
	}

	err := r.client.CreateClusterDefault(cloudinit)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating cloud-init",
			"Could not create cloud-init, unexpected error: "+err.Error(),
		)
		return
	}

	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *cloudinitResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
}

func (r *cloudinitResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
}

func (r *cloudinitResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// once cloud-init set, we cannot unset/delete
}

func (r *cloudinitResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
}
