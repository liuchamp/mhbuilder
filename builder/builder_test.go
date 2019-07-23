package builder

import (
	"go/parser"
	"go/token"
	"io/ioutil"
	"testing"
)

const SEARCHDIR = "../testdata/models"

func TestBuilder_ExtentsFileInfo(t *testing.T) {
	builder := NewBuilder()
	filename := SEARCHDIR + "/user.go"

	src, err := ioutil.ReadFile(filename)
	if err != nil {
		t.Fatal(err.Error())
	}
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, filename, src, parser.ParseComments)
	t.Run("check", func(tc *testing.T) {
		err = builder.ExtentsFileInfo(filename, "models", file)
		if err != nil {
			tc.Fatal(err.Error())
		}
	})
	t.Run("sm", func(ts *testing.T) {
		if builder.FilesMap == nil || len(builder.FilesMap) < 1 {
			return
		}
		t.Log("开始测试")
		for fname := range builder.FilesMap {
			s, err := builder.outAddDtoAndToModel(fname)
			t.Log(s, err)
		}
	})
}
