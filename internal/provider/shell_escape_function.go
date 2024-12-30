package provider

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/function"
	"regexp"
	"strings"
)

var _ function.Function = &ShellEscapeFunction{}

type ShellEscapeFunction struct{}

func NewShellEscapeFunction() function.Function {
	return &ShellEscapeFunction{}
}

func (f *ShellEscapeFunction) Metadata(ctx context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "shell_escape"
}

func (f *ShellEscapeFunction) Definition(ctx context.Context, req function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary:     "Escape a shell metacharacter-containing string",
		Description: "Escapes a string containing shell metacharacters.",
		Parameters: []function.Parameter{
			function.StringParameter{
				Name:        "input",
				Description: "The string in which to escape shell metacharacters",
			},
		},
		Return: function.StringReturn{},
	}
}

func (f *ShellEscapeFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	var input string

	resp.Error = function.ConcatFuncErrors(req.Arguments.Get(ctx, &input))

	output := shellEscape(input)

	resp.Error = function.ConcatFuncErrors(resp.Error, resp.Result.Set(ctx, output))
}

func shellEscape(input string) string {
	// Essentially shlex.quote() in Python
	pattern := regexp.MustCompile(`[^\w@%+=:,./-]`)

	if len(input) == 0 {
		return "''"
	}

	if !pattern.MatchString(input) {
		return input
	}

	return "'" + strings.ReplaceAll(input, "'", "'\"'\"'") + "'"
}
