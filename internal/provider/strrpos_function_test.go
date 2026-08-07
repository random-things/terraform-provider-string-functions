package provider

import (
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
)

func TestStrRPos(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		substring string
		expected  int
	}{
		{
			name:      "found",
			input:     "abcdefghijklmnopqrstuvwxyzdef",
			substring: "def",
			expected:  26,
		},
		{
			name:      "last occurrence",
			input:     "abcdabcd",
			substring: "abcd",
			expected:  4,
		},
		{
			name:      "not found",
			input:     "abcdefghijklmnopqrstuvwxyz",
			substring: "yz1",
			expected:  -1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := strRPos(tt.input, tt.substring)
			expected := strings.LastIndex(tt.input, tt.substring)
			if actual != expected {
				t.Fatalf("expected %d, got %d", expected, actual)
			}
			if actual != tt.expected {
				t.Errorf("expected %d, got %d", tt.expected, actual)
			}
		})
	}
}

func TestAccStrRPos_Known(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		TerraformVersionChecks:   terraform18OrNewer,
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
