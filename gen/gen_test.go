package gen

import (
	"fmt"
	"testing"
)

func TestParser_ParModel(t *testing.T) {
	p := NewParser()
	searchDir := "/Users/mac/Code/go/src/github.com/liuchamp/godemo/models"
	err := p.ParModel(searchDir)
	if err != nil {
		fmt.Println(err)
	}
}
