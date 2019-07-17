package gen

import (
	"go/parser"
	"go/token"
	"io/ioutil"
	"testing"
)

func TestBuilder_ExtentsFileInfo(t *testing.T) {
	builder := NewBuilder()
	filename := "/Users/mac/Code/go/src/github.com/liuchamp/godemo/models/user.go"

	src, err := ioutil.ReadFile(filename)
	if err != nil {
		t.Fatal(err.Error())
	}
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, filename, src, parser.ParseComments)
	err = builder.ExtentsFileInfo(file)
	if err != nil {
		t.Fatal(err.Error())
	}
}
