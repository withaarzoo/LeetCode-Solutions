func findMaxForm(strs []string, m int, n int) int {
    // dp[z][o] = max strings using at most z zeros and o ones
    dp := make([][]int, m+1)
    for i := range dp {
        dp[i] = make([]int, n+1)
    }

    for _, s := range strs {
        z, o := 0, 0
        for _, ch := range s {
            if ch == '0' {
                z++
            } else {
                o++
            }
        }

        // Backwards loops for 0/1 knapsack
        for i := m; i >= z; i-- {
            for j := n; j >= o; j-- {
                if dp[i-z][j-o]+1 > dp[i][j] {
                    dp[i][j] = dp[i-z][j-o] + 1
                }
            }
        }
    }
    return dp[m][n]
}
