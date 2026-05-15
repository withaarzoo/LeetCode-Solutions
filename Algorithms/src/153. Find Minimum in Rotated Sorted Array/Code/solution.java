class Solution {
    public int findMin(int[] nums) {

        // Left pointer at start
        int left = 0;

        // Right pointer at end
        int right = nums.length - 1;

        // Binary Search loop
        while (left < right) {

            // Find middle index
            int mid = left + (right - left) / 2;

            // Minimum is in right half
            if (nums[mid] > nums[right]) {

                // Move left pointer
                left = mid + 1;
            } else {

                // Minimum can be mid itself
                right = mid;
            }
        }

        // Return minimum element
        return nums[left];
    }
}