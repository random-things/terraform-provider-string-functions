package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
)

func TestMultiReplaceSorted(t *testing.T) {
	tests := []struct {
		name         string
		input        string
		replacements map[string]string
		expected     string
	}{
		{
			name:  "standard",
			input: "this is a test",
			replacements: map[string]string{
				" is":   " was",
				" test": " trial",
			},
			expected: "this was a trial",
		},
		{
			name:  "lexical sorting ensures deterministic nested replacement",
			input: "FOO",
			replacements: map[string]string{
				"FOO": "BAR",
				"BAR": "BAZ",
			},
			// "BAR" comes before "FOO" alphabetically.
			// 1. input = strings.ReplaceAll("FOO", "BAR", "BAZ") -> "FOO"
			// 2. input = strings.ReplaceAll("FOO", "FOO", "BAR") -> "BAR"
			expected: "BAR",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := multiReplaceSorted(tt.input, tt.replacements)
			if actual != tt.expected {
				t.Errorf("expected %q, got %q", tt.expected, actual)
			}
		})
	}
}

func TestMultiReplaceSequential(t *testing.T) {
	tests := []struct {
		name         string
		input        string
		replacements []replacementModel
		expected     string
	}{
		{
			name:  "sequential nested replacement",
			input: "FOO",
			replacements: []replacementModel{
				{From: "FOO", To: "BAR"},
				{From: "BAR", To: "BAZ"},
			},
			// 1. "FOO" -> "BAR"
			// 2. "BAR" -> "BAZ"
			expected: "BAZ",
		},
		{
			name:  "reverse sequential nested replacement",
			input: "FOO",
			replacements: []replacementModel{
				{From: "BAR", To: "BAZ"},
				{From: "FOO", To: "BAR"},
			},
			// 1. "FOO" -> "FOO" (no BAR found)
			// 2. "FOO" -> "BAR"
			expected: "BAR",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := multiReplaceSequential(tt.input, tt.replacements)
			if actual != tt.expected {
				t.Errorf("expected %q, got %q", tt.expected, actual)
			}
		})
	}
}

func TestAccMultiReplace_Known(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		TerraformVersionChecks:   terraform18OrNewer,
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: `
				locals {
					input = "this is a test"
                    replacements = {
						" is" = " was"
						" test" = " trial"
                    }
				}
				output "deprecated" {
					value = provider::string-functions::multi_replace(local.input, local.replacements)
				}
				output "sorted" {
					value = provider::string-functions::multi_replace_sorted(local.input, local.replacements)
				}
				output "sequential" {
					value = provider::string-functions::multi_replace_sequential(local.input, [
						{ from = " is", to = " was" },
						{ from = " test", to = " trial" },
					])
				}
				`,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownOutputValue("deprecated", knownvalue.StringExact("this was a trial")),
					statecheck.ExpectKnownOutputValue("sorted", knownvalue.StringExact("this was a trial")),
					statecheck.ExpectKnownOutputValue("sequential", knownvalue.StringExact("this was a trial")),
				},
			},
			{
				Config: `
				output "nested_sequential" {
					value = provider::string-functions::multi_replace_sequential("FOO", [
						{ from = "FOO", to = "BAR" },
						{ from = "BAR", to = "BAZ" },
					])
				}
				output "nested_sorted" {
					value = provider::string-functions::multi_replace_sorted("FOO", {
						"FOO" = "BAR"
						"BAR" = "BAZ"
					})
				}
				`,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownOutputValue("nested_sequential", knownvalue.StringExact("BAZ")),
					statecheck.ExpectKnownOutputValue("nested_sorted", knownvalue.StringExact("BAR")),
				},
			},
		},
	})
}
