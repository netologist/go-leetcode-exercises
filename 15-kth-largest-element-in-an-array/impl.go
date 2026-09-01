package kthlargest

import "math/rand/v2"

// FindKthLargest returns the kth largest element in an unsorted array.
// Uses Quickselect with randomized pivot selection.
// Time Complexity: O(N) average, O(N^2) worst-case; Space Complexity: O(1) iterative
func FindKthLargest(nums []int, k int) int {
	// The kth largest element is at index targetIdx = len(nums) - k in sorted order (0-indexed)
	targetIdx := len(nums) - k
	left := 0
	right := len(nums) - 1

	for left <= right {
		pivotIdx := partition(nums, left, right)
		if pivotIdx == targetIdx {
			return nums[pivotIdx]
		} else if pivotIdx < targetIdx {
			left = pivotIdx + 1
		} else {
			right = pivotIdx - 1
		}
	}

	return nums[left]
}

func partition(nums []int, left, right int) int {
	// Pick random pivot and swap with right
	randIdx := left + rand.IntN(right-left+1)
	nums[randIdx], nums[right] = nums[right], nums[randIdx]

	pivot := nums[right]
	i := left

	for j := left; j < right; j++ {
		if nums[j] <= pivot {
			nums[i], nums[j] = nums[j], nums[i]
			i++
		}
	}

	nums[i], nums[right] = nums[right], nums[i]
	return i
}
