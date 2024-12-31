package api

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/francisco-alonso/key-value-store/internal/kvstore"
)

func TestSetHandler(t *testing.T) {
	kvStore := kvstore.NewKeyValueStore()
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
			statusCode: http.StatusBadRequest,
		},
		{
			body:       `{"value": "bar"}`,
			statusCode: http.StatusBadRequest,
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
	kvStore := kvstore.NewKeyValueStore()
	kvStore.Set("foo", "bar") // Set a key-value pair to test

	handler := GetHandler(kvStore)

	tests := []struct {
		key        string
		expected   string
		statusCode int
	}{
		{
			key:        "foo",
			expected:   `{"value":"bar"}`,
			statusCode: http.StatusOK,
		},
		{
			key:        "nonexistent",
			expected:   `{"value":""}`,
			statusCode: http.StatusNotFound,
		},
	}

	for _, tt := range tests {
		req := httptest.NewRequest("GET", "/get?key="+tt.key, nil)
		rr := httptest.NewRecorder()

		handler.ServeHTTP(rr, req)

		if rr.Code != tt.statusCode {
			t.Errorf("expected status %d, got %d", tt.statusCode, rr.Code)
		}

		if rr.Body.String() != tt.expected {
			t.Errorf("expected body %s, got %s", tt.expected, rr.Body.String())
		}
	}
}

func TestDeleteHandler(t *testing.T) {
	kvStore := kvstore.NewKeyValueStore()
	kvStore.Set("foo", "bar") // Set a key-value pair to test

	handler := DeleteHandler(kvStore)

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
		req := httptest.NewRequest("DELETE", "/delete?key="+tt.key, nil)
		rr := httptest.NewRecorder()

		handler.ServeHTTP(rr, req)

		if rr.Code != tt.statusCode {
			t.Errorf("expected status %d, got %d", tt.statusCode, rr.Code)
		}
	}
}

func TestExistsHandler(t *testing.T) {
	kvStore := kvstore.NewKeyValueStore()
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
