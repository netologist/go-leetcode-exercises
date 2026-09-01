package lrucache

import "testing"

func RunLRUCacheTests(t *testing.T) {
	t.Run("Standard LRU Operations", func(t *testing.T) {
		cache := Constructor(2)

		cache.Put(1, 1)
		cache.Put(2, 2)
		if val := cache.Get(1); val != 1 {
			t.Errorf("expected Get(1) = 1, got %d", val)
		}

		// Evicts key 2
		cache.Put(3, 3)
		if val := cache.Get(2); val != -1 {
			t.Errorf("expected Get(2) = -1 (evicted), got %d", val)
		}

		// Evicts key 1 because key 3 was just added and key 1 was accessed before key 3
		cache.Put(4, 4)
		if val := cache.Get(1); val != -1 {
			t.Errorf("expected Get(1) = -1 (evicted), got %d", val)
		}
		if val := cache.Get(3); val != 3 {
			t.Errorf("expected Get(3) = 3, got %d", val)
		}
		if val := cache.Get(4); val != 4 {
			t.Errorf("expected Get(4) = 4, got %d", val)
		}
	})

	t.Run("Update Existing Key Value and Move to MRU", func(t *testing.T) {
		cache := Constructor(2)
		cache.Put(1, 10)
		cache.Put(2, 20)
		cache.Put(1, 100) // Update key 1, becomes MRU

		cache.Put(3, 30) // Should evict key 2, not 1
		if val := cache.Get(2); val != -1 {
			t.Errorf("expected key 2 to be evicted (-1), got %d", val)
		}
		if val := cache.Get(1); val != 100 {
			t.Errorf("expected key 1 to be 100, got %d", val)
		}
		if val := cache.Get(3); val != 30 {
			t.Errorf("expected key 3 to be 30, got %d", val)
		}
	})

	t.Run("Capacity 1 Edge Case", func(t *testing.T) {
		cache := Constructor(1)
		cache.Put(2, 1)
		if val := cache.Get(2); val != 1 {
			t.Errorf("expected Get(2) = 1, got %d", val)
		}
		cache.Put(3, 2) // Evicts key 2
		if val := cache.Get(2); val != -1 {
			t.Errorf("expected Get(2) = -1, got %d", val)
		}
		if val := cache.Get(3); val != 2 {
			t.Errorf("expected Get(3) = 2, got %d", val)
		}
	})

	t.Run("Non-existent Key Lookup", func(t *testing.T) {
		cache := Constructor(2)
		if val := cache.Get(42); val != -1 {
			t.Errorf("expected Get(42) = -1, got %d", val)
		}
	})
}
