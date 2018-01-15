package strtemplate

import (
	"log"
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
			log.Println(*msgout)
			So(err, ShouldBeNil)
		})
	})
}
