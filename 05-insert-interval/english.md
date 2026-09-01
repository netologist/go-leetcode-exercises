# Insert Interval

## LeetCode Information
- **Number:** 57
- **Difficulty:** Medium
- **FinTech Companies:** Stripe, Bloomberg, Citadel, Robinhood, Two Sigma

---

## FinTech Relevance & Real-World Application
In real-time trading engines and ledger tier management:
1. **Dynamic Fee Bracket Injection:** Inserting a promotional zero-fee volume threshold into an already sorted list of commission tiers.
2. **Settlement / Clearing Window Scheduling:** Injecting an ad-hoc clearing auction window into existing market trading schedules.
3. **Credit Line / Risk Tier Adjustments:** Merging temporary leverage exposure limits into pre-existing risk exposure bounds.

---

## Algorithmic Approach
Because the input `intervals` list is already sorted by start time and non-overlapping, sorting is unnecessary. We execute a **3-Phase Linear Scan ($O(N)$)**:

1. **Phase 1 (Preceding Disjoint Intervals):** Append all intervals that finish strictly before `newInterval` starts (`interval[1] < newInterval[0]`).
2. **Phase 2 (Overlapping Merge):** For all intervals starting before `newInterval` ends (`interval[0] <= newInterval[1]`), merge boundaries:
   - `newInterval[0] = min(newInterval[0], interval[0])`
   - `newInterval[1] = max(newInterval[1], interval[1])`
   Append the merged `newInterval`.
3. **Phase 3 (Succeeding Disjoint Intervals):** Append all remaining intervals.

---

## Complexity Analysis
- **Time Complexity:** $O(N)$ - Single pass over the slice with no sorting overhead.
- **Space Complexity:** $O(N)$ - Allocates the resulting non-overlapping intervals slice.

---

## Critical Edge Cases
- Empty input list `[]` (simply returns `[newInterval]`).
- `newInterval` placed strictly before the head or after the tail without overlaps.
- `newInterval` completely spans and swallows all existing intervals.
