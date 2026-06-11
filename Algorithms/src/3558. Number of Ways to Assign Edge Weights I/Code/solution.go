func assignEdgeWeights(edges [][]int) int {
	const MOD int64 = 1000000007

	n := len(edges) + 1

	// Build adjacency list
	graph := make([][]int, n+1)

	for _, e := range edges {
		u, v := e[0], e[1]

		graph[u] = append(graph[u], v)
		graph[v] = append(graph[v], u)
	}

	type Pair struct {
		node  int
		depth int
	}

	// Iterative DFS
	stack := []Pair{{1, 0}}

	visited := make([]bool, n+1)
	visited[1] = true

	maxDepth := 0

	for len(stack) > 0 {
		cur := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if cur.depth > maxDepth {
			maxDepth = cur.depth
		}

		for _, next := range graph[cur.node] {
			if !visited[next] {
				visited[next] = true
				stack = append(stack, Pair{next, cur.depth + 1})
			}
		}
	}

	// Fast modular exponentiation
	base := int64(2)
	exp := maxDepth - 1
	result := int64(1)

	for exp > 0 {
		// Multiply when current bit is set
		if exp&1 == 1 {
			result = (result * base) % MOD
		}

		// Square the base
		base = (base * base) % MOD
		exp >>= 1
	}

	return int(result)
}