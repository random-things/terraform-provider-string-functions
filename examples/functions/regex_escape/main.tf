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
  input_string = "\\this is a test."
}

output "output_string" {
  value = provider::string-functions::regex_escape(local.input_string)
}
