func maximumSaleItems(items [][]int, budget int) int {
    mult := make([]int, 1501)
    for _, it := range items {
        fac := it[0]
        for f := 1; f <= 1500; f++ {
            if fac%f == 0 {
                mult[f]++
            }
        }
    }
    dp := make([]int, budget+1)
    for _, it := range items {
        f, p := it[0], it[1]
        ndp := make([]int, budget+1)
        copy(ndp, dp)
        for b := budget; b >= p; b-- {
            if dp[b-p]+mult[f] > ndp[b] {
                ndp[b] = dp[b-p] + mult[f]
            }
        }
        for b := p; b <= budget; b++ {
            if ndp[b-p]+1 > ndp[b] {
                ndp[b] = ndp[b-p] + 1
            }
        }
        dp = ndp
    }
    ans := 0
    for _, v := range dp {
        if v > ans {
            ans = v
        }
    }
    return ans
}