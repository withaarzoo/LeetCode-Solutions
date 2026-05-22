func search(nums []int, target int) int {
    
    // Start pointer
    left := 0

    // End pointer
    right := len(nums) - 1

    // Continue until search space becomes empty
    for left <= right {

        // Find middle index safely
        mid := left + (right-left)/2

        // If target is found
        if nums[mid] == target {
            return mid
        }

        // Check if left half is sorted
        if nums[left] <= nums[mid] {

            // Check whether target lies inside left sorted half
            if nums[left] <= target && target < nums[mid] {

                // Search left side
                right = mid - 1
            } else {

                // Search right side
                left = mid + 1
            }

        } else {

            // Right half is sorted

            // Check whether target lies inside right sorted half
            if nums[mid] < target && target <= nums[right] {

                // Search right side
                left = mid + 1
            } else {

                // Search left side
                right = mid - 1
            }
        }
    }

    // Target not found
    return -1
}