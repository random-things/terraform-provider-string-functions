package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
)

func TestChunkStrings(t *testing.T) {
	tests := []struct {
		name      string
		inputs    []string
		chunkSize int
		delimiter string
		expected  []string
		expectErr bool
	}{
		{
			name:      "standard",
			inputs:    []string{"a", "b", "c"},
			chunkSize: 5,
			delimiter: "|",
			expected:  []string{"a|b|c"},
		},
		{
			name:      "multiple chunks",
			inputs:    []string{"a", "b", "c"},
			chunkSize: 1,
			delimiter: "|",
			expected:  []string{"a", "b", "c"},
		},
		{
			name:      "exact fit",
			inputs:    []string{"aa", "bb"},
			chunkSize: 5,
			delimiter: "|",
			expected:  []string{"aa|bb"},
		},
		{
			name:      "just over fit",
			inputs:    []string{"aa", "bb", "cc"},
			chunkSize: 5,
			delimiter: "|",
			expected:  []string{"aa|bb", "cc"},
		},
		{
			name:      "multi-byte delimiter",
			inputs:    []string{"a", "b", "c"},
			chunkSize: 20,
			delimiter: "🔥🔥",
			expected:  []string{"a🔥🔥b🔥🔥c"},
		},
		{
			name:      "oversized string",
			inputs:    []string{"abcde"},
			chunkSize: 3,
			delimiter: "|",
			expectErr: true,
		},
		{
			name:      "empty input",
			inputs:    []string{},
			chunkSize: 10,
			delimiter: "|",
			expected:  []string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual, err := chunkStrings(tt.inputs, tt.chunkSize, tt.delimiter)
			if (err != nil) != tt.expectErr {
				t.Fatalf("expected error: %v, got: %v", tt.expectErr, err)
			}
			if err != nil {
				return
			}
			if len(actual) != len(tt.expected) {
				t.Fatalf("expected %d chunks, got %d", len(tt.expected), len(actual))
			}
			for i := range actual {
				if actual[i] != tt.expected[i] {
					t.Errorf("chunk %d: expected %q, got %q", i, tt.expected[i], actual[i])
				}
			}
		})
	}
}

func TestAccChunkStrings_Known(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		TerraformVersionChecks:   terraform18OrNewer,
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: `
				locals {
					strings = ["this", "is", "a", "test", "1234567890", "12345678901", "123456789012", "1234567890123", "12345678901234", "123456789012345", "1234567890123456", "12345678901234567", "123456789012345678", "1234567890123456789", "12345678901234567890"]
				}
				output "chunks" {
					value = provider::string-functions::chunk_strings(local.strings, 100, "|")
				}
				`,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownOutputValue("chunks", knownvalue.ListExact([]knownvalue.Check{
						knownvalue.StringExact("this|is|a|test|1234567890|12345678901|123456789012|1234567890123|12345678901234|123456789012345"),
						knownvalue.StringExact("1234567890123456|12345678901234567|123456789012345678|1234567890123456789|12345678901234567890"),
					})),
				},
			},
		},
	})
}
