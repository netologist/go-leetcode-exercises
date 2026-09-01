# Two Sum

## LeetCode Information
- **Number:** 1
- **Difficulty:** Easy
- **FinTech Companies:** Stripe, Bloomberg, Plaid, Robinhood, Adyen

---

## FinTech Relevance & Real-World Application
While fundamental, Two Sum maps directly to real-world financial operations:
1. **Ledger Reconciliation & Netting:** Matching debit/credit transactions to settle a target balance or net to zero.
2. **Payment Split Matching:** Identifying two sub-transactions that aggregate to a single customer charge.
3. **Forex/Arbitrage Pair Matching:** Locating currency conversions that offset slippage against a fixed target spread.

---

## Algorithmic Approach
Instead of a naive $O(N^2)$ brute-force double loop, we use a **One-Pass Hash Map**:
- For each item `num` at index `i`, compute the required `complement = target - num`.
- If `complement` already exists in the map, return indices `[seen[complement], i]`.
- Otherwise, record `seen[num] = i` and proceed.

---

## Complexity Analysis
- **Time Complexity:** $O(N)$ - Single pass through the array with $O(1)$ average hash map lookups.
- **Space Complexity:** $O(N)$ - Stores at most $N$ elements in the hash map.

---

## Critical Edge Cases
- Elements cannot be used twice (enforced by checking `seen` before inserting current index).
- Negative numbers and zero targets.
- Duplicate values forming the target pair (e.g., `[3, 3]` with `target = 6`).
