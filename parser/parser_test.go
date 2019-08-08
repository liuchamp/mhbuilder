package parser

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

const SEARCHDIR = "../testdata/models"

func TestParser_ParModel(t *testing.T) {
	p := NewParser()
	searchDir := SEARCHDIR

	Convey("解析目标目录", t, func() {
		err := p.ParModel(searchDir)
		So(err, ShouldEqual, nil)
		So(len(p.files), ShouldEqual, 1)
	})
}
