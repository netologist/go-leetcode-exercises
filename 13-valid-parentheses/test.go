package validparentheses

import "testing"

func RunValidParenthesesTests(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{
			name:     "Simple parentheses",
			input:    "()",
			expected: true,
		},
		{
			name:     "Mixed valid brackets",
			input:    "()[]{}",
			expected: true,
		},
		{
			name:     "Mismatched bracket type",
			input:    "(]",
			expected: false,
		},
		{
			name:     "Cross overlapping brackets",
			input:    "([)]",
			expected: false,
		},
		{
			name:     "Nested valid brackets",
			input:    "{[()]}",
			expected: true,
		},
		{
			name:     "Odd length string",
			input:    "({[",
			expected: false,
		},
		{
			name:     "Starts with closing bracket",
			input:    "]",
			expected: false,
		},
		{
			name:     "Empty string",
			input:    "",
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IsValid(tt.input)
			if got != tt.expected {
				t.Errorf("IsValid(%q) = %v, want %v", tt.input, got, tt.expected)
			}
		})
	}
}
