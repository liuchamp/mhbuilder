package builder

import "github.com/liuchamp/mhbuilder/utils"

/**
产生更新代码
更新代码格式
由于更新服务代码，复杂程度更大， 有跨文件变量。 在生成文件过程中，不再是post,filter简单的代码，而是各种复杂内部变量配合的。完成scope过滤，validate校验。最终生成更新update。
当然这里还有一个问题，不能更新复杂结构体。复杂结构体还是需要用户自己编码，或者后续更新支持
*/
// 导出整哥文件
func (builder *Builder) outPut() (string, error) {
	return "", nil
}

type putFileHeadTemp struct {
	ModelNames []string // 本文件的model Name列表
}

var _putFileHeadTemp = `
import (
	"errors"
	"github.com/wxnacy/wgo/arrays"
	"go.mongodb.org/mongo-driver/bson"
	"windplatform/webbackend/server/models"
)
var (
{{range $element := .ModelNames}}
	{{$element}}ScopeMap map[int][]string
	{{$element}}JBMap map[string]string
	{{$element}}ValidatorMap map[string]string
{{end}}	
)
`

func getPutHeader(pfht putFileHeadTemp) (string, error) {
	f, err := utils.ParserName(_filterFileTemp, pfht)
	if err != nil {
		return "", err
	}
	return f.String(), nil
}

var _putFileInitTemp = `
func init() {
	{{.InitString}}
}
`

type putModelFilterInitTemp struct {
	ModelName     string
	ScopeToFields map[int][]string
	JBMap         map[string]string
	ValiMap       map[string]string
}

var _putModelFilterInitTemp = `
{{$moName:=.ModelName}}
{{range $scope,$fields := .ScopeToFields}}
	{{$moName}}ScopeMap[{{$scope}}] = {{ $length := lenfxs .Data }} []string{ {{range $index,$field := .fields}}"{{$field}}"{{ if gt $length $index }},{{end}}{{end}}}
{{end}}
{{range $jtag,$btag := .JBMap}}
	{{$moName}}JBMap["{{$jtag}}"] = "{{$btag}}"
{{end}}
{{range $btag,$vali := .ValiMap}}
	{{$moName}}ValidatorMap["{{$btag}}"] = "{{$vali}}"
{{end}}
`

func getPutInit(pfht putModelFilterInitTemp) (string, error) {
	f, err := utils.ParserName(_putModelFilterInitTemp, pfht)
	if err != nil {
		return "", err
	}
	return f.String(), nil
}

type putFunc4ModelTemp struct {
	ModelNameF string // model 原名
	ModelName  string // model名首字母小写
}

var _putFunc4ModelTemp = `
func {{.ModelNameF}}UpdateDTO(values map[string]interface{}, scope int) (updater interface{}, valiErr []models.ValidateErr, err error) {
	if values == nil || len(values) == 0 {
		return nil, nil, errors.New("value nil")
	}

	up := bson.M{}
	for jTag, value := range values {
		// 检查值是否存在scope中
		if check{{.ModelNameF}}ValueOptInScope(jTag, scope) {
			// 值校验
			err := bsVali.Var(value, {{.ModelName}}ValidatorMap[jTag])
			if err != nil {
				valiErr = append(valiErr, models.ValidateErr{FieldName: jTag, ValidatorMsg: err.Error()})
			} else {
				up[{{.ModelName}}JBMap[jTag]] = value
			}
		}
	}

	if valiErr != nil || len(valiErr) > 0 {
		return nil, valiErr, errors.New("validate error")
	}

	return bson.M{"$set": up}, nil, nil
}
`
