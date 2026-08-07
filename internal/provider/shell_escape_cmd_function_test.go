package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
)

func TestShellEscapeCmd(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		expected string
	}{
		{
			name:     "standard",
			input:    []string{"echo", "this is a test"},
			expected: `echo 'this is a test'`,
		},
		{
			name:     "multiple args",
			input:    []string{"ls", "-l", "/tmp"},
			expected: `ls -l /tmp`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := shellEscapeCmd(tt.input)
			if actual != tt.expected {
				t.Errorf("expected %q, got %q", tt.expected, actual)
			}
		})
	}
}

func TestAccShellEscapeCmd_Known(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		TerraformVersionChecks:   terraform18OrNewer,
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
