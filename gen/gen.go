package gen

import (
	"fmt"
	"os"
	"path/filepath"
)

var BUILDMODEFILE = false

type Gen struct {
}

func New() *Gen {
	return &Gen{}
}

type Config struct {
	SearchDir          string
	OutputDir          string
	PropNamingStrategy string
	ParseVendor        bool
	ParseDependency    bool
	Files              []string
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
	p := &Parser{}
	p.PropNamingStrategy = config.PropNamingStrategy
	p.ParseVendor = config.ParseVendor
	p.ParseDependency = config.ParseVendor

	if err := p.ParModel(config.SearchDir); err != nil {
		return err
	}

	// 将parer中数据写入对应的结构体，然后生成go代码
	if err := os.MkdirAll(config.OutputDir, os.ModePerm); err != nil {
		return err
	}
	return nil
}
