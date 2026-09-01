# Coin Change

## LeetCode Information
- **Number:** 322
- **Difficulty:** Medium
- **FinTech Companies:** Citadel, Bloomberg, Robinhood, Goldman Sachs, Two Sigma

---

## FinTech Relevance & Real-World Application
The quintessential dynamic programming benchmark in financial cash management:
1. **ATM & POS Cash Dispensation:** Calculating the minimal number of physical currency bills/notes needed to fulfill an arbitrary withdrawal amount.
2. **Crypto Wallet UTXO Selection:** Selecting the minimal set of Unspent Transaction Outputs (UTXOs) to satisfy a target transfer while minimizing blockchain gas fees.
3. **Foreign Exchange Breakdown:** Assembling target foreign currency denominations with minimum banknote count.

---

## Algorithmic Approach
A greedy heuristic fails on non-canonical coin systems (e.g. `coins = [1, 3, 4, 5]` and `amount = 7`; greedy chooses `5+1+1` [3 coins], whereas the true optimum is `4+3` [2 coins]).

We apply **Bottom-Up Dynamic Programming (Unbounded Knapsack)**:
- Define `dp[i]` as the minimum coin count required to assemble amount `i`.
- Base state: `dp[0] = 0`; initialize `dp[1...amount]` to `math.MaxInt32`.
- Transition for each $i \in [1..\text{amount}]$ across each coin $c$:
  - If $i - c \ge 0$, then `dp[i] = min(dp[i], dp[i-c] + 1)`.
- If `dp[amount] == math.MaxInt32`, the target amount cannot be composed; return `-1`. Otherwise return `dp[amount]`.

---

## Complexity Analysis
- **Time Complexity:** $O(\text{Amount} \times C)$ where $C$ is the number of coin denominations.
- **Space Complexity:** $O(\text{Amount})$ - 1D array of size $\text{amount} + 1$.

---

## Critical Edge Cases
- Zero target amount (`amount = 0` yields `0`).
- Unattainable targets (returns `-1`).
- Non-canonical coin systems where greedy selection produces suboptimal solutions.
