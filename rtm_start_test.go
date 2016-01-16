package slacker

import (
	"net/http"
	"net/http/httptest"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestRTM(t *testing.T) {
	Convey("RTM Methods", t, func() {
		Convey("rtm.start success", func() {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				_, err := w.Write([]byte(`{"ok": true, "url": "wss:\/\/ms100.slack-msgs.com\/websocket\/bunk"}`))
				if err != nil {
					panic(err)
				}
			}))

			client := NewAPIClient("my-token", server.URL)
			result, err := client.RTMStart()
			So(err, ShouldBeNil)
			So(result.URL, ShouldEqual, "wss://ms100.slack-msgs.com/websocket/bunk")
		})
	})
}
