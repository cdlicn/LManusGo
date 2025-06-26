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

func NewLManus(name, systemMessage string, toolList []tools.Tool) *LManus {
	// 读取配置文件
	conf := config.Conf

	opts := []openai.Option{
		openai.WithBaseURL(conf.LLM.BaseUrl),
		openai.WithModel(conf.LLM.Model),
		openai.WithToken(conf.LLM.ApiKey),
	}

	newLLM, err := openai.New(opts...)
	if err != nil {
		log.Fatal(err)
	}

	agent := NewToolCallAgent(name, systemMessage, conf.LLM.Temperature, newLLM, 20, toolList)
	return &LManus{
		ToolCallAgent: agent,
	}

}
