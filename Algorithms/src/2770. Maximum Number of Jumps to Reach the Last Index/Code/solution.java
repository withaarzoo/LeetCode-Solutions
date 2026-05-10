class Solution {
    public int maximumJumps(int[] nums, int target) {

        int n = nums.length;

        // dp[i] = maximum jumps needed to reach index i
        int[] dp = new int[n];

        // Fill with -1 meaning unreachable
        Arrays.fill(dp, -1);

        // Starting index
        dp[0] = 0;

        // Try every current index
        for (int i = 0; i < n; i++) {

            // Skip unreachable indices
            if (dp[i] == -1)
                continue;

            // Try every next index
            for (int j = i + 1; j < n; j++) {

                // Calculate difference
                long diff = (long) nums[j] - nums[i];

                // Check valid jump
                if (diff >= -target && diff <= target) {

                    // Update maximum jumps
                    dp[j] = Math.max(dp[j], dp[i] + 1);
                }
            }
        }

        // Return answer
        return dp[n - 1];
    }
}