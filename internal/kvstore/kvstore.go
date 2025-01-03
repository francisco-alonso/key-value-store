package kvstore

import (
	"encoding/json"
	"errors"
	"log"
	"os"
	"sync"
)

type KeyValueStore struct {
	store map[string]string
	mu    sync.RWMutex
	filePath string
}

// NewKeyValueStore initializes a new key-value store
func NewKeyValueStore(filePath string) (*KeyValueStore, error) {
	kv := &KeyValueStore{
		store: make(map[string]string),
		filePath: filePath,
	}
	
	if err := kv.load(); err != nil {
		return nil, err
	}
	
	return kv, nil
}

// Set stores or updates a key-value pair
func (kv *KeyValueStore) Set(key, value string) error {
	kv.mu.Lock()
	defer kv.mu.Unlock()
	kv.store[key] = value
	
	err := kv.Save()
	
	if err != nil {
		log.Printf("unable to save data: %v", err)
		return err
	}
	return kv.Save()
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
	err := kv.Save()
	
	if err != nil {
		log.Printf("unable to save data: %v", err)
		return err
	}
	return kv.Save()
}

// Exists checks if a key exists in the store
func (kv *KeyValueStore) Exists(key string) bool {
	kv.mu.RLock()
	defer kv.mu.RUnlock()
	_, exists := kv.store[key]
	return exists
}

// load reads the key-value store from the file
func (kv *KeyValueStore) load() error {
	// Open the file
	file, err := os.OpenFile(kv.filePath, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	defer file.Close()

	// If the file is empty, just return
	if stat, _ := file.Stat(); stat.Size() == 0 {
		return nil
	}

	// Deserialize data from the file into the store
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&kv.store); err != nil {
		return err
	}
	return nil
}

func (kv *KeyValueStore) Save() error {
	// Open the file
	file, err := os.OpenFile(kv.filePath, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	defer file.Close()
	
	// Serialize the store into JSON and write it to the file
	encoder := json.NewEncoder(file)
	if err := encoder.Encode(kv.store); err != nil {
		return err
	}
	return nil
}
