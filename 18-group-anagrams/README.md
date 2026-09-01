# Group Anagrams

## LeetCode Information
- **Number:** 49
- **Difficulty:** Medium
- **FinTech Companies:** Stripe, Plaid, Bloomberg, Robinhood, Revolut

---

## FinTech Relevance & Real-World Application
Critical in transaction data enrichment and merchant resolution pipelines:
1. **Merchant Name Canonicalization:** Clustering noisy billing descriptor strings from debit/credit feeds that share identical token signatures.
2. **Transaction Tagging & Pattern Deduplication:** Grouping permutated payload attributes for automated accounting reconciliation.
3. **Synthetic Fraud Fingerprinting:** Detecting permutations of registration strings across identity verification funnels.

---

## Algorithmic Approach
- Rather than sorting every individual string ($O(N \times K \log K)$), we construct a **26-element character frequency array** signature ($O(N \times K)$).
- **Go Specific Advantage:** In Go, fixed-size arrays such as `[26]byte` are value types and strictly comparable. This allows them to be used directly as map keys (`map[[26]byte][]string`) without string allocation or hashing overhead!
- Iterate through each string, populate its `[26]byte` frequency array, and append the string to the corresponding map slice.

---

## Complexity Analysis
- **Time Complexity:** $O(N \times K)$ where $N$ is the number of strings and $K$ is the maximum string length.
- **Space Complexity:** $O(N \times K)$ - Map storage holding all strings grouped by signature.

---

## Critical Edge Cases
- Empty string input `[""]`.
- Single character strings.
- Sets of identical strings (e.g. `["uber", "uber"]`).

---

## 💻 Production Go Implementation

```go
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
```

---

## 🧪 Running Unit Tests

```bash
go test -v ./18-group-anagrams/...
go test -race ./18-group-anagrams/...
```

---

[⬅️ Back to All Exercises](../README.md)
