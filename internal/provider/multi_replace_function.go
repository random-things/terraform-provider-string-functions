package provider

import (
	"context"
	"slices"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ function.Function = &MultiReplaceFunction{}
var _ function.Function = &MultiReplaceSortedFunction{}
var _ function.Function = &MultiReplaceSequentialFunction{}

type MultiReplaceFunction struct{}

func NewMultiReplaceFunction() function.Function {
	return &MultiReplaceFunction{}
}

func (*MultiReplaceFunction) Metadata(ctx context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "multi_replace"
}

func (*MultiReplaceFunction) Definition(ctx context.Context, req function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary:            "Replace multiple substrings in a string (Deprecated)",
		Description:        "Use multi_replace_sorted or multi_replace_sequential instead. This function uses a map to replace multiple substrings in a string, but the replacement order is non-deterministic.",
		DeprecationMessage: "Use multi_replace_sorted or multi_replace_sequential instead. This function uses a map to replace multiple substrings in a string, but the replacement order is non-deterministic.",
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
	if resp.Error != nil {
		return
	}

	output := multiReplace(input, replacements)

	resp.Error = function.ConcatFuncErrors(resp.Error, resp.Result.Set(ctx, output))
}

func multiReplace(input string, replacements map[string]string) string {
	for oldStr, newStr := range replacements {
		input = strings.ReplaceAll(input, oldStr, newStr)
	}
	return input
}

type MultiReplaceSortedFunction struct{}

func NewMultiReplaceSortedFunction() function.Function {
	return &MultiReplaceSortedFunction{}
}

func (*MultiReplaceSortedFunction) Metadata(ctx context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "multi_replace_sorted"
}

func (*MultiReplaceSortedFunction) Definition(ctx context.Context, req function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary:     "Replace multiple substrings in a string (Sorted)",
		Description: "Uses a map to replace multiple substrings in a string. Keys are sorted lexically before replacement to ensure deterministic behavior.",
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

func (f *MultiReplaceSortedFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	var input string
	var replacements map[string]string

	resp.Error = function.ConcatFuncErrors(req.Arguments.Get(ctx, &input, &replacements))
	if resp.Error != nil {
		return
	}

	output := multiReplaceSorted(input, replacements)

	resp.Error = function.ConcatFuncErrors(resp.Error, resp.Result.Set(ctx, output))
}

func multiReplaceSorted(input string, replacements map[string]string) string {
	keys := make([]string, 0, len(replacements))
	for oldStr := range replacements {
		keys = append(keys, oldStr)
	}

	slices.Sort(keys)

	for _, oldStr := range keys {
		input = strings.ReplaceAll(input, oldStr, replacements[oldStr])
	}
	return input
}

type MultiReplaceSequentialFunction struct{}

func NewMultiReplaceSequentialFunction() function.Function {
	return &MultiReplaceSequentialFunction{}
}

func (*MultiReplaceSequentialFunction) Metadata(ctx context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "multi_replace_sequential"
}

func (*MultiReplaceSequentialFunction) Definition(ctx context.Context, req function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary:     "Replace multiple substrings in a string (Sequential)",
		Description: "Uses a list of objects to replace multiple substrings in a string sequentially, preserving the order of the list.",
		Parameters: []function.Parameter{
			function.StringParameter{
				Name:        "input",
				Description: "The string in which to replace substrings",
			},
			function.ListParameter{
				Name:        "replacements",
				Description: "The list of replacements to apply in order. Each replacement is an object with 'from' and 'to' attributes.",
				ElementType: types.ObjectType{
					AttrTypes: map[string]attr.Type{
						"from": types.StringType,
						"to":   types.StringType,
					},
				},
			},
		},
		Return: function.StringReturn{},
	}
}

type replacementModel struct {
	From string `tfsdk:"from"`
	To   string `tfsdk:"to"`
}

func (f *MultiReplaceSequentialFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	var input string
	var replacements []replacementModel

	resp.Error = function.ConcatFuncErrors(req.Arguments.Get(ctx, &input, &replacements))
	if resp.Error != nil {
		return
	}

	output := multiReplaceSequential(input, replacements)

	resp.Error = function.ConcatFuncErrors(resp.Error, resp.Result.Set(ctx, output))
}

func multiReplaceSequential(input string, replacements []replacementModel) string {
	for _, r := range replacements {
		input = strings.ReplaceAll(input, r.From, r.To)
	}
	return input
}
