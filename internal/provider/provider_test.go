package provider

import (
	"github.com/hashicorp/go-version"
	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-testing/tfversion"
)

var testAccProtoV6ProviderFactories = map[string]func() (tfprotov6.ProviderServer, error){
	"string-functions": providerserver.NewProtocol6WithError(New("test")()),
}

var terraform18OrNewer = []tfversion.TerraformVersionCheck{
	tfversion.SkipBelow(version.Must(version.NewVersion("1.8.0"))),
}
