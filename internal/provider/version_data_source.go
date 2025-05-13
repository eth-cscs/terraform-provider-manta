package provider

// Temporary data source definition to test the provider

import (
	"context"
	"fmt"
	"terraform-provider-manta/manta"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// mantaDataSource is the data source implementation.
type versionDataSource struct {
	wrapper *manta.Wrapper
}

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &versionDataSource{}
	_ datasource.DataSourceWithConfigure = &versionDataSource{}
)

func NewVersionDataSource() datasource.DataSource {
	return &versionDataSource{}
}

// versionModel maps version schema data.
type versionDataSourceModel struct {
	Version types.String `tfsdk:"version"`
}

// Configure adds the provider configured client to the data source.
func (d *versionDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	// Add a nil check when handling ProviderData because Terraform
	// sets that data after it calls the ConfigureProvider RPC.
	if req.ProviderData == nil {
		return
	}

	wrapper, ok := req.ProviderData.(*manta.Wrapper)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			fmt.Sprintf("Expected *hmanta.Wrapper, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	d.wrapper = wrapper
}

func (d *versionDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_version"
}

// Schema defines the schema for the data source.
func (d *versionDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"version": schema.StringAttribute{
				Computed: true,
			},
		},
	}
}

func (d *versionDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state versionDataSourceModel
	// var beta betaReleaseModel

	out, err := d.wrapper.Version()
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Read Manta CLI Version",
			err.Error(),
		)
		return
	}

	// Parse version string
	state.Version = types.StringValue(out)

	// Set state
	diags := resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}
