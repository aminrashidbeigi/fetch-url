package fetch

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSaveFile(t *testing.T) {
	content := []byte("Test content")

	url := "https://example.com"
	err := SaveFile(content, url)
	assert.NoError(t, err)

	fileName := sanitizeFileName(filepath.Base(url)) + ".html"
	filePath := filepath.Join(".", fileName)
	defer os.Remove(filePath)

	fileContent, err := os.ReadFile(filePath)
	assert.NoError(t, err)
	assert.Equal(t, string(content), string(fileContent))
}

func TestSanitizeFileName(t *testing.T) {
	fileName := "example.com.html"
	sanitized := sanitizeFileName(fileName)
	assert.Equal(t, fileName, sanitized)

	fileNameWithInvalidChars := "example.com?.html"
	sanitized = sanitizeFileName(fileNameWithInvalidChars)
	assert.Equal(t, "example.com_.html", sanitized)
}
