package provider

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ function.Function = &ChunkStringsFunction{}

type ChunkStringsFunction struct{}

func NewChunkStringsFunction() function.Function {
	return &ChunkStringsFunction{}
}

func (*ChunkStringsFunction) Metadata(ctx context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "chunk_strings"
}

func (*ChunkStringsFunction) Definition(ctx context.Context, req function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary:     "Chunk a string into an array of smaller strings",
		Description: "Takes a list of strings, joins them with a delimiter, and splits them into chunks of a specified size.",
		Parameters: []function.Parameter{
			function.ListParameter{
				Name:        "inputs",
				Description: "The list of strings to chunk",
				ElementType: types.StringType,
			},
			function.Int64Parameter{
				Name:        "chunk_size",
				Description: "The maximum size of each chunk",
			},
			function.StringParameter{
				Name:        "delimiter",
				Description: "The delimiter to use when joining the strings",
			},
		},
		Return: function.ListReturn{
			ElementType: types.StringType,
		},
	}
}

func (*ChunkStringsFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	var inputs []string
	var chunkSize int64
	var delimiter string

	resp.Error = function.ConcatFuncErrors(req.Arguments.Get(ctx, &inputs, &chunkSize, &delimiter))
	if resp.Error != nil {
		return
	}

	chunks, err := chunkStrings(inputs, int(chunkSize), delimiter)
	if err != nil {
		resp.Error = function.ConcatFuncErrors(resp.Error, function.NewFuncError(err.Error()))
		return
	}

	resp.Error = function.ConcatFuncErrors(resp.Error, resp.Result.Set(ctx, chunks))
}

func chunkStrings(strs []string, chunkSize int, delimiter string) ([]string, error) {
	if chunkSize <= 0 {
		return nil, nil
	}

	chunks := make([]string, 0)
	current := ""

	for _, str := range strs {
		if len(str) > chunkSize {
			return nil, function.NewFuncError("input string is longer than chunk size")
		}

		if current == "" {
			current = str
			continue
		}

		candidate := current + delimiter + str
		if len(candidate) > chunkSize {
			chunks = append(chunks, current)
			current = str
			continue
		}

		current = candidate
	}

	if current != "" {
		chunks = append(chunks, current)
	}

	return chunks, nil
}
