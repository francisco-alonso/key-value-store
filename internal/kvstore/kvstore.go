package kvstore

import (
	"errors"
	"sync"
)

type KeyValueStore struct {
	store map[string]string
	mu    sync.RWMutex
}

// NewKeyValueStore initializes a new key-value store
func NewKeyValueStore() *KeyValueStore {
	return &KeyValueStore{
		store: make(map[string]string),
	}
}

// Set stores or updates a key-value pair
func (kv *KeyValueStore) Set(key, value string) {
	kv.mu.Lock()
	defer kv.mu.Unlock()
	kv.store[key] = value
}

// Get retrieves the value associated with a key
func (kv *KeyValueStore) Get(key string) (string, error) {
	kv.mu.RLock()
	defer kv.mu.RUnlock()
	value, exists := kv.store[key]
	if !exists {
		return "", errors.New("key not found")
	}
	return value, nil
}

// Delete removes a key-value pair
func (kv *KeyValueStore) Delete(key string) error {
	kv.mu.Lock()
	defer kv.mu.Unlock()
	if _, exists := kv.store[key]; !exists {
		return errors.New("key not found")
	}
	delete(kv.store, key)
	return nil
}

// Exists checks if a key exists in the store
func (kv *KeyValueStore) Exists(key string) bool {
	kv.mu.RLock()
	defer kv.mu.RUnlock()
	_, exists := kv.store[key]
	return exists
}
