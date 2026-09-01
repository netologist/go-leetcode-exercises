package groupanagrams

// GroupAnagrams groups an array of strings into anagram clusters.
// In FinTech: Used for merchant string normalization and transaction tagging.
// Time Complexity: O(N * K) where N is number of strings and K is max string length.
// Space Complexity: O(N * K)
func GroupAnagrams(strs []string) [][]string {
	// [26]byte is comparable and can be used directly as a Go map key without string allocation
	groups := make(map[[26]byte][]string)

	for _, s := range strs {
		var count [26]byte
		for i := range len(s) {
			count[s[i]-'a']++
		}
		groups[count] = append(groups[count], s)
	}

	result := make([][]string, 0, len(groups))
	for _, group := range groups {
		result = append(result, group)
	}

	return result
}
