package fetch

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFetchWebPage(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response := "Mock HTML Content"
		w.Header().Set("Content-Type", "text/html")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(response))
	}))
	defer testServer.Close()

	testURL := testServer.URL

	content, err := FetchWebPage(testURL)
	assert.NoError(t, err)
	assert.Equal(t, []byte("Mock HTML Content"), content)
}
