package utils

import (
	"bytes"
	"text/template"
)

// tp 模版， ts是要解析的数据
func ParserName(tp string, ts interface{}) (*bytes.Buffer, error) {
	t, err := template.New("").Funcs(template.FuncMap{"lenfxs": lenfxs, "lenfxi": lenfxi}).Parse(tp)
	if err != nil {
		return nil, err
	}
	var buf bytes.Buffer
	if err = t.Execute(&buf, ts); err != nil {
		return nil, err
	}
	return &buf, nil
}

// 算出数组len-1的使用场景 string
func lenfxs(v []string) int {
	return len(v) - 1
}
func lenfxi(v []int) int {
	return len(v) - 1
}
