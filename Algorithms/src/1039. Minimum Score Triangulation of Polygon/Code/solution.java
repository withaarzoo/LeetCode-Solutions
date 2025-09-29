class Solution {
    public int minScoreTriangulation(int[] values) {
        int n = values.length;
        int[][] dp = new int[n][n]; // default 0
        
        for (int len = 3; len <= n; len++) {
            for (int i = 0; i + len - 1 < n; i++) {
                int j = i + len - 1;
                int best = Integer.MAX_VALUE;
                for (int k = i + 1; k < j; k++) {
                    int cost = dp[i][k] + dp[k][j] + values[i] * values[k] * values[j];
                    if (cost < best) best = cost;
                }
                dp[i][j] = best;
            }
        }
        return n == 0 ? 0 : dp[0][n-1];
    }
}
