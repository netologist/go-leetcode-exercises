package islands

import "testing"

func RunIslandsTests(t *testing.T) {
	tests := []struct {
		name     string
		grid     [][]byte
		expected int
	}{
		{
			name: "Single large island",
			grid: [][]byte{
				{'1', '1', '1', '1', '0'},
				{'1', '1', '0', '1', '0'},
				{'1', '1', '0', '0', '0'},
				{'0', '0', '0', '0', '0'},
			},
			expected: 1,
		},
		{
			name: "Multiple distinct islands",
			grid: [][]byte{
				{'1', '1', '0', '0', '0'},
				{'1', '1', '0', '0', '0'},
				{'0', '0', '1', '0', '0'},
				{'0', '0', '0', '1', '1'},
			},
			expected: 3,
		},
		{
			name: "All water",
			grid: [][]byte{
				{'0', '0', '0'},
				{'0', '0', '0'},
			},
			expected: 0,
		},
		{
			name: "All land",
			grid: [][]byte{
				{'1', '1'},
				{'1', '1'},
			},
			expected: 1,
		},
		{
			name: "Diagonal islands (not connected horizontally or vertically)",
			grid: [][]byte{
				{'1', '0'},
				{'0', '1'},
			},
			expected: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Make copy of grid to avoid mutating test definition across runs
			gridCopy := make([][]byte, len(tt.grid))
			for i := range tt.grid {
				gridCopy[i] = make([]byte, len(tt.grid[i]))
				copy(gridCopy[i], tt.grid[i])
			}

			got := NumIslands(gridCopy)
			if got != tt.expected {
				t.Errorf("NumIslands = %d, want %d", got, tt.expected)
			}
		})
	}
}
