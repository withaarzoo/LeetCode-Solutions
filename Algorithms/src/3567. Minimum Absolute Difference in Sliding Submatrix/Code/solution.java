import java.util.*;

class Solution {
    public int[][] minAbsDiff(int[][] grid, int k) {
        int m = grid.length;
        int n = grid[0].length;
        int[][] ans = new int[m - k + 1][n - k + 1];

        for (int i = 0; i + k <= m; i++) {
            for (int j = 0; j + k <= n; j++) {
                ArrayList<Integer> vals = new ArrayList<>();

                // Collect all values from the current k x k submatrix
                for (int r = i; r < i + k; r++) {
                    for (int c = j; c < j + k; c++) {
                        vals.add(grid[r][c]);
                    }
                }

                Collections.sort(vals);

                int best = Integer.MAX_VALUE;

                // Check only consecutive different values
                for (int x = 1; x < vals.size(); x++) {
                    if (!vals.get(x).equals(vals.get(x - 1))) {
                        best = Math.min(best, vals.get(x) - vals.get(x - 1));
                    }
                }

                ans[i][j] = (best == Integer.MAX_VALUE ? 0 : best);
            }
        }

        return ans;
    }
}