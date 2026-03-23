class Solution {
    public int maxProductPath(int[][] grid) {
        int m = grid.length, n = grid[0].length;
        long MOD = (long)1e9 + 7;

        long[][] maxDp = new long[m][n];
        long[][] minDp = new long[m][n];

        maxDp[0][0] = minDp[0][0] = grid[0][0];

        for (int i = 1; i < m; i++) {
            maxDp[i][0] = minDp[i][0] = maxDp[i-1][0] * grid[i][0];
        }

        for (int j = 1; j < n; j++) {
            maxDp[0][j] = minDp[0][j] = maxDp[0][j-1] * grid[0][j];
        }

        for (int i = 1; i < m; i++) {
            for (int j = 1; j < n; j++) {
                long val = grid[i][j];

                long a = maxDp[i-1][j] * val;
                long b = minDp[i-1][j] * val;
                long c = maxDp[i][j-1] * val;
                long d = minDp[i][j-1] * val;

                maxDp[i][j] = Math.max(Math.max(a, b), Math.max(c, d));
                minDp[i][j] = Math.min(Math.min(a, b), Math.min(c, d));
            }
        }

        long res = maxDp[m-1][n-1];
        return res < 0 ? -1 : (int)(res % MOD);
    }
}