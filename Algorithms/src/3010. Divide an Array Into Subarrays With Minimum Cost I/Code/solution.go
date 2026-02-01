func minimumCost(nums []int) int {
    first := nums[0]

    // Sort remaining elements
    sort.Ints(nums[1:])

    return first + nums[1] + nums[2]
}
