func resultArray(nums []int, k int, queries [][]int) []int {
    type Node struct {
        prod int
        freq []int
    }
    
    n := len(nums)
    tree := make([]Node, 4*n)
    for i := range tree {
        tree[i] = Node{prod: 1, freq: make([]int, k)}
    }
    
    merge := func(L, R Node) Node {
        res := Node{prod: 1, freq: make([]int, k)}
        res.prod = (L.prod * R.prod) % k
        copy(res.freq, L.freq)
        for r := 0; r < k; r++ {
            if R.freq[r] != 0 {
                nr := (L.prod * r) % k
                res.freq[nr] += R.freq[r]
            }
        }
        return res
    }
    
    var build func(v, tl, tr int)
    build = func(v, tl, tr int) {
        if tl == tr {
            tree[v].prod = nums[tl] % k
            tree[v].freq[tree[v].prod] = 1
            return
        }
        tm := (tl + tr) / 2
        build(v*2, tl, tm)
        build(v*2+1, tm+1, tr)
        tree[v] = merge(tree[v*2], tree[v*2+1])
    }
    
    var update func(v, tl, tr, pos, val int)
    update = func(v, tl, tr, pos, val int) {
        if tl == tr {
            tree[v].prod = val % k
            for i := range tree[v].freq {
                tree[v].freq[i] = 0
            }
            tree[v].freq[tree[v].prod] = 1
            return
        }
        tm := (tl + tr) / 2
        if pos <= tm {
            update(v*2, tl, tm, pos, val)
        } else {
            update(v*2+1, tm+1, tr, pos, val)
        }
        tree[v] = merge(tree[v*2], tree[v*2+1])
    }
    
    var query func(v, tl, tr, l, r int) Node
    query = func(v, tl, tr, l, r int) Node {
        if l > r {
            return Node{prod: 1, freq: make([]int, k)}
        }
        if l == tl && r == tr {
            return tree[v]
        }
        tm := (tl + tr) / 2
        return merge(query(v*2, tl, tm, l, min(r, tm)),
                     query(v*2+1, tm+1, tr, max(l, tm+1), r))
    }
    
    build(1, 0, n-1)
    ans := make([]int, 0, len(queries))
    for _, q := range queries {
        idx, val, start, x := q[0], q[1], q[2], q[3]
        update(1, 0, n-1, idx, val)
        res := query(1, 0, n-1, start, n-1)
        ans = append(ans, res.freq[x])
    }
    return ans
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