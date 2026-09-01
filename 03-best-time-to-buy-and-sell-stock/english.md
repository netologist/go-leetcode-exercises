# Best Time to Buy and Sell Stock

## LeetCode Information
- **Number:** 121
- **Difficulty:** Easy
- **FinTech Companies:** Citadel, Bloomberg, Robinhood, Two Sigma, Goldman Sachs, Jane Street

---

## FinTech Relevance & Real-World Application
Foundational to algorithmic execution and risk management:
1. **Optimal Execution Benchmark:** Measuring the upper bound of theoretical profit achievable from a single buy-sell trade cycle over a tick window.
2. **Streaming Market Feed Analytics:** Operates in $O(1)$ auxiliary space, making it directly suitable for processing continuous UDP/WebSocket price streams without buffer allocations.
3. **Relation to Maximum Drawdown (MDD):** The symmetric counterpart of this algorithm is standard in computing peak-to-trough risk metrics for portfolio health.

---

## Algorithmic Approach
Instead of testing every pair ($O(N^2)$), we maintain a single-pass greedy invariant:
- Track `minPrice` observed so far.
- Track `maxProfit` achievable so far.
- At each tick:
  - If `price < minPrice`, reset `minPrice = price`.
  - Else if `price - minPrice > maxProfit`, update `maxProfit`.

---

## Complexity Analysis
- **Time Complexity:** $O(N)$ - Exactly one pass over the prices slice.
- **Space Complexity:** $O(1)$ - Constant memory using two scalar variables.

---

## Critical Edge Cases
- Strictly falling asset prices (returns `0`, as negative profits / losses are avoided by not trading).
- Slice length $< 2$ (cannot complete a trade, returns `0`).
- Constant flat asset price (returns `0`).
