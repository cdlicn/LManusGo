package agent

import (
	"LManusGo/message"
	"LManusGo/tools"
	"context"
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/tmc/langchaingo/llms"
	"sync"
)

type ToolCallAgent struct {
	*ReActAgent
	AvailableTools       []llms.Tool
	ToolsManager         tools.ToolsMap
	ToolCallChatResponse *llms.ContentChoice
}

// NewToolCallAgent 创建新的ReActAgent
func NewToolCallAgent(reactAgent *ReActAgent, toolList []tools.Tool) *ToolCallAgent {
	mp, err := tools.NewToolsMap(toolList...)
	if err != nil {
		logrus.Panicln(err)
	}
	var availableTools []llms.Tool
	for _, tool := range toolList {
		availableTools = append(availableTools, tool.GetTool())
	}
	agent := &ToolCallAgent{
		ReActAgent:     reactAgent,
		AvailableTools: availableTools,
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
		messageList = append(messageList, message.AIMessage("an error was encountered while processing, err: "+err.Error()))
		return false, errors.New(fmt.Sprintf("%s's thought process ran into a problem: %s", agent.Name, err.Error()))
	}
	// 记录响应，用于等下 Act
	aiMessage := content.Choices[0]
	agent.ToolCallChatResponse = aiMessage

	// 解析工具调用结果，获取要调用的工具
	// 如果不需要调用工具，返回 false
	if len(aiMessage.ToolCalls) == 0 {
		// 没有需要调用的工具
		messageList = append(messageList, message.AIMessage(aiMessage.Content))
		return false, nil
	} else {
		// 获取要调用的工具列表
		toolCalls := aiMessage.ToolCalls
		logrus.Infof("%s selects %d tools to call", agent.Name, len(toolCalls))
		for _, tool := range toolCalls {
			logrus.Infof("tool name: %s, parameters: %s", tool.FunctionCall.Name, tool.FunctionCall.Arguments)
		}
		return true, nil
	}
}

func (agent *ToolCallAgent) Act() (string, error) {
	if len(agent.ToolCallChatResponse.ToolCalls) == 0 {
		return "there are no tools to call", nil
	}

	// 调用工具
	const maxConcurrent = 3
	var wg sync.WaitGroup
	sem := make(chan struct{}, maxConcurrent)

	for _, toolCall := range agent.ToolCallChatResponse.ToolCalls {
		wg.Add(1)
		agent.work(toolCall, &wg, sem)
	}

	wg.Wait()

	ctx := context.Background()
	options := llms.CallOptions{
		Tools:       agent.AvailableTools,
		MaxTokens:   agent.BaseAgent.MaxTokens,
		Temperature: agent.Temperature,
	}
	content, err := agent.LLM.GenerateContent(ctx, agent.MessageList, llms.WithOptions(options))
	if err != nil {
		return "", errors.New(fmt.Sprintf("%s's thought process is running into a problem: %s", agent.Name, err.Error()))
	}

	//logrus.Infoln(content)

	var result string
	if len(content.Choices[0].ToolCalls) == 0 {
		agent.State = FINISHED
		result = content.Choices[0].Content
	} else {
		result = "next call the tool" + content.Choices[0].ToolCalls[0].FunctionCall.Name
	}
	//fmt.Println(result)
	return result, nil
}

func (agent *ToolCallAgent) work(toolCall llms.ToolCall, wg *sync.WaitGroup, sem chan struct{}) {
	sem <- struct{}{}
	defer func() {
		<-sem
		wg.Done()
	}()

	// 加入AIMessage
	agent.MessageList = append(agent.MessageList, message.ToolCallMessage(toolCall))
	res, err := agent.ToolsManager.ExecuteTool(toolCall.FunctionCall.Name, toolCall.FunctionCall.Arguments)
	if err != nil {
		res = err.Error()
	}

	logrus.Infoln("tool call: %s, tool response: %s", toolCall.FunctionCall.Name, res)

	agent.MessageList = append(agent.MessageList, message.ToolResponseMessage(toolCall.ID, toolCall.FunctionCall.Name, res))
}
