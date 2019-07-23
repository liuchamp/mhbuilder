package gen

import (
	"fmt"
	"testing"
)

const SEARCHDIR = "../testdata/models"

func TestParser_ParModel(t *testing.T) {
	p := NewParser()
	searchDir := SEARCHDIR
	err := p.ParModel(searchDir)
	if err != nil {
		fmt.Println(err)
	}
}

func TestNew(t *testing.T) {
	gen := New()
	c := &Config{
		SearchDir:          SEARCHDIR,
		OutputDir:          SEARCHDIR,
		PropNamingStrategy: "camelcase",
		Files:              nil,
		ParseVendor:        false,
		ParseDependency:    false,
	}
	t.Run("withNotParse", func(t *testing.T) {
		if err := gen.Build(c); err != nil {
			t.Error(err)
		}
	})
	t.Run("withParse", func(t *testing.T) {
		c.ParseDependency = true
		c.ParseVendor = true
		if err := gen.Build(c); err != nil {
			t.Error(err)
		}
	})

}
