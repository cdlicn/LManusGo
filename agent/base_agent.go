package agent

import (
	"LManusGo/message"
	"errors"
	"github.com/sirupsen/logrus"
	"github.com/tmc/langchaingo/llms"
	"strconv"
	"strings"
)

type AgentState string

const (
	// IDLE 空闲状态
	IDLE AgentState = "IDLE"

	// RUNNING 运行中状态
	RUNNING AgentState = "RUNNING"

	// FINISHED 已完成状态
	FINISHED AgentState = "FINISHED"

	// ERROR 错误状态
	ERROR AgentState = "ERROR"
)

type ReActFunc interface {
	Step() (string, error)
	Cleanup()
}

type BaseAgent struct {
	Name          string
	SystemMessage string
	State         AgentState
	CurrentStep   int
	MaxSteps      int
	LLM           llms.Model
	MaxTokens     int
	Temperature   float64
	MessageList   []llms.MessageContent
	ReActFunc
}

// NewBaseAgent 创建基础代理
func NewBaseAgent(name, systemMessage string, llm llms.Model, maxSteps, maxTokens int, temperature float64) *BaseAgent {
	systemMessage = strings.TrimSpace(systemMessage)
	return &BaseAgent{
		Name:          name,
		SystemMessage: systemMessage,
		State:         IDLE,
		MaxSteps:      maxSteps,
		LLM:           llm,
		MaxTokens:     maxTokens,
		Temperature:   temperature,
		MessageList:   make([]llms.MessageContent, 0),
	}
}

func (agent *BaseAgent) Run(userMessage string) (string, error) {
	// 基础校验
	userMessage = strings.TrimSpace(userMessage)
	if agent.State != IDLE {
		return "", errors.New("cannot run agent from state: " + string(agent.State))
	}
	if userMessage == "" {
		return "", errors.New("cannot run agent with empty user message")
	}
	// 执行，更改状态
	agent.State = RUNNING
	if agent.SystemMessage != "" {
		agent.MessageList = append(agent.MessageList, message.SystemMessage(agent.SystemMessage))
	}
	// 添加用户消息上下文
	agent.MessageList = append(agent.MessageList, message.UserMessage(userMessage))

	// 清理资源
	defer func() {
		agent.Cleanup()
	}()

	// 保存结果的列表
	var results []string

	// 执行循环
	for i := 0; i < agent.MaxSteps && agent.State != FINISHED; i++ {
		stepNo := i + 1
		agent.CurrentStep = stepNo
		logrus.Printf("Executing step %d/%d", stepNo, agent.MaxSteps)

		// 执行
		stepResult, err := agent.Step()
		if err != nil {
			agent.State = ERROR
			return "", errors.New("error running agent: " + err.Error())
		}
		results = append(results, "Step "+strconv.Itoa(agent.CurrentStep)+": "+stepResult)
	}
	if agent.CurrentStep >= agent.MaxSteps {
		agent.State = FINISHED
		logrus.Info("Terminated: Reach max steps (" + strconv.Itoa(agent.MaxSteps) + ")")
		results = append(results, "Terminated: Reach max steps ("+strconv.Itoa(agent.MaxSteps)+")")
	}

	return strings.Join(results, "\n"), nil
}
