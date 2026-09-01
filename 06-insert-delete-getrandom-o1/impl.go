package randomizedset

import "math/rand/v2"

// RandomizedSet supports Insert, Remove, and GetRandom in average O(1) time.
type RandomizedSet struct {
	nums []int       // Holds the values for O(1) random index lookup
	pos  map[int]int // Maps value -> index in nums slice
}

// Constructor initializes the RandomizedSet.
func Constructor() RandomizedSet {
	return RandomizedSet{
		nums: make([]int, 0),
		pos:  make(map[int]int),
	}
}

// Insert inserts a value to the set. Returns true if the set did not already contain the specified element.
func (r *RandomizedSet) Insert(val int) bool {
	if _, exists := r.pos[val]; exists {
		return false
	}
	r.pos[val] = len(r.nums)
	r.nums = append(r.nums, val)
	return true
}

// Remove removes a value from the set using swap-and-pop technique.
// Returns true if the set contained the specified element.
func (r *RandomizedSet) Remove(val int) bool {
	idx, exists := r.pos[val]
	if !exists {
		return false
	}

	lastIdx := len(r.nums) - 1
	lastVal := r.nums[lastIdx]

	// Swap target with last element
	r.nums[idx] = lastVal
	r.pos[lastVal] = idx

	// Pop last element
	r.nums = r.nums[:lastIdx]
	delete(r.pos, val)

	return true
}

// GetRandom returns a random element from the current set with uniform probability.
func (r *RandomizedSet) GetRandom() int {
	return r.nums[rand.IntN(len(r.nums))]
}
