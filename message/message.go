package message

import "github.com/tmc/langchaingo/llms"

func AIMessage(message string) llms.MessageContent {
	return llms.MessageContent{
		Role: llms.ChatMessageTypeAI,
		Parts: []llms.ContentPart{
			llms.TextPart(message),
		},
	}
}

func UserMessage(message string) llms.MessageContent {
	return llms.MessageContent{
		Role: llms.ChatMessageTypeHuman,
		Parts: []llms.ContentPart{
			llms.TextPart(message),
		},
	}
}

func SystemMessage(message string) llms.MessageContent {
	return llms.MessageContent{
		Role: llms.ChatMessageTypeSystem,
		Parts: []llms.ContentPart{
			llms.TextPart(message),
		},
	}
}

func ToolCallMessage(toolCall llms.ToolCall) llms.MessageContent {
	return llms.MessageContent{
		Role: llms.ChatMessageTypeAI,
		Parts: []llms.ContentPart{
			toolCall,
		},
	}
}

func ToolResponseMessage(toolCallID, name, content string) llms.MessageContent {
	return llms.MessageContent{
		Role: llms.ChatMessageTypeTool,
		Parts: []llms.ContentPart{
			llms.ToolCallResponse{
				ToolCallID: toolCallID,
				Name:       name,
				Content:    content,
			},
		},
	}
}
