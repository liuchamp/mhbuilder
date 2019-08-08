package gen

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

const SEARCHDIR = "/Users/mac/Code/go/src/github.com/liuchamp/mhbuilder/testdata/models"

func TestNew(t *testing.T) {
	gen := New()
	c := &Config{
		SearchDir:            SEARCHDIR,
		OutputDir:            SEARCHDIR,
		PropNamingStrategy:   "camelcase",
		RelathionPutAndPatch: "None",
		Files:                nil,
		ParseVendor:          false,
		ParseDependency:      false,
	}
	Convey("构建不包含vendor", t, func() {
		err := gen.Build(c)
		So(err, ShouldEqual, nil)
	})
	Convey("构建包含vendor", t, func() {
		c.ParseDependency = true
		c.ParseVendor = true
		err := gen.Build(c)
		So(err, ShouldEqual, nil)
	})
}
