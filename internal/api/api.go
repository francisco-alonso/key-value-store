package api

import (
	"encoding/json"

	"github.com/francisco-alonso/key-value-store/internal/kvstore"

	"net/http"
)

// SetHandler handles POST /set to store or update a key-value pair
func SetHandler(kv *kvstore.KeyValueStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req map[string]string
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid input", http.StatusBadRequest)
			return
		}

		key, exists := req["key"]
		if !exists {
			http.Error(w, "Key is required", http.StatusBadRequest)
			return
		}

		kv.Set(key, req["value"])
		w.WriteHeader(http.StatusOK)
	}
}

// GetHandler handles GET /get to retrieve the value for a key
func GetHandler(kv *kvstore.KeyValueStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		key := r.URL.Query().Get("key")
		if key == "" {
			http.Error(w, "Key is required", http.StatusBadRequest)
			return
		}

		value, err := kv.Get(key)
		if err != nil {
			http.Error(w, "Key not found", http.StatusNotFound)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"value": value})
	}
}

// DeleteHandler handles DELETE /delete to remove a key-value pair
func DeleteHandler(kv *kvstore.KeyValueStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		key := r.URL.Query().Get("key")
		if key == "" {
			http.Error(w, "Key is required", http.StatusBadRequest)
			return
		}

		err := kv.Delete(key)
		if err != nil {
			http.Error(w, "Key not found", http.StatusNotFound)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}

// ExistsHandler handles GET /exists to check if a key exists
func ExistsHandler(kv *kvstore.KeyValueStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		key := r.URL.Query().Get("key")
		if key == "" {
			http.Error(w, "Key is required", http.StatusBadRequest)
			return
		}

		if kv.Exists(key) {
			w.WriteHeader(http.StatusOK)
		} else {
			http.Error(w, "Key not found", http.StatusNotFound)
		}
	}
}
