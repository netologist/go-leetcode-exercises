package twosum

import (
	"reflect"
	"testing"
)

func RunTwoSumTests(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		target   int
		expected []int
	}{
		{
			name:     "Standard positive numbers",
			nums:     []int{2, 7, 11, 15},
			target:   9,
			expected: []int{0, 1},
		},
		{
			name:     "Middle elements",
			nums:     []int{3, 2, 4},
			target:   6,
			expected: []int{1, 2},
		},
		{
			name:     "Duplicate numbers",
			nums:     []int{3, 3},
			target:   6,
			expected: []int{0, 1},
		},
		{
			name:     "Negative numbers",
			nums:     []int{-1, -2, -3, -4, -5},
			target:   -8,
			expected: []int{2, 4},
		},
		{
			name:     "Target zero with opposite signs",
			nums:     []int{10, -10, 5, 2},
			target:   0,
			expected: []int{0, 1},
		},
		{
			name:     "No matching pair",
			nums:     []int{1, 2, 3},
			target:   10,
			expected: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := TwoSum(tt.nums, tt.target)
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("TwoSum(%v, %d) = %v, want %v", tt.nums, tt.target, got, tt.expected)
			}
		})
	}
}
