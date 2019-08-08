package builder

import (
	"flag"
	"github.com/liuchamp/mhbuilder/log"
	"github.com/liuchamp/mhbuilder/utils"
	. "github.com/smartystreets/goconvey/convey"
	"go/parser"
	"go/token"
	"io/ioutil"
	"os"
	"path"
	"strings"
	"testing"
)

const SEARCHDIR = "../testdata/models"

var builder *Builder

func TestMain(m *testing.M) {

	filename := SEARCHDIR + "/user.go"
	pkg, _ := utils.GetPkgName(path.Dir(filename))
	src, _ := ioutil.ReadFile(filename)

	fset := token.NewFileSet()
	file, _ := parser.ParseFile(fset, filename, src, parser.ParseComments)
	builder = NewBuilder(pkg, filename, file)
	flag.Parse()
	exitCode := m.Run()

	// 退出
	os.Exit(exitCode)
}

func TestBuilder_ExtentsFileInfo(t *testing.T) {

	Convey("builder", t, func() {
		So(builder.file, ShouldNotBeNil)
		So(builder.PkgName, ShouldNotBeNil)
	})
}

func TestBuilder_WirteFile(t *testing.T) {
	Convey("builder wiite file", t, func() {
		bu, err := builder.WirteFile()
		So(err, ShouldBeNil)
		So(bu, ShouldNotBeNil)
		So(bu.Post, ShouldNotBeNil)
		So(bu.Filter, ShouldNotBeNil)
		log.Debug(builder.fm.Models[0].Name)
		So(strings.Contains(bu.Filter, builder.fm.Models[0].Name), ShouldBeTrue)
	})
}
