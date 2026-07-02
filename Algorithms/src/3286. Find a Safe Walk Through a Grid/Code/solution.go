func findSafeWalk(grid [][]int, health int) bool {
	m := len(grid)
	n := len(grid[0])

	const INF = int(1e9)

	// Store the minimum health lost for every cell
	dist := make([][]int, m)
	for i := 0; i < m; i++ {
		dist[i] = make([]int, n)
		for j := 0; j < n; j++ {
			dist[i][j] = INF
		}
	}

	// Simple deque implementation
	type Pair struct {
		x, y int
	}

	deque := make([]Pair, 0)
	head := 0

	// Starting cost includes the starting cell
	dist[0][0] = grid[0][0]
	deque = append(deque, Pair{0, 0})

	dx := []int{-1, 0, 1, 0}
	dy := []int{0, 1, 0, -1}

	for head < len(deque) {
		cur := deque[head]
		head++

		// Explore all four directions
		for k := 0; k < 4; k++ {
			nx := cur.x + dx[k]
			ny := cur.y + dy[k]

			// Skip cells outside the grid
			if nx < 0 || ny < 0 || nx >= m || ny >= n {
				continue
			}

			// Cost after entering the next cell
			newCost := dist[cur.x][cur.y] + grid[nx][ny]

			// Keep only the best cost
			if newCost < dist[nx][ny] {
				dist[nx][ny] = newCost

				// Weight 0 should be processed first
				if grid[nx][ny] == 0 {
					deque = append([]Pair{{nx, ny}}, deque[head:]...)
					head = 0
				} else {
					deque = append(deque, Pair{nx, ny})
				}
			}
		}
	}

	// Health must remain at least 1
	return dist[m-1][n-1] < health
}