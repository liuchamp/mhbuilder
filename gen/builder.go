package gen

import (
	"go/ast"
)

const (
	POSTTOSUFFIX = "AddDTO"
	POSTTOMODEL  = "toModel"
	PUTSUFFIX    = "UpdateDTO"
	PUTMETHMOD   = "Update"
	PUTCHSUFFIX  = "Match"
	SCOPESUFFIX  = "Filter"
)

type Outer interface {
	// 代码生成策略
	// 1 先将代码生成于 tmp 中，
	// 2 代码生成后，tmp目录删除
	out() (string, error)
}
type Builder struct {
	AddDTO map[string]AddDTOMap
}

func NewBuilder() *Builder {
	return &Builder{
		AddDTO: make(map[string]AddDTOMap),
	}
}

type AddDTOMap struct {
	Name    string
	Comment string
	Fields  map[string]DTOMap
}

type DTOMap struct {
	Comment string
	Fields  []FieldMap
}
type FieldMap struct {
	FieldName string
	Tags      map[string]string
	Comment   string
}

func (dto *AddDTOMap) out() (string, error) {
	return "", nil
}

type UpdateDTOMap struct {
}

func (builder *Builder) ExtentsFileInfo(file *ast.File) error {
	structsMap := make(map[string]*ast.StructType)

	collectStructs := func(x ast.Node) bool {
		ts, ok := x.(*ast.TypeSpec)
		if !ok || ts.Type == nil {
			return true
		}

		// 获取结构体名称
		structName := ts.Name.Name
		s, ok := ts.Type.(*ast.StructType)
		if !ok {
			return true
		}
		structsMap[structName] = s
		return false
	}
	ast.Inspect(file, collectStructs)
	return nil
}
