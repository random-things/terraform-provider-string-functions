package provider

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/function"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"regexp"
	"strings"
)

var _ function.Function = &KebabCaseFunction{}
var _ function.Function = &SnakeCaseFunction{}

type KebabCaseFunction struct{}
type SnakeCaseFunction struct{}

func NewKebabCaseFunction() function.Function {
	return &KebabCaseFunction{}
}

func NewSnakeCaseFunction() function.Function {
	return &SnakeCaseFunction{}
}

func (f *KebabCaseFunction) Metadata(ctx context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "kebab_case"
}

func (f *SnakeCaseFunction) Metadata(ctx context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "snake_case"
}

func (f *KebabCaseFunction) Definition(ctx context.Context, req function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary:     "Kebab case a string",
		Description: "Kebab case a string",
		Parameters: []function.Parameter{
			function.StringParameter{
				Name:        "input",
				Description: "The string to kebab case",
			},
		},
		Return: function.StringReturn{},
	}
}

func (f *SnakeCaseFunction) Definition(ctx context.Context, req function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary:     "Snake case a string",
		Description: "Snake case a string",
		Parameters: []function.Parameter{
			function.StringParameter{
				Name:        "input",
				Description: "The string to snake case",
			},
		},
		Return: function.StringReturn{},
	}
}

func (f *KebabCaseFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	var input string

	resp.Error = function.ConcatFuncErrors(req.Arguments.Get(ctx, &input))

	output := toKebabCase(input)

	resp.Error = function.ConcatFuncErrors(resp.Error, resp.Result.Set(ctx, output))
}

func (f *SnakeCaseFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	var input string

	resp.Error = function.ConcatFuncErrors(req.Arguments.Get(ctx, &input))

	output := toSnakeCase(input)

	resp.Error = function.ConcatFuncErrors(resp.Error, resp.Result.Set(ctx, output))
}

func toSeparatedCase(s string, delimiter string) string {
	if s == "" {
		return s
	}

	var words []string
	var lower = cases.Lower(language.Und)

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
	for i, word := range words {
		words[i] = lower.String(word)
	}

	return strings.Join(words, delimiter)
}

func toKebabCase(input string) string {
	return toSeparatedCase(input, "-")
}

func toSnakeCase(input string) string {
	return toSeparatedCase(input, "_")
}
