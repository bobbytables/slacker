package slacker

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestSlacker(t *testing.T) {
	Convey("Slacker", t, func() {
		Convey("NewAPIClient()", func() {
			client := NewAPIClient("my-token", DefaultAPIURL)

			So(client.SlackURL, ShouldEqual, DefaultAPIURL)
		})
	})
}
