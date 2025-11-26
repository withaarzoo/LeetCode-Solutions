class Solution {
    private static final int MOD = 1_000_000_007;

    public int numberOfPaths(int[][] grid, int k) {
        int m = grid.length;
        int n = grid[0].length;

        // prev[j][r] = paths to (i-1, j) with sum % k == r
        // cur[j][r] = paths to (i, j) with sum % k == r
        int[][] prev = new int[n][k];
        int[][] cur = new int[n][k];

        for (int i = 0; i < m; i++) {
            // clear current row
            for (int j = 0; j < n; j++) {
                Arrays.fill(cur[j], 0);
            }

            for (int j = 0; j < n; j++) {
                int val = grid[i][j] % k;

                // starting cell
                if (i == 0 && j == 0) {
                    cur[0][val] = 1;
                    continue;
                }

                // from top
                if (i > 0) {
                    for (int r = 0; r < k; r++) {
                        if (prev[j][r] == 0)
                            continue;
                        int nr = (r + val) % k;
                        cur[j][nr] = (cur[j][nr] + prev[j][r]) % MOD;
                    }
                }

                // from left
                if (j > 0) {
                    for (int r = 0; r < k; r++) {
                        if (cur[j - 1][r] == 0)
                            continue;
                        int nr = (r + val) % k;
                        cur[j][nr] = (cur[j][nr] + cur[j - 1][r]) % MOD;
                    }
                }
            }

            // move current row into prev
            int[][] tmp = prev;
            prev = cur;
            cur = tmp;
        }

        return prev[n - 1][0];
    }
}
