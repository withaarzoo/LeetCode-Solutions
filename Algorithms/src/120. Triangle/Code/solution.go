func minimumTotal(triangle [][]int) int {
    n := len(triangle)
    // initialize dp with last row
    dp := make([]int, n)
    copy(dp, triangle[n-1])

    // bottom-up updates
    for i := n - 2; i >= 0; i-- {
        for j := 0; j <= i; j++ {
            if dp[j] < dp[j+1] {
                dp[j] = triangle[i][j] + dp[j]
            } else {
                dp[j] = triangle[i][j] + dp[j+1]
            }
        }
    }
    return dp[0]
}
