class Solution {
    public long maximumScore(int[][] grid) {
        int n = grid.length;
        if (n == 1)
            return 0L;

        // pref[c][k] = sum of first k cells in column c
        long[][] pref = new long[n][n + 1];
        for (int c = 0; c < n; c++) {
            for (int r = 0; r < n; r++) {
                pref[c][r + 1] = pref[c][r] + grid[r][c];
            }
        }

        final long NEG = -(1L << 60);

        // dp[a][b] = best score after processing up to current column,
        // with previous height = a and current height = b.
        long[][] dp = new long[n + 1][n + 1];
        for (int a = 0; a <= n; a++) {
            for (int b = 0; b <= n; b++) {
                dp[a][b] = Math.max(0L, pref[0][b] - pref[0][a]);
            }
        }

        for (int col = 1; col < n; col++) {
            long[][] ndp = new long[n + 1][n + 1];
            for (int i = 0; i <= n; i++) {
                for (int j = 0; j <= n; j++) {
                    ndp[i][j] = NEG;
                }
            }

            for (int mid = 0; mid <= n; mid++) {
                long[] q = new long[n + 1];
                for (int x = 0; x <= n; x++) {
                    q[x] = Math.max(0L, pref[col][x] - pref[col][mid]);
                }

                long[] prefixBest = new long[n + 1];
                prefixBest[0] = dp[0][mid];
                for (int a = 1; a <= n; a++) {
                    prefixBest[a] = Math.max(prefixBest[a - 1], dp[a][mid]);
                }

                long[] suffixBest = new long[n + 2];
                for (int i = 0; i <= n + 1; i++)
                    suffixBest[i] = NEG;
                suffixBest[n] = dp[n][mid] + q[n];
                for (int a = n - 1; a >= 0; a--) {
                    suffixBest[a] = Math.max(suffixBest[a + 1], dp[a][mid] + q[a]);
                }

                int limit = (col == n - 1 ? 0 : n);
                for (int nxt = 0; nxt <= limit; nxt++) {
                    long best = NEG;

                    if (prefixBest[nxt] != NEG) {
                        best = Math.max(best, prefixBest[nxt] + q[nxt]);
                    }
                    if (suffixBest[nxt + 1] != NEG) {
                        best = Math.max(best, suffixBest[nxt + 1]);
                    }

                    ndp[mid][nxt] = Math.max(ndp[mid][nxt], best);
                }
            }

            dp = ndp;
        }

        long ans = 0;
        for (int a = 0; a <= n; a++) {
            for (int b = 0; b <= n; b++) {
                ans = Math.max(ans, dp[a][b]);
            }
        }
        return ans;
    }
}