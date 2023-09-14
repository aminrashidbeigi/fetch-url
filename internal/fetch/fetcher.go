package fetch

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"

	customErrors "example.com/fetch-web/pkg/errors"
)

func FetchWebPage(url string) ([]byte, error) {
	response, err := http.Get(url)
	if err != nil {
		log.Printf("Error fetching %s: %v", url, err)
		return nil, fmt.Errorf("%w:%v", customErrors.ErrFetchError, err)
	}
	defer response.Body.Close()

	buf := new(bytes.Buffer)
	_, err = io.Copy(buf, response.Body)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
