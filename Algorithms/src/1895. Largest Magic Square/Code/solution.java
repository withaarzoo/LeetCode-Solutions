class Solution {
    public int largestMagicSquare(int[][] grid) {
        int m = grid.length;
        int n = grid[0].length;

        int[][] row = new int[m][n + 1];
        int[][] col = new int[m + 1][n];

        for (int i = 0; i < m; i++) {
            for (int j = 0; j < n; j++) {
                row[i][j + 1] = row[i][j] + grid[i][j];
                col[i + 1][j] = col[i][j] + grid[i][j];
            }
        }

        for (int k = Math.min(m, n); k >= 2; k--) {
            for (int i = 0; i + k <= m; i++) {
                for (int j = 0; j + k <= n; j++) {

                    int target = row[i][j + k] - row[i][j];
                    boolean ok = true;

                    for (int r = i; r < i + k && ok; r++)
                        if (row[r][j + k] - row[r][j] != target)
                            ok = false;

                    for (int c = j; c < j + k && ok; c++)
                        if (col[i + k][c] - col[i][c] != target)
                            ok = false;

                    int d1 = 0, d2 = 0;
                    for (int x = 0; x < k; x++) {
                        d1 += grid[i + x][j + x];
                        d2 += grid[i + x][j + k - 1 - x];
                    }

                    if (ok && d1 == target && d2 == target)
                        return k;
                }
            }
        }
        return 1;
    }
}
