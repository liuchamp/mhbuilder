package builder

type headerTemplate struct {
	PkgName       string
	UseStrConv    bool
	EnableBatch   bool
	ImportPackage string
}

var _headerTemplate = `
// Code generated by champ tool dtogen. DO NOT EDIT.
/* 
  Package {{.PkgName}} is a generated mc cache package.
  It is generated from:
  ARGS
*/

package {{.PkgName}}

import (
	"context"
	"fmt"
	{{if .UseStrConv}}"strconv"{{end}}
	{{if .EnableBatch }}"sync"{{end}}
	"{{.ImportPackage}}"
)

`

type addFile struct {
	FileHeader string
	Body       []string
}

var _addFile = `
{{.FileHeader}}

{{range $element := .Body}}
{{$element}}
{{end}}

`

type addDtoTemplate struct {
	StructName string
	Feilds     []string
}

var _addDtoTemplate = `
// {{.StructName}} created
type {{.StructName}} struct {
{{range $element := .Feilds}}
	{{$element}}
{{end}}
}
`

type addDtoToModelTemplate struct {
	StructName string
	Model      string
	Fields     []string
}

var _addDtoToModelTemplate = `
	// {{.StructName}} function To model
func (dto *{{.StructName}})toModel() *{{.Model}} {
	model:=&{{.Model}}{}
{{range $element := .Fields}}
	{{$element}}
{{end}}
}
`

type fieldAddTemplate struct {
	FiledName string
	Types     string
	Tags      string
}

var _fieldAddTemplate = "{{.FiledName}} {{.Types}} `{{.Tags}}`"

var _fieldAtdmTemplate = "model.{{.Field}} = dto.{{.Field}}"
