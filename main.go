package main

import (
	"LManusGo/agent"
	"LManusGo/tools"
	"fmt"
)

func main() {
	name := "LManus"
	systemMessage := `
		You are LManus, a versatile AI assistant designed to solve any task requested by users.
		You can use various tools to efficiently complete complex requests.
		Proactively select the most suitable tool or tool combination based on user needs.
		For complex tasks, you can break down the problem and gradually use different tools to solve it.
		After using each tool, clearly explain the execution results and suggest the next steps.
		If you want to stop interaction at any time, use the 'terminate' tool/function call.
		You only work with a single conversation, and you don't need to ask the user for any action after you end the conversation.
	`
	toolList := []tools.Tool{
		tools.DoTerminate{},
		tools.CurrentDate{},
		tools.CurrentTime{},
		tools.SearchWeb{},
		tools.SaveFile{},
	}

	manus := agent.NewLManus(name, systemMessage, toolList)

	for {
		var userMessage string
		fmt.Print("enter your question: ")
		fmt.Scanln(&userMessage)
		resp, err := manus.Run(userMessage)
		if err != nil {
			panic(err)
		}
		fmt.Println(resp)
	}
}
