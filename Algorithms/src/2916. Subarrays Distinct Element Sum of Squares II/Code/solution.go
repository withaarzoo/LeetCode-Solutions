func sumCounts(nums []int) int {
    const MOD = 1000000007
    n := len(nums)
    type Node struct {
        sum, sumsq, lazy int64
    }
    tree := make([]Node, 4*n)
    var push func(v, tl, tr int)
    push = func(v, tl, tr int) {
        if tree[v].lazy != 0 && tl != tr {
            mid := (tl + tr) >> 1
            llen := int64(mid - tl + 1)
            rlen := int64(tr - mid)
            add := tree[v].lazy
            tree[v<<1].sumsq = (tree[v<<1].sumsq + 2*add*tree[v<<1].sum + add*add%MOD*llen) % MOD
            tree[v<<1].sum = (tree[v<<1].sum + add*llen) % MOD
            tree[v<<1].lazy = (tree[v<<1].lazy + add) % MOD
            tree[v<<1|1].sumsq = (tree[v<<1|1].sumsq + 2*add*tree[v<<1|1].sum + add*add%MOD*rlen) % MOD
            tree[v<<1|1].sum = (tree[v<<1|1].sum + add*rlen) % MOD
            tree[v<<1|1].lazy = (tree[v<<1|1].lazy + add) % MOD
            tree[v].lazy = 0
        }
    }
    var update func(v, tl, tr, l, r int, add int64)
    update = func(v, tl, tr, l, r int, add int64) {
        if l > r {
            return
        }
        if l == tl && r == tr {
            length := int64(tr - tl + 1)
            tree[v].sumsq = (tree[v].sumsq + 2*add*tree[v].sum + add*add%MOD*length) % MOD
            tree[v].sum = (tree[v].sum + add*length) % MOD
            tree[v].lazy = (tree[v].lazy + add) % MOD
            return
        }
        push(v, tl, tr)
        mid := (tl + tr) >> 1
        update(v<<1, tl, mid, l, min(r, mid), add)
        update(v<<1|1, mid+1, tr, max(l, mid+1), r, add)
        tree[v].sum = (tree[v<<1].sum + tree[v<<1|1].sum) % MOD
        tree[v].sumsq = (tree[v<<1].sumsq + tree[v<<1|1].sumsq) % MOD
    }
    var query func(v, tl, tr, l, r int) int64
    query = func(v, tl, tr, l, r int) int64 {
        if l > r {
            return 0
        }
        if l == tl && r == tr {
            return tree[v].sumsq
        }
        push(v, tl, tr)
        mid := (tl + tr) >> 1
        return (query(v<<1, tl, mid, l, min(r, mid)) + query(v<<1|1, mid+1, tr, max(l, mid+1), r)) % MOD
    }
    last := make([]int, 100001)
    for i := range last {
        last[i] = -1
    }
    var ans int64
    for r := 0; r < n; r++ {
        prv := last[nums[r]]
        if prv+1 <= r-1 {
            update(1, 0, n-1, prv+1, r-1, 1)
        }
        update(1, 0, n-1, r, r, 1)
        ans = (ans + query(1, 0, n-1, 0, r)) % MOD
        last[nums[r]] = r
    }
    return int(ans)
}
func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}
func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}