func minCost(n int, edges [][]int) int {
    graph := make([][][2]int, n)

    for _, e := range edges {
        u, v, w := e[0], e[1], e[2]
        graph[u] = append(graph[u], [2]int{v, w})
        graph[v] = append(graph[v], [2]int{u, 2 * w})
    }

    const INF = int(1e18)
    dist := make([]int, n)
    for i := range dist {
        dist[i] = INF
    }
    dist[0] = 0

    pq := [][]int{{0, 0}} // cost, node

    for len(pq) > 0 {
        // extract min
        minIdx := 0
        for i := range pq {
            if pq[i][0] < pq[minIdx][0] {
                minIdx = i
            }
        }
        cur := pq[minIdx]
        pq = append(pq[:minIdx], pq[minIdx+1:]...)

        cost, node := cur[0], cur[1]
        if cost > dist[node] {
            continue
        }

        for _, e := range graph[node] {
            next, w := e[0], e[1]
            if dist[next] > cost+w {
                dist[next] = cost + w
                pq = append(pq, []int{dist[next], next})
            }
        }
    }

    if dist[n-1] == INF {
        return -1
    }
    return dist[n-1]
}
