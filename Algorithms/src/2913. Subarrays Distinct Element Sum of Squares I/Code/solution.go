func sumCounts(nums []int) int {
    n := len(nums)
    res := 0
    for i := 0; i < n; i++ {
        s := make(map[int]bool)
        for j := i; j < n; j++ {
            s[nums[j]] = true
            d := len(s)
            res += d * d
        }
    }
    return res
}