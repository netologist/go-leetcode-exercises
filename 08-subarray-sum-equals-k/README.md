# Subarray Sum Equals K

## LeetCode Information
- **Number:** 560
- **Difficulty:** Medium
- **FinTech Companies:** Stripe, Robinhood, Citadel, Bloomberg, Two Sigma

---

## FinTech Relevance & Real-World Application
Foundational in financial ledger analysis and AML (Anti-Money Laundering) detection:
1. **Ledger Sub-segment Reconciliation:** Locating consecutive transaction sequences that net out to a specific balance $K$ (e.g. balancing zero-sum cycles).
2. **Financial Fraud & Structuring Detection:** Detecting smurfing patterns where consecutive micro-transfers aggregate exactly to a regulatory compliance threshold $K$.
3. **P&L Target Windowing:** Identifying contiguous trade intervals that hit a designated profit/loss target.

---

## Algorithmic Approach
Because the input array can contain negative values, a traditional Two-Pointer / Sliding Window will fail. We use a **Prefix Sum with Frequency Hash Map**:
- The sum of subarray `nums[i..j]` is equal to `PrefixSum[j] - PrefixSum[i-1]`.
- For `nums[i..j] == k`, the condition `PrefixSum[j] - k == PrefixSum[i-1]` must hold.
- As we iterate, we maintain `currentSum` (representing `PrefixSum[j]`).
- We check if `currentSum - k` exists in `prefixCount` map and add its frequency to `totalCount`.
- Increment `prefixCount[currentSum]`.
- **Base Case:** Initialize `prefixCount[0] = 1` to account for valid subarrays starting from index 0 where `currentSum == k`.

---

## Complexity Analysis
- **Time Complexity:** $O(N)$ - Single pass over the slice with $O(1)$ average hash map operations.
- **Space Complexity:** $O(N)$ - Hash map storing prefix sum frequencies.

---

## Critical Edge Cases
- Streams with negative debits, positive credits, and zero balances.
- Target `k` is negative or zero.
- Crucial base case initialization `prefixCount[0] = 1`.

---

## 💻 Production Go Implementation

```go
package subarraysum

// SubarraySum finds the total number of continuous subarrays whose sum equals k.
// Time Complexity: O(N), Space Complexity: O(N)
func SubarraySum(nums []int, k int) int {
    prefixCount := make(map[int]int, len(nums)+1)
    // Base case: a prefix sum of 0 has occurred once before processing elements
    prefixCount[0] = 1

    currentSum := 0
    totalCount := 0

    for _, num := range nums {
        currentSum += num
        // If (currentSum - k) exists in map, it means a sub-segment summing to k was found
        if count, exists := prefixCount[currentSum-k]; exists {
            totalCount += count
        }
        prefixCount[currentSum]++
    }

    return totalCount
}
```

---

## 🧪 Running Unit Tests

```bash
go test -v ./08-subarray-sum-equals-k/...
go test -race ./08-subarray-sum-equals-k/...
```

---

[⬅️ Back to All Exercises](../README.md)
