func firstStableIndex(nums []int, k int) int {
    n := len(nums)

    // suffixMin[i] stores the minimum value from i to the end.
    suffixMin := make([]int, n)

    // The last suffix contains only the last element.
    suffixMin[n-1] = nums[n-1]

    // Build suffix minimums from right to left.
    for i := n - 2; i >= 0; i-- {
        // Store the smaller value between nums[i]
        // and the minimum value already calculated to its right.
        if nums[i] < suffixMin[i+1] {
            suffixMin[i] = nums[i]
        } else {
            suffixMin[i] = suffixMin[i+1]
        }
    }

    // Store the largest value seen from index 0 to the current index.
    prefixMax := nums[0]

    // Check indices from smallest to largest.
    for i := 0; i < n; i++ {
        // Update the maximum value in nums[0..i].
        if nums[i] > prefixMax {
            prefixMax = nums[i]
        }

        // Calculate the instability score at index i.
        instability := prefixMax - suffixMin[i]

        // Return immediately because this is the first stable index.
        if instability <= k {
            return i
        }
    }

    // No stable index exists.
    return -1
}