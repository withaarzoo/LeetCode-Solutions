class Solution {
    public boolean check(int[] nums) {

        int n = nums.length;

        // Counts how many times order breaks
        int count = 0;

        // Traverse the array
        for (int i = 0; i < n; i++) {

            // Compare current element with next element
            // % n helps compare last element with first
            if (nums[i] > nums[(i + 1) % n]) {
                count++;
            }

            // More than one break means invalid
            if (count > 1) {
                return false;
            }
        }

        // Array is valid
        return true;
    }
}