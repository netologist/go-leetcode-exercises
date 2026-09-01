package ratelimiter

import (
	"sync"
	"testing"
	"time"
)

func RunRateLimiterTests(t *testing.T) {
	t.Run("Basic Capacity and Burst Exhaustion", func(t *testing.T) {
		tb := NewTokenBucket(3, 1) // Capacity 3, 1 token/sec refill

		if !tb.Allow() {
			t.Errorf("1st token should be allowed")
		}
		if !tb.Allow() {
			t.Errorf("2nd token should be allowed")
		}
		if !tb.Allow() {
			t.Errorf("3rd token should be allowed")
		}
		if tb.Allow() {
			t.Errorf("4th token should be rejected (bucket exhausted)")
		}
	})

	t.Run("Lazy Refill After Elapsed Time", func(t *testing.T) {
		tb := NewTokenBucket(2, 10) // 10 tokens per second (1 token every 100ms)

		// Consume all tokens
		tb.Allow()
		tb.Allow()
		if tb.Allow() {
			t.Errorf("should be exhausted")
		}

		// Wait 250ms -> should refill ~2.5 tokens capped at 2
		time.Sleep(250 * time.Millisecond)

		if !tb.Allow() {
			t.Errorf("token should be allowed after refill sleep")
		}
		if !tb.Allow() {
			t.Errorf("second refilled token should be allowed")
		}
		if tb.Allow() {
			t.Errorf("third token should be rejected (capped at capacity 2)")
		}
	})

	t.Run("Concurrent Goroutine Access Safety", func(t *testing.T) {
		tb := NewTokenBucket(50, 100)
		var wg sync.WaitGroup
		allowedCount := 0
		var mu sync.Mutex

		for range 100 {
			wg.Add(1)
			go func() {
				defer wg.Done()
				if tb.Allow() {
					mu.Lock()
					allowedCount++
					mu.Unlock()
				}
			}()
		}

		wg.Wait()

		if allowedCount > 100 || allowedCount < 50 {
			t.Errorf("expected allowed tokens between 50 and 100, got %d", allowedCount)
		}
	})
}
