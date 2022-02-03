package util

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"path/filepath"
)

// NewMockServerFromFixture create a new HTTP test server using given fixture as JSON response.
func NewMockServerFromFixture(dir, file string) *httptest.Server {
	data, _ := ioutil.ReadFile(filepath.Join("test", dir, "fixtures", file))
	return httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
			w.Header().Set("content-type", "application/json")
			w.Write(data)
		}),
	)
}
