func minimumCost(cost1 int, cost2 int, costBoth int, need1 int, need2 int) int64 {
    c1, c2, cb := int64(cost1), int64(cost2), int64(costBoth)
    n1, n2 := int64(need1), int64(need2)
    pairs := n1
    if n2 < pairs {
        pairs = n2
    }
    costPairs := pairs * min(cb, c1+c2)
    rem1 := n1 - pairs
    rem2 := n2 - pairs
    costRem := rem1*min(c1, cb) + rem2*min(c2, cb)
    return costPairs + costRem
}

func min(a, b int64) int64 {
    if a < b {
        return a
    }
    return b
}