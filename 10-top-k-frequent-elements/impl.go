package topkfrequent

// TopKFrequent returns the k most frequent elements in nums.
// Time Complexity: O(N) using Bucket Sort, Space Complexity: O(N)
func TopKFrequent(nums []int, k int) []int {
	counts := make(map[int]int)
	for _, num := range nums {
		counts[num]++
	}

	// Buckets where index represents frequency (0 to len(nums))
	buckets := make([][]int, len(nums)+1)
	for num, count := range counts {
		buckets[count] = append(buckets[count], num)
	}

	result := make([]int, 0, k)
	// Iterate from highest frequency bucket downwards
	for freq := len(buckets) - 1; freq >= 0 && len(result) < k; freq-- {
		if len(buckets[freq]) > 0 {
			for _, num := range buckets[freq] {
				result = append(result, num)
				if len(result) == k {
					break
				}
			}
		}
	}

	return result
}
