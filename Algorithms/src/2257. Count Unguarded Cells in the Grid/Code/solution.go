package main

// Note: LeetCode's Go function signature is usually inside package main
// This function uses the signature requested by the user.
func countUnguarded(m int, n int, guards [][]int, walls [][]int) int {
    // 0 = empty, 1 = guard, 2 = wall, 3 = guarded
    grid := make([][]int, m)
    for i := 0; i < m; i++ {
        grid[i] = make([]int, n)
    }
    for _, w := range walls {
        grid[w[0]][w[1]] = 2
    }
    for _, g := range guards {
        grid[g[0]][g[1]] = 1
    }

    dirs := [4][2]int{{-1,0},{1,0},{0,-1},{0,1}}
    for _, g := range guards {
        r, c := g[0], g[1]
        for _, d := range dirs {
            nr, nc := r + d[0], c + d[1]
            for nr >= 0 && nr < m && nc >= 0 && nc < n {
                if grid[nr][nc] == 2 || grid[nr][nc] == 1 {
                    break
                }
                if grid[nr][nc] == 0 {
                    grid[nr][nc] = 3
                }
                nr += d[0]
                nc += d[1]
            }
        }
    }

    ans := 0
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if grid[i][j] == 0 {
                ans++
            }
        }
    }
    return ans
}
