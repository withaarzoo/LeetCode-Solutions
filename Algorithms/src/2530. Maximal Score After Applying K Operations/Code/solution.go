import (
    "container/heap"
    "math"
)

type MaxHeap []int64

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] } // Max-heap property
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MaxHeap) Push(x interface{}) {
    *h = append(*h, x.(int64))
}

func (h *MaxHeap) Pop() interface{} {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[0 : n-1]
    return x
}

func maxKelements(nums []int, k int) int64 {
    maxHeap := &MaxHeap{}
    heap.Init(maxHeap)
    
    for _, num := range nums {
        heap.Push(maxHeap, int64(num))
    }
    
    var score int64 = 0
    
    for i := 0; i < k; i++ {
        maxVal := heap.Pop(maxHeap).(int64)
        score += maxVal
        heap.Push(maxHeap, int64(math.Ceil(float64(maxVal) / 3)))
    }
    
    return score
}