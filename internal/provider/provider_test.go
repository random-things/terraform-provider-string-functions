package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
)

const providerConfig = `{}`

var testAccProtoV6ProviderFactories = map[string]func() (tfprotov6.ProviderServer, error){
	"string-functions": providerserver.NewProtocol6WithError(New("test")()),
}
