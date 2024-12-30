package provider

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/function"
	"regexp"
)

var _ function.Function = &RegExEscapeFunction{}

type RegExEscapeFunction struct{}

func NewRegExEscapeFunction() function.Function {
	return &RegExEscapeFunction{}
}

func (f *RegExEscapeFunction) Metadata(ctx context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "regex_escape"
}

func (f *RegExEscapeFunction) Definition(ctx context.Context, req function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary:     "Escape a regular expression-containing string",
		Description: "Escapes a string containing regular expressions.",
		Parameters: []function.Parameter{
			function.StringParameter{
				Name:        "input",
				Description: "The string in which to escape regular expressions",
			},
		},
		Return: function.StringReturn{},
	}
}

func (f *RegExEscapeFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	var input string

	resp.Error = function.ConcatFuncErrors(req.Arguments.Get(ctx, &input))

	output := regExEscape(input)

	resp.Error = function.ConcatFuncErrors(resp.Error, resp.Result.Set(ctx, output))
}

func regExEscape(input string) string {
	return regexp.QuoteMeta(input)
}
