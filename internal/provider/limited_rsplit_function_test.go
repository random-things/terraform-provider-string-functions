package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
)

func TestLimitedRSplit(t *testing.T) {
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
			n:         3,
			expected:  []string{"this is", "a", "test"},
		},
		{
			name:      "no split",
			input:     "a,b,c",
			delimiter: ",",
			n:         1,
			expected:  []string{"a,b,c"},
		},
		{
			name:      "full split",
			input:     "a,b,c",
			delimiter: ",",
			n:         -1,
			expected:  []string{"a", "b", "c"},
		},
		{
			name:      "non-palindromic delimiter",
			input:     "a<->b<->c",
			delimiter: "<->",
			n:         2,
			expected:  []string{"a<->b", "c"},
		},
		{
			name:      "multi-byte delimiter",
			input:     "a🔥🔥b🔥🔥c",
			delimiter: "🔥🔥",
			n:         2,
			expected:  []string{"a🔥🔥b", "c"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := limitedRSplit(tt.input, tt.delimiter, tt.n)
			if len(actual) != len(tt.expected) {
				t.Fatalf("expected %d parts, got %d", len(tt.expected), len(actual))
			}
			for i := range actual {
				if actual[i] != tt.expected[i] {
					t.Errorf("part %d: expected %q, got %q", i, tt.expected[i], actual[i])
				}
			}
		})
	}
}

func TestAccLimitedRSplit_Known(t *testing.T) {
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
