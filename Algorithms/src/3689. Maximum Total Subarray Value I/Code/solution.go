func maxTotalValue(nums []int, k int) int64 {
    // Initialize minimum and maximum
    mn := nums[0]
    mx := nums[0]

    // Find global minimum and maximum
    for _, num := range nums {
        if num < mn {
            mn = num
        }
        if num > mx {
            mx = num
        }
    }

    // Best subarray value
    best := int64(mx - mn)

    // Choose the same best subarray k times
    return best * int64(k)
}