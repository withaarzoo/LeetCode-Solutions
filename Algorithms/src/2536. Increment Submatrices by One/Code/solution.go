func rangeAddQueries(n int, queries [][]int) [][]int {
    // diff: (n+1) x (n+1)
    diff := make([][]int, n+1)
    for i := range diff {
        diff[i] = make([]int, n+1)
    }

    // Apply corner updates for each query
    for _, q := range queries {
        r1, c1, r2, c2 := q[0], q[1], q[2], q[3]
        diff[r1][c1] += 1
        diff[r1][c2+1] -= 1
        diff[r2+1][c1] -= 1
        diff[r2+1][c2+1] += 1
    }

    // Build result using 2D prefix sums
    res := make([][]int, n)
    for i := range res {
        res[i] = make([]int, n)
    }

    for i := 0; i < n; i++ {
        for j := 0; j < n; j++ {
            up := 0
            left := 0
            diag := 0
            if i > 0 {
                up = diff[i-1][j]
            }
            if j > 0 {
                left = diff[i][j-1]
            }
            if i > 0 && j > 0 {
                diag = diff[i-1][j-1]
            }
            diff[i][j] = diff[i][j] + up + left - diag
            res[i][j] = diff[i][j]
        }
    }
    return res
}
