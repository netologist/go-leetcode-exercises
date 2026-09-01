package twosum

// TwoSum finds two indices such that their values add up to target.
// Returns an empty slice if no such pair exists.
// Time Complexity: O(N), Space Complexity: O(N)
func TwoSum(nums []int, target int) []int {
	seen := make(map[int]int, len(nums)) // value -> index

	for i, num := range nums {
		complement := target - num
		if idx, exists := seen[complement]; exists {
			return []int{idx, i}
		}
		seen[num] = i
	}

	return nil
}
