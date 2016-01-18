package slacker

import (
	"encoding/json"
	"time"

	"github.com/gorilla/websocket"
)

// RTMBroker handles incoming and outgoing messages to a Slack RTM Websocket
type RTMBroker struct {
	url      string
	incoming chan []byte
	outgoing chan []byte
	events   chan RTMEvent
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

// Connect connects to the RTM Websocket
func (b *RTMBroker) Connect() error {
	conn, _, err := websocket.DefaultDialer.Dial(b.url, nil)
	if err != nil {
		return err
	}

	b.conn = conn
	b.incoming = make(chan []byte, 0)
	b.events = make(chan RTMEvent, 0)

	go b.startRecv()
	go b.handleEvents()

	return nil
}

// Close Closes the connection to Slack RTM
func (b *RTMBroker) Close() error {
	b.closed = true
	return b.conn.Close()
}

// Events returns a receive-only channel for all Events RTM API pushes
// to the broker.
func (b *RTMBroker) Events() <-chan RTMEvent {
	return b.events
}

// Publish pushes an event to the RTM Websocket
func (b *RTMBroker) Publish(e Publishable) error {
	d, err := e.Publishable()
	if err != nil {
		return err
	}

	return b.conn.WriteMessage(websocket.TextMessage, d)
}

func (b *RTMBroker) startRecv() {
	for !b.closed {
		msgType, message, _ := b.conn.ReadMessage()
		if msgType == websocket.TextMessage {

			b.incoming <- message
		}

		time.Sleep(25 * time.Millisecond)
	}
}

func (b *RTMBroker) handleEvents() {
	for !b.closed {
		raw := json.RawMessage(<-b.incoming)

		rtmEvent := RTMEvent{
			RawMessage: raw,
		}

		if err := json.Unmarshal(raw, &rtmEvent); err != nil {
			panic(err)
		}

		b.events <- rtmEvent
	}
}
