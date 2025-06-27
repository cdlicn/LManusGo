package tools

import (
	"LManusGo/tools/search"
	"context"
	"encoding/json"
	"fmt"
	"github.com/tmc/langchaingo/llms"
)

var _ Tool = &SearchWeb{}

type SearchWeb struct{}

func (t SearchWeb) Name() string {
	return "SearchWeb"
}
func (t SearchWeb) GetTool() llms.Tool {
	return llms.Tool{
		Type: "function",
		Function: &llms.FunctionDefinition{
			Name:        "SearchWeb",
			Description: `Search for information from Baidu search engine. Don't use Markdown syntax, use plain text.`,
			Parameters: map[string]any{
				"type": "object",
				"properties": map[string]any{
					"query": map[string]any{
						"type":        "string",
						"description": "keyword that need to be searched",
					},
				},
				"required": []string{"query"},
			},
		},
	}
}

func (t SearchWeb) Call(ctx context.Context, input string) string {
	var mp map[string]string
	err := json.Unmarshal([]byte(input), &mp)
	if err != nil {
		return "failed to unmarshal JSON" + err.Error()
	}

	items, err := search.SearchEngine.Call(mp["query"])

	return fmt.Sprintln(items)
}
