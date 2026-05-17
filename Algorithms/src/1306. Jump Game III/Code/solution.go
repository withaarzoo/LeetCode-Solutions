func canReach(arr []int, start int) bool {
    
    // Visited array to prevent revisiting indexes
    visited := make([]bool, len(arr))

    // DFS function
    var dfs func(int) bool

    dfs = func(index int) bool {

        // Invalid index
        if index < 0 || index >= len(arr) {
            return false
        }

        // Skip already visited indexes
        if visited[index] {
            return false
        }

        // Found value 0
        if arr[index] == 0 {
            return true
        }

        // Mark current index as visited
        visited[index] = true

        // Explore forward and backward jumps
        return dfs(index+arr[index]) || dfs(index-arr[index])
    }

    // Start DFS from given index
    return dfs(start)
}