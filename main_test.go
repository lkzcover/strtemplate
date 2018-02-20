package strtemplate

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestMain(t *testing.T) {
	Convey("Main packege test", t, func() {
		Convey("Test 1", func() {
			data := make(map[string]interface{})

			data["params"] = "Hello"

			templates := "Templates example {{params}}"

			msgout, err := GetString(data, templates)
			So(err, ShouldBeNil)
			So(*msgout, ShouldEqual, "Templates example Hello")
		})
	})
}
