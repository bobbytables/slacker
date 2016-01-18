package slacker

import (
	"net/http"
	"net/http/httptest"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

var (
	usersListResp = `{ "ok": true, "members": [ { "team_id": "HURFFDURF", "id": "AKDOGURTUE", "name": "bobbytables", "deleted": false, "status": null, "color": "3c989f", "real_name": "Bobby Tables", "tz": "America\/Los_Angeles", "tz_label": "Pacific Standard Time", "tz_offset": -28800, "profile": { "first_name": "Bobby", "last_name": "Tables", "avatar_hash": "gd1bb47f21fb", "real_name": "Bobby Tables", "real_name_normalized": "Bobby Tables", "email": "bobbytables@dropstudents.com", "image_24": "https://fake.com/image.png", "image_32": "https://fake.com/image.png", "image_48": "https://fake.com/image.png", "image_72": "https://fake.com/image.png", "image_192": "https://fake.com/image.png", "image_512": "https://fake.com/image.png" }, "is_admin": false, "is_owner": false, "is_primary_owner": false, "is_restricted": false, "is_ultra_restricted": false, "is_bot": false, "has_2fa": false } ], "cache_ts": 1453140218 }`
)

func TestUsersAPI(t *testing.T) {
	Convey("Slack Users API Endpoints", t, func() {
		Convey("users.list", func() {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Write([]byte(usersListResp))
			}))
			defer server.Close()

			client := NewAPIClient("bunk", server.URL)
			users, err := client.UsersList()
			So(err, ShouldBeNil)
			So(users, ShouldHaveLength, 1)

			user := users[0]

			So(user.ID, ShouldEqual, "AKDOGURTUE")
			So(user.TeamID, ShouldEqual, "HURFFDURF")
			So(user.Name, ShouldEqual, "bobbytables")
			So(user.Deleted, ShouldBeFalse)
			So(user.Color, ShouldEqual, "3c989f")
			So(user.RealName, ShouldEqual, "Bobby Tables")

			profile := user.Profile

			So(profile.FirstName, ShouldEqual, "Bobby")
			So(profile.LastName, ShouldEqual, "Tables")
			So(profile.AvatarHash, ShouldEqual, "gd1bb47f21fb")
			So(profile.RealName, ShouldEqual, "Bobby Tables")
			So(profile.RealNameNormalized, ShouldEqual, "Bobby Tables")
			So(profile.Email, ShouldEqual, "bobbytables@dropstudents.com")
			So(profile.Image24, ShouldEqual, "https://fake.com/image.png")
			So(profile.Image32, ShouldEqual, "https://fake.com/image.png")
			So(profile.Image48, ShouldEqual, "https://fake.com/image.png")
			So(profile.Image72, ShouldEqual, "https://fake.com/image.png")
			So(profile.Image192, ShouldEqual, "https://fake.com/image.png")
			So(profile.Image512, ShouldEqual, "https://fake.com/image.png")
		})
	})
}
