terraform {
  required_providers {
    string-functions = {
      source = "registry.hashicorp.io/random-things/string-functions"
    }
  }
  required_version = ">= 1.8.0"
}

provider "string-functions" {}

locals {
  strings = [
    "string1",
    "string2",
    "string3",
    "string4",
    "string5",
    "string6",
    "string7",
    "string8",
    "string9",
    "string10",
    "string11",
    "string12",
    "string13",
    "string14",
    "string15",
    "string16",
    "string17",
    "string18",
    "string19",
    "string20",
    "string21",
    "string22",
    "string23",
    "string24",
    "string25",
    "string26",
    "string27",
    "string28",
    "string29",
    "string30",
    "string31",
    "string32",
    "string33",
    "string34",
    "string35",
    "string36",
  ]
}

output "chunked_string" {
  value = provider::string-functions::chunk_strings(local.strings, 100, "|")
}