class Solution {
    public int closedIsland(int[][] grid) {
        int m = grid.length, n = grid[0].length;
        for (int i = 0; i < m; i++) {
            if (grid[i][0] == 0)
                dfs(grid, i, 0);
            if (grid[i][n - 1] == 0)
                dfs(grid, i, n - 1);
        }
        for (int j = 0; j < n; j++) {
            if (grid[0][j] == 0)
                dfs(grid, 0, j);
            if (grid[m - 1][j] == 0)
                dfs(grid, m - 1, j);
        }
        int ans = 0;
        for (int i = 1; i < m - 1; i++) {
            for (int j = 1; j < n - 1; j++) {
                if (grid[i][j] == 0) {
                    ans++;
                    dfs(grid, i, j);
                }
            }
        }
        return ans;
    }

    private void dfs(int[][] grid, int i, int j) {
        if (i < 0 || i >= grid.length || j < 0 || j >= grid[0].length || grid[i][j] == 1)
            return;
        grid[i][j] = 1;
        dfs(grid, i + 1, j);
        dfs(grid, i - 1, j);
        dfs(grid, i, j + 1);
        dfs(grid, i, j - 1);
    }
}