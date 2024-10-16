import (
    "container/heap"
    "strings"
)

type CharCount struct {
    count int
    ch    byte
}

// Priority queue implementation for Go
type MaxHeap []CharCount

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i].count > h[j].count }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MaxHeap) Push(x interface{}) {
    *h = append(*h, x.(CharCount))
}

func (h *MaxHeap) Pop() interface{} {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[0 : n-1]
    return x
}

func longestDiverseString(a int, b int, c int) string {
    pq := &MaxHeap{}
    heap.Init(pq)

    if a > 0 {
        heap.Push(pq, CharCount{a, 'a'})
    }
    if b > 0 {
        heap.Push(pq, CharCount{b, 'b'})
    }
    if c > 0 {
        heap.Push(pq, CharCount{c, 'c'})
    }

    var result strings.Builder

    for pq.Len() > 0 {
        first := heap.Pop(pq).(CharCount)

        if result.Len() >= 2 && result.String()[result.Len()-1] == first.ch && result.String()[result.Len()-2] == first.ch {
            if pq.Len() == 0 {
                break
            }
            second := heap.Pop(pq).(CharCount)
            result.WriteByte(second.ch)
            second.count--
            if second.count > 0 {
                heap.Push(pq, second)
            }
            heap.Push(pq, first)
        } else {
            result.WriteByte(first.ch)
            first.count--
            if first.count > 0 {
                heap.Push(pq, first)
            }
        }
    }

    return result.String()
}