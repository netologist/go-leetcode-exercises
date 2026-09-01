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
