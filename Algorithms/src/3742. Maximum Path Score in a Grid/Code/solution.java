import java.util.Arrays;

class Solution {
    public int maxPathScore(int[][] grid, int k) {
        int m = grid.length; // Number of rows in the grid.
        int n = grid[0].length; // Number of columns in the grid.
        final int NEG = -1_000_000_000; // Sentinel for impossible states.

        // prev[j][c] = best score at column j in the previous row using exact cost c.
        int[][] prev = new int[n][k + 1];
        for (int j = 0; j < n; j++) {
            Arrays.fill(prev[j], NEG); // Mark every state as impossible first.
        }

        for (int i = 0; i < m; i++) {
            // Rebuild the current row from scratch so old values do not interfere.
            int[][] curr = new int[n][k + 1];
            for (int j = 0; j < n; j++) {
                Arrays.fill(curr[j], NEG); // Reset the row before filling new states.
            }

            for (int j = 0; j < n; j++) {
                int gain = grid[i][j]; // Score gained by taking this cell.
                int need = gain > 0 ? 1 : 0; // Cost spent by this cell: 0 for 0, 1 for 1/2.

                // A path to (i, j) cannot spend more than i + j budget points.
                int limit = Math.min(k, i + j);

                // The starting cell is always 0, so it is the base state.
                if (i == 0 && j == 0) {
                    curr[0][0] = 0; // Zero score, zero cost at the start.
                    continue;
                }

                for (int c = need; c <= limit; c++) {
                    int best = NEG;

                    // Coming from above means using the finished previous row.
                    if (i > 0 && prev[j][c - need] != NEG) {
                        best = Math.max(best, prev[j][c - need] + gain);
                    }

                    // Coming from the left means using the current row already built.
                    if (j > 0 && curr[j - 1][c - need] != NEG) {
                        best = Math.max(best, curr[j - 1][c - need] + gain);
                    }

                    curr[j][c] = best; // Save the best exact-cost answer here.
                }
            }

            prev = curr; // Move this row up so it becomes the previous row.
        }

        int ans = NEG; // Best score among all valid costs.
        for (int c = 0; c <= k; c++) {
            ans = Math.max(ans, prev[n - 1][c]);
        }

        return ans < 0 ? -1 : ans; // Return -1 when no valid path exists.
    }
}