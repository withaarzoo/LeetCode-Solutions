// Depth-First Search (DFS) function to traverse the graph
func dfs(node int, adj [][]int, visited map[int]bool) {
    // Mark the current node as visited
    visited[node] = true
    
    // Explore all neighbors of the current node
    for _, neighbor := range adj[node] {
        // If the neighbor hasn't been visited yet, perform DFS on it
        if !visited[neighbor] {
            dfs(neighbor, adj, visited)
        }
    }
}

// Function to determine the maximum number of stones that can be removed
func removeStones(stones [][]int) int {
    // Get the number of stones
    n := len(stones)
    
    // Initialize an adjacency list to represent the graph
    adj := make([][]int, n)
    
    // Build the graph by connecting stones that are in the same row or column
    for i := 0; i < n; i++ {
        for j := i + 1; j < n; j++ {
            // If two stones share the same row or column, they are connected
            if stones[i][0] == stones[j][0] || stones[i][1] == stones[j][1] {
                // Add an edge between the two stones in the adjacency list
                adj[i] = append(adj[i], j)
                adj[j] = append(adj[j], i)
            }
        }
    }

    // Map to keep track of visited nodes
    visited := make(map[int]bool)
    
    // Counter for the number of connected components in the graph
    numComponents := 0

    // Iterate through each stone and perform DFS to find connected components
    for i := 0; i < n; i++ {
        // If the stone hasn't been visited, it's a new component
        if !visited[i] {
            // Perform DFS starting from this stone
            dfs(i, adj, visited)
            // Increment the number of components found
            numComponents++
        }
    }

    // The maximum number of stones that can be removed is the total number of stones
    // minus the number of connected components (each component leaves one stone behind)
    return n - numComponents
}
