package reorganizestring

import (
	"container/heap"
	"strings"
)

// CharFreq represents a character and its count.
type CharFreq struct {
	char  byte
	count int
}

// MaxHeap implements heap.Interface for CharFreq based on count descending.
type MaxHeap []CharFreq

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i].count > h[j].count }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MaxHeap) Push(x any) {
	*h = append(*h, x.(CharFreq))
}

func (h *MaxHeap) Pop() any {
	old := *h
	n := len(old)
	item := old[n-1]
	*h = old[:n-1]
	return item
}

// ReorganizeString rearranges characters so that no two adjacent characters are identical.
// Returns an empty string if impossible.
// Time Complexity: O(N log A) where A is alphabet size (A <= 26), Space Complexity: O(A)
func ReorganizeString(s string) string {
	counts := make(map[byte]int)
	for i := range len(s) {
		counts[s[i]]++
	}

	h := &MaxHeap{}
	heap.Init(h)

	maxAllowed := (len(s) + 1) / 2
	for char, count := range counts {
		if count > maxAllowed {
			return ""
		}
		heap.Push(h, CharFreq{char: char, count: count})
	}

	var result strings.Builder
	result.Grow(len(s))

	for h.Len() >= 2 {
		first := heap.Pop(h).(CharFreq)
		second := heap.Pop(h).(CharFreq)

		result.WriteByte(first.char)
		result.WriteByte(second.char)

		first.count--
		second.count--

		if first.count > 0 {
			heap.Push(h, first)
		}
		if second.count > 0 {
			heap.Push(h, second)
		}
	}

	if h.Len() > 0 {
		last := heap.Pop(h).(CharFreq)
		result.WriteByte(last.char)
	}

	return result.String()
}
