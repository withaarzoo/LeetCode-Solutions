class Solution {
    public boolean predictTheWinner(int[] nums) {
        int n = nums.length;

        // dp[i][j] stores the maximum score difference
        // current player can achieve for subarray [i...j].
        int[][] dp = new int[n][n];

        // Base case:
        // One element means the current player takes it.
        for (int i = 0; i < n; i++) {
            dp[i][i] = nums[i];
        }

        // Fill the table from smaller ranges to larger ranges.
        for (int len = 2; len <= n; len++) {
            for (int i = 0; i + len - 1 < n; i++) {
                int j = i + len - 1;

                // Pick the left number.
                int takeLeft = nums[i] - dp[i + 1][j];

                // Pick the right number.
                int takeRight = nums[j] - dp[i][j - 1];

                // Store the better result.
                dp[i][j] = Math.max(takeLeft, takeRight);
            }
        }

        // Player 1 wins or ties if score difference is non-negative.
        return dp[0][n - 1] >= 0;
    }
}