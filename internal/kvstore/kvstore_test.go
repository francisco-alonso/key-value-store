package kvstore

import (
	"testing"
)

func TestKeyValueStore_SetAndGet(t *testing.T) {
	kv := NewKeyValueStore()

	// Test setting a key-value pair
	kv.Set("key1", "value1")

	// Test getting the value
	value, err := kv.Get("key1")
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	if value != "value1" {
		t.Errorf("expected value 'value1', got %s", value)
	}
}

func TestKeyValueStore_GetNonExistentKey(t *testing.T) {
	kv := NewKeyValueStore()

	// Test getting a non-existent key
	_, err := kv.Get("nonexistent")
	if err == nil {
		t.Errorf("expected error, got nil")
	}
}

func TestKeyValueStore_Delete(t *testing.T) {
	kv := NewKeyValueStore()

	// Set and delete a key
	kv.Set("key1", "value1")
	err := kv.Delete("key1")
	if err != nil {
		t.Errorf("expected no error on delete, got %v", err)
	}

	// Test getting a deleted key
	_, err = kv.Get("key1")
	if err == nil {
		t.Errorf("expected error for deleted key, got nil")
	}
}

func TestKeyValueStore_DeleteNonExistentKey(t *testing.T) {
	kv := NewKeyValueStore()

	// Attempt to delete a non-existent key
	err := kv.Delete("nonexistent")
	if err == nil {
		t.Errorf("expected error, got nil")
	}
}

func TestKeyValueStore_Exists(t *testing.T) {
	kv := NewKeyValueStore()

	// Check existence of a non-existent key
	if kv.Exists("key1") {
		t.Errorf("expected key1 to not exist")
	}

	// Set and check existence
	kv.Set("key1", "value1")
	if !kv.Exists("key1") {
		t.Errorf("expected key1 to exist")
	}

	// Delete and check existence
	kv.Delete("key1")
	if kv.Exists("key1") {
		t.Errorf("expected key1 to not exist after delete")
	}
}