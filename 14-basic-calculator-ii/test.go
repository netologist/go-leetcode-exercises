package calculator

import "testing"

func RunCalculatorTests(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int
	}{
		{
			name:     "Multiplication with addition",
			input:    "3+2*2",
			expected: 7,
		},
		{
			name:     "Division with spaces",
			input:    " 3/2 ",
			expected: 1,
		},
		{
			name:     "Multiple mixed operators",
			input:    " 3+5 / 2 ",
			expected: 5,
		},
		{
			name:     "Consecutive subtractions and additions",
			input:    "1-1+1",
			expected: 1,
		},
		{
			name:     "Multi-digit numbers with multiplication",
			input:    "100 * 2 + 50 / 2 - 25",
			expected: 200,
		},
		{
			name:     "Single number",
			input:    "42",
			expected: 42,
		},
		{
			name:     "Chained divisions",
			input:    "14-3/2",
			expected: 13,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Calculate(tt.input)
			if got != tt.expected {
				t.Errorf("Calculate(%q) = %d, want %d", tt.input, got, tt.expected)
			}
		})
	}
}
