func countCompleteComponents(n int, edges [][]int) int {

	// Build adjacency list
	graph := make([][]int, n)

	for _, edge := range edges {
		u, v := edge[0], edge[1]
		graph[u] = append(graph[u], v)
		graph[v] = append(graph[v], u)
	}

	// Track visited nodes
	visited := make([]bool, n)

	// DFS for one connected component
	var dfs func(int, *int, *int)

	dfs = func(node int, vertices *int, degreeSum *int) {

		// Mark node as visited
		visited[node] = true

		// Count current vertex
		*vertices++

		// Add current node's degree
		*degreeSum += len(graph[node])

		// Visit neighbors
		for _, next := range graph[node] {
			if !visited[next] {
				dfs(next, vertices, degreeSum)
			}
		}
	}

	answer := 0

	// Process every component
	for i := 0; i < n; i++ {

		if visited[i] {
			continue
		}

		vertices := 0
		degreeSum := 0

		dfs(i, &vertices, &degreeSum)

		// Every edge is counted twice
		edgeCount := degreeSum / 2

		// Check whether this component is complete
		if edgeCount == vertices*(vertices-1)/2 {
			answer++
		}
	}

	return answer
}