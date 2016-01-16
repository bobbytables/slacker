# Slacker

Slacker is a Golang package to easily integrate with the Slack RTM API. (Real Time Messaging).

It provides a simple interface to create a connection and start receiving events.

## Quick Start

```golang
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
```
