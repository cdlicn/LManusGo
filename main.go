package main

import (
	"LManusGo/agent"
	"fmt"
	"github.com/sirupsen/logrus"
	"strings"
)

// test prompt: 告诉我现在的时间，帮我制定南昌一日游的旅游攻略，并保存下来

func main() {
	manus := agent.NewLManus()
	for {
		var prompt string
		fmt.Print("Enter your prompt:")
		fmt.Scanln(&prompt)

		prompt = strings.TrimSpace(prompt)

		if prompt == "" {
			logrus.Warn("Empty prompt provided.")
			return
		}
		logrus.Warn("Processing your request...")
		resp, err := manus.Run(prompt)
		if err != nil {
			logrus.Panic(err)
		}
		fmt.Println(resp)
		logrus.Info("Request processing completed.")
	}
}
