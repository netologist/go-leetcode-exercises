# Basic Calculator II (Arithmetic Precedence Engine)

## LeetCode Information
- **Number:** 227
- **Difficulty:** Medium
- **FinTech Companies:** Bloomberg, Stripe, Robinhood, Citadel, Goldman Sachs

---

## FinTech Relevance & Real-World Application
Foundational in financial formula engines, dynamic billing, and secure DSLs:
1. **Dynamic Fee Calculation Engines:** Evaluating parameterized commission formulas without relying on unsafe script evaluators (`eval`).
2. **Interest & Yield Compounding Calculators:** Parsing expressions such as `principal * rate + fee / days` from configuration files or client inputs.
3. **Sandboxed Arithmetic Security:** Preventing arbitrary code execution vulnerabilities when executing financial formulas supplied by third-party integrations.

---

## Algorithmic Approach
Using an **Explicit Operator Precedence Stack**:
- Iterate character by character across the string.
- Build multi-digit numbers: `currentNum = currentNum * 10 + int(ch - '0')`.
- When encountering an operator or the end of the string (ignoring whitespace), resolve the **previous operator (`lastOp`)**:
  - `+`: Push `currentNum` onto the stack.
  - `-`: Push `-currentNum` onto the stack.
  - `*`: Pop the top value, multiply by `currentNum`, and push the product back.
  - `/`: Pop the top value, perform integer division by `currentNum`, and push back.
- Update `lastOp = ch` and reset `currentNum = 0`.
- Sum all accumulated values in the stack at the conclusion of the scan.

---

## Complexity Analysis
- **Time Complexity:** $O(N)$ - Exactly one pass across string $s$ with constant-time stack operations.
- **Space Complexity:** $O(N)$ - Stack capacity proportional to the number of terms in expression.

---

## Critical Edge Cases
- Multi-digit numbers (e.g. `100 * 2`).
- Arbitrary whitespace padding (e.g. ` 3+5 / 2 `).
- Consecutive lower-precedence operators (e.g. `1-1+1`).
- Processing the terminal number when `i == len(s)-1`.

---

## 💻 Production Go Implementation

```go
package calculator

import "unicode"

// Calculate evaluates a mathematical string expression containing non-negative integers and '+', '-', '*', '/' operators.
// Uses operator precedence without relying on standard eval engines.
// Time Complexity: O(N), Space Complexity: O(N)
func Calculate(s string) int {
    stack := make([]int, 0)
    currentNum := 0
    lastOp := '+'

    for i, ch := range s {
        if unicode.IsDigit(ch) {
            currentNum = currentNum*10 + int(ch-'0')
        }

        // If operator or last character reached (ignoring spaces)
        if (!unicode.IsDigit(ch) && ch != ' ') || i == len(s)-1 {
            switch lastOp {
            case '+':
                stack = append(stack, currentNum)
            case '-':
                stack = append(stack, -currentNum)
            case '*':
                top := stack[len(stack)-1]
                stack[len(stack)-1] = top * currentNum
            case '/':
                top := stack[len(stack)-1]
                stack[len(stack)-1] = top / currentNum
            }
            lastOp = ch
            currentNum = 0
        }
    }

    result := 0
    for _, val := range stack {
        result += val
    }

    return result
}
```

---

## 🧪 Running Unit Tests

```bash
go test -v ./14-basic-calculator-ii/...
go test -race ./14-basic-calculator-ii/...
```

---

[⬅️ Back to All Exercises](../README.md)
