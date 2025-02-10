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
  input_string = "hello_world"
}

output "output_string" {
  value = provider::string-functions::collapse_start(local.input_string, "", 6)
}
