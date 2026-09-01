# Reorganize String (Task & Throttle Dispatcher)

## LeetCode Information
- **Number:** 767
- **Difficulty:** Medium
- **FinTech Companies:** Stripe, Citadel, Robinhood, Bloomberg

---

## FinTech Relevance & Real-World Application
Asked frequently as a **"Fair Payment / Task Throttle Dispatcher"**:
1. **Merchant / Card Velocity Throttling:** Spreading out consecutive charges from the same card or merchant across the processing pipeline to prevent downstream rate-limiting or false fraud trips.
2. **Order Matching Engine Fair Dispatch:** Preventing high-frequency quotes for a single asset symbol from monopolizing execution pipelines over other liquid instruments.
3. **Outbound Banking API Schedule Smoothing:** Interleaving HTTP requests destined for multiple external partner banks to ensure smooth traffic distribution.

---

## Algorithmic Approach
1. **Frequency Mapping:** Calculate counts for every unique symbol or character.
2. **Feasibility Validation (Pigeonhole Principle):** If the highest frequency exceeds `(len(s) + 1) / 2`, it is mathematically impossible to arrange without adjacent collisions; immediately return `""`.
3. **Max-Heap Priority Queue:**
   - Extract the two most frequent remaining elements (`first`, `second`).
   - Append both to the result string builder and decrement their counts.
   - Re-insert them into the heap if remaining counts $> 0$.
   - This greedy approach guarantees the most critical bottlenecks are depleted first.

---

## Complexity Analysis
- **Time Complexity:** $O(N \log A)$ where $A$ is the alphabet size ($A \le 26$, effectively $O(N)$).
- **Space Complexity:** $O(A)$ - Priority queue and frequency map storage bounded by unique symbols.

---

## Idiomatic Go: `container/heap`
Demonstrating clean implementation of Go's `heap.Interface` (`Len`, `Less`, `Swap`, `Push`, `Pop`) with pointer receivers is a prized evaluation point in FinTech senior live coding interviews.

---

## 💻 Production Go Implementation

```go
package reorganizestring

import (
    "container/heap"
    "strings"
)

// CharFreq represents a character and its count.
type CharFreq struct {
    char  byte
    count int
}

// MaxHeap implements heap.Interface for CharFreq based on count descending.
type MaxHeap []CharFreq

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i].count > h[j].count }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MaxHeap) Push(x any) {
    *h = append(*h, x.(CharFreq))
}

func (h *MaxHeap) Pop() any {
    old := *h
    n := len(old)
    item := old[n-1]
    *h = old[:n-1]
    return item
}

// ReorganizeString rearranges characters so that no two adjacent characters are identical.
// Returns an empty string if impossible.
// Time Complexity: O(N log A) where A is alphabet size (A <= 26), Space Complexity: O(A)
func ReorganizeString(s string) string {
    counts := make(map[byte]int)
    for i := range len(s) {
        counts[s[i]]++
    }

    h := &MaxHeap{}
    heap.Init(h)

    maxAllowed := (len(s) + 1) / 2
    for char, count := range counts {
        if count > maxAllowed {
            return ""
        }
        heap.Push(h, CharFreq{char: char, count: count})
    }

    var result strings.Builder
    result.Grow(len(s))

    for h.Len() >= 2 {
        first := heap.Pop(h).(CharFreq)
        second := heap.Pop(h).(CharFreq)

        result.WriteByte(first.char)
        result.WriteByte(second.char)

        first.count--
        second.count--

        if first.count > 0 {
            heap.Push(h, first)
        }
        if second.count > 0 {
            heap.Push(h, second)
        }
    }

    if h.Len() > 0 {
        last := heap.Pop(h).(CharFreq)
        result.WriteByte(last.char)
    }

    return result.String()
}
```

---

## 🧪 Running Unit Tests

```bash
go test -v ./09-reorganize-string/...
go test -race ./09-reorganize-string/...
```

---

[⬅️ Back to All Exercises](../README.md)
