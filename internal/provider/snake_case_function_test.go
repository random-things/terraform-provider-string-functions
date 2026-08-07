package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
)

func TestKebabCase(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"", ""},
		{"hello_world", "hello-world"},
		{"helloWorld", "hello-world"},
		{"hello_world_", "hello-world"},
		{"hello_world_1", "hello-world-1"},
		{"hello_world_1_", "hello-world-1"},
		{"hello_world_1_2", "hello-world-1-2"},
		{"HelloWorld", "hello-world"},
		{" hello_World\n", "hello-world"},
		{" hello-World\t", "hello-world"},
		{" hello world\r", "hello-world"},
		{"HELLO   world", "hello-world"},
		{"HELLO---WORLD", "hello-world"},
		{"hello___WORLD", "hello-world"},
		{"XMLHttpRequest", "xml-http-request"},
		{"XRequestId", "x-request-id"},
		{"config2_test", "config2-test"},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			actual := toKebabCase(tt.input)
			if actual != tt.expected {
				t.Errorf("expected %q, got %q", tt.expected, actual)
			}
		})
	}
}

func TestSnakeCase(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"", ""},
		{"hello_world", "hello_world"},
		{"helloWorld", "hello_world"},
		{"hello_world_", "hello_world"},
		{"hello_world_1", "hello_world_1"},
		{"hello_world_1_", "hello_world_1"},
		{"hello_world_1_2", "hello_world_1_2"},
		{"HelloWorld", "hello_world"},
		{" hello_World\n", "hello_world"},
		{" hello-World\t", "hello_world"},
		{" hello world\r", "hello_world"},
		{"HELLO   world", "hello_world"},
		{"HELLO---WORLD", "hello_world"},
		{"hello___WORLD", "hello_world"},
		{"XMLHttpRequest", "xml_http_request"},
		{"XRequestId", "x_request_id"},
		{"config2_test", "config2_test"},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			actual := toSnakeCase(tt.input)
			if actual != tt.expected {
				t.Errorf("expected %q, got %q", tt.expected, actual)
			}
		})
	}
}

func TestAccKebabCase(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		TerraformVersionChecks:   terraform18OrNewer,
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: `
				locals {
					input = "HelloWorld"
				}
				output "kebab_cased" {
					value = provider::string-functions::kebab_case(local.input)
				}
				`,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownOutputValue("kebab_cased", knownvalue.StringExact("hello-world")),
				},
			},
		},
	})
}

func TestAccSnakeCase(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		TerraformVersionChecks:   terraform18OrNewer,
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: `
				locals {
					input = "HelloWorld"
				}
				output "snake_cased" {
					value = provider::string-functions::snake_case(local.input)
				}
				`,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownOutputValue("snake_cased", knownvalue.StringExact("hello_world")),
				},
			},
		},
	})
}
