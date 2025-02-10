package provider

import (
	"github.com/hashicorp/go-version"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfversion"
	"testing"
)

type FuncArgs struct {
	Input     string
	Delimiter string
	MaxLength int64
}

func TestCollapseEarlyExit(t *testing.T) {
	var inputs = map[FuncArgs]string{
		// Max length = input length
		{"abcdefghijklmnopqrstuvwxyz", "", 26}:    "abcdefghijklmnopqrstuvwxyz",
		{"abcdefghijklmnopqrstuvwxyz", "...", 26}: "abcdefghijklmnopqrstuvwxyz",
		// Max length = 0
		{"abcdefghijklmnopqrstuvwxyz", "", 0}:    "",
		{"abcdefghijklmnopqrstuvwxyz", "...", 0}: "",
		// Max length = 1
		{"abcdefghijklmnopqrstuvwxyz", "", 1}:    "…",
		{"abcdefghijklmnopqrstuvwxyz", "...", 1}: ".",
		// Max length > input length
		{"abcdefghij", "", 26}:    "abcdefghij",
		{"abcdefghij", "...", 26}: "abcdefghij",
		// Input length + delimiter length = max length
		{"abcdefghij", "", 11}: "abcdefghij",
		// Input length + delimiter length > max length, input length < max_length
		{"abcdefghij", "...", 12}: "abcdefghij",
	}

	for input, expected := range inputs {
		output, _ := collapseString(input.Input, input.Delimiter, input.MaxLength, Start)
		output2, _ := collapseString(output, input.Delimiter, input.MaxLength, End)
		output3, _ := collapseString(output2, input.Delimiter, input.MaxLength, Middle)
		if output != expected {
			t.Errorf("Expected %s, got %s", expected, output)
		}
		if output != output2 {
			t.Errorf("Expected %s, got %s", output, output2)
		}
		if output != output3 {
			t.Errorf("Expected %s, got %s", output, output3)
		}
	}
}

func TestCollapseStart(t *testing.T) {
	var inputs = map[FuncArgs]string{
		// Normal cases
		{"abcdefghijklmnopqrstuvwxyz", "", 10}:    "…rstuvwxyz",
		{"abcdefghijklmnopqrstuvwxyz", "...", 10}: "...tuvwxyz",
		// Delimiter length > input length > max length
		{"abcdef", ".......", 5}: ".....",
	}

	for input, expected := range inputs {
		output, _ := collapseString(input.Input, input.Delimiter, input.MaxLength, Start)
		output2, _ := collapseString(output, input.Delimiter, input.MaxLength, Start)
		if output != expected {
			t.Errorf("Expected %s, got %s", expected, output)
		}
		if output != output2 {
			t.Errorf("Expected %s, got %s", output, output2)
		}
	}
}

func TestCollapseMiddle(t *testing.T) {
	var inputs = map[FuncArgs]string{
		// Normal cases, even input length with even delimiter
		{"abcdefghijklmnopqrstuvwxyz", "..", 10}: "abcd..wxyz",
		// Normal cases, even input length with odd delimiter
		{"abcdefghijklmnopqrstuvwxyz", "", 10}:    "abcde…wxyz",
		{"abcdefghijklmnopqrstuvwxyz", "...", 10}: "abcd...xyz",
		// Normal cases, odd input length with even delimiter
		{"abcdefghijklmnopqrstuvwxy", "..", 10}: "abcd..vwxy",
		// Normal cases, odd input length with odd delimiter
		{"abcdefghijklmnopqrstuvwxy", "", 10}: "abcde…vwxy",
		// Delimiter length > input length > max length
		{"abcdef", ".......", 5}: ".....",
	}

	for input, expected := range inputs {
		output, _ := collapseString(input.Input, input.Delimiter, input.MaxLength, Middle)
		output2, _ := collapseString(output, input.Delimiter, input.MaxLength, Middle)
		if output != expected {
			t.Errorf("Expected %s, got %s", expected, output)
		}
		if output != output2 {
			t.Errorf("Expected %s, got %s", output, output2)
		}
	}
}

func TestCollapseEnd(t *testing.T) {
	var inputs = map[FuncArgs]string{
		// Normal cases
		{"abcdefghijklmnopqrstuvwxyz", "", 10}:    "abcdefghi…",
		{"abcdefghijklmnopqrstuvwxyz", "...", 10}: "abcdefg...",
		// Delimiter length > input length > max length
		{"abcdef", ".......", 5}: ".....",
	}

	for input, expected := range inputs {
		output, _ := collapseString(input.Input, input.Delimiter, input.MaxLength, End)
		output2, _ := collapseString(output, input.Delimiter, input.MaxLength, End)
		if output != expected {
			t.Errorf("Expected %s, got %s", expected, output)
		}
		if output != output2 {
			t.Errorf("Expected %s, got %s", output, output2)
		}
	}
}

func TestAccCollapseStart(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(version.Must(version.NewVersion("1.8.0"))),
		},
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: `
				locals {
					input = "hello_world"
				}
				output "collapsed_start" {
					value = provider::string-functions::collapse_start(local.input, "", 6)
				}
				`,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownOutputValue("collapsed_start", knownvalue.StringExact("…world")),
				},
			},
		},
	})
}

func TestAccCollapseMiddle(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(version.Must(version.NewVersion("1.8.0"))),
		},
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: `
				locals {
					input = "hello_world"
				}
				output "collapsed_middle" {
					value = provider::string-functions::collapse_middle(local.input, "", 6)
				}
				`,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownOutputValue("collapsed_middle", knownvalue.StringExact("hel…ld")),
				},
			},
		},
	})
}

func TestAccCollapseEnd(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(version.Must(version.NewVersion("1.8.0"))),
		},
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: `
				locals {
					input = "hello_world"
				}
				output "collapsed_end" {
					value = provider::string-functions::collapse_end(local.input, "", 6)
				}
				`,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownOutputValue("collapsed_end", knownvalue.StringExact("hello…")),
				},
			},
		},
	})
}
