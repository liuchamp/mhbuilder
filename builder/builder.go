package builder

import (
	"errors"
	"github.com/fatih/structtag"
	"github.com/liuchamp/mhbuilder/utils"
	"os"
	"path"
	"path/filepath"

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

	TAG_BUILD = "build"
	TAG_SCOPE = "scope"
	TAG_JSON  = "json"
	TAG_BSON  = "bson"

	BUILD_POST  = "post"
	BUILD_PUT   = "put"
	BUILD_PATCH = "patch"
)

var (
	NOTAG  = errors.New("Not find Tag\n")
	NOBODY = errors.New("No Body for File\n")
)

type Outer interface {
	// 代码生成策略
	// 1 先将代码生成于 tmp 中，
	// 2 代码生成后，tmp目录删除
	out() (string, error)
}
type Builder struct {
	FilesMap map[string]FileMap
	// 输出文件的基础路径，一般是OutputDir的绝对路径
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
	Types     string
	Tags      *structtag.Tags
	Comment   string
}

func (dto *AddDTOMaps) out() (string, error) {
	return "", nil
}

type UpdateDTOMap struct {
}

// 解析go model源文件
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

// add model 的字段和toModel方法需要的字段
// @return field,toModel
func addFeildString(field *FieldMap) (string, string, error) {
	sSet, err := fieldCollectionBuild(field)
	if err != nil {
		return "", "", err
	}
	if !sSet.Has(BUILD_POST) {
		return "", "", NOTAG
	}

	fieldoutadd := fieldAddTemplate{
		FiledName: field.FieldName,
		Types:     field.Types,
		Tags:      field.Tags.String(),
	}

	code, err := utils.ParserName(_fieldAddTemplate, fieldoutadd)
	if err != nil {
		return "", "", err
	}

	dtomodel, err := utils.ParserName(_fieldAtdmTemplate, struct {
		Field string
	}{
		Field: field.FieldName,
	})
	if err != nil {
		return "", "", nil
	}
	return code.String(), dtomodel.String(), nil

}

// 获取build tag的所有值
func fieldCollectionBuild(field *FieldMap) (*utils.StringSet, error) {
	tags, err := field.Tags.Get(TAG_BUILD)
	if err != nil {
		return nil, err
	}
	stringSet := utils.NewStringSet()
	stringSet.Add(tags.Name)
	for _, v := range tags.Options {
		stringSet.Add(v)
	}
	return stringSet, nil
}

// 将数据写入对应文件夹
func (builder *Builder) WirteFile(outDir, opt string) error {
	// 文件输出的实际目录
	fileP := filepath.Join(outDir, opt)
	if err := os.MkdirAll(fileP, os.ModePerm); err != nil {
		return err
	}

	for fname := range builder.FilesMap {
		fileName := path.Base(fname)
		s, err := builder.selectOutChange(fname, opt)
		if err != nil {
			return err
		}
		file, err := os.Create(path.Join(fileP, fileName))
		if err != nil {
			return err
		}
		var data []byte = []byte(s)
		file.Write(data)
		file.Close()
	}
	return nil
}
func (builder *Builder) selectOutChange(sourceFile, opt string) (string, error) {
	if opt == BUILD_POST {
		return builder.outAddDtoAndToModel(sourceFile)
	}
	return "", errors.New("can not find opt")
}

// 需要拓展的数据，按照文件/结构的方式展开
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
