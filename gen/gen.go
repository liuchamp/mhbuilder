package gen

import (
	"fmt"
	"github.com/liuchamp/mhbuilder/builder"
	"github.com/liuchamp/mhbuilder/log"
	"github.com/liuchamp/mhbuilder/parser"
	"github.com/liuchamp/mhbuilder/utils"
	"os"
	"path"
	"path/filepath"
)

var BUILDMODEFILE = false

type Gen struct {
}

func New() *Gen {
	return &Gen{}
}

type Config struct {
	SearchDir            string
	OutputDir            string
	PropNamingStrategy   string
	ParseVendor          bool
	ParseDependency      bool
	RelathionPutAndPatch string
	Files                []string
}

func (g *Gen) Build(config *Config) error {
	if len(config.Files) > 0 {
		BUILDMODEFILE = true

	}
	if _, err := os.Stat(config.SearchDir); os.IsNotExist(err) {
		return fmt.Errorf("dir: %s is not exist", config.SearchDir)
	}
	if BUILDMODEFILE {
		for _, f := range config.Files {
			filePathString := filepath.Join(config.SearchDir, f)
			if _, err := os.Stat(filePathString); os.IsNotExist(err) {
				return fmt.Errorf("file: %s is not exist, \n path: %s", f, filePathString)
			}
		}
	}

	// 解析文件夹，文档，并写入Parser对象中
	p := parser.NewParser()
	p.PropNamingStrategy = config.PropNamingStrategy
	p.ParseVendor = config.ParseVendor
	p.ParseDependency = config.ParseVendor

	if err := p.ParModel(config.SearchDir); err != nil {
		return err
	}

	// 将parer中数据写入对应的结构体，然后生成go代码
	if err := os.MkdirAll(filepath.Join(config.OutputDir, builder.BUILD_PUT), os.ModePerm); err != nil {
		return err
	}
	if err := os.MkdirAll(filepath.Join(config.OutputDir, builder.BUILD_PATCH), os.ModePerm); err != nil {
		return err
	}
	if err := os.MkdirAll(filepath.Join(config.OutputDir, builder.BUILD_FILTER), os.ModePerm); err != nil {
		return err
	}
	if err := os.MkdirAll(filepath.Join(config.OutputDir, builder.BUILD_POST), os.ModePerm); err != nil {
		return err
	}

	// 将数据写入文件
	for k, v := range p.GetFileMap() {
		// 解析文件，导入自定义的代码中
		b := builder.NewBuilder(p.PkgName, k, v)
		// 到对应目录文件中
		uor, err := b.WirteFile()
		if err != nil {
			log.Error(err)
			return err
		}
		fileName := path.Base(k)
		err = utils.FileOuter(filepath.Join(config.OutputDir, builder.BUILD_POST, fileName), uor.Post)
		if err != nil {
			log.Error("写入文件错误", filepath.Join(config.OutputDir, builder.BUILD_POST, fileName))
			return err
		}
		err = utils.FileOuter(filepath.Join(config.OutputDir, builder.BUILD_FILTER, fileName), uor.Filter)
		if err != nil {
			log.Error("写入文件错误", filepath.Join(config.OutputDir, builder.BUILD_POST, fileName))
			return err
		}
	}

	return nil
}
