package groupanagrams

import (
	"sort"
	"strings"
	"testing"
)

func RunGroupAnagramsTests(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		expected [][]string
	}{
		{
			name:     "Standard mixed anagram words",
			input:    []string{"eat", "tea", "tan", "ate", "nat", "bat"},
			expected: [][]string{{"eat", "tea", "ate"}, {"tan", "nat"}, {"bat"}},
		},
		{
			name:     "Single empty string",
			input:    []string{""},
			expected: [][]string{{""}},
		},
		{
			name:     "Single character",
			input:    []string{"a"},
			expected: [][]string{{"a"}},
		},
		{
			name:     "All identical words",
			input:    []string{"uber", "uber", "uber"},
			expected: [][]string{{"uber", "uber", "uber"}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GroupAnagrams(tt.input)
			if len(got) != len(tt.expected) {
				t.Fatalf("GroupAnagrams returned %d groups, want %d", len(got), len(tt.expected))
			}

			// Canonicalize groups for order-independent comparison
			canonicalGot := canonicalize(got)
			canonicalExpected := canonicalize(tt.expected)

			for i := range canonicalGot {
				if canonicalGot[i] != canonicalExpected[i] {
					t.Errorf("Group %d = %s, want %s", i, canonicalGot[i], canonicalExpected[i])
				}
			}
		})
	}
}

func canonicalize(groups [][]string) []string {
	res := make([]string, len(groups))
	for i, group := range groups {
		sortedGroup := make([]string, len(group))
		copy(sortedGroup, group)
		sort.Strings(sortedGroup)
		res[i] = strings.Join(sortedGroup, ",")
	}
	sort.Strings(res)
	return res
}
