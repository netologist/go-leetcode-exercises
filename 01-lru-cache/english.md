# LRU Cache (Least Recently Used Cache)

## LeetCode Information
- **Number:** 146
- **Difficulty:** Medium
- **FinTech Companies:** Stripe, Robinhood, Citadel, Bloomberg, Two Sigma, Brex, Revolut

---

## FinTech Relevance & Real-World Application
In financial distributed systems and ultra-low-latency architectures:
1. **Market Tick Data & Price Caching:** Caching the latest level-1 and level-2 quotes in-memory with strictly deterministic $O(1)$ lookup and eviction.
2. **Payment Idempotency Keys:** Payment processors (e.g., Stripe) cache idempotency keys for fast deduplication before committing to relational ledgers.
3. **Session & Balance Snapshots:** Retaining frequently active accounts and high-volume trading counterparties in low-overhead RAM.

---

## Algorithmic Architecture & Intuition
To achieve true $O(1)$ time complexity for both `Get` and `Put`:
1. **Hash Map (`map[int]*Node`):** Direct key-to-node pointer lookup in $O(1)$ average time.
2. **Doubly Linked List (`Node{prev, next}`):** Maintains chronological recency. Most Recently Used (MRU) nodes reside at the head; Least Recently Used (LRU) nodes reside at the tail.

### The Power of Sentinel (Dummy) Nodes
By maintaining dummy `head` and `tail` nodes, edge cases such as inserting into an empty list or deleting the only element are handled cleanly without repetitive conditional checks or `nil` pointer panics.

---

## Complexity Analysis
- **Time Complexity:**
  - `Get(key)`: $O(1)$ - Hash map lookup followed by unlinking and relinking at head.
  - `Put(key, value)`: $O(1)$ - Insertion or update plus eviction of tail predecessor when capacity overflows.
- **Space Complexity:**
  - $O(\text{Capacity})$ - Dedicated nodes and hash map buckets bounded by capacity $C$.

---

## Concurrency Note in Go
In production FinTech services, callers access the cache concurrently across goroutines. Note that `Get` is **not** a read-only operation because it mutates the linked list order (`moveToHead`). Hence, a full `sync.Mutex` or write lock with `sync.RWMutex` is required across both `Get` and `Put`.
