package accountsmerge

import "sort"

// UnionFind manages disjoint set operations with path compression and union by rank.
type UnionFind struct {
	parent []int
	rank   []int
}

func newUnionFind(n int) *UnionFind {
	parent := make([]int, n)
	rank := make([]int, n)
	for i := range n {
		parent[i] = i
		rank[i] = 1
	}
	return &UnionFind{parent: parent, rank: rank}
}

func (uf *UnionFind) Find(x int) int {
	if uf.parent[x] != x {
		uf.parent[x] = uf.Find(uf.parent[x]) // Path compression
	}
	return uf.parent[x]
}

func (uf *UnionFind) Union(x, y int) {
	rootX := uf.Find(x)
	rootY := uf.Find(y)

	if rootX != rootY {
		// Union by rank
		if uf.rank[rootX] < uf.rank[rootY] {
			uf.parent[rootX] = rootY
		} else if uf.rank[rootX] > uf.rank[rootY] {
			uf.parent[rootY] = rootX
		} else {
			uf.parent[rootY] = rootX
			uf.rank[rootX]++
		}
	}
}

// AccountsMerge merges user accounts sharing common email addresses.
// Time Complexity: O(N * K * log(N * K) + N * K * alpha(N)), Space Complexity: O(N * K)
func AccountsMerge(accounts [][]string) [][]string {
	uf := newUnionFind(len(accounts))
	emailToAccountIdx := make(map[string]int)

	// Step 1: Union accounts sharing at least one common email
	for i, account := range accounts {
		for j := 1; j < len(account); j++ {
			email := account[j]
			if existingIdx, exists := emailToAccountIdx[email]; exists {
				uf.Union(i, existingIdx)
			} else {
				emailToAccountIdx[email] = i
			}
		}
	}

	// Step 2: Group emails by their root parent account ID
	rootToEmails := make(map[int]map[string]bool)
	for email, idx := range emailToAccountIdx {
		root := uf.Find(idx)
		if rootToEmails[root] == nil {
			rootToEmails[root] = make(map[string]bool)
		}
		rootToEmails[root][email] = true
	}

	// Step 3: Format output with account name and sorted emails
	result := make([][]string, 0, len(rootToEmails))
	for root, emailsMap := range rootToEmails {
		name := accounts[root][0]
		emails := make([]string, 0, len(emailsMap))
		for email := range emailsMap {
			emails = append(emails, email)
		}
		sort.Strings(emails)

		mergedAccount := append([]string{name}, emails...)
		result = append(result, mergedAccount)
	}

	return result
}
