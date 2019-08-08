package utils

import (
	"os"
)

func FileOuter(filePath string, content string) error {

	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	var data []byte = []byte(content)
	file.Write(data)
	file.Close()

	return nil
}
