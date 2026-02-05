func constructTransformedArray(nums []int) []int {
    n := len(nums)
    result := make([]int, n)

    for i := 0; i < n; i++ {
        if nums[i] == 0 {
            result[i] = nums[i]
        } else {
            target := (i + nums[i]) % n
            if target < 0 {
                target += n
            }
            result[i] = nums[target]
        }
    }
    return result
}
