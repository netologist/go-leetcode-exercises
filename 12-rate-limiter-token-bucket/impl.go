package ratelimiter

import (
	"sync"
	"time"
)

// TokenBucket implements a thread-safe token bucket rate limiter.
type TokenBucket struct {
	mu           sync.Mutex
	capacity     float64   // Maximum tokens the bucket can hold (burst limit)
	tokens       float64   // Current number of available tokens
	refillRate   float64   // Tokens added per second
	lastRefillAt time.Time // Timestamp of last refill calculation
}

// NewTokenBucket creates a new TokenBucket with capacity and refillRate (tokens/sec).
func NewTokenBucket(capacity float64, refillRate float64) *TokenBucket {
	return &TokenBucket{
		capacity:     capacity,
		tokens:       capacity, // Bucket starts full
		refillRate:   refillRate,
		lastRefillAt: time.Now(),
	}
}

// Allow is shorthand for AllowN(1).
func (tb *TokenBucket) Allow() bool {
	return tb.AllowN(1)
}

// AllowN checks if n tokens can be consumed right now.
// Returns true and deducts tokens if available; returns false otherwise.
// Refills lazily based on elapsed time since last request.
func (tb *TokenBucket) AllowN(n float64) bool {
	tb.mu.Lock()
	defer tb.mu.Unlock()

	tb.refill()

	if tb.tokens >= n {
		tb.tokens -= n
		return true
	}

	return false
}

// refill calculates new tokens generated since lastRefillAt and caps at capacity.
func (tb *TokenBucket) refill() {
	now := time.Now()
	elapsed := now.Sub(tb.lastRefillAt).Seconds()
	tb.lastRefillAt = now

	tb.tokens += elapsed * tb.refillRate
	if tb.tokens > tb.capacity {
		tb.tokens = tb.capacity
	}
}

// AvailableTokens returns the current number of available tokens (for inspection/testing).
func (tb *TokenBucket) AvailableTokens() float64 {
	tb.mu.Lock()
	defer tb.mu.Unlock()
	tb.refill()
	return tb.tokens
}
