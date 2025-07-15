package tools

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/tmc/langchaingo/llms"
)

var _ Tool = &AskHuman{}

type AskHuman struct{}

func (t AskHuman) Name() string {
	return "AskHuman"
}
func (t AskHuman) GetTool() llms.Tool {
	return llms.Tool{
		Type: "function",
		Function: &llms.FunctionDefinition{
			Name:        "AskHuman",
			Description: "Use this tool to ask human for help.",
			Parameters: map[string]any{
				"type": "object",
				"properties": map[string]any{
					"inquire": map[string]any{
						"type":        "string",
						"description": "The question you want to ask human.",
					},
				},
				"required": []string{"inquire"},
			},
		},
	}
}
func (t AskHuman) Call(ctx context.Context, input string) string {
	mp, err := unmarshallJson(input)
	if err != nil {
		return err.Error()
	}
	inquire := mp["inquire"]
	fmt.Println("Please answer: " + inquire)
	fmt.Println("Press enter to continue")
	var answer string
	_, err = fmt.Scanln(&answer)
	if err != nil {
		logrus.Panic("scanln failed:", err.Error())
	}
	return answer
}
