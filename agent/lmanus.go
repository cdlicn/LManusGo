package agent

import (
	"LManusGo/config"
	"LManusGo/tools"
	"github.com/tmc/langchaingo/llms/openai"
	"log"
)

type LManus struct {
	*ToolCallAgent
}

func NewLManus() *LManus {
	// 读取配置文件
	conf := config.Conf

	// 基本配置
	name := "LManus"
	systemMessage := `
		You are LManus, a versatile AI assistant designed to solve any task requested by users.
		You can use various tools to efficiently complete complex requests.
		Proactively select the most suitable tool or tool combination based on user needs.
		For complex tasks, you can break down the problem and gradually use different tools to solve it.
		After using each tool, clearly explain the execution results and suggest the next steps.
		If you want to stop interaction at any time, use the 'terminate' tool/function call.
		You only work with a single conversation, and you don't need to ask the user for any action after you end the conversation.
		If there is an error in the invoking tool, you can try again with a different parameter.
	`
	// 工具
	toolList := []tools.Tool{
		tools.DoTerminate{},
		tools.CurrentDate{},
		tools.CurrentTime{},
		tools.SaveFile{},
		tools.AccessURL{},
		tools.AskHuman{},
	}

	// 搜索引擎
	if conf.Base.SearchEngine != "" {
		toolList = append(toolList, tools.SearchWeb{})
	}

	// 创建LLM
	opts := []openai.Option{
		openai.WithBaseURL(conf.LLM.BaseUrl),
		openai.WithModel(conf.LLM.Model),
		openai.WithToken(conf.LLM.ApiKey),
	}
	newLLM, err := openai.New(opts...)
	if err != nil {
		log.Fatal(err)
	}

	// 最多步数
	maxSteps := 20

	// base agent
	baseAgent := NewBaseAgent(name, systemMessage, newLLM, maxSteps, config.Conf.MaxTokens, conf.LLM.Temperature)
	//reason act agent
	actAgent := NewReActAgent(baseAgent)
	// tool call agent
	agent := NewToolCallAgent(actAgent, toolList)

	return &LManus{
		ToolCallAgent: agent,
	}

}
