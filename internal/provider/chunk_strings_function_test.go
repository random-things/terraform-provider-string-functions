package provider

import (
	"github.com/hashicorp/go-version"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfversion"
	"testing"
)

func TestChunkStrings(t *testing.T) {
	inputs := []string{"this", "is", "a", "test", "1234567890", "12345678901", "123456789012", "1234567890123", "12345678901234", "123456789012345", "1234567890123456", "12345678901234567", "123456789012345678", "1234567890123456789", "12345678901234567890"}
	chunkSize := 100
	delimiter := "|"

	chunks := chunkStrings(inputs, chunkSize, delimiter)
	for _, chunk := range chunks {
		if len(chunk) > chunkSize {
			t.Errorf("chunk size %d is greater than chunkSize %d", len(chunk), chunkSize)
		}
	}
}

func TestAccChunkStrings_Known(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(version.Must(version.NewVersion("1.8.0"))),
		},
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
						knownvalue.StringExact("1234567890123456|12345678901234567|123456789012345678|1234567890123456789"),
					})),
				},
			},
		},
	})
}
