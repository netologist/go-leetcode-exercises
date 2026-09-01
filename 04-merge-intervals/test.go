package mergeintervals

import (
	"reflect"
	"testing"
)

func RunMergeIntervalsTests(t *testing.T) {
	tests := []struct {
		name      string
		intervals [][]int
		expected  [][]int
	}{
		{
			name:      "Standard overlapping intervals",
			intervals: [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}},
			expected:  [][]int{{1, 6}, {8, 10}, {15, 18}},
		},
		{
			name:      "Touching border intervals",
			intervals: [][]int{{1, 4}, {4, 5}},
			expected:  [][]int{{1, 5}},
		},
		{
			name:      "Nested interval inside larger interval",
			intervals: [][]int{{1, 10}, {2, 6}, {3, 5}},
			expected:  [][]int{{1, 10}},
		},
		{
			name:      "Unsorted intervals input",
			intervals: [][]int{{8, 10}, {1, 3}, {15, 18}, {2, 6}},
			expected:  [][]int{{1, 6}, {8, 10}, {15, 18}},
		},
		{
			name:      "Completely disjoint intervals",
			intervals: [][]int{{1, 2}, {4, 5}, {7, 8}},
			expected:  [][]int{{1, 2}, {4, 5}, {7, 8}},
		},
		{
			name:      "Single interval / Empty input",
			intervals: [][]int{{1, 5}},
			expected:  [][]int{{1, 5}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Merge(tt.intervals)
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("Merge(%v) = %v, want %v", tt.intervals, got, tt.expected)
			}
		})
	}
}
