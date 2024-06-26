package provider

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/function"
	"strings"
)

var _ function.Function = &StrPosFunction{}

type StrPosFunction struct{}

func NewStrPosFunction() function.Function {
	return &StrPosFunction{}
}

func (f *StrPosFunction) Metadata(ctx context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "strpos"
}

func (f *StrPosFunction) Definition(ctx context.Context, req function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary:     "Find the position of a substring in a string",
		Description: "Returns the position of a substring in a string. If the substring is not found, returns -1.",
		Parameters: []function.Parameter{
			function.StringParameter{
				Name:        "input",
				Description: "The string in which to find the substring position",
			},
			function.StringParameter{
				Name:        "substring",
				Description: "The substring whose position to find in the input string",
			},
		},
		Return: function.Int64Return{},
	}
}

func (f *StrPosFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	var input string
	var substring string

	resp.Error = function.ConcatFuncErrors(req.Arguments.Get(ctx, &input, &substring))

	pos := strPos(input, substring)

	resp.Error = function.ConcatFuncErrors(resp.Error, resp.Result.Set(ctx, pos))
}

func strPos(input string, substring string) int {
	return strings.Index(input, substring)
}
