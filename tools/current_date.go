package tools

import (
	"context"
	"github.com/tmc/langchaingo/llms"
	"time"
)

var _ Tool = &CurrentDate{}

type CurrentDate struct{}

func (t CurrentDate) Name() string {
	return "CurrentDate"
}

func (t CurrentDate) GetTool() llms.Tool {
	return llms.Tool{
		Type: "function",
		Function: &llms.FunctionDefinition{
			Name:        "CurrentDate",
			Description: "This is a method to obtain the current year, month, and day in the format yyyy-MM-DD, without any input parameters required.",
			Parameters:  nil,
		},
	}
}

func (t CurrentDate) Call(ctx context.Context, input string) string {
	return time.Now().Format("2006-01-02")
}
