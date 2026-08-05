func remainingMethods(n int, k int, invocations [][]int) []int {

	// Build adjacency list
	graph := make([][]int, n)

	for _, edge := range invocations {
		u := edge[0]
		v := edge[1]
		graph[u] = append(graph[u], v)
	}

	// Marks suspicious methods
	vis := make([]bool, n)

	// DFS from method k
	var dfs func(int)

	dfs = func(u int) {
		vis[u] = true

		for _, v := range graph[u] {
			if !vis[v] {
				dfs(v)
			}
		}
	}

	dfs(k)

	// Check whether removal is allowed
	for _, edge := range invocations {
		u := edge[0]
		v := edge[1]

		if !vis[u] && vis[v] {
			ans := make([]int, n)

			for i := 0; i < n; i++ {
				ans[i] = i
			}

			return ans
		}
	}

	// Collect remaining methods
	ans := []int{}

	for i := 0; i < n; i++ {
		if !vis[i] {
			ans = append(ans, i)
		}
	}

	return ans
}