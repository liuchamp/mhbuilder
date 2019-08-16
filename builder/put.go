package builder

import (
	"errors"
	"github.com/liuchamp/mhbuilder/log"
	"github.com/liuchamp/mhbuilder/utils"
	"strings"
)

/**
产生更新代码
更新代码格式
由于更新服务代码，复杂程度更大， 有跨文件变量。 在生成文件过程中，不再是post,filter简单的代码，而是各种复杂内部变量配合的。完成scope过滤，validate校验。最终生成更新update。
当然这里还有一个问题，不能更新复杂结构体。复杂结构体还是需要用户自己编码，或者后续更新支持
*/
// 导出整个文件
func (builder *Builder) outPut() (string, error) {
	models := builder.fm.Models
	// 收集model名称
	var mns []string
	for _, v := range models {
		mns = append(mns, strings.ToLower(v.Name[0:1])+v.Name[1:])
	}
	// 获取文件头部
	header, err := getPutHeader(putFileHeadTemp{ModelNames: mns})
	if err != nil {
		return "", err
	}
	var allModelInit string
	var modelsf string
	for _, v := range models {
		minit, err := getModelInitInfo(v)
		if err != nil {
			return "", err
		}
		allModelInit += minit
		sf, err := getModelPutInfo(v)
		if err != nil {
			return "", err
		}
		modelsf += sf
	}
	init, err := getPutInitString(map[string]string{"InitString": allModelInit})

	return header + init + modelsf, nil
}

type putFileHeadTemp struct {
	ModelNames []string // 本文件的model Name列表
}

var _putFileHeadTemp = `package put
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
	f, err := utils.ParserName(_putFileHeadTemp, pfht)
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

func getPutInitString(ni map[string]string) (string, error) {
	f, err := utils.ParserName(_putFileInitTemp, ni)
	if err != nil {
		return "", err
	}
	return f.String(), nil
}

type putModelFilterInitTemp struct {
	ModelName     string
	ScopeToFields map[int][]string
	JBMap         map[string]string
	ValiMap       map[string]string
}

var _putModelFilterInitTemp = `
{{- $moName:=.ModelName}}
{{- $slen := len .ScopeToFields}}
{{- if gt $slen 0}}
	{{$moName}}ScopeMap = make(map[int][]string)
{{- end}}
{{- range $scope,$fields := .ScopeToFields}}
	{{$moName}}ScopeMap[{{$scope}}] = {{ $length := lenfxs $fields }} []string{ {{range $index,$field := $fields}}"{{$field}}"{{ if gt $length $index }},{{end}}{{end}} }
{{- end}}
{{- $ljm := len .JBMap}}
{{- if gt $ljm 0}}
	{{$moName}}JBMap = make(map[string]string)
{{- end}}
{{- range $jtag,$btag := .JBMap}}
	{{$moName}}JBMap["{{$jtag}}"] = "{{$btag}}"
{{- end}}

{{- $le:= len .ValiMap}}
{{- if gt $le 0}}
	{{$moName}}ValidatorMap = make(map[string]string)
{{- end}}
{{- range $btag,$vali := .ValiMap}}
	{{$moName}}ValidatorMap["{{$btag}}"] = "{{$vali}}"
{{- end}}
`

func getPutInit(pfht putModelFilterInitTemp) (string, error) {
	f, err := utils.ParserName(_putModelFilterInitTemp, pfht)
	if err != nil {
		return "", err
	}
	return f.String(), nil
}

func getModelInitInfo(extend ModelExtend) (string, error) {
	pmfit := new(putModelFilterInitTemp)
	// 对field按照scope分组
	s, _, err := colScopeToMap(&extend)
	if err != nil {
		return "", err
	}
	pmfit.ModelName = strings.ToLower(extend.Name[0:1]) + extend.Name[1:]
	pmfit.ScopeToFields = s

	// 将json tag与bson tag映射
	jbm, err := jsonMapBson(extend)
	if err != nil {
		log.Warn(extend.Name, "no Feild")
	}
	pmfit.JBMap = jbm

	// 将json tag 与validate 做映射
	jmv, _ := jsonMapVali(extend)
	pmfit.ValiMap = jmv
	return getPutInit(*pmfit)
}

func jsonMapVali(extend ModelExtend) (jmv map[string]string, err error) {
	if len(extend.Fields) < 1 {
		return nil, errors.New("model no fieid")
	}
	smk := make(map[string]string)
	for _, v := range extend.Fields {
		var jtString string
		jt, err := v.Tags.Get(TAG_JSON)
		if err != nil {
			jtString = v.FieldName
		} else {
			jtString = jt.Name
		}
		var btString string
		bt, err := v.Tags.Get(TAG_BINDING)
		if err != nil {
			continue
		} else {
			btString = bt.Value()
		}
		smk[jtString] = btString
	}
	return smk, nil
}
func jsonMapBson(extend ModelExtend) (jmb map[string]string, err error) {
	if len(extend.Fields) < 1 {
		return nil, errors.New("model no fieid")
	}
	smk := make(map[string]string)
	for _, v := range extend.Fields {
		var jtString string
		jt, err := v.Tags.Get(TAG_JSON)
		if err != nil {
			jtString = v.FieldName
		} else {
			jtString = jt.Name
		}
		var btString string
		bt, err := v.Tags.Get(TAG_BSON)
		if err != nil {
			btString = strings.ToLower(v.FieldName)
		} else {
			btString = bt.Name
		}
		smk[jtString] = btString
	}
	return smk, nil
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
var _putCheckFInScopeTemp = `
func check{{.ModelNameF}}ValueOptInScope(valueKey string, scope int) bool {
	for k, v := range {{.ModelName}}ScopeMap {
		if k < scope {
			if arrays.ContainsString(v, valueKey) != -1 {
				return true
			}
		}
	}
	return false
}
`

func getModelPutInfo(extend ModelExtend) (string, error) {
	modelNameF := extend.Name
	modelName := strings.ToLower(modelNameF[0:1]) + modelNameF[1:]

	mn := putFunc4ModelTemp{ModelNameF: modelNameF, ModelName: modelName}
	pfmt, err := utils.ParserName(_putFunc4ModelTemp, mn)
	if err != nil {
		return "", err
	}

	pcit, err := utils.ParserName(_putCheckFInScopeTemp, mn)
	if err != nil {
		return "", err
	}
	return pfmt.String() + pcit.String(), nil
}
