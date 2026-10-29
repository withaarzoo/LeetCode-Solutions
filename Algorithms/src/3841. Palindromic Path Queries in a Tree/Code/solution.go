func palindromePath(n int, edges [][]int, s string, queries []string) []bool {
    g := make([][]int, n)
    for _, e := range edges {
        g[e[0]] = append(g[e[0]], e[1])
        g[e[1]] = append(g[e[1]], e[0])
    }
    const LOG = 16
    up := make([][]int, LOG)
    for i := range up {
        up[i] = make([]int, n)
        for j := range up[i] {
            up[i][j] = -1
        }
    }
    depth := make([]int, n)
    inn := make([]int, n)
    out := make([]int, n)
    parent := make([]int, n)
    for i := range parent {
        parent[i] = -1
    }
    mask := make([]int, n)
    for i := 0; i < n; i++ {
        mask[i] = 1 << (s[i] - 'a')
    }
    timer := 1
    var dfs func(int, int)
    dfs = func(u, p int) {
        inn[u] = timer
        timer++
        parent[u] = p
        up[0][u] = p
        for k := 1; k < LOG; k++ {
            if up[k-1][u] != -1 {
                up[k][u] = up[k-1][up[k-1][u]]
            }
        }
        for _, v := range g[u] {
            if v != p {
                depth[v] = depth[u] + 1
                dfs(v, u)
            }
        }
        out[u] = timer
    }
    dfs(0, -1)
    bit := make([]int, n+2)
    upd := func(i, val int) {
        for ; i <= n; i += i & -i {
            bit[i] ^= val
        }
    }
    qry := func(i int) int {
        res := 0
        for ; i > 0; i -= i & -i {
            res ^= bit[i]
        }
        return res
    }
    rangeXor := func(l, r, val int) {
        if l > r {
            return
        }
        upd(l, val)
        upd(r+1, val)
    }
    for i := 0; i < n; i++ {
        rangeXor(inn[i], out[i]-1, mask[i])
    }
    getLca := func(u, v int) int {
        if depth[u] < depth[v] {
            u, v = v, u
        }
        diff := depth[u] - depth[v]
        for k := 0; k < LOG; k++ {
            if (diff>>k)&1 == 1 {
                u = up[k][u]
            }
        }
        if u == v {
            return u
        }
        for k := LOG - 1; k >= 0; k-- {
            if up[k][u] != up[k][v] {
                u = up[k][u]
                v = up[k][v]
            }
        }
        return parent[u]
    }
    pathXor := func(u, v int) int {
        l := getLca(u, v)
        return qry(inn[u]) ^ qry(inn[v]) ^ mask[l]
    }
    ans := []bool{}
    for _, q := range queries {
        if q[0] == 'u' {
            u := 0
            i := 7
            for q[i] != ' ' {
                u = u*10 + int(q[i]-'0')
                i++
            }
            c := q[len(q)-1]
            nm := 1 << (c - 'a')
            delta := mask[u] ^ nm
            mask[u] = nm
            rangeXor(inn[u], out[u]-1, delta)
        } else {
            u := 0
            v := 0
            i := 6
            for q[i] != ' ' {
                u = u*10 + int(q[i]-'0')
                i++
            }
            i++
            for i < len(q) {
                v = v*10 + int(q[i]-'0')
                i++
            }
            m := pathXor(u, v)
            cnt := 0
            for m > 0 {
                cnt += m & 1
                m >>= 1
            }
            ans = append(ans, cnt <= 1)
        }
    }
    return ans
}