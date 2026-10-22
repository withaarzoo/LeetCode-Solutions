func minOperations(nums []int) int {
    n := len(nums)
    starts := make([][2]int, 0, 2)
    inc, dec := true, true
    for i := 0; i < n; i++ {
        if nums[(i+1)%n] != (nums[i]+1)%n {
            inc = false
        }
        if nums[(i+1)%n] != (nums[i]-1+n)%n {
            dec = false
        }
    }
    if inc {
        starts = append(starts, [2]int{0, nums[0]})
    }
    if dec {
        starts = append(starts, [2]int{1, (n - 1 - nums[0] + n) % n})
    }
    if len(starts) == 0 {
        return -1
    }

    dist := make([][]int, 2)
    for i := range dist {
        dist[i] = make([]int, n)
        for j := range dist[i] {
            dist[i][j] = -1
        }
    }
    type state struct{ t, s int }
    q := make([]state, 0, 2)
    for _, p := range starts {
        dist[p[0]][p[1]] = 0
        q = append(q, state{p[0], p[1]})
    }
    for len(q) > 0 {
        cur := q[0]
        q = q[1:]
        t, s := cur.t, cur.s
        if t == 0 && s == 0 {
            return dist[t][s]
        }
        d := dist[t][s]
        nt, ns := t, (s+1)%n
        if dist[nt][ns] == -1 {
            dist[nt][ns] = d + 1
            q = append(q, state{nt, ns})
        }
        nt, ns = 1-t, (n-s)%n
        if dist[nt][ns] == -1 {
            dist[nt][ns] = d + 1
            q = append(q, state{nt, ns})
        }
    }
    return -1
}