terraform {
  required_providers {
    string-functions = {
      source = "registry.hashicorp.io/random-things/string-functions"
    }
  }
  required_version = ">= 1.8.0"
}

provider "string-functions" {}

output "split_string" {
  value = provider::string-functions::limited_rsplit("this is a test string", " ", 3)
}
