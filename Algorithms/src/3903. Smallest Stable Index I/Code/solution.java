class Solution {
    public int firstStableIndex(int[] nums, int k) {
        int n = nums.length;

        // suffixMin[i] stores the minimum value from i to n - 1.
        int[] suffixMin = new int[n];

        // For the last index, the suffix contains only nums[n - 1].
        suffixMin[n - 1] = nums[n - 1];

        // Build suffix minimums from right to left.
        for (int i = n - 2; i >= 0; i--) {
            // Keep the smaller value between the current element
            // and the minimum value already calculated to its right.
            suffixMin[i] = Math.min(nums[i], suffixMin[i + 1]);
        }

        // Store the largest value seen from index 0 to the current index.
        int prefixMax = nums[0];

        // Check indices from smallest to largest.
        for (int i = 0; i < n; i++) {
            // Update the maximum value in nums[0..i].
            prefixMax = Math.max(prefixMax, nums[i]);

            // Calculate the instability score at index i.
            int instability = prefixMax - suffixMin[i];

            // The first valid index is the smallest stable index.
            if (instability <= k) {
                return i;
            }
        }

        // No stable index was found.
        return -1;
    }
}