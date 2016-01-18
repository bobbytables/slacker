package slacker

import (
	"net/http"
	"net/http/httptest"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

var (
	channelsListResp = `{"ok":true,"channels":[{"id":"C0HH54ZNV","name":"engineering","is_channel":true,"created":1451536266,"creator":"AKJSHFHFU","is_archived":false,"is_general":false,"is_member":true,"members":["AKJSHFHFU"],"topic":{"value":"","creator":"","last_set":0},"purpose":{"value":"All things engineering","creator":"AKJSHFHFU","last_set":1451536266},"num_members":1}]}`
)

func TestChannelsAPI(t *testing.T) {
	Convey("Slack Channels API Endpoints", t, func() {
		Convey("channels.list", func() {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Write([]byte(channelsListResp))
			}))
			defer server.Close()

			client := NewAPIClient("bunk", server.URL)
			channels, err := client.ChannelsList()
			So(err, ShouldBeNil)
			So(channels, ShouldHaveLength, 1)

			channel := channels[0]

			So(channel.ID, ShouldEqual, "C0HH54ZNV")
			So(channel.Name, ShouldEqual, "engineering")
			So(channel.IsChannel, ShouldEqual, true)
			So(channel.Created, ShouldEqual, 1451536266)
			So(channel.Creator, ShouldEqual, "AKJSHFHFU")
			So(channel.IsArchived, ShouldEqual, false)
			So(channel.IsGeneral, ShouldEqual, false)
			So(channel.IsMember, ShouldEqual, true)
			So(channel.NumMembers, ShouldEqual, 1)

			So(channel.Members[0], ShouldEqual, "AKJSHFHFU")

			So(channel.Topic.Value, ShouldEqual, "")
			So(channel.Topic.Creator, ShouldEqual, "")
			So(channel.Topic.LastSet, ShouldEqual, 0)

			So(channel.Purpose.Value, ShouldEqual, "All things engineering")
			So(channel.Purpose.Creator, ShouldEqual, "AKJSHFHFU")
			So(channel.Purpose.LastSet, ShouldEqual, 1451536266)
		})
	})
}
