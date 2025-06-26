package tools

import (
	"context"
	"github.com/tmc/langchaingo/llms"
	"time"
)

var _ Tool = &CurrentTime{}

type CurrentTime struct{}

func (t CurrentTime) Name() string {
	return "CurrentTime"
}

func (t CurrentTime) GetTool() llms.Tool {
	return llms.Tool{
		Type: "function",
		Function: &llms.FunctionDefinition{
			Name:        "CurrentTime",
			Description: "Get the current year month day hour minute second",
			Parameters:  nil,
		},
	}
}

func (t CurrentTime) Call(ctx context.Context, input string) string {
	return time.Now().Format("2006-01-02 15:04:05")
}
