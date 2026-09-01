# ⚡ FinTech & Systems Live Coding Suite in Go

A curated, production-grade reference implementation of **20 mission-critical algorithmic and systems coding problems** frequently featured in technical interviews and live coding assessments at premier FinTech, Quantitative Trading, and Infrastructure companies (**Stripe, Robinhood, Citadel, Bloomberg, Two Sigma, Brex, Revolut, Plaid, Coinbase, Adyen**).

Every problem includes:
- 📖 **FinTech Domain Context & Real-World Systems Application**
- 🧠 **Algorithmic Intuition & Data Structure Architecture**
- 💻 **Idiomatic, Thread-Safe Go Implementation** (`impl.go`)
- 🧪 **Comprehensive Unit & Concurrency Test Suites** (`test.go`, `impl_test.go`)
- ⏱️ **Rigorous Time & Space Complexity Analysis** ($O(N)$, $O(1)$)

Created and maintained by **Hasan Özgan** ([netologist.org](https://netologist.org) · [@netologist](https://github.com/netologist)).

---

## 🧭 Problem Index & FinTech Domain Mapping

Click on any problem name or folder link to navigate to the detailed architectural explanation, algorithm walkthrough, and Go implementation:

| # | Problem | LeetCode | Difficulty | Real-World FinTech Domain & Target Companies | Time | Space | Package Path |
|:---:|:---|:---:|:---:|:---|:---:|:---:|:---|
| **01** | [**LRU Cache**](./01-lru-cache/README.md) | #146 | Medium | *Market Tick Caching, Idempotency Keys, Session Caches* (Stripe, Robinhood, Citadel) | $O(1)$ | $O(C)$ | [`01-lru-cache`](./01-lru-cache/README.md) |
| **02** | [**Two Sum**](./02-two-sum/README.md) | #1 | Easy | *Double-Entry Ledger Balancing, FX Currency Pair Matching* (Stripe, Plaid) | $O(N)$ | $O(N)$ | [`02-two-sum`](./02-two-sum/README.md) |
| **03** | [**Stock Trading Max Profit**](./03-best-time-to-buy-and-sell-stock/README.md) | #121 | Easy | *Algorithmic Trading, Maximum Drawdown, Real-Time PnL* (Citadel, Two Sigma) | $O(N)$ | $O(1)$ | [`03-best-time-to-buy-and-sell-stock`](./03-best-time-to-buy-and-sell-stock/README.md) |
| **04** | [**Merge Intervals**](./04-merge-intervals/README.md) | #56 | Medium | *Order Book Price Tiers, Interest Calculation Periods* (Stripe, Bloomberg) | $O(N \log N)$ | $O(N)$ | [`04-merge-intervals`](./04-merge-intervals/README.md) |
| **05** | [**Insert Interval**](./05-insert-interval/README.md) | #57 | Medium | *Dynamic Fee Tier Adjustments, Settlement Window Merging* (Citadel, Robinhood) | $O(N)$ | $O(N)$ | [`05-insert-interval`](./05-insert-interval/README.md) |
| **06** | [**RandomizedSet $O(1)$**](./06-insert-delete-getrandom-o1/README.md) | #380 | Medium | *Liquidity Pool Sampling, Real-Time Fraud Audit Inspection* (Robinhood) | $O(1)$ | $O(N)$ | [`06-insert-delete-getrandom-o1`](./06-insert-delete-getrandom-o1/README.md) |
| **07** | [**Meeting Rooms II**](./07-meeting-rooms-ii/README.md) | #253 | Medium | *Concurrent Payment Server Allocation, Core Banking Scalability* (Bloomberg, Stripe) | $O(N \log N)$ | $O(N)$ | [`07-meeting-rooms-ii`](./07-meeting-rooms-ii/README.md) |
| **08** | [**Subarray Sum Equals K**](./08-subarray-sum-equals-k/README.md) | #560 | Medium | *Transaction Window Reconciliation, Anti-Money Laundering (AML)* (Stripe, Citadel) | $O(N)$ | $O(N)$ | [`08-subarray-sum-equals-k`](./08-subarray-sum-equals-k/README.md) |
| **09** | [**Reorganize String**](./09-reorganize-string/README.md) | #767 | Medium | *Merchant Rate-Limit Scheduling, Symbol Dispatcher* (Stripe, Citadel) | $O(N)$ | $O(1)$ | [`09-reorganize-string`](./09-reorganize-string/README.md) |
| **10** | [**Top K Frequent Elements**](./10-top-k-frequent-elements/README.md) | #347 | Medium | *Most Actively Traded Equities, High-Velocity Fraud Hotspots* (Bloomberg) | $O(N)$ | $O(N)$ | [`10-top-k-frequent-elements`](./10-top-k-frequent-elements/README.md) |
| **11** | [**Container With Most Water**](./11-container-with-most-water/README.md) | #11 | Medium | *Liquidity Depth Maximization, Spread Arbitrage Windows* (Robinhood) | $O(N)$ | $O(1)$ | [`11-container-with-most-water`](./11-container-with-most-water/README.md) |
| **12** | [**Token Bucket Rate Limiter**](./12-rate-limiter-token-bucket/README.md) | Custom | Medium | *Stripe/Revolut Public API Throttling, Burst Capacity Control* (Stripe, Revolut) | $O(1)$ | $O(1)$ | [`12-rate-limiter-token-bucket`](./12-rate-limiter-token-bucket/README.md) |
| **13** | [**Valid Parentheses**](./13-valid-parentheses/README.md) | #20 | Easy | *FIX / SWIFT / ISO-8583 Message Syntax Parsing* (Stripe, Plaid) | $O(N)$ | $O(N)$ | [`13-valid-parentheses`](./13-valid-parentheses/README.md) |
| **14** | [**Basic Calculator II**](./14-basic-calculator-ii/README.md) | #227 | Medium | *Fee & Interest Rule Engine, Arithmetic Precedence* (Bloomberg) | $O(N)$ | $O(N)$ | [`14-basic-calculator-ii`](./14-basic-calculator-ii/README.md) |
| **15** | [**Kth Largest Element in Array**](./15-kth-largest-element-in-an-array/README.md) | #215 | Medium | *Latency Percentile Computations (p99), Dark Pool Crossing* (Citadel) | $O(N)$ | $O(1)$ | [`15-kth-largest-element-in-an-array`](./15-kth-largest-element-in-an-array/README.md) |
| **16** | [**Number of Islands**](./16-number-of-islands/README.md) | #200 | Medium | *Fraud Ring Detection, Connected Account Clustering* (Bloomberg, Stripe) | $O(M \times N)$ | $O(M \times N)$ | [`16-number-of-islands`](./16-number-of-islands/README.md) |
| **17** | [**Coin Change (DP)**](./17-coin-change/README.md) | #322 | Medium | *ATM Cash Dispensing Optimization, Crypto UTXO Selection* (Citadel, Robinhood) | $O(A \times C)$ | $O(A)$ | [`17-coin-change`](./17-coin-change/README.md) |
| **18** | [**Group Anagrams**](./18-group-anagrams/README.md) | #49 | Medium | *Merchant Name Normalization, Bank Statement Enrichment* (Stripe, Plaid) | $O(N \times K)$ | $O(N \times K)$ | [`18-group-anagrams`](./18-group-anagrams/README.md) |
| **19** | [**Design Hit Counter**](./19-design-hit-counter/README.md) | #362 | Medium | *Real-Time Transaction Per Second (TPS), Error Rate Breakers* (Stripe, Datadog) | $O(1)$ | $O(1)$ | [`19-design-hit-counter`](./19-design-hit-counter/README.md) |
| **20** | [**Accounts Merge (Union-Find)**](./20-accounts-merge/README.md) | #721 | Medium | *KYC Single Customer View, Sybil & Multi-Account Detection* (Plaid, Stripe) | $O(NK \log NK)$ | $O(NK)$ | [`20-accounts-merge`](./20-accounts-merge/README.md) |

---

## 🚀 Getting Started & Running Tests

### Prerequisites
- **Go 1.22+** (tested up to Go 1.26+)

### Running All Test Suites
```bash
# Run tests across all 20 packages
go test ./...

# Run tests with verbose output
go test -v ./...

# Run tests with data race detector enabled
go test -race ./...
```

### Running Tests for a Specific Problem
```bash
# Test LRU Cache
go test -v ./01-lru-cache/...

# Test Token Bucket Rate Limiter
go test -v ./12-rate-limiter-token-bucket/...

# Test Accounts Merge (Union-Find)
go test -v ./20-accounts-merge/...
```

---

## 🌐 Author & Engineering Notes

Deep-dive architecture notes, distributed system mental models, and software engineering principles are published on:
👉 **[netologist.org](https://netologist.org)** — Digital Garden & Software Engineering Notes

---

## 📄 License

This repository is licensed under the [MIT License](LICENSE).
