package main

func pacificAtlantic(heights [][]int) [][]int {
    if len(heights) == 0 || len(heights[0]) == 0 {
        return [][]int{}
    }
    m, n := len(heights), len(heights[0])
    pac := make([][]bool, m)
    atl := make([][]bool, m)
    for i := 0; i < m; i++ {
        pac[i] = make([]bool, n)
        atl[i] = make([]bool, n)
    }

    bfs := func(starts [][2]int, visited [][]bool) {
        head := 0
        queue := make([][2]int, len(starts))
        copy(queue, starts)
        for head < len(queue) {
            cur := queue[head]
            head++
            r, c := cur[0], cur[1]
            dirs := [][2]int{{1,0},{-1,0},{0,1},{0,-1}}
            for _, d := range dirs {
                nr, nc := r + d[0], c + d[1]
                if nr < 0 || nr >= m || nc < 0 || nc >= n {
                    continue
                }
                if visited[nr][nc] {
                    continue
                }
                if heights[nr][nc] < heights[r][c] {
                    continue
                }
                visited[nr][nc] = true
                queue = append(queue, [2]int{nr, nc})
            }
        }
    }

    // Pacific starts
    var pacStarts [][2]int
    for j := 0; j < n; j++ {
        pacStarts = append(pacStarts, [2]int{0, j})
        pac[0][j] = true
    }
    for i := 1; i < m; i++ {
        pacStarts = append(pacStarts, [2]int{i, 0})
        pac[i][0] = true
    }
    bfs(pacStarts, pac)

    // Atlantic starts
    var atlStarts [][2]int
    for j := 0; j < n; j++ {
        atlStarts = append(atlStarts, [2]int{m - 1, j})
        atl[m - 1][j] = true
    }
    for i := 0; i < m - 1; i++ {
        atlStarts = append(atlStarts, [2]int{i, n - 1})
        atl[i][n - 1] = true
    }
    bfs(atlStarts, atl)

    var res [][]int
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if pac[i][j] && atl[i][j] {
                res = append(res, []int{i, j})
            }
        }
    }
    return res
}
