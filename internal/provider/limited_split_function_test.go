package provider

import (
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
)

func TestLimitedSplit(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		delimiter string
		n         int
		expected  []string
	}{
		{
			name:      "standard",
			input:     "this is a test",
			delimiter: " ",
			n:         2,
			expected:  []string{"this", "is a test"},
		},
		{
			name:      "no limit",
			input:     "a,b,c",
			delimiter: ",",
			n:         -1,
			expected:  []string{"a", "b", "c"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := limitedSplit(tt.input, tt.delimiter, tt.n)
			expected := strings.SplitN(tt.input, tt.delimiter, tt.n)
			if len(actual) != len(expected) {
				t.Fatalf("expected %d parts, got %d", len(expected), len(actual))
			}
			for i := range actual {
				if actual[i] != expected[i] {
					t.Errorf("part %d: expected %q, got %q", i, expected[i], actual[i])
				}
			}
		})
	}
}

func TestAccLimitedSplit_Known(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		TerraformVersionChecks:   terraform18OrNewer,
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: `
				locals {
					input = "this is a test"
				}
				output "parts" {
					value = provider::string-functions::limited_split(local.input, " ", 3)
				}
				`,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownOutputValue("parts", knownvalue.ListExact([]knownvalue.Check{
						knownvalue.StringExact("this"),
						knownvalue.StringExact("is"),
						knownvalue.StringExact("a test"),
					})),
				},
			},
		},
	})
}
