func minCost(grid [][]int) int {
    m, n := len(grid), len(grid[0])
    directions := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
    cost := make([][]int, m)
    for i := range cost {
        cost[i] = make([]int, n)
        for j := range cost[i] {
            cost[i][j] = 1 << 30 // Max int
        }
    }
    cost[0][0] = 0
    dq := [][]int{{0, 0}}
    
    for len(dq) > 0 {
        x, y := dq[0][0], dq[0][1]
        dq = dq[1:]
        
        for i, d := range directions {
            nx, ny := x+d[0], y+d[1]
            newCost := cost[x][y]
            if grid[x][y] != i+1 {
                newCost++
            }
            
            if nx >= 0 && ny >= 0 && nx < m && ny < n && newCost < cost[nx][ny] {
                cost[nx][ny] = newCost
                if grid[x][y] == i+1 {
                    dq = append([][]int{{nx, ny}}, dq...)
                } else {
                    dq = append(dq, []int{nx, ny})
                }
            }
        }
    }
    return cost[m-1][n-1]
}
