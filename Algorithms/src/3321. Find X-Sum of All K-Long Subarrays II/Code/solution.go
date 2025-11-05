package main

import "container/heap"

type item struct {
	val   int
	freq  int
	idx   int
	inTop bool
}

// min-heap for TOP (worst at top): (freq asc, value asc)
type hotHeap []*item
func (h hotHeap) Len() int { return len(h) }
func (h hotHeap) Less(i, j int) bool {
	if h[i].freq != h[j].freq { return h[i].freq < h[j].freq }
	return h[i].val < h[j].val
}
func (h hotHeap) Swap(i, j int){ h[i], h[j] = h[j], h[i]; h[i].idx, h[j].idx = i, j }
func (h *hotHeap) Push(x interface{}) { it := x.(*item); it.idx = len(*h); *h = append(*h, it) }
func (h *hotHeap) Pop() interface{} { old := *h; it := old[len(old)-1]; *h = old[:len(old)-1]; return it }

// max-heap for REST (best at top): (freq desc, value desc)
type restHeap []*item
func (h restHeap) Len() int { return len(h) }
func (h restHeap) Less(i, j int) bool {
	if h[i].freq != h[j].freq { return h[i].freq > h[j].freq }
	return h[i].val > h[j].val
}
func (h restHeap) Swap(i, j int){ h[i], h[j] = h[j], h[i]; h[i].idx, h[j].idx = i, j }
func (h *restHeap) Push(x interface{}) { it := x.(*item); it.idx = len(*h); *h = append(*h, it) }
func (h *restHeap) Pop() interface{} { old := *h; it := old[len(old)-1]; *h = old[:len(old)-1]; return it }

func findXSum(nums []int, k int, x int) []int64 {
	n := len(nums)
	ans := make([]int64, n-k+1)

	freq := map[int]*item{}
	hot := &hotHeap{}   // TOP
	rest := &restHeap{} // REST
	heap.Init(hot); heap.Init(rest)

	var sum int64 = 0

	removeFromCurrent := func(it *item){
		if it.inTop {
			sum -= int64(it.val) * int64(it.freq)
			heap.Remove(hot, it.idx)
			it.inTop = false
		} else if it.freq > 0 && it.idx >= 0 && it.idx < rest.Len() && (*rest)[it.idx] == it {
			heap.Remove(rest, it.idx)
		}
	}

	promoteIfPossible := func(){
		for hot.Len() < x && rest.Len() > 0 {
			best := heap.Pop(rest).(*item)
			best.inTop = true
			sum += int64(best.val) * int64(best.freq)
			heap.Push(hot, best)
		}
	}

	insertVal := func(v int){
		it, ok := freq[v]
		if !ok {
			it = &item{val: v, idx: -1}
			freq[v] = it
		} else {
			removeFromCurrent(it)
		}
		it.freq++
		it.inTop = true
		sum += int64(it.val) * int64(it.freq)
		heap.Push(hot, it)
		if hot.Len() > x {
			worst := heap.Pop(hot).(*item)
			sum -= int64(worst.val) * int64(worst.freq)
			worst.inTop = false
			heap.Push(rest, worst)
		}
	}

	eraseVal := func(v int){
		it := freq[v]
		removeFromCurrent(it)
		it.freq--
		if it.freq == 0 {
			delete(freq, v)
			it.idx, it.inTop = -1, false
		} else {
			heap.Push(rest, it)
		}
		promoteIfPossible()
	}

	for i := 0; i < k; i++ { insertVal(nums[i]) }
	ans[0] = sum
	for i := k; i < n; i++ {
		eraseVal(nums[i-k])
		insertVal(nums[i])
		ans[i-k+1] = sum
	}
	return ans
}
