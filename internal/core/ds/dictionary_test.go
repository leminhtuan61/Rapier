package ds

import (
	"testing"
	"time"
)

func TestDictionary_SetGet(t *testing.T) {
	dict := NewDictionary()
	
	// Test basic set/get
	dict.Set("key1", "value1", 0)
	val, exists := dict.Get("key1")
	if !exists || val != "value1" {
		t.Errorf("Expected value1, got %v", val)
	}
}

func TestDictionary_Expire(t *testing.T) {
	dict := NewDictionary()
	
	// Set with expire time (1 second from now)
	expireTime := uint64(time.Now().Add(1*time.Second).UnixMilli())
	dict.Set("key2", "value2", expireTime)
	
	// Should exist immediately
	val, exists := dict.Get("key2")
	if !exists || val != "value2" {
		t.Errorf("Expected value2, got %v", val)
	}
	
	// Wait for expiration
	time.Sleep(1100 * time.Millisecond)
	
	// Should be expired
	if !dict.HasExpired("key2") {
		t.Error("Key should be expired")
	}
}

func TestDictionary_Delete(t *testing.T) {
	dict := NewDictionary()
	
	dict.Set("key3", "value3", 0)
	dict.Delete("key3")
	
	_, exists := dict.Get("key3")
	if exists {
		t.Error("Key should be deleted")
	}
}