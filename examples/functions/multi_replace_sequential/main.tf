terraform {
  required_providers {
    string-functions = {
      source = "registry.terraform.io/random-things/string-functions"
    }
  }
  required_version = ">= 1.8.0"
}

provider "string-functions" {}

locals {
  input_string = "FOO"
  replacements = [
    { from = "FOO", to = "BAR" },
    { from = "BAR", to = "BAZ" },
  ]
}

output "output_string" {
  value = provider::string-functions::multi_replace_sequential(local.input_string, local.replacements)
}

# output_string = "BAZ"
