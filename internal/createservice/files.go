package createservice

import (
	"fmt"
	"os"
)

func createFile(path, content string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}

	defer func() {
		if err = file.Close(); err != nil {
			fmt.Printf("error closing file: %v", err)
		}
	}()

	_, err = file.WriteString(content)
	return err
}

func createFiles(pathsContent map[string]string) error {
	for path, content := range pathsContent {
		err := createFile(path, content)
		if err != nil {
			return err
		}
	}

	return nil
}
