func minScoreTriangulation(values []int) int {
    n := len(values)
    if n < 3 {
        return 0
    }
    dp := make([][]int, n)
    for i := 0; i < n; i++ {
        dp[i] = make([]int, n)
    }
    // Max int
    maxInt := int(^uint(0) >> 1)
    
    for length := 3; length <= n; length++ {
        for i := 0; i+length-1 < n; i++ {
            j := i + length - 1
            best := maxInt
            for k := i + 1; k < j; k++ {
                cost := dp[i][k] + dp[k][j] + values[i]*values[k]*values[j]
                if cost < best {
                    best = cost
                }
            }
            dp[i][j] = best
        }
    }
    return dp[0][n-1]
}
