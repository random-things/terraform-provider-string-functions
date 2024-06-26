package provider

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"strings"
)

var _ function.Function = &LimitedSplitFunction{}

type LimitedSplitFunction struct{}

func NewLimitedSplitFunction() function.Function {
	return &LimitedSplitFunction{}
}

func (f *LimitedSplitFunction) Metadata(ctx context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "limited_split"
}

func (f *LimitedSplitFunction) Definition(ctx context.Context, req function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary:     "Splits a string using a delimiter a specified number of times",
		Description: "Splits a string using a delimiter a specified number of times. The result is an array of strings.",
		Parameters: []function.Parameter{
			function.StringParameter{
				Name:        "input",
				Description: "The string to split",
			},
			function.StringParameter{
				Name:        "delimiter",
				Description: "The delimiter to use when splitting the string",
			},
			function.Int64Parameter{
				Name:        "n",
				Description: "The maximum number of items to return in the result array. If n is less than 1, the entire string is returned as the first element of the result array.",
			},
		},
		Return: function.ListReturn{
			ElementType: types.StringType,
		},
	}
}

func (f *LimitedSplitFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	var input string
	var delimiter string
	var timesToSplit int

	resp.Error = function.ConcatFuncErrors(req.Arguments.Get(ctx, &input, &delimiter, &timesToSplit))

	splitStrings := limitedSplit(input, delimiter, timesToSplit)

	resp.Error = function.ConcatFuncErrors(resp.Error, resp.Result.Set(ctx, splitStrings))
}

func limitedSplit(input string, delimiter string, timesToSplit int) []string {
	return strings.SplitN(input, delimiter, timesToSplit)
}
