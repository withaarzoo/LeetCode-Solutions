package main

import (
	"container/heap"
)

type DSU struct {
	p  []int
	sz []int
}

func NewDSU(n int) *DSU {
	p := make([]int, n+1)
	sz := make([]int, n+1)
	for i := 0; i <= n; i++ {
		p[i] = i
		sz[i] = 1
	}
	return &DSU{p, sz}
}
func (d *DSU) Find(x int) int {
	if d.p[x] != x {
		d.p[x] = d.Find(d.p[x])
	}
	return d.p[x]
}
func (d *DSU) Unite(a, b int) {
	ra, rb := d.Find(a), d.Find(b)
	if ra == rb {
		return
	}
	if d.sz[ra] < d.sz[rb] {
		ra, rb = rb, ra
	}
	d.p[rb] = ra
	d.sz[ra] += d.sz[rb]
}

// Min-heap of ints
type IntHeap []int

func (h IntHeap) Len() int            { return len(h) }
func (h IntHeap) Less(i, j int) bool  { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}
func (h IntHeap) Peek() int { return h[0] }

func processQueries(c int, connections [][]int, queries [][]int) []int {
	dsu := NewDSU(c)
	for _, e := range connections {
		dsu.Unite(e[0], e[1])
	}

	// root -> heap index in arrays
	heaps := make(map[int]*IntHeap, c*2)
	for i := 1; i <= c; i++ {
		r := dsu.Find(i)
		h, ok := heaps[r]
		if !ok {
			tmp := &IntHeap{}
			heaps[r] = tmp
		}
		heap.Push(heaps[r], i)
	}
	// heap.Init not required because we pushed in order; but to be safe:
	for _, h := range heaps {
		heap.Init(h)
	}

	offline := make([]bool, c+1)
	ans := make([]int, 0, len(queries))

	for _, q := range queries {
		t, x := q[0], q[1]
		if t == 2 {
			offline[x] = true
		} else {
			if !offline[x] {
				ans = append(ans, x)
			} else {
				r := dsu.Find(x)
				h := heaps[r]
				if h == nil {
					ans = append(ans, -1)
					continue
				}
				for h.Len() > 0 && offline[h.Peek()] {
					heap.Pop(h)
				}
				if h.Len() == 0 {
					ans = append(ans, -1)
				} else {
					ans = append(ans, h.Peek())
				}
			}
		}
	}
	return ans
}
