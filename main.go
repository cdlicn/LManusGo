package main

import (
	"LManusGo/agent"
	"bufio"
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"strings"
	"time"
)

// test prompt 1: I need to do a final assignment for the student management system, help me generate a simple demo, save the code locally, and ask for the python language.
// test prompt 2: I would like to travel to Beijing in the near future and help me designate a travel guide and keep the strategy locally.

func main() {
	for {
		manus := agent.NewLManus()
		var prompt string
		fmt.Print("Enter your prompt (or 'exit' to quit): ")
		scanner := bufio.NewScanner(os.Stdin)

		if scanner.Scan() {
			prompt = scanner.Text()
		}

		prompt = strings.TrimSpace(prompt)

		if prompt == "exit" {
			break
		}

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
