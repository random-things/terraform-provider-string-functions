package provider

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"strings"
)

var _ function.Function = &ShellEscapeCmdFunction{}

type ShellEscapeCmdFunction struct{}

func NewShellEscapeCmdFunction() function.Function {
	return &ShellEscapeCmdFunction{}
}

func (f *ShellEscapeCmdFunction) Metadata(ctx context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "shell_escape_cmd"
}

func (f *ShellEscapeCmdFunction) Definition(ctx context.Context, req function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary:     "Escape a shell command provided as a list of strings",
		Description: "Escapes a list of strings containing a shell command and its arguments.",
		Parameters: []function.Parameter{
			function.ListParameter{
				Name:        "input",
				Description: "The shell command and its arguments to escape",
				ElementType: types.StringType,
			},
		},
		Return: function.StringReturn{},
	}
}

func (f *ShellEscapeCmdFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	var input []string

	resp.Error = function.ConcatFuncErrors(req.Arguments.Get(ctx, &input))

	output := shellEscapeCmd(input)

	resp.Error = function.ConcatFuncErrors(resp.Error, resp.Result.Set(ctx, output))
}

func shellEscapeCmd(input []string) string {
	args := make([]string, len(input))

	for i, arg := range input {
		args[i] = shellEscape(arg)
	}

	return strings.Join(args, " ")
}
