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
  input = ["echo", "test string"]
}

output "output_string" {
  value = provider::string-functions::shell_escape_cmd(local.input)
}
