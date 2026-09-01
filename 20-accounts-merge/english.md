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
