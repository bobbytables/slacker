# Slacker

[![GoDoc](https://godoc.org/github.com/bobbytables/slacker?status.svg)](https://godoc.org/github.com/bobbytables/slacker)
[![Build Status](https://travis-ci.org/bobbytables/slacker.svg?branch=master)](https://travis-ci.org/bobbytables/slacker)

Slacker is a Golang package to interface with Slack's API and Real Time Messaging API.

For full documentation, always check [godoc](https://godoc.org/github.com/bobbytables/slacker).

## Simple Examples

It's always fun to see quick ways to use a package. Here are some examples of
how to use slacker for simple things.

### Getting all channels for a team

```go
c := slacker.NewAPIClient("your-slack-token", "")
channels, err := c.ChannelsList()
if err != nil {
	panic(err)
}

// Map channels so we can easily retrieve a channel by name.
mappedChannels := map[string]*slacker.Channel{}
for _, channel := range channels {
	mappedChannels[channel.Name] = channel
}

fmt.Printf("Channels: %+v", mappedChannels)
```

### Getting all members for a team

```go
c := slacker.NewAPIClient("your-slack-token", "")
users, err := c.UsersList()
if err != nil {
	panic(err)
}

mappedUsers := map[string]*slacker.User{}
for _, user := range users {
	mappedUsers[user.ID] = user
}
```

### Starting an RTM broker (real time messaging)

This example starts a websocket to Slack's RTM API and displays events as they
come in.

```go
c := slacker.NewAPIClient("your-slack-token", "")
rtmStart, err := c.RTMStart()
if err != nil {
	panic(err)
}

broker := slacker.NewRTMBroker(rtmStart)
broker.Connect()

for {
	event := <-broker.Events()
	fmt.Println(event.Type)

	if event.Type == "message" {
		msg, err := event.Message()
		if err != nil {
			panic(err)
		}

		fmt.Println(msg.Text)
	}
}
```


# License

Slacker is released under the [MIT License](LICENSE.md).
