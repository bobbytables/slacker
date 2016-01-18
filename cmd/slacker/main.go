package main

import (
	"fmt"
	"os"

	"github.com/bobbytables/slacker"
)

// This command line tool is mostly used to just see if you're wired up
// correctly. If you're making contributions this is a good tool to modify
// to do smoke testing.
func main() {
	token := os.Getenv("SLACK_TOKEN")
	if token == "" {
		fmt.Println("You must set SLACK_TOKEN")
		os.Exit(1)
	}

	slackerAPI := slacker.NewAPIClient(token, "")
	startResult, err := slackerAPI.RTMStart()
	if err != nil {
		fmt.Printf("Error: %q\n", err.Error())
		os.Exit(1)
	}

	users := map[string]string{}
	slackUsers, err := slackerAPI.UsersList()
	if err != nil {
		fmt.Printf("Error: %q\n", err.Error())
		os.Exit(1)
	}

	for _, user := range slackUsers {
		users[user.ID] = user.Profile.RealName
	}

	broker := slacker.NewRTMBroker(startResult)
	if err := broker.Connect(); err != nil {
		fmt.Printf("Error: %q\n", err.Error())
		os.Exit(1)
	}

	msg := slacker.RTMMessage{
		Text: "Hello",
	}

	broker.Publish(msg)
}

func printMessage(e slacker.RTMEvent, users map[string]string) {
	msg, err := e.Message()
	if err != nil {
		return
	}

	var uName string
	uName, ok := users[msg.User]
	if !ok {
		uName = "unknown"
	}

	fmt.Printf("%s: %s\n", uName, msg.Text)
}
