func minimumDifference(nums []int, k int) int {
    // If k == 1, difference is always 0
    if k == 1 {
        return 0
    }

    // Step 1: Sort the array
    sort.Ints(nums)

    minDiff := math.MaxInt32

    // Step 2: Sliding window
    for i := 0; i + k - 1 < len(nums); i++ {
        diff := nums[i + k - 1] - nums[i]
        if diff < minDiff {
            minDiff = diff
        }
    }

    return minDiff
}
