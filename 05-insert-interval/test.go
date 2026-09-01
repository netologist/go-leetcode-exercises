package insertinterval

import (
	"reflect"
	"testing"
)

func RunInsertIntervalTests(t *testing.T) {
	tests := []struct {
		name        string
		intervals   [][]int
		newInterval []int
		expected    [][]int
	}{
		{
			name:        "Insert and merge in middle",
			intervals:   [][]int{{1, 3}, {6, 9}},
			newInterval: []int{2, 5},
			expected:    [][]int{{1, 5}, {6, 9}},
		},
		{
			name:        "Insert merging multiple intervals",
			intervals:   [][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}},
			newInterval: []int{4, 8},
			expected:    [][]int{{1, 2}, {3, 10}, {12, 16}},
		},
		{
			name:        "Insert into empty intervals",
			intervals:   [][]int{},
			newInterval: []int{5, 7},
			expected:    [][]int{{5, 7}},
		},
		{
			name:        "Insert before all intervals (non-overlapping)",
			intervals:   [][]int{{3, 5}, {6, 9}},
			newInterval: []int{1, 2},
			expected:    [][]int{{1, 2}, {3, 5}, {6, 9}},
		},
		{
			name:        "Insert after all intervals (non-overlapping)",
			intervals:   [][]int{{1, 2}, {3, 5}},
			newInterval: []int{6, 8},
			expected:    [][]int{{1, 2}, {3, 5}, {6, 8}},
		},
		{
			name:        "Insert engulfing all existing intervals",
			intervals:   [][]int{{2, 3}, {4, 5}, {6, 7}},
			newInterval: []int{1, 10},
			expected:    [][]int{{1, 10}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Insert(tt.intervals, tt.newInterval)
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("Insert(%v, %v) = %v, want %v", tt.intervals, tt.newInterval, got, tt.expected)
			}
		})
	}
}
