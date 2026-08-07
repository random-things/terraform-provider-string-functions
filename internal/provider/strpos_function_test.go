package provider

import (
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
)

func TestStrPos(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		substring string
		expected  int
	}{
		{
			name:      "found",
			input:     "abcdefghijklmnopqrstuvwxyz",
			substring: "def",
			expected:  3,
		},
		{
			name:      "not found",
			input:     "abcdefghijklmnopqrstuvwxyz",
			substring: "yz1",
			expected:  -1,
		},
		{
			name:      "empty input",
			input:     "",
			substring: "a",
			expected:  -1,
		},
		{
			name:      "empty substring",
			input:     "abc",
			substring: "",
			expected:  0,
		},
		{
			name:      "unicode",
			input:     "hello 🔥 world",
			substring: "🔥",
			expected:  6, // bytes
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := strPos(tt.input, tt.substring)
			expected := strings.Index(tt.input, tt.substring)
			if actual != expected {
				t.Fatalf("expected %d, got %d", expected, actual)
			}
			if actual != tt.expected {
				t.Errorf("expected %d, got %d", tt.expected, actual)
			}
		})
	}
}

func TestAccStrPos_Known(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		TerraformVersionChecks:   terraform18OrNewer,
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
