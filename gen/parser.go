package gen

import (
	"fmt"
	"github.com/KyleBanks/depth"
	"github.com/liuchamp/mhbuilder/builder"
	"github.com/liuchamp/mhbuilder/log"
	"github.com/liuchamp/mhbuilder/utils"
	"go/ast"
	goparser "go/parser"
	"go/token"
	"io/ioutil"
	"os"
	"path/filepath"
)

type Parser struct {
	files map[string]*ast.File
	Build *builder.Builder
	// registerTypes is a map that stores [refTypeName][*ast.TypeSpec]
	registerTypes map[string]*ast.TypeSpec

	PropNamingStrategy string
	PkgName            string

	ParseVendor bool

	// ParseDependencies whether swag should be parse outside dependency folder
	ParseDependency bool

	// structStack stores full names of the structures that were already parsed or are being parsed now
	structStack []string
}

func NewParser() *Parser {
	return &Parser{
		files: make(map[string]*ast.File),
		Build: builder.NewBuilder(),
	}
}

func (parser *Parser) ParModel(searchDir string) error {
	log.Debug("Generate general API Info, search dir: ", searchDir)
	if err := parser.getAllGoFileInfo(searchDir); err != nil {
		return err
	}
	pkgName, err := utils.GetPkgName(searchDir)
	if err != nil {
		return err
	}
	var t depth.Tree

	if err := t.Resolve(pkgName); err != nil {
		return fmt.Errorf("pkg %s cannot find all dependencies, %s", pkgName, err)
	}
	parser.PkgName = pkgName
	if parser.ParseDependency {
		for i := 0; i < len(t.Root.Deps); i++ {
			if err := parser.getAllGoFileInfoFromDeps(&t.Root.Deps[i]); err != nil {
				return err
			}
		}
	}
	if parser.files == nil {
		log.Errorln("not find name")
		return nil
	}
	// 开始解析文件
	err = parser.ExtentsFile()
	if err != nil {
		return err
	}

	return nil
}

// 解析文件，并且将文件中struct写入对应的结构体中
func (parser *Parser) ExtentsFile() error {
	if parser.files == nil || len(parser.files) < 1 {
		return fmt.Errorf("not find file")
	}
	for k, v := range parser.files {
		err := parser.Build.ExtentsFileInfo(k, parser.PkgName, v)
		if err != nil {
			log.Errorf("Can not Extend file: %s in package: %s", k, parser.PkgName)
			return err
		}
	}
	return nil
}

func (parser *Parser) getAllGoFileInfo(searchDir string) error {
	return filepath.Walk(searchDir, parser.visit)
}

func (parser *Parser) visit(path string, f os.FileInfo, err error) error {
	if err := parser.Skip(path, f); err != nil {
		return err
	}
	return parser.parseFile(path)
}

func (parser *Parser) parseFile(path string) error {
	if ext := filepath.Ext(path); ext == ".go" {
		fset := token.NewFileSet() // positions are relative to fset
		astFile, err := goparser.ParseFile(fset, path, nil, goparser.ParseComments)
		if err != nil {
			return fmt.Errorf("ParseFile error:%+v", err)
		}

		parser.files[path] = astFile
	}
	return nil
}

// Skip returns filepath.SkipDir error if match vendor and hidden folder
func (parser *Parser) Skip(path string, f os.FileInfo) error {

	if !parser.ParseVendor { // ignore vendor
		if f.IsDir() && f.Name() == "vendor" {
			return filepath.SkipDir
		}
	}
	// exclude all hidden folder
	if f.IsDir() && len(f.Name()) > 1 && f.Name()[0] == '.' {
		return filepath.SkipDir
	}
	return nil
}

func (parser *Parser) getAllGoFileInfoFromDeps(pkg *depth.Pkg) error {
	if pkg.Internal || !pkg.Resolved { // ignored internal and not resolved dependencies
		return nil
	}

	files, err := ioutil.ReadDir(pkg.SrcDir) // only parsing files in the dir(don't contains sub dir files)
	if err != nil {
		return err
	}

	for _, f := range files {
		if f.IsDir() {
			continue
		}

		path := filepath.Join(pkg.SrcDir, f.Name())
		if err := parser.parseFile(path); err != nil {
			return err
		}
	}

	for i := 0; i < len(pkg.Deps); i++ {
		if err := parser.getAllGoFileInfoFromDeps(&pkg.Deps[i]); err != nil {
			return err
		}
	}

	return nil
}

func (parser *Parser) OutFile() error {
	//build:= parser.Build
	// 导出post

	// 导出 put

	// 导出 patch

	// 导出 filter

	return nil
}
