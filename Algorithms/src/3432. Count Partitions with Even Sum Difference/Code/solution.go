func countPartitions(nums []int) int {
    // Compute the total sum of the array
    total := 0
    for _, x := range nums {
        total += x
    }

    // If total sum is odd, no valid partition
    if total%2 != 0 {
        return 0
    }

    // If total is even, every position between elements is a valid partition
    n := len(nums)
    return n - 1
}
