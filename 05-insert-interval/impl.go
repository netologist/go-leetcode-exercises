package insertinterval

// Insert inserts a new interval into a sorted list of non-overlapping intervals,
// merging if necessary, and returns the updated list.
// Time Complexity: O(N), Space Complexity: O(N)
func Insert(intervals [][]int, newInterval []int) [][]int {
	result := make([][]int, 0, len(intervals)+1)
	i := 0
	n := len(intervals)

	// Step 1: Append all intervals ending before newInterval starts
	for i < n && intervals[i][1] < newInterval[0] {
		result = append(result, intervals[i])
		i++
	}

	// Step 2: Merge all overlapping intervals with newInterval
	for i < n && intervals[i][0] <= newInterval[1] {
		if intervals[i][0] < newInterval[0] {
			newInterval[0] = intervals[i][0]
		}
		if intervals[i][1] > newInterval[1] {
			newInterval[1] = intervals[i][1]
		}
		i++
	}
	result = append(result, newInterval)

	// Step 3: Append all remaining intervals that start after newInterval ends
	for i < n {
		result = append(result, intervals[i])
		i++
	}

	return result
}
