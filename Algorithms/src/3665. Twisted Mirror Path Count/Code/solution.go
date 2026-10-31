func uniquePaths(grid [][]int) int {
    m, n := len(grid), len(grid[0])
    const MOD = 1000000007
    landR := make([][][2]int, m)
    landC := make([][][2]int, m)
    for i := range landR {
        landR[i] = make([][2]int, n)
        landC[i] = make([][2]int, n)
    }
    for x := m - 1; x >= 0; x-- {
        for y := n - 1; y >= 0; y-- {
            for d := 0; d < 2; d++ {
                if grid[x][y] == 0 {
                    landR[x][y][d] = x
                    landC[x][y][d] = y
                } else {
                    nx, ny, nd := x, y, 1-d
                    if d == 0 {
                        nx = x + 1
                    } else {
                        ny = y + 1
                    }
                    if nx < 0 || nx >= m || ny < 0 || ny >= n {
                        landR[x][y][d] = -1
                        landC[x][y][d] = -1
                    } else {
                        landR[x][y][d] = landR[nx][ny][nd]
                        landC[x][y][d] = landC[nx][ny][nd]
                    }
                }
            }
        }
    }
    dp := make([][]int, m)
    for i := range dp {
        dp[i] = make([]int, n)
    }
    dp[0][0] = 1
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if dp[i][j] == 0 {
                continue
            }
            if j+1 < n {
                nr := landR[i][j+1][0]
                nc := landC[i][j+1][0]
                if nr != -1 {
                    dp[nr][nc] = (dp[nr][nc] + dp[i][j]) % MOD
                }
            }
            if i+1 < m {
                nr := landR[i+1][j][1]
                nc := landC[i+1][j][1]
                if nr != -1 {
                    dp[nr][nc] = (dp[nr][nc] + dp[i][j]) % MOD
                }
            }
        }
    }
    return dp[m-1][n-1]
}