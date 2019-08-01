package gen

import (
	"testing"
)

const SEARCHDIR = "/Users/mac/Code/go/src/github.com/liuchamp/mhbuilder/testdata/models"

func TestNew(t *testing.T) {
	gen := New()
	c := &Config{
		SearchDir:            SEARCHDIR,
		OutputDir:            SEARCHDIR,
		PropNamingStrategy:   "camelcase",
		RelathionPutAndPatch: "None",
		Files:                nil,
		ParseVendor:          false,
		ParseDependency:      false,
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
