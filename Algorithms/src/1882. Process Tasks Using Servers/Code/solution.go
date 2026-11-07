func assignTasks(servers []int, tasks []int) []int {
    idle := hp{}
    for i, x := range servers {
        heap.Push(&idle, pair{x, i})
    }
    busy := hp2{}
    ans := make([]int, len(tasks))
    for j, t := range tasks {
        for len(busy) > 0 && busy[0].ft <= int64(j) {
            p := heap.Pop(&busy).(bpair)
            heap.Push(&idle, pair{p.w, p.i})
        }
        if idle.Len() > 0 {
            p := heap.Pop(&idle).(pair)
            ans[j] = p.i
            heap.Push(&busy, bpair{int64(j) + int64(t), p.w, p.i})
        } else {
            p := heap.Pop(&busy).(bpair)
            ans[j] = p.i
            heap.Push(&busy, bpair{p.ft + int64(t), p.w, p.i})
        }
    }
    return ans
}

type pair struct{ w, i int }
type hp []pair
func (h hp) Len() int            { return len(h) }
func (h hp) Less(i, j int) bool  { return h[i].w < h[j].w || (h[i].w == h[j].w && h[i].i < h[j].i) }
func (h hp) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v interface{}) { *h = append(*h, v.(pair)) }
func (h *hp) Pop() interface{}   { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }

type bpair struct {
    ft int64
    w, i int
}
type hp2 []bpair
func (h hp2) Len() int { return len(h) }
func (h hp2) Less(i, j int) bool {
    if h[i].ft != h[j].ft {
        return h[i].ft < h[j].ft
    }
    if h[i].w != h[j].w {
        return h[i].w < h[j].w
    }
    return h[i].i < h[j].i
}
func (h hp2) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *hp2) Push(v interface{}) { *h = append(*h, v.(bpair)) }
func (h *hp2) Pop() interface{}   { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }