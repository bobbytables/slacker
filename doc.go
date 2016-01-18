/*
Package slacker is a Go package for working with the Slack integration tools.
This includes the API and RTM endpoints.

To create a new slacker client, you can run

	client := slacker.NewAPIClient("my-slack-api-token", "")

The first parameter is an OAuth2 token you should have obtained through either
the Slack integrations dashboard, or the 3-legged OAuth2 token flow.

The second parameter is the base URL for the Slack API. If left empty, it will
use https://slack.com/api for all RPC calls.

After you have created a client, you can call methods against the Slack API.
For example, retrieving all user's of a team

	users, err := client.UsersList()
	if err != nil {
		panic(err)
	}

	for _, user := range users {
		fmt.Println(user.RealName)
	}

If you want to run a method that is not supported by this package, you can call
generic methods by running:

	client.RunMethod("users.list")

This will return a []byte of the JSON returned by the Slack API.

*/
package slacker
