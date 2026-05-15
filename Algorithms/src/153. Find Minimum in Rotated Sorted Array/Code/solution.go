func findMin(nums []int) int {
    
    // Left pointer
    left := 0

    // Right pointer
    right := len(nums) - 1

    // Binary Search loop
    for left < right {

        // Middle index
        mid := left + (right-left)/2

        // Minimum is on right side
        if nums[mid] > nums[right] {

            // Move left pointer
            left = mid + 1
        } else {

            // Minimum may be at mid
            right = mid
        }
    }

    // Return minimum element
    return nums[left]
}