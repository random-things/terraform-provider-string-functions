package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
)

func TestShellEscape(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "empty",
			input:    "",
			expected: "''",
		},
		{
			name:     "double quotes",
			input:    `"test string"`,
			expected: `'"test string"'`,
		},
		{
			name:     "single quotes",
			input:    `'test string'`,
			expected: `''"'"'test string'"'"''`,
		},
		{
			name:     "no quotes",
			input:    `test string`,
			expected: `'test string'`,
		},
		{
			name:     "safe characters",
			input:    `abc/def.ghi`,
			expected: `abc/def.ghi`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := shellEscape(tt.input)
			if actual != tt.expected {
				t.Errorf("expected %q, got %q", tt.expected, actual)
			}
		})
	}
}

func TestAccShellEscape_Empty(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		TerraformVersionChecks:   terraform18OrNewer,
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
		TerraformVersionChecks:   terraform18OrNewer,
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
		TerraformVersionChecks:   terraform18OrNewer,
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
		TerraformVersionChecks:   terraform18OrNewer,
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
