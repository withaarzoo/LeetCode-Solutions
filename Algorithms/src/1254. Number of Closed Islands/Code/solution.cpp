class Solution
{
public:
    int closedIsland(vector<vector<int>> &grid)
    {
        int m = grid.size(), n = grid[0].size();
        auto dfs = [&](auto &&self, int i, int j) -> void
        {
            if (i < 0 || i >= m || j < 0 || j >= n || grid[i][j] == 1)
                return;
            grid[i][j] = 1;
            self(self, i + 1, j);
            self(self, i - 1, j);
            self(self, i, j + 1);
            self(self, i, j - 1);
        };
        for (int i = 0; i < m; ++i)
        {
            if (grid[i][0] == 0)
                dfs(dfs, i, 0);
            if (grid[i][n - 1] == 0)
                dfs(dfs, i, n - 1);
        }
        for (int j = 0; j < n; ++j)
        {
            if (grid[0][j] == 0)
                dfs(dfs, 0, j);
            if (grid[m - 1][j] == 0)
                dfs(dfs, m - 1, j);
        }
        int ans = 0;
        for (int i = 1; i < m - 1; ++i)
        {
            for (int j = 1; j < n - 1; ++j)
            {
                if (grid[i][j] == 0)
                {
                    ++ans;
                    dfs(dfs, i, j);
                }
            }
        }
        return ans;
    }
};