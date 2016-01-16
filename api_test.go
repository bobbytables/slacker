package slacker

import (
	"bytes"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestAPIParsing(t *testing.T) {
	Convey("API Response Parsing Utilities", t, func() {
		Convey("When the response has failed", func() {
			badResponse := `{"ok":false,"error":"not_authed"}`
			dest := struct{}{}

			err := ParseResponse(bytes.NewBufferString(badResponse), &dest)
			So(err, ShouldEqual, ErrNotAuthed)
		})

		Convey("When the response has failed with an error we don't know about", func() {
			badResponse := `{"ok":false,"error":"chicken"}`
			dest := struct{}{}

			err := ParseResponse(bytes.NewBufferString(badResponse), &dest)
			So(err.Error(), ShouldEqual, "slacker: unknown error returned: chicken")
		})

		Convey("When the response is successful it decodes onto the destination", func() {
			goodResponse := `{"ok":true,"hello":"world"}`
			dest := &struct {
				Hello string `json:"hello"`
			}{}

			err := ParseResponse(bytes.NewBufferString(goodResponse), dest)
			So(err, ShouldBeNil)
			So(dest.Hello, ShouldEqual, "world")
		})
	})
}
