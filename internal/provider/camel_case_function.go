package provider

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/function"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"regexp"
	"strings"
	"unicode"
	"unicode/utf8"
)

var _ function.Function = &CamelCaseFunction{}
var _ function.Function = &PascalCaseFunction{}

type CamelCaseFunction struct{}
type PascalCaseFunction struct{}

func NewCamelCaseFunction() function.Function {
	return &CamelCaseFunction{}
}

func NewPascalCaseFunction() function.Function {
	return &PascalCaseFunction{}
}

func (f *CamelCaseFunction) Metadata(ctx context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "camel_case"
}

func (f *PascalCaseFunction) Metadata(ctx context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "pascal_case"
}

func (f *CamelCaseFunction) Definition(ctx context.Context, req function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary:     "Camel case a string",
		Description: "Camel case a string",
		Parameters: []function.Parameter{
			function.StringParameter{
				Name:        "input",
				Description: "The string to camel case",
			},
		},
		Return: function.StringReturn{},
	}
}

func (f *PascalCaseFunction) Definition(ctx context.Context, req function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary:     "Pascal case a string",
		Description: "Pascal case a string",
		Parameters: []function.Parameter{
			function.StringParameter{
				Name:        "input",
				Description: "The string to Pascal case",
			},
		},
		Return: function.StringReturn{},
	}
}

func (f *CamelCaseFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	var input string

	resp.Error = function.ConcatFuncErrors(req.Arguments.Get(ctx, &input))

	output := toCamelCase(input)

	resp.Error = function.ConcatFuncErrors(resp.Error, resp.Result.Set(ctx, output))
}

func (f *PascalCaseFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	var input string

	resp.Error = function.ConcatFuncErrors(req.Arguments.Get(ctx, &input))

	output := toPascalCase(input)

	resp.Error = function.ConcatFuncErrors(resp.Error, resp.Result.Set(ctx, output))
}

func splitCamelCase(s string) []string {
	var words []string
	var currentWord []rune
	runes := []rune(s)

	for i, r := range runes {
		if i > 0 && unicode.IsUpper(r) {
			prev := runes[i-1]
			if unicode.IsLower(prev) {
				words = append(words, string(currentWord))
				currentWord = []rune{r}
				continue
			}
			if i+1 < len(runes) && unicode.IsLower(runes[i+1]) {
				words = append(words, string(currentWord))
				currentWord = []rune{r}
				continue
			}
		}
		currentWord = append(currentWord, r)
	}
	if len(currentWord) > 0 {
		words = append(words, string(currentWord))
	}
	return words
}

func toCamelCase(s string) string {
	if s == "" {
		return s
	}

	var words []string
	var lower = cases.Lower(language.Und)
	var title = cases.Title(language.Und)

	if strings.ContainsAny(s, "_- ") {
		re := regexp.MustCompile(`[A-Za-z0-9]+`)
		words = re.FindAllString(s, -1)
	} else {
		words = splitCamelCase(s)
	}

	if len(words) == 0 {
		return ""
	}

	words[0] = lower.String(words[0])
	for i := 1; i < len(words); i++ {
		words[i] = title.String(lower.String(words[i]))
	}

	return strings.Join(words, "")
}

func toPascalCase(input string) string {
	if input == "" {
		return input
	}

	camelCased := toCamelCase(input)
	upper := cases.Upper(language.Und)
	r, size := utf8.DecodeRuneInString(camelCased)

	return upper.String(string(r)) + camelCased[size:]
}
