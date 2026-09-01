package subarraysum

import "testing"

func RunSubarraySumTests(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		k        int
		expected int
	}{
		{
			name:     "Standard positive numbers",
			nums:     []int{1, 1, 1},
			k:        2,
			expected: 2,
		},
		{
			name:     "Subarrays with multiple combinations",
			nums:     []int{1, 2, 3},
			k:        3,
			expected: 2, // [1, 2] and [3]
		},
		{
			name:     "Negative numbers in transaction stream",
			nums:     []int{1, -1, 0},
			k:        0,
			expected: 3, // [1, -1], [0], [1, -1, 0]
		},
		{
			name:     "Negative target sum",
			nums:     []int{-1, -1, 1},
			k:        -2,
			expected: 1, // [-1, -1]
		},
		{
			name:     "Single element matching k",
			nums:     []int{5},
			k:        5,
			expected: 1,
		},
		{
			name:     "No matching subarray",
			nums:     []int{1, 2, 3},
			k:        10,
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SubarraySum(tt.nums, tt.k)
			if got != tt.expected {
				t.Errorf("SubarraySum(%v, %d) = %d, want %d", tt.nums, tt.k, got, tt.expected)
			}
		})
	}
}
