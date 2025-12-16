class Solution {
    int n, K;
    int[] p;
    Long[][][] dp;
    final long NEG = (long) -1e18;

    long dfs(int i, int t, int s) {
        if (i == n)
            return s == 0 ? 0 : NEG;

        if (dp[i][t][s] != null)
            return dp[i][t][s];

        long res = dfs(i + 1, t, s);

        if (s == 0) {
            res = Math.max(res, dfs(i + 1, t, 1) - p[i]);
            res = Math.max(res, dfs(i + 1, t, 2) + p[i]);
        } else if (s == 1 && t < K) {
            res = Math.max(res, dfs(i + 1, t + 1, 0) + p[i]);
        } else if (s == 2 && t < K) {
            res = Math.max(res, dfs(i + 1, t + 1, 0) - p[i]);
        }

        return dp[i][t][s] = res;
    }

    public long maximumProfit(int[] prices, int k) {
        p = prices;
        n = prices.length;
        K = k;
        dp = new Long[n + 1][k + 1][3];
        return dfs(0, 0, 0);
    }
}
