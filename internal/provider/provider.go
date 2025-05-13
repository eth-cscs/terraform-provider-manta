package provider

import (
	"context"
	"os"
	"terraform-provider-manta/manta"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ provider.Provider = &mantaProvider{}
)

// mantaProviderModel maps provider schema data to a Go type.
type mantaProviderModel struct {
	BaseURL     types.String `tfsdk:"base_url"`
	AccessToken types.String `tfsdk:"access_token"`
}

// New is a helper function to simplify provider server and testing implementation.
func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &mantaProvider{
			version: version,
		}
	}
}

// mantaProvider is the provider implementation.
type mantaProvider struct {
	// version is set to the provider version on release, "dev" when the
	// provider is built and ran locally, and "test" when running acceptance
	// testing.
	version string
}

// Metadata returns the provider type name.
func (p *mantaProvider) Metadata(_ context.Context, _ provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "manta"
	resp.Version = p.version
}

// Schema defines the provider-level schema for configuration data.
func (p *mantaProvider) Schema(_ context.Context, _ provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"base_url": schema.StringAttribute{
				Required: true,
			},
			"access_token": schema.StringAttribute{
				Optional:  true,
				Sensitive: true,
			},
		},
	}
}

// Configure prepares a manta API client for data sources and resources.
func (p *mantaProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	// Retrieve provider data from configuration
	var config mantaProviderModel
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// If practitioner provided a configuration value for any of the
	// attributes, it must be a known value.

	if config.BaseURL.IsUnknown() {
		resp.Diagnostics.AddAttributeError(
			path.Root("base_url"),
			"Unknown Manta API Base URL",
			"The provider cannot create the Manta API client as there is an unknown configuration value for the Manta API base URL. "+
				"Either target apply the source of the value first, set the value statically in the configuration, or use the MANTA_BASE_URL environment variable.",
		)
	}

	if config.AccessToken.IsUnknown() {
		resp.Diagnostics.AddAttributeError(
			path.Root("access_token"),
			"Unknown Manta API Access Token",
			"The provider cannot create the Manta API client as there is an unknown configuration value for the Manta API access token. "+
				"Either target apply the source of the value first, set the value statically in the configuration, or use the MANTA_ACCESS_TOKEN environment variable.",
		)
	}

	if resp.Diagnostics.HasError() {
		return
	}

	// Default values to environment variables, but override
	// with Terraform configuration value if set.

	base_url := os.Getenv("MANTA_BASE_URL")
	access_token := os.Getenv("MANA_ACCESS_TOKEN")

	if !config.BaseURL.IsNull() {
		base_url = config.BaseURL.ValueString()
	}

	if !config.AccessToken.IsNull() {
		access_token = config.AccessToken.ValueString()
	}

	// If any of the expected configurations are missing, return
	// errors with provider-specific guidance.

	if base_url == "" {
		resp.Diagnostics.AddAttributeError(
			path.Root("base_url"),
			"Missing Manta API Base URL",
			"The provider cannot create the Manta API client as there is a missing or empty value for the Manta API base URL. "+
				"Set the base URL value in the configuration or use the MANTA_BASE_URL environment variable. "+
				"If either is already set, ensure the value is not empty.",
		)
	}

	if access_token == "" {
		resp.Diagnostics.AddAttributeError(
			path.Root("access_token"),
			"Missing Manta API Access Token",
			"The provider cannot create the Manta API client as there is a missing or empty value for the Manta API access token. "+
				"Set the access token value in the configuration or use the MANTA_ACCESS_TOKEN environment variable. "+
				"If either is already set, ensure the value is not empty.",
		)
	}

	if resp.Diagnostics.HasError() {
		return
	}
	tflog.Info(ctx, "Manta API Base URL: "+base_url)
	ctx = tflog.SetField(ctx, "manta_base_url", base_url)
	ctx = tflog.SetField(ctx, "manta_access_token", access_token)
	ctx = tflog.MaskFieldValuesWithFieldKeys(ctx, "manta_access_token")
	tflog.Debug(ctx, "Creating HashiCups client")

	// Create a new Manta client using the configuration values
	w, err := manta.NewWrapper(base_url, access_token)
	if err != nil {
		tflog.Error(ctx, "Unable to Create Manta CLI Wrapper: "+err.Error())
		resp.Diagnostics.AddError(
			"Unable to Create Manta CLI Wrapper",
			"An unexpected error occurred when creating the Manta Wrapper. "+
				"Make sure Manta CLI is properly installed in your setup.\n\n"+
				"Manta CLI Error: "+err.Error(),
		)
		return
	}

	// Make the HashiCups client available during DataSource and Resource
	// type Configure methods.
	resp.DataSourceData = w
	resp.ResourceData = w
}

// DataSources defines the data sources implemented in the provider.
func (p *mantaProvider) DataSources(_ context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		NewVersionDataSource,
	}
}

// Resources defines the resources implemented in the provider.
func (p *mantaProvider) Resources(_ context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		NewrfeResource,
		NewnodeResource,
	}
}
