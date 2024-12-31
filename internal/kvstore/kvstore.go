package kvstore

import "sync"

type KeyValueStore struct {
	store map[string]string
	mu    sync.RWMutex
}

func NewKeyValueStore() *KeyValueStore {
	return &KeyValueStore{
		store: make(map[string]string),
	}
}