func maximumProfit(prices []int, k int) int64 {
    n := len(prices)
    const NEG int64 = -1e18

    // dp[i][state][t]
    dp := make([][][]int64, n+1)
    for i := 0; i <= n; i++ {
        dp[i] = make([][]int64, 3)
        for s := 0; s < 3; s++ {
            dp[i][s] = make([]int64, k+1)
            for t := 0; t <= k; t++ {
                dp[i][s][t] = NEG
            }
        }
    }

    // Base case: end of days
    for t := 0; t <= k; t++ {
        dp[n][0][t] = 0        // free is valid
        dp[n][1][t] = NEG     // holding long is invalid
        dp[n][2][t] = NEG     // holding short is invalid
    }

    for i := n - 1; i >= 0; i-- {
        for t := 0; t <= k; t++ {

            // state 0: free
            dp[i][0][t] = dp[i+1][0][t]
            dp[i][0][t] = max(dp[i][0][t], -int64(prices[i])+dp[i+1][1][t])
            dp[i][0][t] = max(dp[i][0][t], int64(prices[i])+dp[i+1][2][t])

            if t < k {
                // state 1: holding long
                dp[i][1][t] = max(
                    dp[i+1][1][t],
                    int64(prices[i])+dp[i+1][0][t+1],
                )

                // state 2: holding short
                dp[i][2][t] = max(
                    dp[i+1][2][t],
                    -int64(prices[i])+dp[i+1][0][t+1],
                )
            }
        }
    }

    return dp[0][0][0]
}

func max(a, b int64) int64 {
    if a > b {
        return a
    }
    return b
}
