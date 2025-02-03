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
  input_string = "HelloWorld"
}

output "output_string" {
  value = provider::string-functions::snake_case(local.input_string)
}
