func assignEdgeWeights(edges [][]int, queries [][]int) []int {
	const MOD = 1000000007

	n := len(edges) + 1

	// Compute required binary lifting height
	LOG := 1
	for (1 << LOG) <= n {
		LOG++
	}

	// Build adjacency list
	graph := make([][]int, n+1)

	for _, e := range edges {
		u, v := e[0], e[1]

		graph[u] = append(graph[u], v)
		graph[v] = append(graph[v], u)
	}

	depth := make([]int, n+1)

	// up[node][j] = 2^j-th ancestor
	up := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		up[i] = make([]int, LOG)
		for j := 0; j < LOG; j++ {
			up[i][j] = 1
		}
	}

	// DFS preprocessing
	var dfs func(int, int)

	dfs = func(node, parent int) {
		up[node][0] = parent

		for j := 1; j < LOG; j++ {
			up[node][j] = up[up[node][j-1]][j-1]
		}

		for _, next := range graph[node] {
			if next == parent {
				continue
			}

			depth[next] = depth[node] + 1
			dfs(next, node)
		}
	}

	dfs(1, 1)

	// LCA using binary lifting
	lca := func(a, b int) int {
		if depth[a] < depth[b] {
			a, b = b, a
		}

		diff := depth[a] - depth[b]

		for j := LOG - 1; j >= 0; j-- {
			if ((diff >> j) & 1) == 1 {
				a = up[a][j]
			}
		}

		if a == b {
			return a
		}

		for j := LOG - 1; j >= 0; j-- {
			if up[a][j] != up[b][j] {
				a = up[a][j]
				b = up[b][j]
			}
		}

		return up[a][0]
	}

	// Precompute powers of 2
	pow2 := make([]int, n+1)
	pow2[0] = 1

	for i := 1; i <= n; i++ {
		pow2[i] = int((int64(pow2[i-1]) * 2) % MOD)
	}

	ans := make([]int, len(queries))

	for i, q := range queries {
		u, v := q[0], q[1]

		ancestor := lca(u, v)

		dist := depth[u] + depth[v] - 2*depth[ancestor]

		if dist == 0 {
			ans[i] = 0
		} else {
			ans[i] = pow2[dist-1]
		}
	}

	return ans
}