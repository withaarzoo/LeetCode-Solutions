type item struct {
    d, u, cnt int
}

type hp []item

func (h hp) Len() int            { return len(h) }
func (h hp) Less(i, j int) bool  { return h[i].d < h[j].d }
func (h hp) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(x interface{}) { *h = append(*h, x.(item)) }
func (h *hp) Pop() interface{} {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[0 : n-1]
    return x
}

func shortestPath(n int, edges [][]int, labels string, k int) int {
    g := make([][][2]int, n)
    for _, e := range edges {
        g[e[0]] = append(g[e[0]], [2]int{e[1], e[2]})
    }
    const INF = 1000000000
    dist := make([][]int, n)
    for i := range dist {
        dist[i] = make([]int, k+1)
        for j := range dist[i] {
            dist[i][j] = INF
        }
    }
    pq := &hp{}
    heap.Init(pq)
    dist[0][1] = 0
    heap.Push(pq, item{0, 0, 1})
    for pq.Len() > 0 {
        cur := heap.Pop(pq).(item)
        d, u, cnt := cur.d, cur.u, cur.cnt
        if d > dist[u][cnt] {
            continue
        }
        if u == n-1 {
            return d
        }
        for _, e := range g[u] {
            v, w := e[0], e[1]
            ncnt := 1
            if labels[u] == labels[v] {
                ncnt = cnt + 1
            }
            if ncnt > k {
                continue
            }
            if d+w < dist[v][ncnt] {
                dist[v][ncnt] = d + w
                heap.Push(pq, item{dist[v][ncnt], v, ncnt})
            }
        }
    }
    return -1
}