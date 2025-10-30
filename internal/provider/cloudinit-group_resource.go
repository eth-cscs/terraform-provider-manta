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
	_ resource.Resource                = &cloudinitGroupResource{}
	_ resource.ResourceWithImportState = &cloudinitGroupResource{}
)

func NewcloudinitGroupResource() resource.Resource {
	return &cloudinitGroupResource{}
}

type cloudinitGroupResource struct {
	client *manta.Wrapper
}

func (r *cloudinitGroupResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *cloudinitGroupResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_cloudinit_group"
}

type cloudinitGroupResourceModel struct {
	Name        types.String `tfsdk:"name"`
	Description types.String `tfsdk:"description"`
	File        types.Map    `tfsdk:"file"`
}

func (r *cloudinitGroupResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"name": schema.StringAttribute{
				Required: true,
			},
			"description": schema.StringAttribute{
				Optional: true,
			},
			"file": schema.MapAttribute{
				Required:    true,
				ElementType: types.StringType,
			},
		},
	}
}

func (r *cloudinitGroupResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan cloudinitGroupResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	fileMap := tfMapToMap(ctx, plan.File)

	for _, key := range []string{"content", "name", "encoding"} {
		_, ok := fileMap[key]
		if ok == false {
			fileMap[key] = ""
		}
	}

	file := manta.CloudConfigFile{
		Content:  []byte(fileMap["content"]),
		Name:     fileMap["name"],
		Encoding: fileMap["encoding"],
	}

	cloudinitGroup := manta.GroupData{
		Name:        string(plan.Name.ValueString()),
		Description: string(plan.Description.ValueString()),
		File:        file,
	}

	err := r.client.CreateGroupData(cloudinitGroup)
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

func (r *cloudinitGroupResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
}

func (r *cloudinitGroupResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
}

func (r *cloudinitGroupResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state cloudinitGroupResourceModel

	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	err := r.client.DeleteGroupData(string(state.Name.ValueString()))
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Deleting group",
			"Could not delete group, unexpected error: "+err.Error(),
		)
		return
	}
}

func (r *cloudinitGroupResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
}
