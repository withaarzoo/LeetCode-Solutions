func latestDayToCross(row int, col int, cells [][]int) int {
	n := row * col
	top, bottom := n, n+1

	parent := make([]int, n+2)
	rank := make([]int, n+2)
	grid := make([]bool, n)

	for i := 0; i < n+2; i++ {
		parent[i] = i
	}

	var find func(int) int
	find = func(x int) int {
		if parent[x] != x {
			parent[x] = find(parent[x])
		}
		return parent[x]
	}

	union := func(a, b int) {
		a = find(a)
		b = find(b)
		if a == b {
			return
		}
		if rank[a] < rank[b] {
			parent[a] = b
		} else {
			parent[b] = a
			if rank[a] == rank[b] {
				rank[a]++
			}
		}
	}

	dr := []int{1, -1, 0, 0}
	dc := []int{0, 0, 1, -1}

	for d := n - 1; d >= 0; d-- {
		r := cells[d][0] - 1
		c := cells[d][1] - 1
		id := r*col + c
		grid[id] = true

		if r == 0 {
			union(id, top)
		}
		if r == row-1 {
			union(id, bottom)
		}

		for k := 0; k < 4; k++ {
			nr := r + dr[k]
			nc := c + dc[k]
			if nr >= 0 && nr < row && nc >= 0 && nc < col {
				nid := nr*col + nc
				if grid[nid] {
					union(id, nid)
				}
			}
		}

		if find(top) == find(bottom) {
			return d
		}
	}
	return 0
}
