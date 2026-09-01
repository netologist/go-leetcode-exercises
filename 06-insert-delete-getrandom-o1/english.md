# Insert Delete GetRandom O(1)

## LeetCode Information
- **Number:** 380
- **Difficulty:** Medium
- **FinTech Companies:** Robinhood, Citadel, Two Sigma, Bloomberg, Stripe

---

## FinTech Relevance & Real-World Application
Critical for market simulation, token pools, and stochastic audit systems:
1. **Dynamic Liquidity / Order Sampling:** Selecting random counterparty bids/asks from active matching pools with strictly uniform probability.
2. **Real-time Fraud Inspection Sampling:** Extracting random transactions from live payment processing streams for out-of-band deep verification without performance overhead.
3. **Lottery / Rewards Allocation:** Managing active eligible participant pools with dynamic additions/cancellations while maintaining zero-bias random selection.

---

## Algorithmic Architecture
- A Hash Map alone allows $O(1)$ insertions and deletions, but lacks continuous indexable memory for uniform random draws.
- A Slice alone allows $O(1)$ random indexing, but arbitrary element deletion requires $O(N)$ memory shifting.

**Hybrid Strategy:**
- `nums []int`: Stores elements contiguously for instant random access via index.
- `pos map[int]int`: Maps each value to its current index in `nums`.

### The Swap-and-Pop Technique
To delete `val` in $O(1)$:
1. Find `idx = pos[val]`.
2. Overwrite `nums[idx]` with `lastVal = nums[len(nums)-1]`.
3. Update `pos[lastVal] = idx`.
4. Truncate `nums = nums[:len(nums)-1]` and `delete(pos, val)`.

---

## Complexity Analysis
- **Time Complexity:**
  - `Insert(val)`: $O(1)$ average.
  - `Remove(val)`: $O(1)$ average.
  - `GetRandom()`: $O(1)$ deterministic.
- **Space Complexity:** $O(N)$ - Auxiliary storage for slice and hash map.
