class Solution {
    public String stoneGameIII(int[] stoneValue) {

        int n = stoneValue.length;

        // dp[i] stores the maximum score difference from index i
        int[] dp = new int[n + 1];

        // Fill DP from right to left
        for (int i = n - 1; i >= 0; i--) {

            dp[i] = Integer.MIN_VALUE;
            int sum = 0;

            // Try taking 1, 2 and 3 stones
            for (int j = i; j < Math.min(n, i + 3); j++) {

                // Add current stone value
                sum += stoneValue[j];

                // Update the best possible score difference
                dp[i] = Math.max(dp[i], sum - dp[j + 1]);
            }
        }

        // Decide the winner
        if (dp[0] > 0)
            return "Alice";
        if (dp[0] < 0)
            return "Bob";
        return "Tie";
    }
}