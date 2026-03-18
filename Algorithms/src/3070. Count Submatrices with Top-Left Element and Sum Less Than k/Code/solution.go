func countSubmatrices(grid [][]int, k int) int {
    m, n := len(grid), len(grid[0])
    count := 0

    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {

            if i > 0 {
                grid[i][j] += grid[i-1][j]
            }
            if j > 0 {
                grid[i][j] += grid[i][j-1]
            }
            if i > 0 && j > 0 {
                grid[i][j] -= grid[i-1][j-1]
            }

            if grid[i][j] <= k {
                count++
            }
        }
    }

    return count
}