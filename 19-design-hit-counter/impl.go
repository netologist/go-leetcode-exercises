package hitcounter

import "sync"

// HitCounter counts the number of hits received in the past 5 minutes (300 seconds).
// Uses a circular bucket buffer for strictly O(1) space and O(1) time operations.
type HitCounter struct {
	mu    sync.Mutex
	times [300]int // Timestamps associated with bucket index (timestamp % 300)
	hits  [300]int // Hit count accumulated in bucket index
}

// Constructor initializes the HitCounter.
func Constructor() HitCounter {
	return HitCounter{}
}

// Hit records a hit at given timestamp (in seconds granularity).
// Time Complexity: O(1), Space Complexity: O(1)
func (c *HitCounter) Hit(timestamp int) {
	c.mu.Lock()
	defer c.mu.Unlock()

	idx := timestamp % 300
	if c.times[idx] != timestamp {
		// New second window: overwrite old stale bucket
		c.times[idx] = timestamp
		c.hits[idx] = 1
	} else {
		// Same second window: increment hits
		c.hits[idx]++
	}
}

// GetHits returns the number of hits in the past 300 seconds from given timestamp.
// Time Complexity: O(1) (iterates exactly 300 buckets), Space Complexity: O(1)
func (c *HitCounter) GetHits(timestamp int) int {
	c.mu.Lock()
	defer c.mu.Unlock()

	total := 0
	for i := range 300 {
		if timestamp-c.times[i] < 300 && c.times[i] > 0 {
			total += c.hits[i]
		}
	}

	return total
}
