package server

import (
	"fmt"

	"github.com/francisco-alonso/key-value-store/internal/api"

	"github.com/francisco-alonso/key-value-store/internal/kvstore"

	"net/http"
)

// StartAPI starts the HTTP server for the API
func StartAPI(kv *kvstore.KeyValueStore) {
	http.HandleFunc("/set", api.SetHandler(kv))
	http.HandleFunc("/get", api.GetHandler(kv))
	http.HandleFunc("/delete", api.DeleteHandler(kv))
	http.HandleFunc("/exists", api.ExistsHandler(kv))

	fmt.Println("Starting server on :8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Failed to start server:", err)
	}
}
