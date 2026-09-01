package coinchange

import "testing"

func RunCoinChangeTests(t *testing.T) {
	tests := []struct {
		name     string
		coins    []int
		amount   int
		expected int
	}{
		{
			name:     "Standard coins",
			coins:    []int{1, 2, 5},
			amount:   11,
			expected: 3, // 5 + 5 + 1
		},
		{
			name:     "Impossible amount",
			coins:    []int{2},
			amount:   3,
			expected: -1,
		},
		{
			name:     "Zero amount",
			coins:    []int{1},
			amount:   0,
			expected: 0,
		},
		{
			name:     "Single coin match",
			coins:    []int{5},
			amount:   5,
			expected: 1,
		},
		{
			name:     "Greedy failure case where DP is required",
			coins:    []int{1, 3, 4, 5},
			amount:   7,
			expected: 2, // 4 + 3 (Greedy 5+1+1 = 3 coins is suboptimal)
		},
		{
			name:     "Large amount with small denominations",
			coins:    []int{1, 2, 5, 10, 20, 50, 100},
			amount:   289,
			expected: 8, // 100*2 + 50 + 20 + 10 + 5 + 2*2 = 8 coins
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CoinChange(tt.coins, tt.amount)
			if got != tt.expected {
				t.Errorf("CoinChange(%v, %d) = %d, want %d", tt.coins, tt.amount, got, tt.expected)
			}
		})
	}
}
