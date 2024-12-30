package provider

import (
	"github.com/hashicorp/go-version"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfversion"
	"testing"
)

func TestShellEscapeCmd_Known(t *testing.T) {
	input := []string{"echo", "this is a test"}

	output := shellEscapeCmd(input)
	if output != `echo 'this is a test'` {
		t.Errorf(`expected escaped string is "echo 'this is a test'", got %s`, output)
	}
}

func TestAccShellEscapeCmd_Known(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(version.Must(version.NewVersion("1.8.0"))),
		},
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: `
				locals {
					input = ["echo", "this is a test"]
				}
				output "new_string" {
					value = provider::string-functions::shell_escape_cmd(local.input)
				}
				`,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownOutputValue("new_string", knownvalue.StringExact(`echo 'this is a test'`)),
				},
			},
		},
	})
}
