class Solution {
    public int stoneGameV(int[] stoneValue) {
        int n = stoneValue.length;

        // prefix[i] stores the sum of the first i stones.
        // This makes every interval sum available in O(1).
        long[] prefix = new long[n + 1];

        for (int i = 0; i < n; i++) {
            // Build the prefix sum array.
            prefix[i + 1] = prefix[i] + stoneValue[i];
        }

        // dp[l][r] stores Alice's maximum score for interval [l, r].
        int[][] dp = new int[n][n];

        // leftBest[l][r] stores:
        // max(dp[l][k] + sum(l, k)) for k in [l, r].
        int[][] leftBest = new int[n][n];

        // rightBest[l][r] stores:
        // max(dp[k][r] + sum(k, r)) for k in [l, r].
        int[][] rightBest = new int[n][n];

        // leftPtr[l] tracks the last split where leftSum <= rightSum.
        int[] leftPtr = new int[n];

        // rightPtr[l] tracks the first split where leftSum >= rightSum.
        int[] rightPtr = new int[n];

        for (int i = 0; i < n; i++) {
            // A single stone cannot be split, so its dp value remains 0.
            // The helper tables still need the single stone's value.
            leftBest[i][i] = stoneValue[i];
            rightBest[i][i] = stoneValue[i];

            // There is no split before the starting index.
            leftPtr[i] = i - 1;

            // The first possible split starts at the starting index.
            rightPtr[i] = i;
        }

        // Process intervals from shorter to longer.
        // Smaller intervals are therefore already solved.
        for (int len = 2; len <= n; len++) {
            for (int l = 0; l + len <= n; l++) {
                int r = l + len - 1;

                // Get the total sum of [l, r] in O(1).
                long total = prefix[r + 1] - prefix[l];

                // Move the left boundary while the left side is
                // still smaller than or equal to the right side.
                while (leftPtr[l] + 1 <= r - 1) {
                    int k = leftPtr[l] + 1;
                    long leftSum = prefix[k + 1] - prefix[l];

                    // 2 * leftSum <= total means leftSum <= rightSum.
                    if (2 * leftSum > total) {
                        break;
                    }

                    leftPtr[l]++;
                }

                // Move the right boundary until leftSum >= rightSum.
                while (rightPtr[l] <= r - 1) {
                    int k = rightPtr[l];
                    long leftSum = prefix[k + 1] - prefix[l];

                    // This is the first position where the left side
                    // becomes at least as large as the right side.
                    if (2 * leftSum >= total) {
                        break;
                    }

                    rightPtr[l]++;
                }

                int best = 0;

                // When the left side is smaller or equal, Alice keeps it.
                if (leftPtr[l] >= l) {
                    best = leftBest[l][leftPtr[l]];
                }

                // When the right side is smaller or equal, Alice keeps it.
                if (rightPtr[l] <= r - 1) {
                    best = Math.max(best, rightBest[rightPtr[l] + 1][r]);
                }

                // Save the best possible score for this interval.
                dp[l][r] = best;

                // Include [l, r] as a possible left part of a larger interval.
                leftBest[l][r] = Math.max(
                        leftBest[l][r - 1],
                        dp[l][r] + (int) total);

                // Include [l, r] as a possible right part of a larger interval.
                rightBest[l][r] = Math.max(
                        rightBest[l + 1][r],
                        dp[l][r] + (int) total);
            }
        }

        // The full array is the interval [0, n - 1].
        return dp[0][n - 1];
    }
}