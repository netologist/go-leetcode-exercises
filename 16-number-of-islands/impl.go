package islands

// NumIslands counts the number of connected components of '1's (land) surrounded by '0's (water).
// In FinTech: Used for identifying isolated fraud clusters or transaction networks.
// Time Complexity: O(M * N), Space Complexity: O(M * N) worst-case call stack or queue
func NumIslands(grid [][]byte) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	rows := len(grid)
	cols := len(grid[0])
	count := 0

	// Directions: up, right, down, left
	dirs := [][2]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

	var dfs func(r, c int)
	dfs = func(r, c int) {
		// Out of bounds or water
		if r < 0 || r >= rows || c < 0 || c >= cols || grid[r][c] != '1' {
			return
		}

		// Mark visited in-place
		grid[r][c] = '0'

		for _, d := range dirs {
			dfs(r+d[0], c+d[1])
		}
	}

	for r := range rows {
		for c := range cols {
			if grid[r][c] == '1' {
				count++
				dfs(r, c)
			}
		}
	}

	return count
}
