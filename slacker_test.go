package slacker

import (
	"testing"

	"golang.org/x/oauth2"

	. "github.com/smartystreets/goconvey/convey"
)

func TestSlacker(t *testing.T) {
	Convey("Slacker", t, func() {
		Convey("NewAPIClient()", func() {
			client := NewAPIClient("my-token", DefaultAPIURL)

			So(client.SlackURL, ShouldEqual, DefaultAPIURL)
			So(client.client.Transport, ShouldHaveSameTypeAs, &oauth2.Transport{})
		})
	})
}
