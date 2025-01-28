class Solution {
    public int findMaxFish(int[][] grid) {
        int m = grid.length, n = grid[0].length;
        boolean[][] visited = new boolean[m][n];
        int maxFish = 0;

        int dfs(int r, int c) {
            if (r < 0 || c < 0 || r >= m || c >= n || visited[r][c] || grid[r][c] == 0)
                return 0;
            visited[r][c] = true;
            int fish = grid[r][c];
            fish += dfs(r + 1, c);
            fish += dfs(r - 1, c);
            fish += dfs(r, c + 1);
            fish += dfs(r, c - 1);
            return fish;
        }

        for (int i = 0; i < m; i++) {
            for (int j = 0; j < n; j++) {
                if (!visited[i][j] && grid[i][j] > 0) {
                    maxFish = Math.max(maxFish, dfs(i, j));
                }
            }
        }
        return maxFish;
    }
}
