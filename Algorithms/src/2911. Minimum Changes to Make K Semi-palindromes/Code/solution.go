func minimumChanges(s string, k int) int {
    n := len(s)
    const INF = 1000000000
    cost := make([][]int, n)
    for i := range cost {
        cost[i] = make([]int, n)
        for j := range cost[i] {
            cost[i][j] = INF
        }
    }
    for i := 0; i < n; i++ {
        for j := i + 1; j < n; j++ {
            L := j - i + 1
            best := INF
            for d := 1; d < L; d++ {
                if L%d != 0 {
                    continue
                }
                cur := 0
                m := L / d
                for r := 0; r < d; r++ {
                    for t := 0; t < m/2; t++ {
                        if s[i+r+t*d] != s[i+r+(m-1-t)*d] {
                            cur++
                        }
                    }
                }
                if cur < best {
                    best = cur
                }
            }
            cost[i][j] = best
        }
    }
    dp := make([][]int, n+1)
    for i := range dp {
        dp[i] = make([]int, k+1)
        for p := range dp[i] {
            dp[i][p] = INF
        }
    }
    dp[0][0] = 0
    for i := 1; i <= n; i++ {
        for p := 1; p <= k; p++ {
            for prev := 0; prev <= i-2; prev++ {
                if dp[prev][p-1] != INF && cost[prev][i-1] != INF {
                    if dp[prev][p-1]+cost[prev][i-1] < dp[i][p] {
                        dp[i][p] = dp[prev][p-1] + cost[prev][i-1]
                    }
                }
            }
        }
    }
    return dp[n][k]
}