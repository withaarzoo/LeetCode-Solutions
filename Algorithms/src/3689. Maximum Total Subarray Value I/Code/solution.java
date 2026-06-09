class Solution {
    public long maxTotalValue(int[] nums, int k) {
        // Initialize minimum and maximum using first element
        long mn = nums[0];
        long mx = nums[0];

        // Find global minimum and maximum
        for (int num : nums) {
            mn = Math.min(mn, num);
            mx = Math.max(mx, num);
        }

        // Best possible subarray value
        long best = mx - mn;

        // Choose that subarray k times
        return best * k;
    }
}