func resultsArray(nums []int, k int) []int {
    n := len(nums)
    result := []int{}

    for i := 0; i <= n-k; i++ {
        subarray := make([]int, k)
        copy(subarray, nums[i:i+k])

        // Sort the subarray
        sortedSubarray := make([]int, k)
        copy(sortedSubarray, subarray)
        sort.Ints(sortedSubarray)

        // Check if elements are consecutive
        isConsecutive := true
        for j := 1; j < k; j++ {
            if sortedSubarray[j]-sortedSubarray[j-1] != 1 {
                isConsecutive = false
                break
            }
        }

        // Add the result based on conditions
        if isConsecutive && equal(subarray, sortedSubarray) {
            result = append(result, sortedSubarray[k-1]) // Max element
        } else {
            result = append(result, -1)
        }
    }

    return result
}

// Helper function to compare slices
func equal(a, b []int) bool {
    if len(a) != len(b) {
        return false
    }
    for i := range a {
        if a[i] != b[i] {
            return false
        }
    }
    return true
}