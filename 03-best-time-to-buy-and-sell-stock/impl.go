package stock

import "math"

// MaxProfit calculates the maximum profit achievable from a single buy and sell transaction.
// Time Complexity: O(N), Space Complexity: O(1)
func MaxProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}

	minPrice := math.MaxInt
	maxProfit := 0

	for _, price := range prices {
		if price < minPrice {
			minPrice = price
		} else if profit := price - minPrice; profit > maxProfit {
			maxProfit = profit
		}
	}

	return maxProfit
}
