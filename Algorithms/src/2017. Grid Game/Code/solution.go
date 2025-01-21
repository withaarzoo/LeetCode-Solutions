func gridGame(grid [][]int) int64 {
    n := len(grid[0])
    topSuffix := make([]int64, n)
    bottomPrefix := make([]int64, n)
    
    // Calculate suffix sum for the top row
    topSuffix[n-1] = int64(grid[0][n-1])
    for i := n - 2; i >= 0; i-- {
        topSuffix[i] = topSuffix[i+1] + int64(grid[0][i])
    }
    
    // Calculate prefix sum for the bottom row
    bottomPrefix[0] = int64(grid[1][0])
    for i := 1; i < n; i++ {
        bottomPrefix[i] = bottomPrefix[i-1] + int64(grid[1][i])
    }
    
    // Find the minimum maximum points Robot 2 can collect
    result := int64(1<<63 - 1)
    for i := 0; i < n; i++ {
        top := int64(0)
        if i+1 < n {
            top = topSuffix[i+1]
        }
        bottom := int64(0)
        if i > 0 {
            bottom = bottomPrefix[i-1]
        }
        result = min(result, max(top, bottom))
    }
    
    return result
}

func min(a, b int64) int64 {
    if a < b {
        return a
    }
    return b
}

func max(a, b int64) int64 {
    if a > b {
        return a
    }
    return b
}
