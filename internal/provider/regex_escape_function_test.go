package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
)

func TestRegExEscape(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "standard",
			input:    `\this is a test.`,
			expected: `\\this is a test\.`,
		},
		{
			name:     "special chars",
			input:    `^$.*+?()[]{}|`,
			expected: `\^\$\.\*\+\?\(\)\[\]\{\}\|`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := regExEscape(tt.input)
			if actual != tt.expected {
				t.Errorf("expected %q, got %q", tt.expected, actual)
			}
		})
	}
}

func TestAccRegExEscape_Known(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		TerraformVersionChecks:   terraform18OrNewer,
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: `
				locals {
					input = "\\this is a test."
				}
				output "new_string" {
					value = provider::string-functions::regex_escape(local.input)
				}
				`,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownOutputValue("new_string", knownvalue.StringExact(`\\this is a test\.`)),
				},
			},
		},
	})
}
