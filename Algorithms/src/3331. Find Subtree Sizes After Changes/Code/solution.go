func findSubtreeSizes(parent []int, s string) []int {
    n := len(parent)
    g := make([][]int, n)
    for i := 1; i < n; i++ {
        g[parent[i]] = append(g[parent[i]], i)
    }
    np := make([]int, n)
    copy(np, parent)
    last := make([]int, 26)
    for i := range last {
        last[i] = -1
    }
    var dfs func(int)
    dfs = func(u int) {
        c := int(s[u] - 'a')
        prev := last[c]
        if prev != -1 {
            np[u] = prev
        }
        last[c] = u
        for _, v := range g[u] {
            dfs(v)
        }
        last[c] = prev
    }
    dfs(0)
    ng := make([][]int, n)
    for i := 1; i < n; i++ {
        ng[np[i]] = append(ng[np[i]], i)
    }
    ans := make([]int, n)
    var size func(int) int
    size = func(u int) int {
        sz := 1
        for _, v := range ng[u] {
            sz += size(v)
        }
        ans[u] = sz
        return sz
    }
    size(0)
    return ans
}