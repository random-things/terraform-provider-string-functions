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

func TestRStrPosFound(t *testing.T) {
	input := "abcdefghijklmnopqrstuvwxyzdef"
	substring := "def"

	pos := strRPos(input, substring)
	if pos != strings.LastIndex(input, substring) {
		t.Errorf("Expected %d, got %d", strings.Index(input, substring), pos)
	}
}

func TestRStrPosNotFound(t *testing.T) {
	input := "abcdefghijklmnopqrstuvwxyz"
	substring := "yz1"

	pos := strRPos(input, substring)
	if pos != strings.LastIndex(input, substring) {
		t.Errorf("Expected %d, got %d", strings.Index(input, substring), pos)
	}
}

func TestAccStrRPos_Known(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(version.Must(version.NewVersion("1.8.0"))),
		},
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: `
				locals {
					input = "abcdefghijklmnopqrstuvwxyzabcdef"
				}
				output "position" {
					value = provider::string-functions::strrpos(local.input, "def")
				}
				`,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownOutputValue("position", knownvalue.Int64Exact(29)),
				},
			},
		},
	})
}
