class Solution {
    public int minimumDifference(int[] nums, int k) {
        // If only one student is picked
        if (k == 1)
            return 0;

        // Step 1: Sort the array
        Arrays.sort(nums);

        int minDiff = Integer.MAX_VALUE;

        // Step 2: Sliding window
        for (int i = 0; i + k - 1 < nums.length; i++) {
            int diff = nums[i + k - 1] - nums[i];
            minDiff = Math.min(minDiff, diff);
        }

        return minDiff;
    }
}
