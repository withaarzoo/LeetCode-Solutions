func minOperations(nums []int, x int) int {
    total := 0
    for _, num := range nums {
        total += num
    }
    if total < x {
        return -1
    }
    target := total - x
    if target == 0 {
        return len(nums)
    }
    n := len(nums)
    left := 0
    sum := 0
    maxLen := -1
    for right := 0; right < n; right++ {
        sum += nums[right]
        for sum > target && left <= right {
            sum -= nums[left]
            left++
        }
        if sum == target {
            if right-left+1 > maxLen {
                maxLen = right - left + 1
            }
        }
    }
    if maxLen == -1 {
        return -1
    }
    return n - maxLen
}