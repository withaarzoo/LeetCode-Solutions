func climbStairs(n int, costs []int) int {
    dp := make([]int64, n+1)
    for i := range dp {
        dp[i] = 1e18
    }
    dp[0] = 0
    for j := 1; j <= n; j++ {
        for k := 1; k <= 3; k++ {
            i := j - k
            if i >= 0 {
                jump := int64(j-i) * int64(j-i)
                if dp[i]+int64(costs[j-1])+jump < dp[j] {
                    dp[j] = dp[i] + int64(costs[j-1]) + jump
                }
            }
        }
    }
    return int(dp[n])
}