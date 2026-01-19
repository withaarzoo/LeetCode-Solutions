func maxSideLength(mat [][]int, threshold int) int {
    m, n := len(mat), len(mat[0])
    pre := make([][]int, m+1)

    for i := 0; i <= m; i++ {
        pre[i] = make([]int, n+1)
    }

    // Prefix sum
    for i := 1; i <= m; i++ {
        for j := 1; j <= n; j++ {
            pre[i][j] = mat[i-1][j-1] +
                pre[i-1][j] +
                pre[i][j-1] -
                pre[i-1][j-1]
        }
    }

    left, right := 0, min(m, n)
    ans := 0

    for left <= right {
        mid := (left + right) / 2
        found := false

        for i := mid; i <= m && !found; i++ {
            for j := mid; j <= n; j++ {
                sum := pre[i][j] -
                    pre[i-mid][j] -
                    pre[i][j-mid] +
                    pre[i-mid][j-mid]

                if sum <= threshold {
                    found = true
                    break
                }
            }
        }

        if found {
            ans = mid
            left = mid + 1
        } else {
            right = mid - 1
        }
    }

    return ans
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}
