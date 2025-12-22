func minDeletionSize(strs []string) int {
    n := len(strs)
    m := len(strs[0])

    dp := make([]int, m)
    for i := 0; i < m; i++ {
        dp[i] = 1
    }

    for i := 0; i < m; i++ {
        for j := 0; j < i; j++ {
            valid := true
            for r := 0; r < n; r++ {
                if strs[r][j] > strs[r][i] {
                    valid = false
                    break
                }
            }
            if valid {
                if dp[j]+1 > dp[i] {
                    dp[i] = dp[j] + 1
                }
            }
        }
    }

    keep := 0
    for _, v := range dp {
        if v > keep {
            keep = v
        }
    }

    return m - keep
}
