package utils

import (
	. "github.com/smartystreets/goconvey/convey"
	"strings"
	"testing"
)

func TestGetPkgName(t *testing.T) {
	Convey("pkgName test ", t, func() {
		pkg, err := GetPkgName("./")
		So(err, ShouldBeNil)
		So(strings.Contains(pkg, "utils"), ShouldBeTrue)
	})
	Convey("pkgName test path err", t, func() {
		pkg, err := GetPkgName("~/")
		So(err, ShouldNotBeNil)
		So(pkg, ShouldEqual, "")
	})
}
