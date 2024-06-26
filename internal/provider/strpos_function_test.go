package provider

import (
	"github.com/hashicorp/go-version"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfversion"
	"strings"
	"testing"
)

func TestStrPosFound(t *testing.T) {
	input := "abcdefghijklmnopqrstuvwxyz"
	substring := "def"

	pos := strPos(input, substring)
	if pos != strings.Index(input, substring) {
		t.Errorf("Expected %d, got %d", strings.Index(input, substring), pos)
	}
}

func TestStrPosNotFound(t *testing.T) {
	input := "abcdefghijklmnopqrstuvwxyz"
	substring := "yz1"

	pos := strPos(input, substring)
	if pos != strings.Index(input, substring) {
		t.Errorf("Expected %d, got %d", strings.Index(input, substring), pos)
	}
}

func TestAccStrPos_Known(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(version.Must(version.NewVersion("1.8.0"))),
		},
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: `
				locals {
					input = "abcdefghijklmnopqrstuvwxyz"
				}
				output "position" {
					value = provider::string-functions::strpos(local.input, "def")
				}
				`,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownOutputValue("position", knownvalue.Int64Exact(3)),
				},
			},
		},
	})
}
