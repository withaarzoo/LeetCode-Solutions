func busiestServers(k int, arrival []int, load []int) []int {
    available := &IntHeap{}
    heap.Init(available)
    for i := 0; i < k; i++ {
        heap.Push(available, i)
    }
    busy := &PairHeap{}
    heap.Init(busy)
    cnt := make([]int, k)
    n := len(arrival)
    for i := 0; i < n; i++ {
        t := arrival[i]
        for busy.Len() > 0 && (*busy)[0].end <= t {
            x := heap.Pop(busy).(Pair).server
            heap.Push(available, i+((x-i)%k+k)%k)
        }
        if available.Len() == 0 {
            continue
        }
        j := heap.Pop(available).(int) % k
        cnt[j]++
        heap.Push(busy, Pair{t + load[i], j})
    }
    mx := 0
    for _, c := range cnt {
        if c > mx {
            mx = c
        }
    }
    ans := []int{}
    for i := 0; i < k; i++ {
        if cnt[i] == mx {
            ans = append(ans, i)
        }
    }
    return ans
}

type IntHeap []int
func (h IntHeap) Len() int            { return len(h) }
func (h IntHeap) Less(i, j int) bool  { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *IntHeap) Pop() interface{} {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[0 : n-1]
    return x
}

type Pair struct{ end, server int }
type PairHeap []Pair
func (h PairHeap) Len() int            { return len(h) }
func (h PairHeap) Less(i, j int) bool  { return h[i].end < h[j].end }
func (h PairHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *PairHeap) Push(x interface{}) { *h = append(*h, x.(Pair)) }
func (h *PairHeap) Pop() interface{} {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[0 : n-1]
    return x
}