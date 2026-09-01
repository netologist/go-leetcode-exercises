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

---

## 💻 Production Go Implementation

```go
package lrucache

// Node represents a doubly linked list node holding key-value pairs.
type Node struct {
    key   int
    value int
    prev  *Node
    next  *Node
}

// LRUCache implements a Least Recently Used (LRU) cache with O(1) Get and Put operations.
type LRUCache struct {
    capacity int
    cache    map[int]*Node
    head     *Node // Dummy head (most recently used side)
    tail     *Node // Dummy tail (least recently used side)
}

// Constructor initializes the LRUCache with positive capacity.
func Constructor(capacity int) LRUCache {
    head := &Node{}
    tail := &Node{}
    head.next = tail
    tail.prev = head

    return LRUCache{
        capacity: capacity,
        cache:    make(map[int]*Node, capacity),
        head:     head,
        tail:     tail,
    }
}

// Get retrieves the value of the key if present, moving it to the front (MRU).
// Returns -1 if the key does not exist.
func (c *LRUCache) Get(key int) int {
    node, exists := c.cache[key]
    if !exists {
        return -1
    }
    c.moveToHead(node)
    return node.value
}

// Put inserts or updates the key-value pair.
// If capacity is exceeded, it evicts the least recently used item (before dummy tail).
func (c *LRUCache) Put(key int, value int) {
    if node, exists := c.cache[key]; exists {
        node.value = value
        c.moveToHead(node)
        return
    }

    newNode := &Node{
        key:   key,
        value: value,
    }
    c.cache[key] = newNode
    c.addToHead(newNode)

    if len(c.cache) > c.capacity {
        lru := c.removeTail()
        delete(c.cache, lru.key)
    }
}

// addToHead inserts a node right after the dummy head.
func (c *LRUCache) addToHead(node *Node) {
    node.prev = c.head
    node.next = c.head.next
    c.head.next.prev = node
    c.head.next = node
}

// removeNode unlinks an existing node from the list.
func (c *LRUCache) removeNode(node *Node) {
    node.prev.next = node.next
    node.next.prev = node.prev
}

// moveToHead moves an existing node to the MRU position (after dummy head).
func (c *LRUCache) moveToHead(node *Node) {
    c.removeNode(node)
    c.addToHead(node)
}

// removeTail removes and returns the LRU node right before the dummy tail.
func (c *LRUCache) removeTail() *Node {
    node := c.tail.prev
    c.removeNode(node)
    return node
}
```

---

## 🧪 Running Unit Tests

```bash
go test -v ./01-lru-cache/...
go test -race ./01-lru-cache/...
```

---

[⬅️ Back to All Exercises](../README.md)
