import (
    "container/heap"
    "math"
)

type Element struct {
    value, row, col int
}

type MinHeap []Element

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i].value < h[j].value }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
    *h = append(*h, x.(Element))
}

func (h *MinHeap) Pop() interface{} {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[0 : n-1]
    return x
}

func smallestRange(nums [][]int) []int {
    minHeap := &MinHeap{}
    maxValue := math.MinInt32
    
    // Initialize heap with first element from each list
    for i := range nums {
        heap.Push(minHeap, Element{nums[i][0], i, 0})
        if nums[i][0] > maxValue {
            maxValue = nums[i][0]
        }
    }
    
    rangeStart, rangeEnd := 0, math.MaxInt32
    
    for minHeap.Len() > 0 {
        minElement := heap.Pop(minHeap).(Element)
        
        // Update the smallest range
        if maxValue - minElement.value < rangeEnd - rangeStart {
            rangeStart = minElement.value
            rangeEnd = maxValue
        }
        
        // Move to the next element in the current list
        if minElement.col + 1 < len(nums[minElement.row]) {
            nextValue := nums[minElement.row][minElement.col + 1]
            heap.Push(minHeap, Element{nextValue, minElement.row, minElement.col + 1})
            if nextValue > maxValue {
                maxValue = nextValue
            }
        } else {
            break // One list is exhausted
        }
    }
    
    return []int{rangeStart, rangeEnd}
}