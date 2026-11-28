func maximumSaleItems(items [][]int, budget int) int {
    n := len(items)
    cnt := make([]int, n+1)
    for _, it := range items {
        cnt[it[0]]++
    }
    mult := make([]int, n+1)
    for f := 1; f <= n; f++ {
        for m := f; m <= n; m += f {
            mult[f] += cnt[m]
        }
    }
    minP := int(^uint(0) >> 1)
    type pair struct{ p, g int }
    cand := make([]pair, 0)
    for _, it := range items {
        f, p := it[0], it[1]
        if p < minP {
            minP = p
        }
        g := mult[f] - 1
        if g > 0 {
            cand = append(cand, pair{p, g})
        }
    }
    sort.Slice(cand, func(i, j int) bool { return cand[i].p < cand[j].p })
    spent, total, B := int64(0), int64(0), int64(budget)
    for _, c := range cand {
        if int64(c.p) > 2*int64(minP) {
            break
        }
        can := int64(c.g)
        if rem := (B - spent) / int64(c.p); rem < can {
            can = rem
        }
        total += can * 2
        spent += can * int64(c.p)
    }
    total += (B - spent) / int64(minP)
    return int(total)
}