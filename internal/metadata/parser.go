package metadata

import (
	"fmt"
	"time"

	"example.com/fetch-web/internal/regex"
)

type Metadata struct {
	LastFetchTime string
	NumLinks      int
	NumImages     int
}

func ParseMetadata(content []byte) Metadata {
	contentString := string(content)

	return Metadata{
		LastFetchTime: time.Now().Format("Mon Jan 02 2006 15:04 MST"),
		NumLinks:      regex.CountUrls(contentString),
		NumImages:     regex.CountImages(contentString),
	}
}

func FormatMetadata(metadata Metadata) string {
	return fmt.Sprintf(
		"last_fetch: %s\nnum_links: %d\nnum_images: %d",
		metadata.LastFetchTime,
		metadata.NumLinks,
		metadata.NumImages,
	)

}
