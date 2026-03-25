func canPartitionGrid(grid [][]int) bool {
    m, n := len(grid), len(grid[0])
    
    total := 0
    
    // Step 1: Total sum
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            total += grid[i][j]
        }
    }
    
    // Step 2: Odd check
    if total%2 != 0 {
        return false
    }
    
    target := total / 2
    
    // Step 3: Horizontal cut
    rowSum := 0
    for i := 0; i < m-1; i++ {
        for j := 0; j < n; j++ {
            rowSum += grid[i][j]
        }
        if rowSum == target {
            return true
        }
    }
    
    // Step 4: Column sums
    colSum := make([]int, n)
    for j := 0; j < n; j++ {
        for i := 0; i < m; i++ {
            colSum[j] += grid[i][j]
        }
    }
    
    // Step 5: Vertical cut
    curr := 0
    for j := 0; j < n-1; j++ {
        curr += colSum[j]
        if curr == target {
            return true
        }
    }
    
    return false
}