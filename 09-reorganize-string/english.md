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
