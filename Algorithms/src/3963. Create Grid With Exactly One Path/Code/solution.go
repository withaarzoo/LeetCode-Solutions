func createGrid(m int, n int) []string {
    grid := make([][]byte, m)
    for i := 0; i < m; i++ {
        grid[i] = make([]byte, n)
        for j := 0; j < n; j++ {
            grid[i][j] = '#'
        }
    }
    for j := 0; j < n; j++ {
        grid[0][j] = '.'
    }
    for i := 0; i < m; i++ {
        grid[i][n-1] = '.'
    }
    res := make([]string, m)
    for i := 0; i < m; i++ {
        res[i] = string(grid[i])
    }
    return res
}