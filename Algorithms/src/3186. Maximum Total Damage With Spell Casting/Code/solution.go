func maximumTotalDamage(power []int) int64 {
    if len(power) == 0 { return 0 }
    freq := make(map[int64]int64)
    for _, v := range power {
        freq[int64(v)]++
    }
    vals := make([]int64, 0, len(freq))
    for k := range freq { vals = append(vals, k) }
    sort.Slice(vals, func(i, j int) bool { return vals[i] < vals[j] })
    m := len(vals)
    valueSum := make([]int64, m)
    for i := 0; i < m; i++ { valueSum[i] = vals[i] * freq[vals[i]] }
    dp := make([]int64, m)
    dp[0] = valueSum[0]
    for i := 1; i < m; i++ {
        need := vals[i] - 3
        // binary search last index j with vals[j] <= need
        lo, hi, j := 0, i-1, -1
        for lo <= hi {
            mid := (lo + hi) / 2
            if vals[mid] <= need { j = mid; lo = mid + 1 } else { hi = mid - 1 }
        }
        take := valueSum[i]
        if j >= 0 { take += dp[j] }
        skip := dp[i-1]
        if take > skip { dp[i] = take } else { dp[i] = skip }
    }
    return dp[m-1]
}
