func hasValidPath(grid [][]byte) bool {
    m, n := len(grid), len(grid[0])
    if (m+n-1)%2 != 0 {
        return false
    }
    maxBal := m + n
    dp := make([][][]bool, m)
    for i := range dp {
        dp[i] = make([][]bool, n)
        for j := range dp[i] {
            dp[i][j] = make([]bool, maxBal)
        }
    }
    start := 1
    if grid[0][0] == ')' {
        start = -1
    }
    if start < 0 {
        return false
    }
    dp[0][0][start] = true
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            for bal := 0; bal < maxBal; bal++ {
                if !dp[i][j][bal] {
                    continue
                }
                if i+1 < m {
                    nb := bal
                    if grid[i+1][j] == '(' {
                        nb++
                    } else {
                        nb--
                    }
                    if nb >= 0 && nb < maxBal {
                        dp[i+1][j][nb] = true
                    }
                }
                if j+1 < n {
                    nb := bal
                    if grid[i][j+1] == '(' {
                        nb++
                    } else {
                        nb--
                    }
                    if nb >= 0 && nb < maxBal {
                        dp[i][j+1][nb] = true
                    }
                }
            }
        }
    }
    return dp[m-1][n-1][0]
}