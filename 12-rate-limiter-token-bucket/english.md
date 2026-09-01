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
