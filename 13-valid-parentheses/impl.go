package validparentheses

// IsValid determines if the input string containing brackets '()', '{}', '[]' is valid.
// Time Complexity: O(N), Space Complexity: O(N)
func IsValid(s string) bool {
	if len(s)%2 != 0 {
		return false
	}

	stack := make([]rune, 0, len(s))
	matching := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}

	for _, ch := range s {
		if open, isClose := matching[ch]; isClose {
			// If closing bracket, top of stack must match corresponding opening bracket
			if len(stack) == 0 || stack[len(stack)-1] != open {
				return false
			}
			stack = stack[:len(stack)-1] // Pop
		} else {
			// Opening bracket
			stack = append(stack, ch) // Push
		}
	}

	return len(stack) == 0
}
