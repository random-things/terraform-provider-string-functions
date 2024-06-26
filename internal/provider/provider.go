package provider

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

var _ provider.ProviderWithFunctions = &stringFunctionsProvider{}

func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &stringFunctionsProvider{version: version}
	}
}

type stringFunctionsProvider struct {
	version string
}

func (p *stringFunctionsProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	tflog.Info(ctx, "Setting up string-functions metadata...")
	resp.TypeName = "string-functions"
	resp.Version = p.version
}

func (p *stringFunctionsProvider) Schema(_ context.Context, _ provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{}
}

func (*stringFunctionsProvider) Configure(_ context.Context, _ provider.ConfigureRequest, _ *provider.ConfigureResponse) {
}
func (*stringFunctionsProvider) DataSources(_ context.Context) []func() datasource.DataSource {
	return nil
}
func (*stringFunctionsProvider) Functions(_ context.Context) []func() function.Function {
	return []func() function.Function{
		NewChunkStringsFunction,
		NewLimitedSplitFunction,
		NewLimitedRSplitFunction,
		NewStrRPosFunction,
		NewStrPosFunction,
	}
}
func (p *stringFunctionsProvider) Resources(_ context.Context) []func() resource.Resource {
	return nil
}
