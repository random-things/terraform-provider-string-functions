terraform {
  required_providers {
    string-functions = {
      source = "registry.terraform.io/random-things/string-functions"
    }
  }
  required_version = ">= 1.8.0"
}

provider "string-functions" {}

output "position" {
  value = provider::string-functions::strrpos("abcdefghijklmnopqrstuvwxyzabcdef", "def")
}
