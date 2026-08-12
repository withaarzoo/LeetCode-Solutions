class Solution {
    public int maxSubarrayLength(int[] nums, int k) {
        HashMap<Integer, Integer> freq = new HashMap<>(); // Stores frequencies of values in the current window.
        int left = 0; // Left boundary of the sliding window.
        int ans = 0; // Stores the longest valid window length.

        for (int right = 0; right < nums.length; right++) {
            freq.put(nums[right], freq.getOrDefault(nums[right], 0) + 1); // Add nums[right] to the window.

            while (freq.get(nums[right]) > k) {
                freq.put(nums[left], freq.get(nums[left]) - 1); // Remove nums[left] from the current window.
                left++; // Move the left boundary forward.
            }

            ans = Math.max(ans, right - left + 1); // Update the answer using the current valid window.
        }

        return ans; // Return the longest valid subarray length.
    }
}