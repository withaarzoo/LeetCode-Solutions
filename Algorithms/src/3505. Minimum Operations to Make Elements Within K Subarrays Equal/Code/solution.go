func minOperations(nums []int, x int, k int) int64 {
    n := len(nums)
    vals := make([]int, n)
    copy(vals, nums)
    sort.Ints(vals)
    m := 0
    for i := 0; i < n; i++ {
        if m == 0 || vals[i] != vals[m-1] {
            vals[m] = vals[i]
            m++
        }
    }
    vals = vals[:m]
    rk := make([]int, n)
    for i, v := range nums {
        rk[i] = sort.SearchInts(vals, v) + 1
    }
    cnt := make([]int, m+1)
    sm := make([]int, m+1)
    top := 1
    for top*2 <= m {
        top *= 2
    }
    need := (x + 1) / 2
    w := n - x + 1
    cost := make([]int, w)
    total := 0
    for i := 0; i < n; i++ {
        total += nums[i]
        for p := rk[i]; p <= m; p += p & -p {
            cnt[p]++
            sm[p] += nums[i]
        }
        if i >= x {
            total -= nums[i-x]
            for p := rk[i-x]; p <= m; p += p & -p {
                cnt[p]--
                sm[p] -= nums[i-x]
            }
        }
        if i >= x-1 {
            pos, rem, s0 := 0, need, 0
            for pw := top; pw > 0; pw >>= 1 {
                nxt := pos + pw
                if nxt <= m && cnt[nxt] < rem {
                    pos = nxt
                    rem -= cnt[nxt]
                    s0 += sm[nxt]
                }
            }
            med := vals[pos]
            c0 := need - rem
            cost[i-x+1] = med*c0 - s0 + (total - s0) - med*(x-c0)
        }
    }
    const inf = 1 << 60
    prev := make([]int, n+1)
    cur := make([]int, n+1)
    for j := 0; j < k; j++ {
        for i := 0; i < x; i++ {
            cur[i] = inf
        }
        for i := x; i <= n; i++ {
            a := cur[i-1]
            b := prev[i-x] + cost[i-x]
            if a < b {
                cur[i] = a
            } else {
                cur[i] = b
            }
        }
        prev, cur = cur, prev
    }
    return int64(prev[n])
}