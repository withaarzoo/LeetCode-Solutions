func longestMonotonicSubarray(nums []int) int {
    n := len(nums)
    maxLen, inc, dec := 1, 1, 1
    for i := 1; i < n; i++ {
        if nums[i] > nums[i-1] {
            inc++
            dec = 1
        } else if nums[i] < nums[i-1] {
            dec++
            inc = 1
        } else {
            inc = 1
            dec = 1
        }
        if inc > maxLen {
            maxLen = inc
        }
        if dec > maxLen {
            maxLen = dec
        }
    }
    return maxLen
}