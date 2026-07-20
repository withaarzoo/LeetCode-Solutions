func shiftGrid(grid [][]int, k int) [][]int {

	// Get the grid dimensions
	m := len(grid)
	n := len(grid[0])

	// Total number of elements
	total := m * n

	// Ignore unnecessary complete rotations
	k %= total

	// Create the answer grid
	ans := make([][]int, m)
	for i := 0; i < m; i++ {
		ans[i] = make([]int, n)
	}

	// Move every element directly
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {

			// Flatten the current position
			oldIndex := i*n + j

			// Compute the shifted position
			newIndex := (oldIndex + k) % total

			// Convert back to row and column
			newRow := newIndex / n
			newCol := newIndex % n

			// Place the value
			ans[newRow][newCol] = grid[i][j]
		}
	}

	return ans
}