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
  input_string = "helloWorld"
}

output "output_string" {
  value = provider::string-functions::kebab_case(local.input_string)
}
