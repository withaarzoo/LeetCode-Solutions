func maximumSafenessFactor(grid [][]int) int {

	n := len(grid)

	// Distance from every cell to the nearest thief
	dist := make([][]int, n)
	for i := range dist {
		dist[i] = make([]int, n)
		for j := range dist[i] {
			dist[i][j] = -1
		}
	}

	type Pair struct {
		x, y int
	}

	queue := []Pair{}

	// Push every thief
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				dist[i][j] = 0
				queue = append(queue, Pair{i, j})
			}
		}
	}

	dir := []int{-1, 0, 1, 0, -1}
	head := 0

	// Multi-source BFS
	for head < len(queue) {

		cur := queue[head]
		head++

		for k := 0; k < 4; k++ {

			nx := cur.x + dir[k]
			ny := cur.y + dir[k+1]

			if nx < 0 || ny < 0 || nx >= n || ny >= n || dist[nx][ny] != -1 {
				continue
			}

			dist[nx][ny] = dist[cur.x][cur.y] + 1
			queue = append(queue, Pair{nx, ny})
		}
	}

	canReach := func(limit int) bool {

		if dist[0][0] < limit || dist[n-1][n-1] < limit {
			return false
		}

		vis := make([][]bool, n)
		for i := range vis {
			vis[i] = make([]bool, n)
		}

		bfs := []Pair{{0, 0}}
		vis[0][0] = true
		idx := 0

		for idx < len(bfs) {

			cur := bfs[idx]
			idx++

			if cur.x == n-1 && cur.y == n-1 {
				return true
			}

			for k := 0; k < 4; k++ {

				nx := cur.x + dir[k]
				ny := cur.y + dir[k+1]

				if nx < 0 || ny < 0 || nx >= n || ny >= n {
					continue
				}

				if vis[nx][ny] || dist[nx][ny] < limit {
					continue
				}

				vis[nx][ny] = true
				bfs = append(bfs, Pair{nx, ny})
			}
		}

		return false
	}

	left := 0
	right := 2 * n
	ans := 0

	// Binary search on the answer
	for left <= right {

		mid := left + (right-left)/2

		if canReach(mid) {
			ans = mid
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return ans
}