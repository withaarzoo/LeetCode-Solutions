class Solution {
public:
    int findMaxFish(vector<vector<int>>& grid) {
        int m = grid.size(), n = grid[0].size();
        vector<vector<bool>> visited(m, vector<bool>(n, false));
        int maxFish = 0;

        auto dfs = [&](int r, int c, auto& dfs) -> int {
            if (r < 0 || c < 0 || r >= m || c >= n || visited[r][c] || grid[r][c] == 0)
                return 0;
            visited[r][c] = true;
            int fish = grid[r][c];
            fish += dfs(r + 1, c, dfs);
            fish += dfs(r - 1, c, dfs);
            fish += dfs(r, c + 1, dfs);
            fish += dfs(r, c - 1, dfs);
            return fish;
        };

        for (int i = 0; i < m; ++i) {
            for (int j = 0; j < n; ++j) {
                if (!visited[i][j] && grid[i][j] > 0) {
                    maxFish = max(maxFish, dfs(i, j, dfs));
                }
            }
        }
        return maxFish;
    }
};
