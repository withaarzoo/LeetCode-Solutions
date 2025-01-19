import (
    "container/heap"
)

type Cell struct {
    height, x, y int
}

type MinHeap []Cell

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i].height < h[j].height }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
    *h = append(*h, x.(Cell))
}

func (h *MinHeap) Pop() interface{} {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[:n-1]
    return x
}

func trapRainWater(heightMap [][]int) int {
    m, n := len(heightMap), len(heightMap[0])
    if m < 3 || n < 3 {
        return 0
    }

    visited := make([][]bool, m)
    for i := range visited {
        visited[i] = make([]bool, n)
    }

    h := &MinHeap{}
    heap.Init(h)

    // Add all boundary cells
    for i := 0; i < m; i++ {
        heap.Push(h, Cell{heightMap[i][0], i, 0})
        heap.Push(h, Cell{heightMap[i][n-1], i, n-1})
        visited[i][0], visited[i][n-1] = true, true
    }
    for j := 0; j < n; j++ {
        heap.Push(h, Cell{heightMap[0][j], 0, j})
        heap.Push(h, Cell{heightMap[m-1][j], m-1, j})
        visited[0][j], visited[m-1][j] = true, true
    }

    result := 0
    directions := [][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

    for h.Len() > 0 {
        cell := heap.Pop(h).(Cell)

        for _, dir := range directions {
            nx, ny := cell.x+dir[0], cell.y+dir[1]
            if nx >= 0 && ny >= 0 && nx < m && ny < n && !visited[nx][ny] {
                result += max(0, cell.height-heightMap[nx][ny])
                heap.Push(h, Cell{max(cell.height, heightMap[nx][ny]), nx, ny})
                visited[nx][ny] = true
            }
        }
    }

    return result
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
