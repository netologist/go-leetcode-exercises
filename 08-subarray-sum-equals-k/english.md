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
