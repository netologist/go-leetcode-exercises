package meetingrooms

import "testing"

func RunMeetingRoomsTests(t *testing.T) {
	tests := []struct {
		name      string
		intervals [][]int
		expected  int
	}{
		{
			name:      "Standard overlapping intervals",
			intervals: [][]int{{0, 30}, {5, 10}, {15, 20}},
			expected:  2,
		},
		{
			name:      "Sequential non-overlapping intervals",
			intervals: [][]int{{7, 10}, {2, 4}},
			expected:  1,
		},
		{
			name:      "Multiple simultaneous meetings",
			intervals: [][]int{{1, 10}, {2, 7}, {3, 19}, {8, 12}, {10, 20}, {11, 30}},
			expected:  4,
		},
		{
			name:      "Consecutive meetings where start equals previous end",
			intervals: [][]int{{1, 5}, {5, 10}, {10, 15}},
			expected:  1,
		},
		{
			name:      "Empty intervals",
			intervals: [][]int{},
			expected:  0,
		},
		{
			name:      "Single interval",
			intervals: [][]int{{5, 10}},
			expected:  1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MinMeetingRooms(tt.intervals)
			if got != tt.expected {
				t.Errorf("MinMeetingRooms(%v) = %d, want %d", tt.intervals, got, tt.expected)
			}
		})
	}
}
