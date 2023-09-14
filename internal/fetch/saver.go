package fetch

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"regexp"

	customErrors "example.com/fetch-web/pkg/errors"
)

func SaveFile(content []byte, url string) error {
	fileName := filepath.Base(url)
	fileName = sanitizeFileName(fileName)
	filePath := filepath.Join(".", fileName+".html")

	file, err := os.Create(filePath)
	if err != nil {
		log.Printf("Error creating file %s: %v", fileName, err)
		return fmt.Errorf("%w:%v", customErrors.ErrCreateFileError, err)
	}
	defer file.Close()

	_, err = io.WriteString(file, string(content))
	if err != nil {
		log.Printf("Error saving file %s: %v", fileName, err)
		return fmt.Errorf("%w:%v", customErrors.ErrSaveFileError, err)
	}

	return nil
}

func sanitizeFileName(fileName string) string {
	return regexp.MustCompile(`[^\w.-]`).ReplaceAllString(fileName, "_")
}
