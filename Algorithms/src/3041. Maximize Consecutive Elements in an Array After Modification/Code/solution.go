func maxSelectedElements(nums []int) int {
    sort.Ints(nums)
    dp := make(map[int]int)
    ans := 0
    for _, num := range nums {
        dp[num+1] = dp[num] + 1
        dp[num] = dp[num-1] + 1
        if dp[num] > ans {
            ans = dp[num]
        }
        if dp[num+1] > ans {
            ans = dp[num+1]
        }
    }
    return ans
}