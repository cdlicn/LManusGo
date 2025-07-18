package agent

import (
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
func NewReActAgent(baseAgent *BaseAgent) *ReActAgent {
	agent := &ReActAgent{
		BaseAgent: baseAgent,
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
		agent.State = FINISHED
		return "thinking is done - no action required", nil
	}
	// 行动
	return agent.Act()
}

// Cleanup 清理方法
func (agent *ReActAgent) Cleanup() {
	agent.BaseAgent.MessageList = make([]llms.MessageContent, 0)
	agent.State = IDLE
	agent.CurrentStep = 0
}
