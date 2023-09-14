// metadata/metadata_test.go
package metadata

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestParseMetadata(t *testing.T) {
	htmlContent := []byte(`
		<html>
			<head>
				<title>Example Page</title>
			</head>
			<body>
				<a href="https://example.com">Link 1</a>
				<a href="https://example.com">Link 2</a>
				<img src="image1.jpg" alt="Image 1">
				<img src="image2.jpg" alt="Image 2">
			</body>
		</html>
	`)

	expectedMetadata := Metadata{
		LastFetchTime: time.Now().Format("Mon Jan 02 2006 15:04 MST"),
		NumLinks:      2,
		NumImages:     2,
	}

	actualMetadata := ParseMetadata(htmlContent)

	assert.Equal(t, expectedMetadata.LastFetchTime, actualMetadata.LastFetchTime)
	assert.Equal(t, expectedMetadata.NumLinks, actualMetadata.NumLinks)
	assert.Equal(t, expectedMetadata.NumImages, actualMetadata.NumImages)
}

func TestFormatMetadata(t *testing.T) {
	metadata := Metadata{
		LastFetchTime: "Mon Jan 02 2006 15:04 MST",
		NumLinks:      2,
		NumImages:     2,
	}

	expectedFormattedMetadata := "last_fetch: Mon Jan 02 2006 15:04 MST\nnum_links: 2\nnum_images: 2"

	actualFormattedMetadata := FormatMetadata(metadata)

	assert.Equal(t, expectedFormattedMetadata, actualFormattedMetadata)
}
