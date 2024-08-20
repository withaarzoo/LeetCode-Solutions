class Solution {
    public int stoneGameII(int[] piles) {
        int n = piles.length;

        // dp[i][M] will store the maximum number of stones the current player can get
        // starting from index i with a value of M (where M is the max piles that can be
        // taken).
        int[][] dp = new int[n + 1][n + 1];

        // suffixSum[i] will store the sum of all the elements from index i to the end
        // of the array.
        int[] suffixSum = new int[n + 1];

        // Calculate suffix sums from the end of the array to the beginning.
        for (int i = n - 1; i >= 0; --i) {
            // Suffix sum at index i is the current pile plus the suffix sum at index i + 1.
            suffixSum[i] = suffixSum[i + 1] + piles[i];
        }

        // Fill the dp table starting from the end of the array towards the beginning.
        for (int i = n - 1; i >= 0; --i) {
            // Iterate over possible values of M (1 to n).
            for (int M = 1; M <= n; ++M) {
                // The player can take X piles, where X ranges from 1 to 2 * M.
                // Ensure that taking X piles doesn't exceed the bounds of the array.
                for (int X = 1; X <= 2 * M && i + X <= n; ++X) {
                    // Update dp[i][M] with the maximum possible stones the current player can get.
                    // The current player takes suffixSum[i] stones, and the opponent's maximum
                    // stones is dp[i + X][Math.max(M, X)].
                    dp[i][M] = Math.max(dp[i][M], suffixSum[i] - dp[i + X][Math.max(M, X)]);
                }
            }
        }

        // The answer is the maximum number of stones the first player can get
        // starting from index 0 with M initially set to 1.
        return dp[0][1];
    }
}
