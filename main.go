package main

import (
	"LManusGo/agent"
	"fmt"
	"github.com/sirupsen/logrus"
	"strings"
	"time"
)

// test prompt: 告诉我现在的时间，帮我制定南昌一日游的旅游攻略，并保存下来

func main() {
	manus := agent.NewLManus()
	for {
		var prompt string
		fmt.Print("Enter your prompt:")
		_, err := fmt.Scanln(&prompt)
		if err != nil {
			logrus.Panic("scanln failed:", err.Error())
		}

		prompt = strings.TrimSpace(prompt)

		if prompt == "" {
			logrus.Panic("empty prompt provided.")
		}
		logrus.Warnln("processing your request...")
		resp, err := manus.Run(prompt)
		if err != nil {
			logrus.Panic(err)
		}
		fmt.Println(resp)
		logrus.Infoln("request processing completed.")
		time.Sleep(2 * time.Second)
	}
}
