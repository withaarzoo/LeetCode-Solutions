class Solution {
    private int n;
    private int[] suffix;
    private int[][] dp;

    private int solve(int i, int m) {
        if (i == n) {
            return 0;
        }

        if (dp[i][m] != -1) {
            return dp[i][m];
        }

        int best = 0;

        for (int x = 1; x <= 2 * m && i + x <= n; x++) {
            int nextM = Math.max(m, x);
            int current = suffix[i] - solve(i + x, nextM);
            best = Math.max(best, current);
        }

        return dp[i][m] = best;
    }

    public int stoneGameII(int[] piles) {
        n = piles.length;
        suffix = new int[n + 1];

        for (int i = n - 1; i >= 0; i--) {
            suffix[i] = suffix[i + 1] + piles[i];
        }

        dp = new int[n][n + 1];

        for (int i = 0; i < n; i++) {
            java.util.Arrays.fill(dp[i], -1);
        }

        return solve(0, 1);
    }
}