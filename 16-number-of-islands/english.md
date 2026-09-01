# Number of Islands

## LeetCode Information
- **Number:** 200
- **Difficulty:** Medium
- **FinTech Companies:** Bloomberg, Stripe, Citadel, Amazon, Robinhood, Revolut

---

## FinTech Relevance & Real-World Application
The standard paradigm for connected-components discovery across financial graphs:
1. **Syndicate & Fraud Ring Discovery:** Identifying isolated clusters of colluding accounts, compromised device IDs, and mule banking networks.
2. **Financial Network Contagion Modeling:** Isolating sub-graphs of interconnected counterparty exposures to simulate systemic shock propagation.
3. **Liquidity Partitioning:** Segmenting off-chain vs on-chain asset liquidity subnetworks.

---

## Algorithmic Approach
Graph traversal using **Depth-First Search (DFS)** or **Breadth-First Search (BFS)**:
- Iterate through each cell `(r, c)` of the grid.
- When an unvisited land cell `grid[r][c] == '1'` is discovered:
  - Increment the component counter `count++`.
  - Trigger DFS to flood-fill and visit all 4-directionally connected land cells.
  - Mutate visited cells to `'0'` in-place to avoid duplicate visits without auxiliary memory.

---

## Complexity Analysis
- **Time Complexity:** $O(M \times N)$ - Every cell is inspected a constant number of times ($M$ rows, $N$ columns).
- **Space Complexity:** $O(M \times N)$ - Worst-case DFS call stack depth when the entire matrix is filled with land.

---

## Critical Edge Cases
- Empty grid dimensions ($0 \times 0$).
- Grids containing all water `'0'` or all land `'1'`.
- Diagonal isolation (diagonals do not constitute connectivity).
