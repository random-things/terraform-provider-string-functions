package provider

import (
	"github.com/hashicorp/go-version"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfversion"
	"testing"
)

func TestCamelCase(t *testing.T) {
	var inputs = map[string]string{
		"":                "",
		"hello_world":     "helloWorld",
		"helloWorld":      "helloWorld",
		"hello_world_":    "helloWorld",
		"hello_world_1":   "helloWorld1",
		"hello_world_1_":  "helloWorld1",
		"hello_world_1_2": "helloWorld12",
		"HelloWorld":      "helloWorld",
		" hello_World\n":  "helloWorld",
		" hello-World\t":  "helloWorld",
		" hello world\r":  "helloWorld",
		"HELLO   world":   "helloWorld",
		"HELLO---WORLD":   "helloWorld",
		"hello___WORLD":   "helloWorld",
		"XMLHttpRequest":  "xmlHttpRequest",
		"XRequestId":      "xRequestId",
		"config2_test":    "config2Test",
	}

	for input, expected := range inputs {
		output := toCamelCase(input)
		output2 := toCamelCase(output)
		if output != expected {
			t.Errorf("Expected %s, got %s", expected, output)
		}
		if output != output2 {
			t.Errorf("Expected %s, got %s", output, output2)
		}
	}
}

func TestPascalCase(t *testing.T) {
	var inputs = map[string]string{
		"":                "",
		"hello_world":     "HelloWorld",
		"helloWorld":      "HelloWorld",
		"hello_world_":    "HelloWorld",
		"hello_world_1":   "HelloWorld1",
		"hello_world_1_":  "HelloWorld1",
		"hello_world_1_2": "HelloWorld12",
		"HelloWorld":      "HelloWorld",
		" hello_World\n":  "HelloWorld",
		" hello-World\t":  "HelloWorld",
		" hello world\r":  "HelloWorld",
		"HELLO   world":   "HelloWorld",
		"HELLO---WORLD":   "HelloWorld",
		"hello___WORLD":   "HelloWorld",
		"XMLHttpRequest":  "XmlHttpRequest",
		"XRequestId":      "XRequestId",
		"config2_test":    "Config2Test",
	}

	for input, expected := range inputs {
		output := toPascalCase(input)
		output2 := toPascalCase(output)
		if output != expected {
			t.Errorf("Expected %s, got %s", expected, output)
		}
		if output != output2 {
			t.Errorf("Expected %s, got %s", output, output2)
		}
	}
}

func TestAccCamelCase(t *testing.T) {
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
				output "camel_cased" {
					value = provider::string-functions::camel_case(local.input)
				}
				`,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownOutputValue("camel_cased", knownvalue.StringExact("helloWorld")),
				},
			},
		},
	})
}

func TestAccPascalCase(t *testing.T) {
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
				output "pascal_cased" {
					value = provider::string-functions::pascal_case(local.input)
				}
				`,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownOutputValue("pascal_cased", knownvalue.StringExact("HelloWorld")),
				},
			},
		},
	})
}
