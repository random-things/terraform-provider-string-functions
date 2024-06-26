package provider

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"slices"
	"strings"
)

var _ function.Function = &LimitedRSplitFunction{}

type LimitedRSplitFunction struct{}

func NewLimitedRSplitFunction() function.Function {
	return &LimitedRSplitFunction{}
}

func (f *LimitedRSplitFunction) Metadata(ctx context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "limited_rsplit"
}

func (f *LimitedRSplitFunction) Definition(ctx context.Context, req function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary:     "Splits a string from the end using a delimiter a specified number of times",
		Description: "Splits a string from the end using a delimiter a specified number of times. The result is an array of strings.",
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

func (f *LimitedRSplitFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	var input string
	var delimiter string
	var timesToSplit int

	resp.Error = function.ConcatFuncErrors(req.Arguments.Get(ctx, &input, &delimiter, &timesToSplit))

	splitStrings := limitedRSplit(input, delimiter, timesToSplit)

	resp.Error = function.ConcatFuncErrors(resp.Error, resp.Result.Set(ctx, splitStrings))
}

func limitedRSplit(input string, delimiter string, timesToSplit int) []string {
	reversedInput := reverseString(input)
	splitStrings := strings.SplitN(reversedInput, delimiter, timesToSplit)

	for i, splitString := range splitStrings {
		splitStrings[i] = reverseString(splitString)
	}

	slices.Reverse(splitStrings)

	return splitStrings
}

func reverseString(input string) string {
	runes := []rune(input)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
