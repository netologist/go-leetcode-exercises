package stock

import "testing"

func RunStockTests(t *testing.T) {
	tests := []struct {
		name     string
		prices   []int
		expected int
	}{
		{
			name:     "Standard fluctuating stock prices",
			prices:   []int{7, 1, 5, 3, 6, 4},
			expected: 5, // Buy at 1, sell at 6
		},
		{
			name:     "Monotonically decreasing prices (No profit)",
			prices:   []int{7, 6, 4, 3, 1},
			expected: 0,
		},
		{
			name:     "Monotonically increasing prices",
			prices:   []int{1, 2, 3, 4, 5},
			expected: 4, // Buy at 1, sell at 5
		},
		{
			name:     "Flat prices",
			prices:   []int{3, 3, 3, 3},
			expected: 0,
		},
		{
			name:     "Single price / Empty array",
			prices:   []int{5},
			expected: 0,
		},
		{
			name:     "High volatility spike",
			prices:   []int{10, 2, 8, 1, 9},
			expected: 8, // Buy at 1, sell at 9
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MaxProfit(tt.prices)
			if got != tt.expected {
				t.Errorf("MaxProfit(%v) = %d, want %d", tt.prices, got, tt.expected)
			}
		})
	}
}
