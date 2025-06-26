package tools

import (
	"context"
	"github.com/tmc/langchaingo/llms"
)

var _ Tool = &DoTerminate{}

type DoTerminate struct{}

func (t DoTerminate) Name() string {
	return "DoTerminate"
}
func (t DoTerminate) GetTool() llms.Tool {
	return llms.Tool{
		Type: "function",
		Function: &llms.FunctionDefinition{
			Name: "DoTerminate",
			Description: `Terminate the interaction when the request is met OR if the assistant cannot proceed further with the task.
							When you have finished all the tasks, call this tool to end the work.`,
		},
	}
}

func (t DoTerminate) Call(ctx context.Context, input string) string {
	return "任务结束"
}
