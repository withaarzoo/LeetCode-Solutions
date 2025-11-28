package main

func maxKDivisibleComponents(n int, edges [][]int, values []int, k int) int {
	// Build adjacency list
	adj := make([][]int, n)
	for _, e := range edges {
		u, v := e[0], e[1]
		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
	}

	ans := 0

	var dfs func(u, parent int) int64
	dfs = func(u, parent int) int64 {
		// current node value modulo k (as int64)
		sum := int64(values[u] % k)

		for _, v := range adj[u] {
			if v == parent {
				continue
			}
			childRem := dfs(v, u)
			sum = (sum + childRem) % int64(k)
		}

		if sum%int64(k) == 0 {
			ans++
			return 0
		}
		return sum
	}

	dfs(0, -1)
	return ans
}
