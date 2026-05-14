class Solution {
    public boolean isGood(int[] nums) {

        // Sort the array
        Arrays.sort(nums);

        // Size of array
        int n = nums.length;

        // Largest element
        int mx = nums[n - 1];

        // Valid size should be mx + 1
        if (n != mx + 1) {
            return false;
        }

        // Check sequence 1 to mx
        for (int i = 0; i < n - 1; i++) {

            // Expected number is i + 1
            if (nums[i] != i + 1) {
                return false;
            }
        }

        // Last element should also be mx
        return nums[n - 1] == mx;
    }
}