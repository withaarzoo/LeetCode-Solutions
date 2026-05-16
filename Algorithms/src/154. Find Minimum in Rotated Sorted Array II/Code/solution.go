func findMin(nums []int) int {
    
    // Initialize search boundaries
    left := 0
    right := len(nums) - 1

    // Binary search loop
    for left < right {

        // Calculate middle index
        mid := left + (right-left)/2

        // Minimum is on left side including mid
        if nums[mid] < nums[right] {
            right = mid

        // Minimum is on right side
        } else if nums[mid] > nums[right] {
            left = mid + 1

        // Duplicate case
        // Shrink search space safely
        } else {
            right--
        }
    }

    // Return minimum element
    return nums[left]
}