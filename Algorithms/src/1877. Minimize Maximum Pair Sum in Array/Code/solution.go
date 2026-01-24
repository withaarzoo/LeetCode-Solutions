func minPairSum(nums []int) int {
    // Step 1: Sort the array
    sort.Ints(nums)

    left, right := 0, len(nums)-1
    maxPairSum := 0

    // Step 2: Pair smallest with largest
    for left < right {
        pairSum := nums[left] + nums[right]
        if pairSum > maxPairSum {
            maxPairSum = pairSum
        }
        left++
        right--
    }

    return maxPairSum
}
