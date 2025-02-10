package provider

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/function"
	"unicode/utf8"
)

type StringLocation int

const (
	Start StringLocation = iota
	Middle
	End
)

var _ function.Function = &CollapseStartFunction{}
var _ function.Function = &CollapseMiddleFunction{}
var _ function.Function = &CollapseEndFunction{}

type CollapseStartFunction struct{}
type CollapseMiddleFunction struct{}
type CollapseEndFunction struct{}

func NewCollapseStartFunction() function.Function {
	return &CollapseStartFunction{}
}

func NewCollapseMiddleFunction() function.Function {
	return &CollapseMiddleFunction{}
}

func NewCollapseEndFunction() function.Function {
	return &CollapseEndFunction{}
}

func (f *CollapseStartFunction) Metadata(ctx context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "collapse_start"
}

func (f *CollapseMiddleFunction) Metadata(ctx context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "collapse_middle"
}

func (f *CollapseEndFunction) Metadata(ctx context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "collapse_end"
}

func (f *CollapseStartFunction) Definition(ctx context.Context, req function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary:     "Collapse the start of a string",
		Description: "Collapse the start of a string to a specified delimiter",
		Parameters: []function.Parameter{
			function.StringParameter{
				Name:        "input",
				Description: "The string to collapse",
			},
			function.StringParameter{
				Name:        "delimiter",
				Description: "The delimiter to collapse the string to.",
			},
			function.Int64Parameter{
				Name:        "max_length",
				Description: "The maximum length of the string after collapsing. If the string is longer than this, it will be truncated.",
			},
		},
		Return: function.StringReturn{},
	}
}

func (f *CollapseMiddleFunction) Definition(ctx context.Context, req function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary:     "Collapse the middle of a string",
		Description: "Collapse the middle of a string to a specified delimiter",
		Parameters: []function.Parameter{
			function.StringParameter{
				Name:        "input",
				Description: "The string to collapse",
			},
			function.StringParameter{
				Name:        "delimiter",
				Description: "The delimiter to collapse the string to.",
			},
			function.Int64Parameter{
				Name:        "max_length",
				Description: "The maximum length of the string after collapsing. If the string is longer than this, it will be truncated.",
			},
		},
		Return: function.StringReturn{},
	}
}

func (f *CollapseEndFunction) Definition(ctx context.Context, req function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary:     "Collapse the end of a string",
		Description: "Collapse the end of a string to a specified delimiter",
		Parameters: []function.Parameter{
			function.StringParameter{
				Name:        "input",
				Description: "The string to collapse",
			},
			function.StringParameter{
				Name:        "delimiter",
				Description: "The delimiter to collapse the string to.",
			},
			function.Int64Parameter{
				Name:        "max_length",
				Description: "The maximum length of the string after collapsing. If the string is longer than this, it will be truncated.",
			},
		},
		Return: function.StringReturn{},
	}
}

func (f *CollapseStartFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	var input string
	var delimiter string
	var maxLength int64

	resp.Error = function.ConcatFuncErrors(req.Arguments.Get(ctx, &input, &delimiter, &maxLength))

	output, _ := collapseString(input, delimiter, maxLength, Start)

	resp.Error = function.ConcatFuncErrors(resp.Error, resp.Result.Set(ctx, output))
}

func (f *CollapseMiddleFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	var input string
	var delimiter string
	var maxLength int64

	resp.Error = function.ConcatFuncErrors(req.Arguments.Get(ctx, &input, &delimiter, &maxLength))

	output, _ := collapseString(input, delimiter, maxLength, Middle)

	resp.Error = function.ConcatFuncErrors(resp.Error, resp.Result.Set(ctx, output))
}

func (f *CollapseEndFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	var input string
	var delimiter string
	var maxLength int64

	resp.Error = function.ConcatFuncErrors(req.Arguments.Get(ctx, &input, &delimiter, &maxLength))

	output, _ := collapseString(input, delimiter, maxLength, End)

	resp.Error = function.ConcatFuncErrors(resp.Error, resp.Result.Set(ctx, output))
}

func collapseString(input string, delimiter string, maxLength int64, location StringLocation) (string, error) {
	runes := []rune(input)
	runesLength := int64(len(runes))

	// Set the default delimiter to Unicode ellipsis
	if utf8.RuneCountInString(delimiter) == 0 {
		delimiter = "…"
	}

	delimiterLength := int64(utf8.RuneCountInString(delimiter))

	if maxLength <= 0 {
		return "", nil
	}

	if runesLength <= maxLength {
		return input, nil
	}

	switch location {
	case Start:
		nRunes := maxLength - delimiterLength
		if nRunes < 0 {
			return delimiter[:maxLength], nil
		}
		return delimiter + string(runes[runesLength-nRunes:]), nil
	case Middle:
		nFrontRunes := (maxLength - delimiterLength + 1) / 2
		nBackRunes := maxLength - delimiterLength - nFrontRunes

		if nFrontRunes < 0 {
			nFrontRunes = 0
		}
		if nBackRunes < 0 {
			nBackRunes = 0
		}

		if nFrontRunes == 0 && nBackRunes == 0 {
			return delimiter[:maxLength], nil
		}
		return string(runes[:nFrontRunes]) + delimiter + string(runes[runesLength-nBackRunes:]), nil
	case End:
		nRunes := maxLength - delimiterLength
		if nRunes < 0 {
			return delimiter[:maxLength], nil
		}
		return string(runes[:nRunes]) + delimiter, nil
	}

	return "", function.NewFuncError("Invalid location")
}
