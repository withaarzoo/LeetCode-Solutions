import java.util.*;

class Solution {
    public int minCost(int[][] grid, int k) {
        int m = grid.length, n = grid[0].length;
        final long INF = (long) 4e18;
        long[][] dp = new long[m][n];
        for (int i = 0; i < m; i++)
            Arrays.fill(dp[i], INF);

        // base dp (no teleport)
        dp[0][0] = 0;
        for (int i = 0; i < m; i++) {
            for (int j = 0; j < n; j++) {
                if (i > 0)
                    dp[i][j] = Math.min(dp[i][j], dp[i - 1][j] + grid[i][j]);
                if (j > 0)
                    dp[i][j] = Math.min(dp[i][j], dp[i][j - 1] + grid[i][j]);
            }
        }

        // prepare cells list sorted by value descending
        ArrayList<int[]> cells = new ArrayList<>();
        for (int i = 0; i < m; i++)
            for (int j = 0; j < n; j++)
                cells.add(new int[] { grid[i][j], i, j });
        cells.sort((a, b) -> Integer.compare(b[0], a[0]));

        for (int step = 0; step < k; step++) {
            long[][] start = new long[m][n];
            for (int i = 0; i < m; i++)
                Arrays.fill(start[i], INF);

            long runningMin = INF;
            int idx = 0;
            while (idx < cells.size()) {
                int val = cells.get(idx)[0];
                int j = idx;
                long minGroup = INF;
                while (j < cells.size() && cells.get(j)[0] == val) {
                    int ii = cells.get(j)[1], jj = cells.get(j)[2];
                    minGroup = Math.min(minGroup, dp[ii][jj]);
                    j++;
                }
                runningMin = Math.min(runningMin, minGroup);
                for (int p = idx; p < j; p++) {
                    int ii = cells.get(p)[1], jj = cells.get(p)[2];
                    start[ii][jj] = Math.min(dp[ii][jj], runningMin);
                }
                idx = j;
            }

            long[][] dp2 = new long[m][n];
            for (int i = 0; i < m; i++)
                Arrays.fill(dp2[i], INF);
            for (int i = 0; i < m; i++) {
                for (int j = 0; j < n; j++) {
                    if (start[i][j] < dp2[i][j])
                        dp2[i][j] = start[i][j];
                    if (i + 1 < m && dp2[i][j] < INF)
                        dp2[i + 1][j] = Math.min(dp2[i + 1][j], dp2[i][j] + grid[i + 1][j]);
                    if (j + 1 < n && dp2[i][j] < INF)
                        dp2[i][j + 1] = Math.min(dp2[i][j + 1], dp2[i][j] + grid[i][j + 1]);
                }
            }
            dp = dp2;
        }

        return (int) dp[m - 1][n - 1];
    }
}
