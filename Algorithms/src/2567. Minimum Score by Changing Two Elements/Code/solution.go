func minimizeSum(nums []int) int {
    sort.Ints(nums)
    n := len(nums)
    return min(nums[n-1]-nums[2], min(nums[n-3]-nums[0], nums[n-2]-nums[1]))
}