func eventualSafeNodes(graph [][]int) []int {
    n := len(graph)
    reversedGraph := make([][]int, n)
    inDegree := make([]int, n)
    
    // Reverse the graph and calculate in-degree
    for i := 0; i < n; i++ {
        for _, neighbor := range graph[i] {
            reversedGraph[neighbor] = append(reversedGraph[neighbor], i)
            inDegree[i]++
        }
    }
    
    // Find all terminal nodes
    queue := []int{}
    for i := 0; i < n; i++ {
        if inDegree[i] == 0 {
            queue = append(queue, i)
        }
    }
    
    // Topological sorting to find safe nodes
    safeNodes := []int{}
    for len(queue) > 0 {
        node := queue[0]
        queue = queue[1:]
        safeNodes = append(safeNodes, node)
        
        for _, neighbor := range reversedGraph[node] {
            inDegree[neighbor]--
            if inDegree[neighbor] == 0 {
                queue = append(queue, neighbor)
            }
        }
    }
    
    sort.Ints(safeNodes)
    return safeNodes
}
