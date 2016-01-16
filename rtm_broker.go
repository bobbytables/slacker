package slacker

import (
	"encoding/json"

	"github.com/gorilla/websocket"
)

// RTMBroker handles incoming and outgoing messages to a Slack RTM Websocket
type RTMBroker struct {
	url      string
	incoming chan []byte
	outgoing chan []byte
	conn     *websocket.Conn
	closed   bool
}

// RTMEvent repesents a simple event received from Slack
type RTMEvent struct {
	Type       string `json:"type"`
	RawMessage json.RawMessage
}

// NewRTMBroker returns a connected broker to Slack from a rtm.start result
func NewRTMBroker(s *RTMStartResult) *RTMBroker {
	broker := &RTMBroker{
		url: s.URL,
	}

	return broker
}

func (b *RTMBroker) Connect() error {
	conn, _, err := websocket.DefaultDialer.Dial(b.url, nil)
	if err != nil {
		return err
	}

	b.conn = conn
	b.incoming = make(chan []byte, 0)
	b.outgoing = make(chan []byte, 0)

	go b.startRecv()

	return nil
}

// Close Closes the connection to Slack RTM
func (b *RTMBroker) Close() error {
	b.closed = true
	return b.conn.Close()
}

func (b *RTMBroker) startRecv() {
	for !b.closed {
		msg, message, _ := b.conn.ReadMessage()
		if msg == websocket.TextMessage {
			b.incoming <- message
		}
	}
}

func (b *RTMBroker) Events() <-chan RTMEvent {
	events := make(chan RTMEvent)

	go func() {
		for !b.closed {
			raw := json.RawMessage(<-b.incoming)
			rtmEvent := &RTMEvent{
				RawMessage: raw,
			}

			if err := json.Unmarshal(raw, rtmEvent); err != nil {
				panic(err)
			}

			events <- *rtmEvent
		}
	}()

	return events
}
