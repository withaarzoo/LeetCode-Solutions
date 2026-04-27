func hasValidPath(grid [][]int) bool {
	m, n := len(grid), len(grid[0])

	// Directions: left, right, up, down
	dirs := [4][2]int{
		{0, -1}, // left
		{0, 1},  // right
		{-1, 0}, // up
		{1, 0},  // down
	}

	// For each street type, which directions it supports.
	// 0 = left, 1 = right, 2 = up, 3 = down
	streetDirs := map[int][]int{
		1: {0, 1}, // left-right
		2: {2, 3}, // up-down
		3: {0, 3}, // left-down
		4: {1, 3}, // right-down
		5: {0, 2}, // left-up
		6: {1, 2}, // right-up
	}

	opposite := [4]int{1, 0, 3, 2}

	visited := make([][]bool, m)
	for i := range visited {
		visited[i] = make([]bool, n)
	}

	type Pair struct {
		r, c int
	}

	queue := make([]Pair, 0)
	queue = append(queue, Pair{0, 0})
	visited[0][0] = true
	head := 0

	for head < len(queue) {
		cur := queue[head]
		head++

		r, c := cur.r, cur.c
		if r == m-1 && c == n-1 {
			return true
		}

		for _, d := range streetDirs[grid[r][c]] {
			nr := r + dirs[d][0]
			nc := c + dirs[d][1]

			if nr < 0 || nr >= m || nc < 0 || nc >= n || visited[nr][nc] {
				continue
			}

			nextType := grid[nr][nc]
			ok := false
			for _, nd := range streetDirs[nextType] {
				if nd == opposite[d] {
					ok = true
					break
				}
			}

			if ok {
				visited[nr][nc] = true
				queue = append(queue, Pair{nr, nc})
			}
		}
	}

	return false
}