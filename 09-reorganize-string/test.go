package reorganizestring

import "testing"

func RunReorganizeStringTests(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		possible bool
	}{
		{
			name:     "Standard rearrangable string",
			input:    "aab",
			possible: true,
		},
		{
			name:     "Impossible string (too many duplicates)",
			input:    "aaab",
			possible: false,
		},
		{
			name:     "Multiple characters balanced",
			input:    "vvvlo",
			possible: true,
		},
		{
			name:     "Single character",
			input:    "a",
			possible: true,
		},
		{
			name:     "All unique characters",
			input:    "abcdef",
			possible: true,
		},
		{
			name:     "Two characters alternating",
			input:    "abababab",
			possible: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ReorganizeString(tt.input)
			if !tt.possible {
				if got != "" {
					t.Errorf("ReorganizeString(%q) = %q, want empty string", tt.input, got)
				}
				return
			}

			if len(got) != len(tt.input) {
				t.Errorf("ReorganizeString(%q) output length %d != input length %d", tt.input, len(got), len(tt.input))
			}

			// Validate no two adjacent characters are identical
			for i := 1; i < len(got); i++ {
				if got[i] == got[i-1] {
					t.Errorf("ReorganizeString(%q) produced invalid adjacent characters %c%c at index %d in %q", tt.input, got[i-1], got[i], i, got)
				}
			}
		})
	}
}
