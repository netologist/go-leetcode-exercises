# Design Hit Counter

## LeetCode Information
- **Number:** 362
- **Difficulty:** Medium
- **FinTech Companies:** Stripe, Robinhood, Bloomberg, Citadel, Datadog

---

## FinTech Relevance & Real-World Application
Foundational for telemetry, latency monitoring, and real-time TPS counters:
1. **Real-time TPS (Transactions Per Second) Telemetry:** Measuring throughput volume across payment gateways and exchange feeds over trailing 5-minute rolling windows.
2. **Circuit Breakers & Error Spike Triggers:** Counting consecutive failed authorization calls to trip automated protective failovers.
3. **API Usage Metering:** Enforcing rolling time-window quotas for high-frequency trading clients.

---

## Algorithmic Approach
A queue-based implementation appends every single hit. Under high concurrency (e.g. 100,000 hits/sec), a queue explodes in memory ($O(N)$) and requires frequent garbage collection.

The optimal production pattern is a **300-Bucket Circular Buffer**:
- `times [300]int`: Stores the timestamp associated with bucket index `timestamp % 300`.
- `hits [300]int`: Stores the aggregated hit count for that specific second.
- **`Hit(timestamp)`:**
  - `idx = timestamp % 300`.
  - If `times[idx] != timestamp`, the bucket contains data from a past 5-minute cycle: reset `times[idx] = timestamp` and `hits[idx] = 1`.
  - If `times[idx] == timestamp`, increment `hits[idx]++`.
- **`GetHits(timestamp)`:**
  - Loop through all 300 buckets and accumulate `hits[i]` where `timestamp - times[i] < 300`.

---

## Complexity Analysis
- **Time Complexity:**
  - `Hit`: $O(1)$ constant time.
  - `GetHits`: $O(1)$ constant time (iterates exactly 300 array slots).
- **Space Complexity:** $O(1)$ - Strictly bounded to two 300-element integer arrays regardless of volume.

---

## Concurrency Note
Protected by `sync.Mutex` to ensure data race freedom across concurrent goroutine callers.

---

## 💻 Production Go Implementation

```go
package hitcounter

import "sync"

// HitCounter counts the number of hits received in the past 5 minutes (300 seconds).
// Uses a circular bucket buffer for strictly O(1) space and O(1) time operations.
type HitCounter struct {
    mu    sync.Mutex
    times [300]int // Timestamps associated with bucket index (timestamp % 300)
    hits  [300]int // Hit count accumulated in bucket index
}

// Constructor initializes the HitCounter.
func Constructor() HitCounter {
    return HitCounter{}
}

// Hit records a hit at given timestamp (in seconds granularity).
// Time Complexity: O(1), Space Complexity: O(1)
func (c *HitCounter) Hit(timestamp int) {
    c.mu.Lock()
    defer c.mu.Unlock()

    idx := timestamp % 300
    if c.times[idx] != timestamp {
        // New second window: overwrite old stale bucket
        c.times[idx] = timestamp
        c.hits[idx] = 1
    } else {
        // Same second window: increment hits
        c.hits[idx]++
    }
}

// GetHits returns the number of hits in the past 300 seconds from given timestamp.
// Time Complexity: O(1) (iterates exactly 300 buckets), Space Complexity: O(1)
func (c *HitCounter) GetHits(timestamp int) int {
    c.mu.Lock()
    defer c.mu.Unlock()

    total := 0
    for i := range 300 {
        if timestamp-c.times[i] < 300 && c.times[i] > 0 {
            total += c.hits[i]
        }
    }

    return total
}
```

---

## 🧪 Running Unit Tests

```bash
go test -v ./19-design-hit-counter/...
go test -race ./19-design-hit-counter/...
```

---

[⬅️ Back to All Exercises](../README.md)
