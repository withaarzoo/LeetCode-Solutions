func highestPeak(isWater [][]int) [][]int {
    m, n := len(isWater), len(isWater[0])
    height := make([][]int, m)
    for i := range height {
        height[i] = make([]int, n)
        for j := range height[i] {
            height[i][j] = -1
        }
    }

    queue := [][]int{}
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if isWater[i][j] == 1 {
                height[i][j] = 0
                queue = append(queue, []int{i, j})
            }
        }
    }

    directions := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

    for len(queue) > 0 {
        cell := queue[0]
        queue = queue[1:]
        x, y := cell[0], cell[1]
        for _, dir := range directions {
            nx, ny := x+dir[0], y+dir[1]
            if nx >= 0 && ny >= 0 && nx < m && ny < n && height[nx][ny] == -1 {
                height[nx][ny] = height[x][y] + 1
                queue = append(queue, []int{nx, ny})
            }
        }
    }

    return height
}
