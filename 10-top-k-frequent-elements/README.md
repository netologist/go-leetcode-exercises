# Top K Frequent Elements

## LeetCode Information
- **Number:** 347
- **Difficulty:** Medium
- **FinTech Companies:** Bloomberg, Citadel, Stripe, Robinhood, Two Sigma, Revolut

---

## FinTech Relevance & Real-World Application
Essential in financial monitoring, liquidity analysis, and fraud prevention pipelines:
1. **Most Active Trading Instruments:** Aggregating ticks to rank the top $K$ traded asset symbols or market pairs by order execution frequency.
2. **Fraud & Chargeback Hotspot Identification:** Identifying the top $K$ merchants or IP addresses generating anomalous dispute rates.
3. **Personal Finance Analytics:** Extracting a customer's top spending merchant categories over a trailing statement cycle.

---

## Algorithmic Approaches
Two primary architectures exist:

1. **Min-Heap Approach ($O(N \log K)$):**
   - Populate frequency map in $O(N)$.
   - Maintain a Min-Heap of size bounded by $K$. When the heap exceeds $K$, evict the lowest frequency element.
   - Time: $O(N \log K)$, Space: $O(N + K)$.

2. **Bucket Sort Approach ($O(N)$ - Optimal):**
   - Build frequency map `map[int]int`.
   - Create a bucket array `buckets := make([][]int, len(nums)+1)` where the index corresponds to frequency.
   - Insert numbers into their corresponding frequency bucket.
   - Traverse backwards from highest frequency bucket down to 0, gathering $K$ elements.
   - Time: $O(N)$ strictly linear, Space: $O(N)$.

---

## Complexity Analysis
- **Time Complexity:** $O(N)$ - Achieved via linear Bucket Sort distribution and reverse scan.
- **Space Complexity:** $O(N)$ - Frequency map and bucket slice allocations.

---

## Critical Edge Cases
- $K = 1$ (single mode) or $K = \text{len}(nums)$ (all elements unique).
- Negative numbers and zero values.
- Ties in frequencies (any permutation of tied candidates is valid).

---

## 💻 Production Go Implementation

```go
package topkfrequent

// TopKFrequent returns the k most frequent elements in nums.
// Time Complexity: O(N) using Bucket Sort, Space Complexity: O(N)
func TopKFrequent(nums []int, k int) []int {
    counts := make(map[int]int)
    for _, num := range nums {
        counts[num]++
    }

    // Buckets where index represents frequency (0 to len(nums))
    buckets := make([][]int, len(nums)+1)
    for num, count := range counts {
        buckets[count] = append(buckets[count], num)
    }

    result := make([]int, 0, k)
    // Iterate from highest frequency bucket downwards
    for freq := len(buckets) - 1; freq >= 0 && len(result) < k; freq-- {
        if len(buckets[freq]) > 0 {
            for _, num := range buckets[freq] {
                result = append(result, num)
                if len(result) == k {
                    break
                }
            }
        }
    }

    return result
}
```

---

## 🧪 Running Unit Tests

```bash
go test -v ./10-top-k-frequent-elements/...
go test -race ./10-top-k-frequent-elements/...
```

---

[⬅️ Back to All Exercises](../README.md)
