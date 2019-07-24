package gen

import (
	"fmt"
	"testing"
)

func TestParser_ParModel(t *testing.T) {
	p := NewParser()
	searchDir := SEARCHDIR
	err := p.ParModel(searchDir)
	if err != nil {
		fmt.Println(err)
	}
}
