class Solution {
    public int firstStableIndex(int[] nums, int k) {
        int n = nums.length; // Store the number of elements in the array.

        int[] prefixMax = new int[n]; // prefixMax[i] stores max(nums[0..i]).
        int[] suffixMin = new int[n]; // suffixMin[i] stores min(nums[i..n-1]).

        prefixMax[0] = nums[0]; // For index 0, the prefix contains only nums[0].

        for (int i = 1; i < n; i++) { // Build prefix maximums from left to right.
            prefixMax[i] = Math.max(prefixMax[i - 1], nums[i]); // Keep the largest value seen so far.
        }

        suffixMin[n - 1] = nums[n - 1]; // For the last index, the suffix contains only nums[n-1].

        for (int i = n - 2; i >= 0; i--) { // Build suffix minimums from right to left.
            suffixMin[i] = Math.min(suffixMin[i + 1], nums[i]); // Keep the smallest value in the suffix.
        }

        for (int i = 0; i < n; i++) { // Check indices from smallest to largest.
            int instability = prefixMax[i] - suffixMin[i]; // Calculate the score for index i.

            if (instability <= k) { // The index is stable when its score is at most k.
                return i; // This is the first stable index because we scan left to right.
            }
        }

        return -1; // Return -1 when no stable index exists.
    }
}