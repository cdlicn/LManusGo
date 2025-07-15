package tools

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/tmc/langchaingo/llms"
)

type ToolsMap map[string]Tool

type Tool interface {
	Name() string
	GetTool() llms.Tool
	Call(ctx context.Context, input string) string
}

func NewToolsMap(toolList ...Tool) (toolsMap ToolsMap, err error) {
	toolsMap = make(ToolsMap)
	for _, tool := range toolList {
		if _, ok := toolsMap[tool.Name()]; ok {
			return nil, errors.New("duplicate method names")
		}
		toolsMap[tool.Name()] = tool
	}
	return toolsMap, nil
}

func (toolsMap ToolsMap) ExecuteTool(toolName string, input string) (res string, err error) {
	if _, ok := toolsMap[toolName]; !ok {
		return "", errors.New("tool not found")
	}
	return toolsMap[toolName].Call(context.Background(), input), nil
}

func unmarshallJson(input string) (map[string]string, error) {
	var mp map[string]string
	err := json.Unmarshal([]byte(input), &mp)
	if err != nil {
		return nil, errors.New("there was an error parsing the input " + err.Error())
	}
	return mp, nil
}
