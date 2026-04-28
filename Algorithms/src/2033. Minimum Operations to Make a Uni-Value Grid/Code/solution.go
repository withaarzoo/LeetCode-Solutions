func minOperations(grid [][]int, x int) int {
    nums := []int{}

    // Flatten grid
    for _, row := range grid {
        for _, val := range row {
            nums = append(nums, val)
        }
    }

    // Check feasibility
    rem := nums[0] % x
    for _, num := range nums {
        if num%x != rem {
            return -1
        }
    }

    // Sort
    sort.Ints(nums)

    // Median
    median := nums[len(nums)/2]

    // Count operations
    ops := 0
    for _, num := range nums {
        if num > median {
            ops += (num - median) / x
        } else {
            ops += (median - num) / x
        }
    }

    return ops
}