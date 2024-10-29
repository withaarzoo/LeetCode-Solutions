func maxMoves(grid [][]int) int {
    m, n := len(grid), len(grid[0])
    dp := make([][]int, m)
    for i := range dp {
        dp[i] = make([]int, n)
    }
    maxMoves := 0

    for col := n - 2; col >= 0; col-- {
        for row := 0; row < m; row++ {
            if row > 0 && grid[row][col] < grid[row-1][col+1] {
                dp[row][col] = max(dp[row][col], dp[row-1][col+1]+1)
            }
            if grid[row][col] < grid[row][col+1] {
                dp[row][col] = max(dp[row][col], dp[row][col+1]+1)
            }
            if row < m-1 && grid[row][col] < grid[row+1][col+1] {
                dp[row][col] = max(dp[row][col], dp[row+1][col+1]+1)
            }
        }
    }

    for row := 0; row < m; row++ {
        maxMoves = max(maxMoves, dp[row][0])
    }
    return maxMoves
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}