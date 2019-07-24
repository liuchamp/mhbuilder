package builder

import (
	"errors"
	"github.com/liuchamp/mhbuilder/log"
	"github.com/liuchamp/mhbuilder/utils"
	"path"
)

func (builder *Builder) outAddDtoAndToModel(file string) (string, error) {
	filemap, ok := builder.FilesMap[file]
	pkg, err := utils.GetPkgName(path.Dir(file))
	if err != nil {
		return "", nil
	}

	if !ok {
		return "", errors.New("file not find")
	}
	var bodys []string
	for _, modelDetail := range filemap.Models {
		var fields []string
		var toModels []string
		for _, v := range modelDetail.Fields {
			field, toModel, err := addFeildString(&v)
			if err != nil {
				continue
			}
			log.Debug("Feildname:", v.FieldName)
			fields = append(fields, field)
			toModels = append(toModels, toModel)
		}
		if fields != nil || len(fields) > 0 {
			admot := new(addDtoTemplate)
			admot.StructName = modelDetail.Name + POSTTOSUFFIX
			admot.Feilds = fields

			admotCode, err := utils.ParserName(_addDtoTemplate, admot)
			if err == nil {
				bodys = append(bodys, admotCode.String())
			}

		}
		if toModels != nil {
			adtmt := new(addDtoToModelTemplate)
			adtmt.StructName = modelDetail.Name + POSTTOSUFFIX
			adtmt.Model = modelDetail.Name
			adtmt.Fields = toModels
			adtmtCode, err := utils.ParserName(_addDtoToModelTemplate, adtmt)
			if err == nil {
				bodys = append(bodys, adtmtCode.String())
			}
		}

	}
	if bodys == nil || len(bodys) < 1 {
		return "", NOBODY
	}

	fileHeader := new(headerTemplate)
	fileHeader.PkgName = BUILD_POST
	fileHeader.ImportPackage = pkg
	headerBuffer, err := utils.ParserName(_headerTemplate, fileHeader)
	if err != nil {
		return "", err
	}
	fileOut := new(addFile)
	fileOut.Body = bodys
	fileOut.FileHeader = headerBuffer.String()
	fileBuffer, err := utils.ParserName(_addFile, fileOut)
	if err != nil {
		return "", err
	}

	return fileBuffer.String(), nil
}
