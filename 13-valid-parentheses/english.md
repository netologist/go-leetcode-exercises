# Valid Parentheses

## LeetCode Information
- **Number:** 20
- **Difficulty:** Easy
- **FinTech Companies:** Stripe, Bloomberg, Robinhood, Plaid, Two Sigma

---

## FinTech Relevance & Real-World Application
Foundational for syntax validation in financial protocols, rule engines, and formula evaluators:
1. **Financial Message Syntax Verification (FIX / SWIFT / ISO-8583):** Ensuring balanced enclosing blocks in structured interchange envelopes.
2. **Underwriting & Risk Rule Parsing:** Validating boolean tree expressions in credit scoring and loan origination decision engines.
3. **Ledger Query Language (LQL) Parsing:** Verifying syntactic balance in custom financial DSLs and SQL-like accounting queries.

---

## Algorithmic Approach
A standard **LIFO Stack** pattern:
- Map closing delimiters `)`, `}`, `]` to their opening counterparts `(`, `{`, `[`.
- Iterate through each character in the string:
  - If it is a closing delimiter: verify stack is non-empty and the top element matches the required opening symbol, then pop.
  - If it is an opening delimiter: push onto the stack.
- Return `true` if and only if the stack is completely empty at termination.

---

## Complexity Analysis
- **Time Complexity:** $O(N)$ - Single scan over the string with $O(1)$ stack push and pop operations.
- **Space Complexity:** $O(N)$ - Stack space holding up to $N$ characters in the worst case.

---

## Critical Edge Cases
- Odd string length (instantly invalid via parity check).
- Starting with a closing bracket (e.g. `]`).
- Unclosed trailing openers (e.g. `({[`).
