# Merge Intervals

## LeetCode Information
- **Number:** 56
- **Difficulty:** Medium
- **FinTech Companies:** Stripe, Bloomberg, Robinhood, Citadel, Two Sigma, Brex

---

## FinTech Relevance & Real-World Application
Interval manipulation is core to scheduling, billing, and order management:
1. **Order Book Consolidation:** Consolidating multiple market maker liquidity quote ranges into aggregated depth levels.
2. **Subscription & Billing Periods (Stripe):** Merging overlapping prorated subscription coverage windows into unified active billing spans.
3. **Interest Accrual & Loan Windows:** Consolidating active collateralized debt periods to calculate compounding duration without double-counting.
4. **Exchange Market Trading Sessions:** Computing the union of open trading windows across global clearing houses.

---

## Algorithmic Approach
1. **Sort by Interval Start Time ($O(N \log N)$):** Ensures that any potentially overlapping intervals appear contiguously.
2. **Linear Merge Sweep ($O(N)$):**
   - Push the first interval to the `merged` accumulator.
   - For each subsequent interval `current`, compare with `lastMerged`:
     - If `current[0] <= lastMerged[1]`, an overlap exists: update `lastMerged[1] = max(lastMerged[1], current[1])`.
     - Otherwise, the intervals are disjoint: append `current` as a new entry.

---

## Complexity Analysis
- **Time Complexity:** $O(N \log N)$ - Dominated by the initial sort using `sort.Slice`.
- **Space Complexity:** $O(N)$ - Output slice holding non-overlapping merged intervals.

---

## Critical Edge Cases
- Adjacent/Touching boundaries: `[1, 4]` and `[4, 5]` merge into `[1, 5]`.
- Fully enclosed intervals: `[1, 10]` and `[2, 6]` merge into `[1, 10]`.
- Pre-sorted vs unsorted input sequences.

---

## 💻 Production Go Implementation

```go
package mergeintervals

import "sort"

// Merge merges all overlapping intervals and returns an array of the non-overlapping intervals.
// Time Complexity: O(N log N), Space Complexity: O(N)
func Merge(intervals [][]int) [][]int {
    if len(intervals) <= 1 {
        return intervals
    }

    // Sort intervals by starting time
    sort.Slice(intervals, func(i, j int) bool {
        return intervals[i][0] < intervals[j][0]
    })

    merged := make([][]int, 0, len(intervals))
    merged = append(merged, intervals[0])

    for i := 1; i < len(intervals); i++ {
        current := intervals[i]
        lastIdx := len(merged) - 1
        lastMerged := merged[lastIdx]

        if current[0] <= lastMerged[1] {
            // Overlap detected: merge by taking the max of ending bounds
            if current[1] > lastMerged[1] {
                merged[lastIdx][1] = current[1]
            }
        } else {
            // Disjoint interval: append as a new block
            merged = append(merged, current)
        }
    }

    return merged
}
```

---

## 🧪 Running Unit Tests

```bash
go test -v ./04-merge-intervals/...
go test -race ./04-merge-intervals/...
```

---

[⬅️ Back to All Exercises](../README.md)
