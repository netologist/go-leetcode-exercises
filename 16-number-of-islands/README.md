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

---

## 💻 Production Go Implementation

```go
package islands

// NumIslands counts the number of connected components of '1's (land) surrounded by '0's (water).
// In FinTech: Used for identifying isolated fraud clusters or transaction networks.
// Time Complexity: O(M * N), Space Complexity: O(M * N) worst-case call stack or queue
func NumIslands(grid [][]byte) int {
    if len(grid) == 0 || len(grid[0]) == 0 {
        return 0
    }

    rows := len(grid)
    cols := len(grid[0])
    count := 0

    // Directions: up, right, down, left
    dirs := [][2]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

    var dfs func(r, c int)
    dfs = func(r, c int) {
        // Out of bounds or water
        if r < 0 || r >= rows || c < 0 || c >= cols || grid[r][c] != '1' {
            return
        }

        // Mark visited in-place
        grid[r][c] = '0'

        for _, d := range dirs {
            dfs(r+d[0], c+d[1])
        }
    }

    for r := range rows {
        for c := range cols {
            if grid[r][c] == '1' {
                count++
                dfs(r, c)
            }
        }
    }

    return count
}
```

---

## 🧪 Running Unit Tests

```bash
go test -v ./16-number-of-islands/...
go test -race ./16-number-of-islands/...
```

---

[⬅️ Back to All Exercises](../README.md)
