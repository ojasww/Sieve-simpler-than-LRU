package main

import (
	"testing"
)

// Helper to check cache size
func checkSize(t *testing.T, cache *Cache, expectedSize int) {
	if cache.size != expectedSize {
		t.Errorf("Cache size mismatch: got %d, want %d", cache.size, expectedSize)
	}
}

// Helper to validate cache content
func checkCacheContent(t *testing.T, cache *Cache, expected []int) {
	node := cache.head
	i := 0
	for node != nil {
		if i >= len(expected) {
			t.Errorf("Cache has more elements than expected")
			return
		}
		if node.value != expected[i] {
			t.Errorf("Unexpected value at position %d: got %d, want %d", i, node.value, expected[i])
		}
		node = node.next
		i++
	}
	if i != len(expected) {
		t.Errorf("Cache has fewer elements than expected: got %d, want %d", i, len(expected))
	}
}

func TestCacheInitialization(t *testing.T) {
	cache := NewCache(3)
	if cache.size != 0 || cache.head != nil || cache.tail != nil {
		t.Error("Cache not initialized properly")
	}
}

func TestCacheAdd(t *testing.T) {
	cache := NewCache(3)

	cache.Add(1)
	cache.Add(2)
	cache.Add(3)

	checkSize(t, cache, 3)
	checkCacheContent(t, cache, []int{3, 2, 1})
}

func TestCacheEvict(t *testing.T) {
	cache := NewCache(3)

	// Fill cache to capacity
	cache.Add(1)
	cache.Add(2)
	cache.Add(3)

	// Add one more to trigger eviction
	cache.Add(4)

	checkSize(t, cache, 3)
	checkCacheContent(t, cache, []int{4, 3, 2}) // 1 should be evicted
}

func TestCacheHit(t *testing.T) {
	cache := NewCache(3)

	cache.Hit(1)
	cache.Hit(2)
	cache.Hit(3)
	cache.Hit(1) // 1 is now marked as visited

	checkSize(t, cache, 3)
	checkCacheContent(t, cache, []int{3, 2, 1})
}

func TestCacheEvictionWithVisited(t *testing.T) {
	cache := NewCache(3)

	// Fill cache
	cache.Hit(1)
	cache.Hit(2)
	cache.Hit(3)

	// Mark 1 as visited
	cache.Hit(1)

	// Add 4, should evict an unvisited node
	cache.Hit(4)

	checkSize(t, cache, 3)
	checkCacheContent(t, cache, []int{4, 3, 1}) // 2 should be evicted
}

func TestCacheEmptyEvict(t *testing.T) {
	cache := NewCache(3)

	err := cache.Evict()
	if err == nil {
		t.Error("Expected error when evicting from an empty cache, but got nil")
	}
}
