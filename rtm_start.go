package slacker

// RTMStartResult contains the result of rtm.start in the Slack API
type RTMStartResult struct {
	URL string `json:"url,omitempty"`
}

// RTMStart issues a start command for RTM. This is isually used for retrieving
// a WebSocket URL to start listening / posting messages into Slack.
func (c *APIClient) RTMStart() (*RTMStartResult, error) {
	return nil, nil
}
