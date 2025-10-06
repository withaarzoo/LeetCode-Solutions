package main

import (
    "container/heap"
)

type Cell struct {
    t, r, c int
}

// implement heap for Cell (min-heap by t)
type CellHeap []Cell

func (h CellHeap) Len() int { return len(h) }
func (h CellHeap) Less(i, j int) bool { return h[i].t < h[j].t }
func (h CellHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h *CellHeap) Push(x interface{}) { *h = append(*h, x.(Cell)) }
func (h *CellHeap) Pop() interface{} {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[:n-1]
    return x
}

func swimInWater(grid [][]int) int {
    n := len(grid)
    visited := make([][]bool, n)
    for i := 0; i < n; i++ { visited[i] = make([]bool, n) }
    dirs := [4][2]int{{1,0},{-1,0},{0,1},{0,-1}}

    h := &CellHeap{}
    heap.Init(h)
    heap.Push(h, Cell{grid[0][0], 0, 0})

    for h.Len() > 0 {
        cell := heap.Pop(h).(Cell)
        t, r, c := cell.t, cell.r, cell.c
        if visited[r][c] { continue }
        visited[r][c] = true
        if r == n-1 && c == n-1 { return t }
        for _, d := range dirs {
            nr, nc := r + d[0], c + d[1]
            if nr >= 0 && nr < n && nc >= 0 && nc < n && !visited[nr][nc] {
                nt := t
                if grid[nr][nc] > nt { nt = grid[nr][nc] }
                heap.Push(h, Cell{nt, nr, nc})
            }
        }
    }
    return -1
}
