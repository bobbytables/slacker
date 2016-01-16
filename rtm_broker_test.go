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

func TestRTMBroker(t *testing.T) {

	Convey("RTM Broker", t, func() {
		server := generateServerWithResponse("hello world")
		defer server.Close()
		start := &RTMStartResult{URL: server.URL}

		done := make(chan bool)
		broker := NewRTMBroker(start)
		err := broker.Connect()
		So(err, ShouldBeNil)

		go func() {
			msg := <-broker.incoming
			if exp, got := "hello world", string(msg); exp != got {
				t.Errorf("Expected %q but got %q", exp, got)
			}
			close(done)
		}()

		<-done
		broker.Close()
	})

	Convey("Events() returns RTMEvents on a channel", t, func() {
		server := generateServerWithResponse(`{"type":"message"}`)
		defer server.Close()
		start := &RTMStartResult{URL: server.URL}

		done := make(chan bool)
		broker := NewRTMBroker(start)
		err := broker.Connect()
		So(err, ShouldBeNil)

		go func() {
			i := 0
			for i < 5 {
				select {
				case event := <-broker.Events():
					if exp, got := "message", event.Type; exp != got {
						t.Errorf("Expected %q but got %q", exp, got)
					}
				default:
					i = i + 1
					time.Sleep(10)
				}
			}
			close(done)
		}()

		<-done
	})
}

func generateServerWithResponse(event string) *httptest.Server {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		upgrader := websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool { return true },
		}

		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			http.Error(w, "could not create websock connection", http.StatusInternalServerError)
		}
		conn.WriteMessage(websocket.TextMessage, []byte(event))
	}))
	server.URL = "ws" + strings.TrimPrefix(server.URL, "http")

	return server
}
