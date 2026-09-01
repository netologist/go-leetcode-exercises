package meetingrooms

import "sort"

// MinMeetingRooms calculates the minimum number of conference rooms (or concurrent worker servers)
// required to accommodate all scheduled intervals.
// Time Complexity: O(N log N), Space Complexity: O(N)
func MinMeetingRooms(intervals [][]int) int {
	if len(intervals) == 0 {
		return 0
	}

	n := len(intervals)
	starts := make([]int, n)
	ends := make([]int, n)

	for i, interval := range intervals {
		starts[i] = interval[0]
		ends[i] = interval[1]
	}

	sort.Ints(starts)
	sort.Ints(ends)

	rooms := 0
	endIdx := 0

	for startIdx := range n {
		if starts[startIdx] < ends[endIdx] {
			// A new meeting starts before the earliest ending meeting finishes -> need a new room
			rooms++
		} else {
			// The earliest meeting ended -> room is reused
			endIdx++
		}
	}

	return rooms
}
