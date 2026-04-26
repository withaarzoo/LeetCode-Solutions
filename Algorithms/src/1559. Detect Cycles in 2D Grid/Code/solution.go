func containsCycle(grid [][]byte) bool {
	m := len(grid)
	n := len(grid[0])

	visited := make([][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}

	dr := []int{1, -1, 0, 0}
	dc := []int{0, 0, 1, -1}

	type Node struct {
		r, c int
		pr, pc int
	}

	for r := 0; r < m; r++ {
		for c := 0; c < n; c++ {
			if visited[r][c] {
				continue
			}

			stack := []Node{{r, c, -1, -1}}
			visited[r][c] = true

			for len(stack) > 0 {
				cur := stack[len(stack)-1]
				stack = stack[:len(stack)-1]

				for k := 0; k < 4; k++ {
					nr := cur.r + dr[k]
					nc := cur.c + dc[k]

					if nr < 0 || nr >= m || nc < 0 || nc >= n {
						continue
					}
					if grid[nr][nc] != grid[cur.r][cur.c] {
						continue
					}
					if nr == cur.pr && nc == cur.pc {
						continue
					}

					if visited[nr][nc] {
						return true
					}

					visited[nr][nc] = true
					stack = append(stack, Node{nr, nc, cur.r, cur.c})
				}
			}
		}
	}

	return false
}