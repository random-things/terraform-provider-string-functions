# Changelog

Releases:

* [Unreleased](#unreleased)
* [v0.5.0](#v050)
* [v0.4.0](#v040)
* [v0.3.0](#v030)
* [v0.2.0](#v020)
* [v0.1.0](#v010)

## Unreleased

* ⬆️ Upgrade: golang.org/x/text `0.29.0 -> 0.30.0`
* ⬆️ Upgrade: hashicorp/terraform-plugin-framework `1.16.0 -> 1.16.1`
* ⬆️ Upgrade: hashicorp/terraform-plugin-framework `1.15.1 -> 1.16.0`
* ⬆️ Upgrade: hashicorp/terraform-plugin-go `0.28.0 -> 0.29.0`
* ⬆️ Upgrade: golang.org/x/text `0.28.0 -> 0.29.0`
* ⬆️ Upgrade: actions/setup-go `5.5.0 -> 6.0.0`
* ⬆️ Upgrade: hashicorp/terraform-plugin-testing `1.13.2 -> 1.13.3`
* ⬆️ Upgrade: goreleaser/goreleaser-action `6.3.0 -> 6.4.0`
* ⬆️ Upgrade: actions/checkout `4.2.2 -> 5.0.0`
* ⬆️ Upgrade: golang.org/x/text `0.27.0 -> 0.28.0`
* ⬆️ Upgrade: hashicorp/terraform-plugin-framework `1.15.0 -> 1.15.1`
* 🔧 Configure: Go toolchain version to `1.23.11`
* 🔧 Configure: Minimum Go build version to `1.23.7`
* ⬆️ Upgrade: golang.org/x/text `0.26.0 -> 0.27.0`
* 🧪 Add Terraform v1.12 to the testing matrix
* 🧪 Add Terraform v1.11 to the testing matrix
* ⬆️ Upgrade: hashicorp/terraform-plugin-testing `1.13.1 -> 1.13.2`
* ⬆️ Upgrade: golang.org/x/text `0.25.0 -> 0.26.0`
* ⬆️ Upgrade: hashicorp/terraform-plugin-framework `1.14.1 -> 1.15.0`
* ⬆️ Upgrade: hashicorp/terraform-plugin-testing `1.12.0 -> 1.13.1`
* ⬆️ Upgrade: hashicorp/terraform-plugin-go `0.26.0 -> 0.28.0`
* ⬆️ Upgrade: golang.org/x/text `0.24.0 -> 0.25.0`
* ⬆️ Upgrade: golangci/golangci-lint-action `7.0.0 -> 8.0.0`
* ⬆️ Upgrade: actions/setup-go `5.4.0 -> 5.5.0`
* ⬆️ Upgrade: golang.org/x/text `0.23.0 -> 0.24.0`
* ⬆️ Upgrade: goreleaser/goreleaser-action `6.2.1 -> 6.3.0`
* ⬆️ Upgrade: crazy-max/ghaction-import-gpg `6.2.0 -> 6.3.0`
* 🔧 Migrate: golangci-lint to v2
* ⬆️ Upgrade: golangci/golangci-lint-action `6.5.2 -> 7.0.0`
* ⬆️ Upgrade: hashicorp/terraform-plugin-testing `1.11.0 -> 1.12.0`
* ⬆️ Upgrade: actions/setup-go `5.3.0 -> 5.4.0`
* ⬆️ Upgrade: golangci/golangci-lint-action `6.5.1 -> 6.5.2`
* ⬆️ Upgrade: golangci/golangci-lint-action `6.5.0 -> 6.5.1`
* ⬆️ Upgrade: golang.org/x/text `0.22.0 -> 0.23.0`
* ⬆️ Upgrade: golang.org/x/text `0.21.0 -> 0.22.0`
* ⬆️ Upgrade: hashicorp/terraform-plugin-framework `1.13.0 -> 1.14.1`
* ⬆️ Upgrade: golangci/golangci-lint-action `6.4.1 -> 6.5.0`
* ⬆️ Upgrade: golangci/golangci-lint-action `6.3.2 -> 6.4.1`
* ⬆️ Upgrade: goreleaser/goreleaser-action `6.1.0 -> 6.2.1`
* ⬆️ Upgrade: golangci/golangci-lint-action `6.3.0 -> 6.3.2`

## v0.5.0

* ✨ Function: `collapse_start`
* ✨ Function: `collapse_middle`
* ✨ Function: `collapse_end`
* 🧪 Add Terraform v1.10 to the testing matrix
* 🧹 Exclude docs, examples, and CHANGELOG.md from automatic testing
* ⬆️ Upgrade: golangci/golangci-lint-action `6.2.0 -> 6.3.0`

## v0.4.0

FEATURES:

* ✨ Function: `camel_case`
* ✨ Function: `kebab_case`
* ✨ Function: `pascal_case`
* ✨ Function: `snake_case`
* ⬆️ Upgrade: terraform-plugin-go `0.25.0 -> 0.26.0`
* ⬆️ Upgrade: actions/setup-go `5.2.0 -> 5.3.0`
* ⬆️ Upgrade: golangci/golangci-lint-action `6.1.1 -> 6.2.0`

## v0.3.0

FEATURES:

* ✨ Function: `shell_escape`
* ✨ Function: `shell_escape_cmd`
* ✨ Function: `regex_escape`
* 📝 Documentation cleanup, function descriptions and examples in [README.md](README.md)
* ♻️ Refactor: Update references to `registry.hashicorp.io` to `registry.terraform.io`
* ♻️ Refactor: Rename `timesToSplit` to `maxParts` to clarify the parameter's purpose
in `limited_split` and `limited_rsplit`

## v0.2.0

FEATURES:

* ✨ Function: `multi_replace`

## v0.1.0

FEATURES:

* ✨ Function: `chunk_strings`
* ✨ Function: `limited_split`
* ✨ Function: `limited_rsplit`
* ✨ Function: `strpos`
* ✨ Function: `strrpos`
