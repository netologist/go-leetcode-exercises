package containerwater

import "testing"

func RunContainerWaterTests(t *testing.T) {
	tests := []struct {
		name     string
		height   []int
		expected int
	}{
		{
			name:     "Standard array with peak middle and ends",
			height:   []int{1, 8, 6, 2, 5, 4, 8, 3, 7},
			expected: 49,
		},
		{
			name:     "Two elements",
			height:   []int{1, 1},
			expected: 1,
		},
		{
			name:     "Monotonically increasing heights",
			height:   []int{1, 2, 3, 4, 5, 6, 7, 8},
			expected: 16,
		},
		{
			name:     "Monotonically decreasing heights",
			height:   []int{8, 7, 6, 5, 4, 3, 2, 1},
			expected: 16,
		},
		{
			name:     "High peaks at ends",
			height:   []int{10, 1, 1, 1, 10},
			expected: 40,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MaxArea(tt.height)
			if got != tt.expected {
				t.Errorf("MaxArea(%v) = %d, want %d", tt.height, got, tt.expected)
			}
		})
	}
}
