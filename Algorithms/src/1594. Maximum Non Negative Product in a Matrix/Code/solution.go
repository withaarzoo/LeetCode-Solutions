func maxProductPath(grid [][]int) int {
    m, n := len(grid), len(grid[0])
    MOD := int64(1e9 + 7)

    maxDp := make([][]int64, m)
    minDp := make([][]int64, m)

    for i := range maxDp {
        maxDp[i] = make([]int64, n)
        minDp[i] = make([]int64, n)
    }

    maxDp[0][0] = int64(grid[0][0])
    minDp[0][0] = int64(grid[0][0])

    for i := 1; i < m; i++ {
        maxDp[i][0] = maxDp[i-1][0] * int64(grid[i][0])
        minDp[i][0] = maxDp[i][0]
    }

    for j := 1; j < n; j++ {
        maxDp[0][j] = maxDp[0][j-1] * int64(grid[0][j])
        minDp[0][j] = maxDp[0][j]
    }

    for i := 1; i < m; i++ {
        for j := 1; j < n; j++ {
            val := int64(grid[i][j])

            a := maxDp[i-1][j] * val
            b := minDp[i-1][j] * val
            c := maxDp[i][j-1] * val
            d := minDp[i][j-1] * val

            maxDp[i][j] = max4(a, b, c, d)
            minDp[i][j] = min4(a, b, c, d)
        }
    }

    res := maxDp[m-1][n-1]
    if res < 0 {
        return -1
    }
    return int(res % MOD)
}

func max4(a, b, c, d int64) int64 {
    return max(max(a, b), max(c, d))
}

func min4(a, b, c, d int64) int64 {
    return min(min(a, b), min(c, d))
}

func max(a, b int64) int64 {
    if a > b {
        return a
    }
    return b
}

func min(a, b int64) int64 {
    if a < b {
        return a
    }
    return b
}