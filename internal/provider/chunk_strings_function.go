package provider

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"strings"
)

var _ function.Function = &ChunkStringsFunction{}

type ChunkStringsFunction struct{}

func NewChunkStringsFunction() function.Function {
	return &ChunkStringsFunction{}
}

func (f *ChunkStringsFunction) Metadata(ctx context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "chunk_strings"
}

func (f *ChunkStringsFunction) Definition(ctx context.Context, req function.DefinitionRequest, resp *function.DefinitionResponse) {
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

func (f *ChunkStringsFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	var inputs []string
	var chunkSize int
	var delimiter string

	resp.Error = function.ConcatFuncErrors(req.Arguments.Get(ctx, &inputs, &chunkSize, &delimiter))

	chunks := chunkStrings(inputs, chunkSize, delimiter)

	resp.Error = function.ConcatFuncErrors(resp.Error, resp.Result.Set(ctx, chunks))
}

func chunkStrings(strs []string, chunkSize int, delimiter string) []string {
	chunks := []string{}
	breaks := []int{}

	del := ""
	if len(delimiter) > 0 && delimiter[0:1] != "" {
		del = delimiter[0:1]
	}

	chunkIndex := 0
	currentLength := 0
	for i := 0; i < len(strs); i++ {
		// TODO: Handle case where one string is longer than the chunkSize.
		//if len(strs[i]) > chunkSize {
		//
		//}

		currentLength += len(strs[i])
		if i == len(strs)-1 {
			breaks = append(breaks, i)
			break
		}
		nextLength := currentLength + len(strs[i+1]) + 1

		if nextLength > chunkSize {
			breaks = append(breaks, i+1)
			chunkIndex++
			currentLength = 0
		}
	}

	start := 0
	for i := 0; i < len(breaks); i++ {
		chunks = append(chunks, strings.Join(strs[start:breaks[i]], del))
		start = breaks[i]
	}

	return chunks
}
