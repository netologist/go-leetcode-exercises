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
