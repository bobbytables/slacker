package main

import (
	"fmt"
	"os"
	"time"

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

	broker := slacker.NewRTMBroker(startResult)
	if err := broker.Connect(); err != nil {
		fmt.Printf("Error: %q\n", err.Error())
		os.Exit(1)
	}

	for {
		select {
		case event := <-broker.Events():
			fmt.Println(event.Type)
		default:
			time.Sleep(10 * time.Millisecond)
		}
	}
}
