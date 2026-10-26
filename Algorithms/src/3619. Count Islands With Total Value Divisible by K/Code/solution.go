func countIslands(grid [][]int, k int) int {
    m, n := len(grid), len(grid[0])
    ans := 0
    var dfs func(i, j int) int64
    dfs = func(i, j int) int64 {
        if i < 0 || i >= m || j < 0 || j >= n || grid[i][j] == 0 {
            return 0
        }
        sum := int64(grid[i][j])
        grid[i][j] = 0
        sum += dfs(i+1, j)
        sum += dfs(i-1, j)
        sum += dfs(i, j+1)
        sum += dfs(i, j-1)
        return sum
    }
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if grid[i][j] > 0 {
                if dfs(i, j)%int64(k) == 0 {
                    ans++
                }
            }
        }
    }
    return ans
}