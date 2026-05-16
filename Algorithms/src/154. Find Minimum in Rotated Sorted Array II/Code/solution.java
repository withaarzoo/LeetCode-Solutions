class Solution {
    public int findMin(int[] nums) {

        // Initialize search boundaries
        int left = 0;
        int right = nums.length - 1;

        // Binary search loop
        while (left < right) {

            // Calculate middle index
            int mid = left + (right - left) / 2;

            // Minimum is on left side including mid
            if (nums[mid] < nums[right]) {
                right = mid;
            }

            // Minimum is on right side
            else if (nums[mid] > nums[right]) {
                left = mid + 1;
            }

            // Duplicate case
            // Remove one duplicate safely
            else {
                right--;
            }
        }

        // Return minimum value
        return nums[left];
    }
}