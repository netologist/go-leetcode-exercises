package subarraysum

// SubarraySum finds the total number of continuous subarrays whose sum equals k.
// Time Complexity: O(N), Space Complexity: O(N)
func SubarraySum(nums []int, k int) int {
	prefixCount := make(map[int]int, len(nums)+1)
	// Base case: a prefix sum of 0 has occurred once before processing elements
	prefixCount[0] = 1

	currentSum := 0
	totalCount := 0

	for _, num := range nums {
		currentSum += num
		// If (currentSum - k) exists in map, it means a sub-segment summing to k was found
		if count, exists := prefixCount[currentSum-k]; exists {
			totalCount += count
		}
		prefixCount[currentSum]++
	}

	return totalCount
}
