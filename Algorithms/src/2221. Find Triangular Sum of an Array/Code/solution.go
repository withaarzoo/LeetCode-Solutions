func triangularSum(nums []int) int {
    n := len(nums)
    // reduce length from n down to 1
    for length := n; length > 1; length-- {
        // update in-place; nums[i+1] is still old value during this inner loop
        for i := 0; i < length-1; i++ {
            nums[i] = (nums[i] + nums[i+1]) % 10
        }
    }
    return nums[0]
}
