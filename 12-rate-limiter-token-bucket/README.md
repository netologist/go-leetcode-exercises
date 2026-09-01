# Token Bucket Rate Limiter

## Problem Information
- **Category:** FinTech Systems & Algorithmic Design
- **Difficulty:** Medium
- **FinTech Companies:** Stripe, Revolut, Robinhood, Brex, Adyen, Plaid, Coinbase

---

## FinTech Relevance & Real-World Application
The hallmark live-coding interview question across top tier payment and API platforms:
1. **API Tier Enforcement:** Enforcing contractual requests-per-second (RPS) quotas per API key or merchant ID.
2. **Burst Tolerant Throughput:** Allowing high-throughput spikes (burst capacity $C$) while smoothing continuous consumption to a steady rate $R$.
3. **Card Testing & Bot Attack Defense:** Preventing automated brute-force card testing attacks by gating charge validation endpoints.

---

## Algorithmic Design & Lazy Refill
Rather than spawning background tickers or goroutines (which risk memory and CPU leaks), the optimal production implementation uses **Lazy Time-Delta Refill**:

1. **State Fields:**
   - `capacity`: Maximum token ceiling ($C$).
   - `tokens`: Current available token quantity ($float64$).
   - `refillRate`: Replenishment velocity in tokens per second ($R$).
   - `lastRefillAt`: Timestamp of previous token evaluation (`time.Time`).
   - `mu sync.Mutex`: Protects state against concurrent multi-goroutine mutations.

2. **On Request Evaluation (`AllowN`):**
   - Acquire `tb.mu.Lock()`.
   - Calculate elapsed duration: `elapsed = now.Sub(lastRefillAt).Seconds()`.
   - Replenish tokens: `tokens = min(capacity, tokens + elapsed * refillRate)`.
   - Update `lastRefillAt = now`.
   - If `tokens >= n`, deduct `tokens -= n` and return `true`; else return `false`.

---

## Complexity Analysis
- **Time Complexity:** $O(1)$ - Constant time arithmetic and lock acquisition per request.
- **Space Complexity:** $O(1)$ - Fixed struct footprint with zero dynamic heap allocations.

---

## Concurrency & Go Idioms
- Thread safety is guaranteed via `sync.Mutex`.
- Lazy recalculation eliminates goroutine scheduling overhead and avoids ticker leaks entirely.

---

## 💻 Production Go Implementation

```go
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
```

---

## 🧪 Running Unit Tests

```bash
go test -v ./12-rate-limiter-token-bucket/...
go test -race ./12-rate-limiter-token-bucket/...
```

---

[⬅️ Back to All Exercises](../README.md)
