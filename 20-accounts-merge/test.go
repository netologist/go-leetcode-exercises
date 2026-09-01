package accountsmerge

import (
	"sort"
	"strings"
	"testing"
)

func RunAccountsMergeTests(t *testing.T) {
	tests := []struct {
		name     string
		accounts [][]string
		expected [][]string
	}{
		{
			name: "Standard accounts with overlapping emails",
			accounts: [][]string{
				{"John", "johnsmith@mail.com", "john_newyork@mail.com"},
				{"John", "johnsmith@mail.com", "john00@mail.com"},
				{"Mary", "mary@mail.com"},
				{"John", "johnnybravo@mail.com"},
			},
			expected: [][]string{
				{"John", "john00@mail.com", "john_newyork@mail.com", "johnsmith@mail.com"},
				{"Mary", "mary@mail.com"},
				{"John", "johnnybravo@mail.com"},
			},
		},
		{
			name: "Chained overlapping accounts",
			accounts: [][]string{
				{"David", "d1@mail.com", "d2@mail.com"},
				{"David", "d2@mail.com", "d3@mail.com"},
				{"David", "d3@mail.com", "d4@mail.com"},
			},
			expected: [][]string{
				{"David", "d1@mail.com", "d2@mail.com", "d3@mail.com", "d4@mail.com"},
			},
		},
		{
			name: "Completely disjoint accounts",
			accounts: [][]string{
				{"Alex", "a@mail.com"},
				{"Bob", "b@mail.com"},
			},
			expected: [][]string{
				{"Alex", "a@mail.com"},
				{"Bob", "b@mail.com"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AccountsMerge(tt.accounts)
			if len(got) != len(tt.expected) {
				t.Fatalf("AccountsMerge returned %d accounts, want %d", len(got), len(tt.expected))
			}

			canonicalGot := canonicalizeAccounts(got)
			canonicalExpected := canonicalizeAccounts(tt.expected)

			for i := range canonicalGot {
				if canonicalGot[i] != canonicalExpected[i] {
					t.Errorf("Account %d = %s, want %s", i, canonicalGot[i], canonicalExpected[i])
				}
			}
		})
	}
}

func canonicalizeAccounts(accounts [][]string) []string {
	res := make([]string, len(accounts))
	for i, acc := range accounts {
		name := acc[0]
		emails := make([]string, len(acc)-1)
		copy(emails, acc[1:])
		sort.Strings(emails)
		res[i] = name + ":" + strings.Join(emails, ",")
	}
	sort.Strings(res)
	return res
}
