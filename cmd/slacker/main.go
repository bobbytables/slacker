package main

import (
	"fmt"
	"os"

	"github.com/bobbytables/slacker"
	"github.com/codegangsta/cli"
)

// This command line tool is mostly used to just see if you're wired up
// correctly. If you're making contributions this is a good tool to modify
// to do smoke testing.
func main() {
	app := cli.NewApp()
	app.Usage = "runs methods against the Slack API and returns the result"
	app.Author = "Bobby Tables <me@bobbytables.io>"
	app.Name = "slacker"
	app.Commands = []cli.Command{slackMethod()}

	app.Run(os.Args)
}

func slackMethod() cli.Command {
	return cli.Command{
		Name:  "run",
		Usage: "[method]",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:   "token",
				Usage:  "Your Slack API token",
				EnvVar: "SLACK_TOKEN",
			},
		},
		Description: "Hits the SlackAPI using the format: https://slack.com/api/{method}",
		Action: func(ctx *cli.Context) {
			if len(ctx.Args()) == 0 {
				cli.ShowSubcommandHelp(ctx)
				return
			}

			method := ctx.Args()[0]
			token := ctx.String("token")

			client := slacker.NewAPIClient(token, "")
			b, err := client.RunMethod(method)
			if err != nil {
				fmt.Printf("Error running method: %s", err.Error())
				return
			}

			fmt.Println(string(b))
		},
	}
}
