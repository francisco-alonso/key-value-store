package server

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/francisco-alonso/key-value-store/internal/api"
	"github.com/francisco-alonso/key-value-store/internal/kvstore"
)

func TestStartAPI(t *testing.T) {
	// Initialize the KeyValueStore kvstore
	kvStore, _ := kvstore.NewKeyValueStore("./filePath")

	// Set up the handlers
	http.HandleFunc("/set", api.SetHandler(kvStore))
	http.HandleFunc("/get", api.GetHandler(kvStore))
	http.HandleFunc("/delete", api.DeleteHandler(kvStore))
	http.HandleFunc("/exists", api.ExistsHandler(kvStore))

	// Create a mock HTTP server
	ts := httptest.NewServer(nil)
	defer ts.Close()

	// Test cases with proper request body for POST /set
	tests := []struct {
		method     string
		url        string
		body       string
		statusCode int
	}{
		{
			method:     "POST",
			url:        ts.URL + "/set",
			body:       `{"key": "foo", "value": "bar"}`, // valid request body
			statusCode: http.StatusOK,
		},
		{
			method:     "GET",
			url:        ts.URL + "/get?key=foo",
			body:       "",
			statusCode: http.StatusOK, // Expecting 404 as key "foo" will not be set yet
		},
		{
			method:     "DELETE",
			url:        ts.URL + "/delete?key=foo",
			body:       "",
			statusCode: http.StatusOK, // Expecting 404 as key "foo" will not be set yet
		},
		{
			method:     "GET",
			url:        ts.URL + "/exists?key=foo",
			body:       "",
			statusCode: http.StatusNotFound, // Expecting 404 as key "foo" will not be set yet
		},
	}

	// Loop through each test case
	for _, tt := range tests {
		// Create the request with a body if necessary
		var req *http.Request
		if tt.body != "" {
			req = httptest.NewRequest(tt.method, tt.url, strings.NewReader(tt.body))
			req.Header.Set("Content-Type", "application/json")
		} else {
			req = httptest.NewRequest(tt.method, tt.url, nil)
		}

		// Create a response recorder to capture the response
		rr := httptest.NewRecorder()

		// Call the handler
		http.DefaultServeMux.ServeHTTP(rr, req)

		// Check the response status code
		if rr.Code != tt.statusCode {
			t.Errorf("expected status %d, got %d", tt.statusCode, rr.Code)
		}
	}
}
