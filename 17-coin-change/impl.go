package coinchange

import "math"

// CoinChange calculates the fewest number of coins needed to make up the given amount.
// Returns -1 if that amount of money cannot be made up by any combination of the coins.
// Time Complexity: O(Amount * len(coins)), Space Complexity: O(Amount)
func CoinChange(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}

	// dp[i] represents the minimum coins needed to make amount i
	dp := make([]int, amount+1)
	for i := 1; i <= amount; i++ {
		dp[i] = math.MaxInt32
	}

	dp[0] = 0

	for i := 1; i <= amount; i++ {
		for _, coin := range coins {
			if i-coin >= 0 && dp[i-coin] != math.MaxInt32 {
				if dp[i-coin]+1 < dp[i] {
					dp[i] = dp[i-coin] + 1
				}
			}
		}
	}

	if dp[amount] == math.MaxInt32 {
		return -1
	}

	return dp[amount]
}
