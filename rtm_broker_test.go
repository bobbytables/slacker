package slacker

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/gorilla/websocket"
	. "github.com/smartystreets/goconvey/convey"
)

var voidConnHandler = func(c *websocket.Conn) {}

func TestRTMBroker(t *testing.T) {
	Convey("RTM Broker", t, func() {
		Convey("Events() returns RTMEvents on a channel", func() {
			server := generateServerWithResponse(`{"type":"message"}`, voidConnHandler)
			defer server.Close()
			start := &RTMStartResult{URL: server.URL}

			broker := NewRTMBroker(start)
			err := broker.Connect()

			So(err, ShouldBeNil)

			select {
			case event := <-broker.Events():
				So(event.Type, ShouldEqual, "message")
			case <-time.After(1 * time.Second):
				So(true, ShouldBeFalse)
			}
		})

		Convey("Publish() pushes an event to the server", func() {
			msg := RTMMessage{}
			done := make(chan bool)

			server := generateServerWithResponse(`{"type":"message"}`, func(c *websocket.Conn) {
				c.ReadJSON(&msg)
				done <- true
			})
			defer server.Close()

			start := &RTMStartResult{URL: server.URL}

			broker := NewRTMBroker(start)
			err := broker.Connect()

			So(err, ShouldBeNil)

			broker.Publish(RTMMessage{Text: "hello world!"})

			select {
			case <-done:
				So(msg.Type, ShouldEqual, "message")
				So(msg.Text, ShouldEqual, "hello world!")
				So(msg.ID, ShouldBeGreaterThan, 0)
			case <-time.After(1 * time.Second):
				So(true, ShouldBeFalse)
			}
		})
	})
}

func generateServerWithResponse(event string, f func(c *websocket.Conn)) *httptest.Server {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		upgrader := websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool { return true },
		}

		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			http.Error(w, "could not create websock connection", http.StatusInternalServerError)
		}
		conn.WriteMessage(websocket.TextMessage, []byte(event))
		go f(conn)
	}))
	server.URL = "ws" + strings.TrimPrefix(server.URL, "http")

	return server
}
