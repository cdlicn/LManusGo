package agent

import (
	"LManusGo/message"
	"LManusGo/tools"
	"context"
	"errors"
	"fmt"
	"github.com/tmc/langchaingo/llms"
)

type ToolCallAgent struct {
	*ReActAgent
	AvailableTools       []llms.Tool
	Temperature          float64
	ToolsManager         tools.ToolsMap
	ToolCallChatResponse *llms.ContentChoice
}

// NewToolCallAgent 创建新的ReActAgent
func NewToolCallAgent(name, systemMessage string, temperature float64, llm llms.Model, maxSteps int, toolList []tools.Tool) *ToolCallAgent {
	mp, err := tools.NewToolsMap(toolList...)
	if err != nil {
		panic(err)
	}
	var availableTools []llms.Tool
	for _, tool := range toolList {
		availableTools = append(availableTools, tool.GetTool())
	}
	agent := &ToolCallAgent{
		ReActAgent:     NewReActAgent(name, systemMessage, llm, maxSteps),
		AvailableTools: availableTools,
		Temperature:    temperature,
		ToolsManager:   mp,
	}
	agent.ToolCallFunc = agent
	return agent
}

func (agent *ToolCallAgent) Think() (bool, error) {
	messageList := agent.MessageList
	llm := agent.LLM
	// 创建请求
	ctx := context.Background()
	options := llms.CallOptions{
		Tools:       agent.AvailableTools,
		Temperature: agent.Temperature,
	}
	content, err := llm.GenerateContent(ctx, messageList, llms.WithOptions(options))
	if err != nil {
		messageList = append(messageList, message.AIMessage("处理时遇到了错误："+err.Error()))
		return false, errors.New(fmt.Sprintf("%s的思考过程遇到了问题: %s", agent.Name, err.Error()))
	}
	// 记录响应，用于等下 Act
	aiMessage := content.Choices[0]
	agent.ToolCallChatResponse = aiMessage

	// 3、解析工具调用结果，获取要调用的工具
	// 如果不需要调用工具，返回 false
	if len(aiMessage.ToolCalls) == 0 {
		// 没有需要调用的工具
		messageList = append(messageList, message.AIMessage(aiMessage.Content))
		return false, nil
	} else {
		// 获取要调用的工具列表
		toolCalls := aiMessage.ToolCalls
		fmt.Printf("%s选择了 %d 个工具来调用\n", agent.Name, len(toolCalls))
		for _, tool := range toolCalls {
			fmt.Printf("Tool名称: %s, 参数: %s\n", tool.FunctionCall.Name, tool.FunctionCall.Arguments)
		}
		return true, nil
	}
}

func (agent *ToolCallAgent) Act() (string, error) {
	if len(agent.ToolCallChatResponse.ToolCalls) == 0 {
		return "没有工具需要调用", nil
	}

	// 调用工具
	for _, toolCall := range agent.ToolCallChatResponse.ToolCalls {
		// 加入AIMessage
		agent.MessageList = append(agent.MessageList, message.ToolCallMessage(toolCall))
		res, err := agent.ToolsManager.ExecuteTool(toolCall.FunctionCall.Name, toolCall.FunctionCall.Arguments)
		if err != nil {
			return "", err
		}
		agent.MessageList = append(agent.MessageList, message.ToolResponseMessage(toolCall.ID, toolCall.FunctionCall.Name, res))
	}
	// 判断是否调用了终止工具
	// 再次发送请求，判断是否还需要执行工具
	ctx := context.Background()
	options := llms.CallOptions{
		Tools:       agent.AvailableTools,
		Temperature: agent.Temperature,
	}
	content, err := agent.LLM.GenerateContent(ctx, agent.MessageList, llms.WithOptions(options))
	if err != nil {
		return "", errors.New(fmt.Sprintf("%s的思考过程遇到了问题: %s", agent.Name, err.Error()))
	}

	var result string
	if len(content.Choices[0].ToolCalls) == 0 {
		agent.State = FINISHED
		result = content.Choices[0].Content
	} else {
		result = "接下来调用工具" + content.Choices[0].ToolCalls[0].FunctionCall.Name
	}
	//fmt.Println(result)
	return result, nil
}
