func minCost(grid [][]int) int {
    m, n := len(grid), len(grid[0])
    dp := make([][][1024]bool, m)
    for i := range dp {
        dp[i] = make([][1024]bool, n)
    }
    dp[0][0][grid[0][0]] = true
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if i > 0 {
                for x := 0; x < 1024; x++ {
                    if dp[i-1][j][x] {
                        dp[i][j][x^grid[i][j]] = true
                    }
                }
            }
            if j > 0 {
                for x := 0; x < 1024; x++ {
                    if dp[i][j-1][x] {
                        dp[i][j][x^grid[i][j]] = true
                    }
                }
            }
        }
    }
    for x := 0; x < 1024; x++ {
        if dp[m-1][n-1][x] {
            return x
        }
    }
    return 0
}