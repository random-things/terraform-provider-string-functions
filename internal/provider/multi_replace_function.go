package provider

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"strings"
)

var _ function.Function = &MultiReplaceFunction{}

type MultiReplaceFunction struct{}

func NewMultiReplaceFunction() function.Function {
	return &MultiReplaceFunction{}
}

func (f *MultiReplaceFunction) Metadata(ctx context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "multi_replace"
}

func (f *MultiReplaceFunction) Definition(ctx context.Context, req function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary:     "Replace multiple substrings in a string",
		Description: "Uses a map to replace multiple substrings in a string.",
		Parameters: []function.Parameter{
			function.StringParameter{
				Name:        "input",
				Description: "The string in which to replace substrings",
			},
			function.MapParameter{
				Name:        "replacements",
				Description: "The map of substrings to replace and their replacements",
				ElementType: types.StringType,
			},
		},
		Return: function.StringReturn{},
	}
}

func (f *MultiReplaceFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	var input string
	var replacements map[string]string

	resp.Error = function.ConcatFuncErrors(req.Arguments.Get(ctx, &input, &replacements))

	pos := multiReplace(input, replacements)

	resp.Error = function.ConcatFuncErrors(resp.Error, resp.Result.Set(ctx, pos))
}

func multiReplace(input string, replacements map[string]string) string {
	for oldStr, newStr := range replacements {
		input = strings.ReplaceAll(input, oldStr, newStr)
	}
	return input
}
