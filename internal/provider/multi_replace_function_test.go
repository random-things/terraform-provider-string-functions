package provider

import (
	"github.com/hashicorp/go-version"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfversion"
	"testing"
)

func TestMultiReplace(t *testing.T) {
	input := "this is a test"
	replacements := map[string]string{
		" is":   " was",
		" test": " trial",
	}

	output := multiReplace(input, replacements)
	if output != "this was a trial" {
		t.Errorf("expected replacement string is 'this was a trial', got %s", output)
	}
}

func TestAccMultiReplace_Known(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(version.Must(version.NewVersion("1.8.0"))),
		},
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: `
				locals {
					input = "this is a test"
                    replacements = {
						" is" = " was"
						" test" = " trial"
                    }
				}
				output "new_string" {
					value = provider::string-functions::multi_replace(local.input, local.replacements)
				}
				`,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownOutputValue("new_string", knownvalue.StringExact("this was a trial")),
				},
			},
		},
	})
}
