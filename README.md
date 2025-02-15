# Contents

* [Functions](#functions)
  * [Case conversion](#case-conversion)
    * [camel_case](#camel_case)
    * [kebab_case](#kebab_case)
    * [pascal_case](#pascal_case)
    * [snake_case](#snake_case)
  * [Escaping](#escaping)
    * [regex_escape](#regex_escape)
    * [shell_escape](#shell_escape)
    * [shell_escape_cmd](#shell_escape_cmd)
  * [Replacement](#replacement)
    * [collapse_end](#collapse_end)
    * [collapse_middle](#collapse_middle)
    * [collapse_start](#collapse_start)
    * [multi_replace](#multi_replace)
  * [Splitting and chunking](#splitting-and-chunking)
    * [chunk_strings](#chunk_strings)
    * [limited_rsplit](#limited_rsplit)
    * [limited_split](#limited_split)
  * [Substring position](#substring-position)
    * [strpos](#strpos)
    * [strrpos](#strrpos)
* [Command line actions](#command-line-actions)
  * [Generating documentation](#generating-documentation)
  * [Linting](#linting)
  * [Running tests](#running-tests)
  * [Running acceptance tests](#running-acceptance-tests)
  * [Running an example](#running-an-example)
  * [Building for macOS (Apple Silicon)](#building-for-macos-apple-silicon)
  * [Building for Linux (x86_64)](#building-for-linux-x86_64)
  * [Building for Windows (x86_64)](#building-for-windows-x86_64)

# Functions

## Case conversion

### camel_case

`camel_case(input string) string`

Convert a string to camel case.

Example:

```hcl
locals {
  camel_cased = camel_case("hello_world")
}

output "camel_cased" {
  value = local.camel_cased
}

# camel_cased = "helloWorld"
```

### kebab_case

`kebab_case(input string) string`

Convert a string to kebab case.

Example:

```hcl
locals {
  kebab_cased = kebab_case("helloWorld")
}

output "kebab_cased" {
  value = local.kebab_cased
}

# kebab_cased = "hello-world"
```

### pascal_case

`pascal_case(input string) string`

Convert a string to Pascal case. Also known as Upper Camel Case.

Example:

```hcl
locals {
  pascal_cased = pascal_case("hello-world")
}

output "pascal_cased" {
  value = local.pascal_cased
}

# pascal_cased = "HelloWorld"
```

### snake_case

`snake_case(input string) string`

Convert a string to snake case.

Example:

```hcl
locals {
  snake_cased = snake_case("HelloWorld")
}

output "snake_cased" {
  value = local.snake_cased
}

# snake_cased = "hello_world"
```

## Escaping

### regex_escape

Escape a string containing regular expressions using Go's [`regexp.QuoteMeta`](https://pkg.go.dev/regexp#QuoteMeta).

`regex_escape(input string) string`

Example:

```hcl
locals {
  escaped = regex_escape("a.b.c")
}

output "escaped" {
  value = local.escaped
}

# escaped = "a\.b\.c"
# Without -raw, this looks like "a\\.b\\.c"
```

### shell_escape

Escape a string containing shell metacharacters.

`shell_escape(input string) string`

Example:

```hcl
locals {
  escaped = shell_escape("\"hi\"")
}

output "escaped" {
  value = local.escaped
}

# escaped = "'\"hi\"'"
```

### shell_escape_cmd

Escape a string containing shell metacharacters for use in a shell command.

`shell_escape_cmd(input list of string) string`

Example:

```hcl
locals {
  escaped = shell_escape_cmd(["echo", "hi there"])
}

output "escaped" {
  value = local.escaped
}

# escaped = "echo 'hi there'"
```

## Replacement

### collapse_end

Collapse the end of a string into a delimiter to return a string with a maximum
length. An empty delimiter will be treated as a Unicode ellipsis.

`collapse_end(input string, delimiter string, max_length int64) string`

Example:

```hcl

locals {
  collapsed_end = collapse_end("a,b,c,d,e", "", 6)
}

output "collapsed_end" {
  value = local.collapsed_end
}

# collapsed_end = "a,b,c…"
```

### collapse_middle

Collapse the middle of a string into a delimiter to return a string with a
maximum length. An empty delimiter will be treated as a Unicode ellipsis.

`collapse_middle(input string, delimiter string, max_length int64) string`

Example:

```hcl

locals {
  collapsed_middle = collapse_middle("a,b,c,d,e", "", 6)
}

output "collapsed_middle" {
  value = local.collapsed_middle
}

# collapsed_middle = "a,b…,e"
```

### collapse_start

Collapse the start of a string into a delimiter to return a string with a
maximum length. An empty delimiter will be treated as a Unicode ellipsis.

`collapse_start(input string, delimiter string, max_length int64) string`

Example:

```hcl

locals {
  collapsed_start = collapse_start("a,b,c,d,e", "...", 6)
}

output "collapsed_start" {
  value = local.collapsed_start
}

# collapsed_start = "...d,e"
```

### multi_replace

Replace multiple substrings in a string with other substrings.

`multi_replace(input string, replacements map of string to string) string`

Example:

```hcl

locals {
  replaced = multi_replace("a,b,c,d,e", {
    "," = "|",
    "a" = "z",
  })
}

output "replaced" {
  value = local.replaced
}

# replaced = "z|b|c|d|e"
```


## Splitting and chunking

### chunk_strings

`chunk_strings(inputs list of string, chunk_size number, delimiter string) list of string`

Chunk a string into an array of smaller strings joined by a delimiter. Note that `chunk_size`
represents the maximum character length of each chunk, not the number of items in the chunk.

Example:

```hcl
locals {
  chunked = chunk_strings(["a", "b", "c", "d", "e"], 2, ",")
}

output "chunked" {
  value = local.chunked
}

# chunked = ["a,", "b,", "c,", "d,", "e"]
```

### limited_rsplit

Split a string from the right into a list of strings, limited by a number of resultant parts.

`limited_rsplit(input string, delimiter string, limit number) list of string`

Example:

```hcl
locals {
  split = limited_rsplit("a,b,c,d,e", ",", 3)
}

output "split" {
  value = local.split
}

# split = ["a,b,c", "d", "e"]
```

### limited_split

Split a string into a list of strings, limited by a number of resultant parts.

`limited_split(input string, delimiter string, limit number) list of string`

Example:

```hcl
locals {
  split = limited_split("a,b,c,d,e", ",", 3)
}

output "split" {
  value = local.split
}

# split = ["a", "b", "c,d,e"]
```

## Substring position

### strpos

Find the position of the first occurrence of a substring in a string.

`strpos(input string, substring string) number`

Example:

```hcl
locals {
  position = strpos("a,b,c,d,e", ",")
}

output "position" {
  value = local.position
}

# position = 1
```

### strrpos

Find the position of the last occurrence of a substring in a string.

`strrpos(input string, substring string) number`

Example:

```hcl
locals {
  position = strrpos("a,b,c,d,e", ",")
}

output "position" {
  value = local.position
}

# position = 7
```

# Command line actions

## Generating documentation

```
go generate ./...
```

## Linting

```
golangci-lint run
```

## Running tests

```
go test -v terraform-provider-string-functions/internal/provider
```

## Running acceptance tests

```
cd internal/provider
$env:TF_ACC=1; go test -count=1 -run='TestAccChunkStrings' 
cd ../..
```

## Running an example

```
cd examples/functions/multi_replace
terraform init
terraform plan

Changes to Outputs:
  + output_string = "this was a trial"
```

## Building for macOS (Apple Silicon)

```powershell
$env:GOOS = "darwin"; $env:GOARCH = "arm64"; go build -o terraform-provider-string-functions_darwin_arm64
```

## Building for Linux (x86_64)

```bash
GOOS=linux GOARCH=amd64 go build -o terraform-provider-string-functions_linux_amd64
```

## Building for Windows (x86_64)

```bash
GOOS=windows GOARCH=amd64 go build -o terraform-provider-string-functions_windows_amd64.exe
```