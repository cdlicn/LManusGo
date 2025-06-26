package agent

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
