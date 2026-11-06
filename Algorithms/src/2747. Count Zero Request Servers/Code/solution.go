func countServers(n int, logs [][]int, x int, queries []int) []int {
    sort.Slice(logs, func(i, j int) bool {
        return logs[i][1] < logs[j][1]
    })
    m := len(queries)
    type pair struct{ time, idx int }
    q := make([]pair, m)
    for i := 0; i < m; i++ {
        q[i] = pair{queries[i], i}
    }
    sort.Slice(q, func(i, j int) bool {
        return q[i].time < q[j].time
    })
    ans := make([]int, m)
    freq := make(map[int]int)
    active := 0
    left, right := 0, 0
    L := len(logs)
    for _, p := range q {
        time, idx := p.time, p.idx
        for right < L && logs[right][1] <= time {
            s := logs[right][0]
            freq[s]++
            if freq[s] == 1 {
                active++
            }
            right++
        }
        for left < right && logs[left][1] < time-x {
            s := logs[left][0]
            freq[s]--
            if freq[s] == 0 {
                active--
            }
            left++
        }
        ans[idx] = n - active
    }
    return ans
}