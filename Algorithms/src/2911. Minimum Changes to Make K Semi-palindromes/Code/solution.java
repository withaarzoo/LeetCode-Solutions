class Solution {
    public int minimumChanges(String s, int k) {
        int n = s.length();
        final int INF = 1000000000;
        int[][] cost = new int[n][n];
        for (int i = 0; i < n; i++) {
            for (int j = 0; j < n; j++) cost[i][j] = INF;
        }
        for (int i = 0; i < n; i++) {
            for (int j = i + 1; j < n; j++) {
                int L = j - i + 1;
                int best = INF;
                for (int d = 1; d < L; d++) {
                    if (L % d != 0) continue;
                    int cur = 0;
                    int m = L / d;
                    for (int r = 0; r < d; r++) {
                        for (int t = 0; t < m / 2; t++) {
                            if (s.charAt(i + r + t * d) != s.charAt(i + r + (m - 1 - t) * d))
                                cur++;
                        }
                    }
                    best = Math.min(best, cur);
                }
                cost[i][j] = best;
            }
        }
        int[][] dp = new int[n + 1][k + 1];
        for (int i = 0; i <= n; i++) {
            for (int p = 0; p <= k; p++) dp[i][p] = INF;
        }
        dp[0][0] = 0;
        for (int i = 1; i <= n; i++) {
            for (int p = 1; p <= k; p++) {
                for (int prev = 0; prev <= i - 2; prev++) {
                    if (dp[prev][p - 1] != INF && cost[prev][i - 1] != INF)
                        dp[i][p] = Math.min(dp[i][p], dp[prev][p - 1] + cost[prev][i - 1]);
                }
            }
        }
        return dp[n][k];
    }
}