terraform {
  required_providers {
    string-functions = {
      source = "registry.hashicorp.io/random-things/string-functions"
    }
  }
  required_version = ">= 1.8.0"
}

provider "string-functions" {}

output "position" {
  value = provider::string-functions::strpos("abcdefghijklmnopqrstuvwxyz", "def")
}
