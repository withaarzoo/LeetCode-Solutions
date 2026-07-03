func findMaxPathScore(edges [][]int, online []bool, k int64) int {
	n := len(online)

	type Edge struct {
		to int
		w  int
	}

	// Build graph
	graph := make([][]Edge, n)
	indegree := make([]int, n)

	for _, e := range edges {
		u, v, w := e[0], e[1], e[2]
		graph[u] = append(graph[u], Edge{v, w})
		indegree[v]++
	}

	// Topological order
	queue := make([]int, 0)
	for i := 0; i < n; i++ {
		if indegree[i] == 0 {
			queue = append(queue, i)
		}
	}

	topo := make([]int, 0)

	for head := 0; head < len(queue); head++ {
		u := queue[head]
		topo = append(topo, u)

		for _, e := range graph[u] {
			indegree[e.to]--
			if indegree[e.to] == 0 {
				queue = append(queue, e.to)
			}
		}
	}

	check := func(limit int) bool {
		const INF int64 = 1 << 60

		// Minimum cost to every node
		dp := make([]int64, n)
		for i := range dp {
			dp[i] = INF
		}
		dp[0] = 0

		for _, u := range topo {

			// Skip unreachable nodes
			if dp[u] == INF {
				continue
			}

			// Skip offline intermediate nodes
			if u != 0 && u != n-1 && !online[u] {
				continue
			}

			for _, e := range graph[u] {

				// Ignore small edges
				if e.w < limit {
					continue
				}

				// Cannot enter offline intermediate node
				if e.to != n-1 && !online[e.to] {
					continue
				}

				if dp[u]+int64(e.w) < dp[e.to] {
					dp[e.to] = dp[u] + int64(e.w)
				}
			}
		}

		return dp[n-1] <= k
	}

	left, right := 0, 1000000000
	ans := -1

	for left <= right {
		mid := left + (right-left)/2

		if check(mid) {
			ans = mid
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return ans
}