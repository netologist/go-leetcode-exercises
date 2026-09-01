package mergeintervals

import "sort"

// Merge merges all overlapping intervals and returns an array of the non-overlapping intervals.
// Time Complexity: O(N log N), Space Complexity: O(N)
func Merge(intervals [][]int) [][]int {
	if len(intervals) <= 1 {
		return intervals
	}

	// Sort intervals by starting time
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	merged := make([][]int, 0, len(intervals))
	merged = append(merged, intervals[0])

	for i := 1; i < len(intervals); i++ {
		current := intervals[i]
		lastIdx := len(merged) - 1
		lastMerged := merged[lastIdx]

		if current[0] <= lastMerged[1] {
			// Overlap detected: merge by taking the max of ending bounds
			if current[1] > lastMerged[1] {
				merged[lastIdx][1] = current[1]
			}
		} else {
			// Disjoint interval: append as a new block
			merged = append(merged, current)
		}
	}

	return merged
}
