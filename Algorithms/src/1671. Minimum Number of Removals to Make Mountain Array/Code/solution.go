func minimumMountainRemovals(nums []int) int {
    n := len(nums)
    LIS := make([]int, n)
    LDS := make([]int, n)

    for i := range LIS {
        LIS[i] = 1
        LDS[i] = 1
    }

    // Compute LIS for each index
    for i := 0; i < n; i++ {
        for j := 0; j < i; j++ {
            if nums[i] > nums[j] {
                LIS[i] = max(LIS[i], LIS[j] + 1)
            }
        }
    }

    // Compute LDS from each index
    for i := n - 1; i >= 0; i-- {
        for j := n - 1; j > i; j-- {
            if nums[i] > nums[j] {
                LDS[i] = max(LDS[i], LDS[j] + 1)
            }
        }
    }

    maxMountainLength := 0

    // Find the maximum mountain length
    for i := 1; i < n-1; i++ {
        if LIS[i] > 1 && LDS[i] > 1 {
            maxMountainLength = max(maxMountainLength, LIS[i] + LDS[i] - 1)
        }
    }

    return n - maxMountainLength
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}