package builder

import (
	"errors"
	"github.com/fatih/structtag"
	"github.com/liuchamp/mhbuilder/log"
	"github.com/liuchamp/mhbuilder/utils"
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

	BUILD_POST   = "post"
	BUILD_PUT    = "put"
	BUILD_PATCH  = "patch"
	BUILD_FILTER = "filter"
)

var (
	NOTAG  = errors.New("Not find Tag\n")
	NOBODY = errors.New("No Body for File\n")
)

type Builder struct {
	PkgName  string
	FileName string
	file     *ast.File
	// 输出文件的基础路径，一般是OutputDir的绝对路径
	fm *FileMap
}

type BuilderOut struct {
	Post   string
	Filter string
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

// 解析go model源文件
func NewBuilder(pkg string, fileName string, file *ast.File) *Builder {
	builder := &Builder{
		PkgName:  pkg,
		FileName: fileName,
		file:     file,
	}
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
		return nil
	}
	builder.fm = fm
	return builder
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
func (builder *Builder) WirteFile() (*BuilderOut, error) {
	bot := &BuilderOut{}
	op, err := builder.outAddDtoAndToModel()
	if err != nil {
		return nil, err
	}
	bot.Post = op
	log.Debug(bot.Post)

	flt, err := builder.outFilter()
	if err != nil {
		return nil, err
	}
	bot.Filter = flt
	log.Debug(bot.Filter)
	return bot, nil
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
			if len(vf.Names) < 1 {
				continue
			}
			field := FieldMap{}
			field.FieldName = vf.Names[0].Name
			fType := vf.Type.(*ast.Ident)
			field.Types = fType.Name

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
