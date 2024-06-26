package provider

import (
	"github.com/hashicorp/go-version"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfversion"
	"testing"
)

func TestLimitedRSplit(t *testing.T) {
	input := "this is a test"
	delimiter := " "
	n := 3

	parts := limitedRSplit(input, delimiter, n)
	if len(parts) != 3 {
		t.Errorf("expected 3 parts, got %d", len(parts))
	}
}

func TestAccLimitedRSplit_Known(t *testing.T) {
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
				}
				output "parts" {
					value = provider::string-functions::limited_rsplit(local.input, " ", 3)
				}
				`,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownOutputValue("parts", knownvalue.ListExact([]knownvalue.Check{
						knownvalue.StringExact("this is"),
						knownvalue.StringExact("a"),
						knownvalue.StringExact("test"),
					})),
				},
			},
		},
	})
}
