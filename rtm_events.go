package slacker

import "encoding/json"

// RTMMessage represents a posted message in a channel
type RTMMessage struct {
	Type    string `json:"type"`
	Text    string `json:"text"`
	Channel string `json:"channel"`
	User    string `json:"user"`
	Ts      string `json:"ts"`

	// This should only be set when publishing and is handled by Publishable()
	ID uint64 `json:"id"`
}

// Message converts an event to an RTMMessage. If the event type is not
// "message", it returns an error
func (e RTMEvent) Message() (*RTMMessage, error) {
	var msg RTMMessage
	if err := json.Unmarshal(e.RawMessage, &msg); err != nil {
		return nil, err
	}

	return &msg, nil
}
