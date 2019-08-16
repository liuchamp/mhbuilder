package parser

import (
	"go/ast"
	"go/token"
)

/*
做全局依赖解析，解析一个文件的所有依赖
*/

type DeepSearch interface {
	// 读取一个结构体的所有依赖包
	GetAllDependence(s *ast.File) (token.FileSet, error)
	// 获取一个文件的依赖包
	FindDependent(s *ast.File) ([]ast.Package, error)
}

type DeConfig struct {
	Pkg      string
	IsDonVer bool
}

type deepSearch struct {
	config *DeConfig
}

func NewDeepSearch(config *DeConfig) *deepSearch {
	return &deepSearch{config: config}
}

func (ds deepSearch) GetAllDependence(s *ast.File) (token.FileSet, error) {
	panic("implement me")
}

func (ds deepSearch) FindDependent(s *ast.File) ([]ast.Package, error) {
	panic("implement me")
}
