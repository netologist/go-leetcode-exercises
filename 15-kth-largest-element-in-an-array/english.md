# Kth Largest Element in an Array

## LeetCode Information
- **Number:** 215
- **Difficulty:** Medium
- **FinTech Companies:** Citadel, Two Sigma, Bloomberg, Robinhood, Goldman Sachs

---

## FinTech Relevance & Real-World Application
Crucial in quantitative trading, order routing, and real-time execution analytics:
1. **Price Percentile & Quantile Thresholds:** Computing the 95th or 99th percentile (p99) execution price boundary across millions of live order ticks in linear $O(N)$ time without full $O(N \log N)$ sorting.
2. **Dark Pool & Block Trade Execution:** Identifying the clearing cutoff price level that fulfills a requested cumulative volume quota.
3. **High-Water Mark & Volatility Thresholds:** Selecting threshold benchmark values dynamically from streaming trade books.

---

## Algorithmic Approach
While a Min-Heap of size $K$ operates in $O(N \log K)$, the optimal algorithm is **Quickselect (Hoare's Selection Algorithm)** running in average **$O(N)$**:
- In zero-indexed sorted order, the $K$-th largest element sits at `targetIdx = len(nums) - k`.
- Select a random pivot to avoid worst-case quadratic behavior.
- Partition the sub-array in-place around the pivot.
- **Branch Elimination:**
  - If `pivotIdx == targetIdx`: Element located, return `nums[pivotIdx]`.
  - If `pivotIdx < targetIdx`: Search right partition (`left = pivotIdx + 1`).
  - If `pivotIdx > targetIdx`: Search left partition (`right = pivotIdx - 1`).
- Unlike QuickSort, Quickselect recurses into only one partition half, yielding the geometric series: $N + N/2 + N/4 + \dots = 2N = O(N)$.

---

## Complexity Analysis
- **Time Complexity:** $O(N)$ average. In-place randomized partitioning shields against pathological $O(N^2)$ inputs.
- **Space Complexity:** $O(1)$ - Iterative implementation runs with zero additional heap allocations.

---

## Critical Edge Cases
- Minimal array size ($N = 1, K = 1$).
- Arrays containing all identical values.
- Arrays with negative and zero entries.
