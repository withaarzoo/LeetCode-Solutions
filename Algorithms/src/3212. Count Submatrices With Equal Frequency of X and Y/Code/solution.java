class Solution {
    public int numberOfSubmatrices(char[][] grid) {
        int n = grid.length, m = grid[0].length;

        int[][] sum = new int[2][m + 1];
        int[][] countX = new int[2][m + 1];

        int ans = 0;

        for (int i = 0; i < n; i++) {
            int cur = i % 2;
            int prev = 1 - cur;

            for (int j = 0; j < m; j++) {
                int val = (grid[i][j] == 'X') ? 1 : (grid[i][j] == 'Y' ? -1 : 0);
                int isX = (grid[i][j] == 'X') ? 1 : 0;

                sum[cur][j + 1] = val
                        + sum[cur][j]
                        + sum[prev][j + 1]
                        - sum[prev][j];

                countX[cur][j + 1] = isX
                        + countX[cur][j]
                        + countX[prev][j + 1]
                        - countX[prev][j];

                if (sum[cur][j + 1] == 0 && countX[cur][j + 1] > 0) {
                    ans++;
                }
            }
        }

        return ans;
    }
}