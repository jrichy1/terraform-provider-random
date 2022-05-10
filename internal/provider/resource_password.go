package provider

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
)

type resourcePasswordType struct{}

func (r resourcePasswordType) GetSchema(context.Context) (tfsdk.Schema, diag.Diagnostics) {
	description := "Identical to [random_string](string.html) with the exception that the result is " +
		"treated as sensitive and, thus, _not_ displayed in console output. Read more about sensitive " +
		"data handling in the [Terraform documentation](https://www.terraform.io/docs/language/state/sensitive-data.html).\n" +
		"\n" +
		"This resource *does* use a cryptographic random number generator."
	return getStringSchemaV1(true, description), nil
}

func (r resourcePasswordType) NewResource(_ context.Context, p tfsdk.Provider) (tfsdk.Resource, diag.Diagnostics) {
	return resourcePassword{
		p: *(p.(*provider)),
	}, nil
}

type resourcePassword struct {
	p provider
}

func (r resourcePassword) Create(ctx context.Context, req tfsdk.CreateResourceRequest, resp *tfsdk.CreateResourceResponse) {
	createString(ctx, req, resp, true)
}

// Read does not need to perform any operations as the state in ReadResourceResponse is already populated.
func (r resourcePassword) Read(ctx context.Context, req tfsdk.ReadResourceRequest, resp *tfsdk.ReadResourceResponse) {
}

// Update is intentionally left blank as all required and optional attributes force replacement of the resource
// through the RequiresReplace AttributePlanModifier.
func (r resourcePassword) Update(ctx context.Context, req tfsdk.UpdateResourceRequest, resp *tfsdk.UpdateResourceResponse) {
}

// Delete does not need to explicitly call resp.State.RemoveResource() as this is automatically handled by the
// [framework](https://github.com/hashicorp/terraform-plugin-framework/pull/301).
func (r resourcePassword) Delete(ctx context.Context, req tfsdk.DeleteResourceRequest, resp *tfsdk.DeleteResourceResponse) {
}

func (r resourcePassword) ImportState(ctx context.Context, req tfsdk.ImportResourceStateRequest, resp *tfsdk.ImportResourceStateResponse) {
	importString(ctx, req, resp, true)
}

func (r resourcePassword) ValidateConfig(ctx context.Context, req tfsdk.ValidateResourceConfigRequest, resp *tfsdk.ValidateResourceConfigResponse) {
	validateLength(ctx, req, resp)
}
