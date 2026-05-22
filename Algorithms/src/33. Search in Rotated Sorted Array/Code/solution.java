class Solution {
    public int search(int[] nums, int target) {

        // Start pointer
        int left = 0;

        // End pointer
        int right = nums.length - 1;

        // Continue until search space becomes empty
        while (left <= right) {

            // Find middle index safely
            int mid = left + (right - left) / 2;

            // If target is found, return index
            if (nums[mid] == target) {
                return mid;
            }

            // Check if left half is sorted
            if (nums[left] <= nums[mid]) {

                // Check whether target lies inside left sorted half
                if (nums[left] <= target && target < nums[mid]) {

                    // Move search to left side
                    right = mid - 1;
                } else {

                    // Move search to right side
                    left = mid + 1;
                }
            }
            // Otherwise right half is sorted
            else {

                // Check whether target lies inside right sorted half
                if (nums[mid] < target && target <= nums[right]) {

                    // Move search to right side
                    left = mid + 1;
                } else {

                    // Move search to left side
                    right = mid - 1;
                }
            }
        }

        // Target not found
        return -1;
    }
}