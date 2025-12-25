func maximumHappinessSum(happiness []int, k int) int64 {
    // Sort in descending order
    sort.Slice(happiness, func(i, j int) bool {
        return happiness[i] > happiness[j]
    })
    
    var ans int64 = 0
    
    for i := 0; i < k; i++ {
        curr := happiness[i] - i
        if curr > 0 {
            ans += int64(curr)
        }
    }
    
    return ans
}
