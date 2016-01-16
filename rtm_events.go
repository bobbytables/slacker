package slacker

import "encoding/json"

type RTMMessage struct {
	Type string `json:"type"`
	Text string `json:"text"`
}

func (e RTMEvent) Message() (*RTMMessage, error) {
	var msg RTMMessage
	if err := json.Unmarshal(e.RawMessage, &msg); err != nil {
		return nil, err
	}

	return &msg, nil
}
