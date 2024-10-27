func countSquares(matrix [][]int) int {
    m := len(matrix)
    n := len(matrix[0])
    dp := make([][]int, m)
    for i := range dp {
        dp[i] = make([]int, n)
    }
    totalSquares := 0

    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if matrix[i][j] == 1 {
                if i == 0 || j == 0 {
                    dp[i][j] = 1
                } else {
                    dp[i][j] = min(dp[i-1][j], dp[i][j-1], dp[i-1][j-1]) + 1
                }
                totalSquares += dp[i][j]
            }
        }
    }

    return totalSquares
}

func min(a, b, c int) int {
    if a < b && a < c {
        return a
    }
    if b < c {
        return b
    }
    return c
}