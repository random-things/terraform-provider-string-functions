# Changelog

Releases:

* [Unreleased](#unreleased)
* [v0.4.0](#v040)
* [v0.3.0](#v030)
* [v0.2.0](#v020)
* [v0.1.0](#v010)

## Unreleased

* ⬆️ Upgrade: golangci/golangci-lint-action `6.2.1 -> 6.3.0`

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
