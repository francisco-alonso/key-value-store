package server

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/francisco-alonso/key-value-store/internal/api"

	"github.com/francisco-alonso/key-value-store/internal/kvstore"
)

func TestStartAPI(t *testing.T) {
	// Initialize the KeyValueStore kvstore
	kvStore := kvstore.NewKeyValueStore()
	// Set up the handlers
	http.HandleFunc("/set", api.SetHandler(kvStore))
	http.HandleFunc("/get", api.GetHandler(kvStore))
	http.HandleFunc("/delete", api.DeleteHandler(kvStore))
	http.HandleFunc("/exists", api.ExistsHandler(kvStore))

	// Create a mock HTTP server
	ts := httptest.NewServer(nil)
	defer ts.Close()

	tests := []struct {
		method     string
		url        string
		statusCode int
	}{
		{
			method:     "POST",
			url:        ts.URL + "/set",
			statusCode: http.StatusOK,
		},
		{
			method:     "GET",
			url:        ts.URL + "/get?key=foo",
			statusCode: http.StatusNotFound,
		},
		{
			method:     "DELETE",
			url:        ts.URL + "/delete?key=foo",
			statusCode: http.StatusNotFound,
		},
		{
			method:     "GET",
			url:        ts.URL + "/exists?key=foo",
			statusCode: http.StatusNotFound,
		},
	}

	for _, tt := range tests {
		req := httptest.NewRequest(tt.method, tt.url, nil)
		rr := httptest.NewRecorder()

		http.DefaultServeMux.ServeHTTP(rr, req)

		if rr.Code != tt.statusCode {
			t.Errorf("expected status %d, got %d", tt.statusCode, rr.Code)
		}
	}
}
