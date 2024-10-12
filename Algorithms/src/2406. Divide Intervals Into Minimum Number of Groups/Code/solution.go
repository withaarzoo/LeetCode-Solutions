import (
    "container/heap"
    "sort"
)

type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
    *h = append(*h, x.(int))
}

func (h *MinHeap) Pop() interface{} {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[0 : n-1]
    return x
}

func minGroups(intervals [][]int) int {
    // Sort intervals by start time
    sort.Slice(intervals, func(i, j int) bool {
        return intervals[i][0] < intervals[j][0]
    })
    
    // Min-heap to track the end times of groups
    pq := &MinHeap{}
    heap.Init(pq)
    
    // Traverse through all intervals
    for _, interval := range intervals {
        start, end := interval[0], interval[1]
        
        // If the earliest end time is less than the current start, reuse that group
        if pq.Len() > 0 && (*pq)[0] < start {
            heap.Pop(pq)
        }
        
        // Add the current interval's end time to the heap
        heap.Push(pq, end)
    }
    
    // The size of the heap is the number of groups
    return pq.Len()
}
