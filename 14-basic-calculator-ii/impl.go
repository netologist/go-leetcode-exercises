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
