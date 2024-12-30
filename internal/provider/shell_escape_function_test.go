package provider

import (
	"github.com/hashicorp/go-version"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfversion"
	"testing"
)

func TestShellEscape_Empty(t *testing.T) {
	input := ``

	output := shellEscape(input)
	if output != `''` {
		t.Errorf(`expected escaped string is "''", got %s`, output)
	}
}

func TestShellEscape_DoubleQuotes(t *testing.T) {
	input := `"test string"`

	output := shellEscape(input)
	if output != `'"test string"'` {
		t.Errorf(`expected escaped string is '"test string"', got %s`, output)
	}
}

func TestShellEscape_SingleQuotes(t *testing.T) {
	input := `'test string'`

	output := shellEscape(input)
	if output != `''"'"'test string'"'"''` {
		t.Errorf(`expected escaped string is ''"'"'test string'"'"'', got %s`, output)
	}
}

func TestShellEscape_NoQuotes(t *testing.T) {
	input := `test string`

	output := shellEscape(input)
	if output != `'test string'` {
		t.Errorf(`expected escaped string is "'test string'", got %s`, output)
	}
}

func TestAccShellEscape_Empty(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(version.Must(version.NewVersion("1.8.0"))),
		},
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: `
				locals {
					input = ""
				}
				output "new_string" {
					value = provider::string-functions::shell_escape(local.input)
				}
				`,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownOutputValue("new_string", knownvalue.StringExact(`''`)),
				},
			},
		},
	})
}

func TestAccShellEscape_DoubleQuotes(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(version.Must(version.NewVersion("1.8.0"))),
		},
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: `
				locals {
					input = "\"test string\""
				}
				output "new_string" {
					value = provider::string-functions::shell_escape(local.input)
				}
				`,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownOutputValue("new_string", knownvalue.StringExact(`'"test string"'`)),
				},
			},
		},
	})
}

func TestAccShellEscape_SingleQuotes(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(version.Must(version.NewVersion("1.8.0"))),
		},
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: `
				locals {
					input = "'test string'"
				}
				output "new_string" {
					value = provider::string-functions::shell_escape(local.input)
				}
				`,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownOutputValue("new_string", knownvalue.StringExact(`''"'"'test string'"'"''`)),
				},
			},
		},
	})
}

func TestAccShellEscape_NoQuotes(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(version.Must(version.NewVersion("1.8.0"))),
		},
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: `
				locals {
					input = "test string"
				}
				output "new_string" {
					value = provider::string-functions::shell_escape(local.input)
				}
				`,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownOutputValue("new_string", knownvalue.StringExact(`'test string'`)),
				},
			},
		},
	})
}
