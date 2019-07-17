package gen

import (
	"errors"
	"github.com/fatih/structtag"
	"go/ast"
	"strings"
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
	FilesMap map[string]FileMap
	// filename to addDTO
}

func NewBuilder() *Builder {
	return &Builder{
		FilesMap: make(map[string]FileMap),
	}
}

type FileMap struct {
	PkgName string
	Models  []ModelExtend
}
type ModelExtend struct {
	Name   string
	Fields []FieldMap
}
type AddDTOMaps struct {
	Imports []string
	Comment string
	DTO     map[string]DTOMap
}

type DTOMap struct {
	Comment string
	Fields  []FieldMap
}

type FieldMap struct {
	FieldName string
	Tags      *structtag.Tags
	Comment   string
}

func (dto *AddDTOMaps) out() (string, error) {
	return "", nil
}

type UpdateDTOMap struct {
}

func (builder *Builder) ExtentsFileInfo(fileName string, pkgName string, file *ast.File) error {
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
	fm, errr := builder.extendDTOMap(structsMap)
	if errr != nil {
		return errr
	}
	builder.FilesMap[fileName] = *fm
	return nil
}

func (builder *Builder) extendDTOMap(structsMap map[string]*ast.StructType) (*FileMap, error) {
	if structsMap == nil || len(structsMap) < 1 {
		return nil, errors.New("not find struct")
	}
	adm := &FileMap{}

	for k, v := range structsMap {
		dtom := ModelExtend{}
		dtom.Name = k
		var fields []FieldMap
		for _, vf := range v.Fields.List {
			field := FieldMap{}
			field.FieldName = vf.Names[0].Name
			if vf.Tag != nil {
				tag := vf.Tag.Value
				tag = strings.Trim(tag, "`")
				tags, err := structtag.Parse(string(tag))
				if err != nil {
					return nil, err
				}
				field.Tags = tags
			}
			fields = append(fields, field)
		}
		dtom.Fields = fields
		adm.Models = append(adm.Models, dtom)
	}
	return adm, nil
}
