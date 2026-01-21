func minBitwiseArray(nums []int) []int {
    for i, p := range nums {
        removable := ((p + 1) & ^p) >> 1

        if removable == 0 {
            nums[i] = -1
        } else {
            nums[i] = p ^ removable
        }
    }
    return nums
}
