package main

import (
    "container/heap"
)

// KthLargest is a struct that holds the value of k and a pointer to a min-heap.
type KthLargest struct {
    k       int       // The "k" in KthLargest, representing the position of the largest element.
    minHeap *IntHeap  // A min-heap to maintain the k largest elements.
}

// Constructor initializes the KthLargest struct.
// It takes an integer k and an initial array of integers nums.
func Constructor(k int, nums []int) KthLargest {
    // Create an empty min-heap.
    minHeap := &IntHeap{}
    // Initialize the min-heap.
    heap.Init(minHeap)
    // Initialize the KthLargest struct with k and the empty min-heap.
    kthLargest := KthLargest{k: k, minHeap: minHeap}
    // Add each number in nums to the KthLargest structure.
    for _, num := range nums {
        kthLargest.Add(num)
    }
    // Return the initialized KthLargest struct.
    return kthLargest
}

// Add inserts a new value into the min-heap and ensures that the size of the heap
// does not exceed k. It returns the k-th largest element.
func (this *KthLargest) Add(val int) int {
    // If the heap has fewer than k elements, push the new value onto the heap.
    if this.minHeap.Len() < this.k {
        heap.Push(this.minHeap, val)
    } else if val > (*this.minHeap)[0] {
        // If the new value is greater than the smallest value in the heap (min-heap root),
        // pop the smallest value and push the new value.
        heap.Pop(this.minHeap)
        heap.Push(this.minHeap, val)
    }
    // The root of the heap now contains the k-th largest element.
    return (*this.minHeap)[0]
}

// IntHeap is a type that implements heap.Interface and holds a slice of integers.
// It represents a min-heap.
type IntHeap []int

// Len returns the length of the heap.
func (h IntHeap) Len() int { return len(h) }

// Less compares two elements in the heap. The heap is a min-heap, so the comparison is "less than".
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }

// Swap exchanges the positions of two elements in the heap.
func (h IntHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

// Push adds a new element to the heap. It appends the element to the slice.
func (h *IntHeap) Push(x interface{}) {
    *h = append(*h, x.(int))
}

// Pop removes and returns the smallest element from the heap. 
// It removes the last element of the slice after swapping it with the root.
func (h *IntHeap) Pop() interface{} {
    old := *h           // Save the current slice.
    n := len(old)       // Get the length of the slice.
    x := old[n-1]       // Get the last element (smallest element in the heap).
    *h = old[0 : n-1]   // Resize the slice to remove the last element.
    return x            // Return the smallest element.
}
