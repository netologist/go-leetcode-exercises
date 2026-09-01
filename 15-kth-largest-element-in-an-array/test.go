package kthlargest

import "testing"

func RunKthLargestTests(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		k        int
		expected int
	}{
		{
			name:     "Standard array with k=2",
			nums:     []int{3, 2, 1, 5, 6, 4},
			k:        2,
			expected: 5,
		},
		{
			name:     "Array with duplicates and k=4",
			nums:     []int{3, 2, 3, 1, 2, 4, 5, 5, 6},
			k:        4,
			expected: 4,
		},
		{
			name:     "Single element array",
			nums:     []int{1},
			k:        1,
			expected: 1,
		},
		{
			name:     "Negative numbers",
			nums:     []int{-1, -2, -3, -4, -5},
			k:        2,
			expected: -2,
		},
		{
			name:     "All identical elements",
			nums:     []int{7, 7, 7, 7},
			k:        3,
			expected: 7,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Make copy to preserve input
			numsCopy := make([]int, len(tt.nums))
			copy(numsCopy, tt.nums)

			got := FindKthLargest(numsCopy, tt.k)
			if got != tt.expected {
				t.Errorf("FindKthLargest(%v, %d) = %d, want %d", tt.nums, tt.k, got, tt.expected)
			}
		})
	}
}
