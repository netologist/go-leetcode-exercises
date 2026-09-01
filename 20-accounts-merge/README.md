# Accounts Merge

## LeetCode Information
- **Number:** 721
- **Difficulty:** Medium
- **FinTech Companies:** Stripe, Plaid, Robinhood, Brex, Coinbase, Revolut

---

## FinTech Relevance & Real-World Application
The industry standard for Entity Resolution and KYC (Know Your Customer) systems:
1. **Single Customer View (SCV) Consolidation:** Merging disparate banking, investment, and merchant records into a unified legal entity profile via shared identifiers.
2. **Sybil & Promotional Abuse Detection:** Detecting multi-account farming rings where users register with variations of names but overlap on recovery emails or phone aliases.
3. **Plaid/Open Banking Financial Aggregation:** Linking multiple institutional credentials belonging to the same individual into consolidated balance sheet views.

---

## Algorithmic Approach
This is modeled as finding connected components using **Disjoint Set Union (Union-Find) with Path Compression and Union by Rank**:

1. **Union-Find Structure:**
   - `parent []int`: Tracks root representative for each account index (with $O(\alpha(N))$ inverse Ackermann time via path compression).
   - `rank []int`: Prevents skewed tree heights during union operations.
2. **Map Emails & Union Accounts:**
   - Maintain `emailToAccountIdx map[string]int`.
   - As we iterate through each email, if it was seen under an earlier account index `oldIdx`, execute `uf.Union(currentIdx, oldIdx)`.
3. **Group by Root Representative:**
   - Group all emails under their respective `root = uf.Find(idx)`.
4. **Format & Sort:**
   - Retrieve the account name from `accounts[root][0]`, sort the unique emails lexicographically, and assemble the final record.

---

## Complexity Analysis
- **Time Complexity:** $O(N \cdot K \log(N \cdot K) + N \cdot K \cdot \alpha(N))$ where $N$ is the number of accounts and $K$ is the maximum email count per account. Sorting unique emails dominates.
- **Space Complexity:** $O(N \cdot K)$ - Space for Union-Find arrays, email mapping, and output structures.

---

## Critical Edge Cases
- Disjoint accounts that happen to share identical names (must remain separated if no emails overlap).
- Long transitive linkage chains ($A \leftrightarrow B \leftrightarrow C \dots$).

---

## 💻 Production Go Implementation

```go
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
```

---

## 🧪 Running Unit Tests

```bash
go test -v ./20-accounts-merge/...
go test -race ./20-accounts-merge/...
```

---

[⬅️ Back to All Exercises](../README.md)
