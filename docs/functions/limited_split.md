---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "limited_split function - string-functions"
subcategory: ""
description: |-
  Splits a string using a delimiter a specified number of times
---

# function: limited_split

Splits a string using a delimiter a specified number of times. The result is an array of strings.



## Signature

<!-- signature generated by tfplugindocs -->
```text
limited_split(input string, delimiter string, n number) list of string
```

## Arguments

<!-- arguments generated by tfplugindocs -->
1. `input` (String) The string to split
1. `delimiter` (String) The delimiter to use when splitting the string
1. `n` (Number) The maximum number of items to return in the result array. If n is less than 1, the entire string is returned as the first element of the result array.
