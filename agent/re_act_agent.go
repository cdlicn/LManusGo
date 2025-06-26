package agent

import (
	"fmt"
	"github.com/tmc/langchaingo/llms"
)

type ToolCallFunc interface {
	// Think 处理当前状态并决定下一步行动
	Think() (bool, error)
	// Act 执行当前步骤
	Act() (string, error)
}

// ReActAgent 实现代理
type ReActAgent struct {
	*BaseAgent
	ToolCallFunc
}

// NewReActAgent 创建新的ReActAgent
func NewReActAgent(name, systemMessage string, llm llms.Model, maxSteps int) *ReActAgent {
	agent := &ReActAgent{
		BaseAgent: NewBaseAgent(name, systemMessage, llm, maxSteps),
	}
	agent.BaseAgent.ReActFunc = agent
	return agent
}

// Step 实现具体的步骤逻辑
func (agent *ReActAgent) Step() (string, error) {
	// 思考
	shouldAct, err := agent.Think()
	if err != nil {
		return "", err
	}
	if !shouldAct {
		return "思考完成 - 无需行动", nil
	}
	// 行动
	return agent.Act()
}

// Cleanup 清理方法
func (agent *ReActAgent) Cleanup() {
	// TODO Additional cleanup for MyAgent
	fmt.Println("Clean Up")
}
