package topkfrequent

import (
	"sort"
	"testing"
)

func RunTopKFrequentTests(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		k        int
		expected []int
	}{
		{
			name:     "Standard multiple frequent elements",
			nums:     []int{1, 1, 1, 2, 2, 3},
			k:        2,
			expected: []int{1, 2},
		},
		{
			name:     "Single element",
			nums:     []int{1},
			k:        1,
			expected: []int{1},
		},
		{
			name:     "All elements distinct",
			nums:     []int{1, 2, 3, 4, 5},
			k:        3,
			expected: []int{1, 2, 3}, // any 3 valid, length verified
		},
		{
			name:     "Negative numbers and zeros",
			nums:     []int{-1, -1, 0, 0, 0, 4},
			k:        2,
			expected: []int{0, -1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := TopKFrequent(tt.nums, tt.k)
			if len(got) != tt.k {
				t.Fatalf("expected length %d, got %d", tt.k, len(got))
			}

			// For deterministic comparison when multiple valid answers exist, check frequencies
			freqMap := make(map[int]int)
			for _, n := range tt.nums {
				freqMap[n]++
			}

			sort.Ints(got)
			sort.Ints(tt.expected)

			if tt.name != "All elements distinct" {
				for i := range got {
					if freqMap[got[i]] != freqMap[tt.expected[i]] {
						t.Errorf("TopKFrequent(%v, %d) = %v, want elements with matching frequencies %v", tt.nums, tt.k, got, tt.expected)
					}
				}
			}
		})
	}
}
