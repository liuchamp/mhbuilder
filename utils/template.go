package utils

import (
	"bytes"
	"text/template"
)

// tp 模版， ts是要解析的数据
func ParserName(tp string, ts interface{}) (*bytes.Buffer, error) {
	t, err := template.New("").Parse(tp)
	if err != nil {
		return nil, err
	}
	var buf bytes.Buffer
	if err = t.Execute(&buf, ts); err != nil {
		return nil, err
	}
	return &buf, nil
}
