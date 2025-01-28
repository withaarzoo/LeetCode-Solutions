func findMaxFish(grid [][]int) int {
    m, n := len(grid), len(grid[0])
    visited := make([][]bool, m)
    for i := range visited {
        visited[i] = make([]bool, n)
    }
    maxFish := 0

    var dfs func(r, c int) int
    dfs = func(r, c int) int {
        if r < 0 || c < 0 || r >= m || c >= n || visited[r][c] || grid[r][c] == 0 {
            return 0
        }
        visited[r][c] = true
        fish := grid[r][c]
        fish += dfs(r+1, c)
        fish += dfs(r-1, c)
        fish += dfs(r, c+1)
        fish += dfs(r, c-1)
        return fish
    }

    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if !visited[i][j] && grid[i][j] > 0 {
                maxFish = max(maxFish, dfs(i, j))
            }
        }
    }
    return maxFish
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
