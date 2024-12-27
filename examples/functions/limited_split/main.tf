terraform {
  required_providers {
    string-functions = {
      source = "registry.terraform.io/random-things/string-functions"
    }
  }
  required_version = ">= 1.8.0"
}

provider "string-functions" {}

output "split_string" {
  value = provider::string-functions::limited_split("this is a test string", " ", 3)
}
