# Container With Most Water

## LeetCode Information
- **Number:** 11
- **Difficulty:** Medium
- **FinTech Companies:** Robinhood, Bloomberg, Citadel, Goldman Sachs, Two Sigma

---

## FinTech Relevance & Real-World Application
Serves as an intuitive proxy for dual-parameter economic optimization:
1. **Liquidity Depth & Spread Windowing:** Finding the optimal boundary between bid-ask spread duration (width) and guaranteed matched depth (minimum bounding height).
2. **Margin Volatility Channeling:** Maximizing profit envelopes formed by price fluctuation width multiplied by execution volume limits.
3. **Transaction Batching Window Optimization:** Balancing delayed settlement time windows with batch throughput capacity.

---

## Algorithmic Approach
A naive brute-force checks all pairs in $O(N^2)$.
The optimal **Two-Pointer Greedy** strategy converges in $O(N)$:
- Place `left = 0` and `right = len(height) - 1`.
- Compute `Area = min(height[left], height[right]) * (right - left)`.
- Update `maxWater`.
- **Greedy Invariant:** As pointers move inward, width strictly decreases. To find a larger area, height must increase. Because the area is strictly constrained by the **shorter line**, moving the taller line can never increase the area. Therefore, we advance the pointer at the shorter boundary (`left++` if `height[left] < height[right]`, else `right--`).

---

## Complexity Analysis
- **Time Complexity:** $O(N)$ - Pointers meet after traversing at most $N$ steps.
- **Space Complexity:** $O(1)$ - Constant extra space with zero heap allocations.

---

## Critical Edge Cases
- Minimal length ($N = 2$).
- Monotonically increasing or decreasing heights.
- High boundary peaks with flat interiors.

---

## 💻 Production Go Implementation

```go
package containerwater

// MaxArea finds two lines that together with the x-axis form a container that contains the most water.
// Time Complexity: O(N), Space Complexity: O(1)
func MaxArea(height []int) int {
    left := 0
    right := len(height) - 1
    maxWater := 0

    for left < right {
        hLeft := height[left]
        hRight := height[right]

        // Height of the water is bounded by the shorter line
        currentHeight := hLeft
        if hRight < currentHeight {
            currentHeight = hRight
        }

        currentWidth := right - left
        currentArea := currentHeight * currentWidth

        if currentArea > maxWater {
            maxWater = currentArea
        }

        // Move the pointer of the shorter line inward to search for higher bounds
        if hLeft < hRight {
            left++
        } else {
            right--
        }
    }

    return maxWater
}
```

---

## 🧪 Running Unit Tests

```bash
go test -v ./11-container-with-most-water/...
go test -race ./11-container-with-most-water/...
```

---

[⬅️ Back to All Exercises](../README.md)
