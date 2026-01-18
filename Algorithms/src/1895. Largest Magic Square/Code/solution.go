func largestMagicSquare(grid [][]int) int {
    m := len(grid)
    n := len(grid[0])

    row := make([][]int, m)
    col := make([][]int, m+1)

    for i := range row {
        row[i] = make([]int, n+1)
    }
    for i := range col {
        col[i] = make([]int, n)
    }

    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            row[i][j+1] = row[i][j] + grid[i][j]
            col[i+1][j] = col[i][j] + grid[i][j]
        }
    }

    for k := min(m, n); k >= 2; k-- {
        for i := 0; i+k <= m; i++ {
            for j := 0; j+k <= n; j++ {

                target := row[i][j+k] - row[i][j]
                ok := true

                for r := i; r < i+k; r++ {
                    if row[r][j+k]-row[r][j] != target {
                        ok = false
                    }
                }

                for c := j; c < j+k; c++ {
                    if col[i+k][c]-col[i][c] != target {
                        ok = false
                    }
                }

                d1, d2 := 0, 0
                for x := 0; x < k; x++ {
                    d1 += grid[i+x][j+x]
                    d2 += grid[i+x][j+k-1-x]
                }

                if ok && d1 == target && d2 == target {
                    return k
                }
            }
        }
    }
    return 1
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}
