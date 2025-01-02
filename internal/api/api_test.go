package api

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/francisco-alonso/key-value-store/internal/kvstore"
)

func TestSetHandler(t *testing.T) {
	kvStore, _ := kvstore.NewKeyValueStore("./filePath")
	handler := SetHandler(kvStore)

	tests := []struct {
		body       string
		statusCode int
	}{
		{
			body:       `{"key": "foo", "value": "bar"}`,
			statusCode: http.StatusOK,
		},
		{
			body:       `{"key": "foo"}`,
			statusCode: http.StatusOK,
		},
	}

	for _, tt := range tests {
		req := httptest.NewRequest("POST", "/set", strings.NewReader(tt.body))
		req.Header.Set("Content-Type", "application/json")
		rr := httptest.NewRecorder()

		handler.ServeHTTP(rr, req)

		if rr.Code != tt.statusCode {
			t.Errorf("expected status %d, got %d", tt.statusCode, rr.Code)
		}
	}
}

func TestGetHandler(t *testing.T) {
	kvStore, _ := kvstore.NewKeyValueStore("./filePath")

	// Insert some test data
	kvStore.Set("foo", "bar")

	tests := []struct {
		name           string
		method         string
		url            string
		expectedStatus int
		expectedValue  string
	}{
		{
			name:           "Valid GET request",
			method:         "GET",
			url:            "/get?key=foo",
			expectedStatus: http.StatusOK,
			expectedValue:  "bar",
		},
		{
			name:           "Key not found",
			method:         "GET",
			url:            "/get?key=nonexistent",
			expectedStatus: http.StatusNotFound,
			expectedValue:  "",
		},
		{
			name:           "Missing key parameter",
			method:         "GET",
			url:            "/get",
			expectedStatus: http.StatusBadRequest,
			expectedValue:  "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a new HTTP request for the GET handler
			req, err := http.NewRequest(tt.method, tt.url, nil)
			if err != nil {
				t.Fatalf("Failed to create request: %v", err)
			}

			// Record the HTTP response
			rr := httptest.NewRecorder()

			// Create the handler and invoke it
			handler := GetHandler(kvStore)
			handler.ServeHTTP(rr, req)

			// Check the status code
			if status := rr.Code; status != tt.expectedStatus {
				t.Errorf("Expected status %d, got %v", tt.expectedStatus, status)
			}

			// If the status is OK, check the response body
			if tt.expectedStatus == http.StatusOK {
				var response map[string]string
				if err := json.NewDecoder(rr.Body).Decode(&response); err != nil {
					t.Errorf("Failed to decode response body: %v", err)
				}

				if value, ok := response["value"]; !ok || value != tt.expectedValue {
					t.Errorf("Expected value '%s', got %v", tt.expectedValue, value)
				}
			}
		})
	}
}

func TestExistsHandler(t *testing.T) {
	kvStore, _ := kvstore.NewKeyValueStore("./filePath")
	kvStore.Set("foo", "bar") // Set a key-value pair to test

	handler := ExistsHandler(kvStore)

	tests := []struct {
		key        string
		statusCode int
	}{
		{
			key:        "foo",
			statusCode: http.StatusOK,
		},
		{
			key:        "nonexistent",
			statusCode: http.StatusNotFound,
		},
	}

	for _, tt := range tests {
		req := httptest.NewRequest("GET", "/exists?key="+tt.key, nil)
		rr := httptest.NewRecorder()

		handler.ServeHTTP(rr, req)

		if rr.Code != tt.statusCode {
			t.Errorf("expected status %d, got %d", tt.statusCode, rr.Code)
		}
	}
}
