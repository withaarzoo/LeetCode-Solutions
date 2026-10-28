class Solution
{
public:
    int minCost(vector<vector<int>> &grid)
    {
        int m = grid.size(), n = grid[0].size();
        vector<vector<bitset<1024>>> dp(m, vector<bitset<1024>>(n));
        dp[0][0].set(grid[0][0]);
        for (int i = 0; i < m; i++)
        {
            for (int j = 0; j < n; j++)
            {
                if (i > 0)
                {
                    for (int x = 0; x < 1024; x++)
                    {
                        if (dp[i - 1][j][x])
                            dp[i][j].set(x ^ grid[i][j]);
                    }
                }
                if (j > 0)
                {
                    for (int x = 0; x < 1024; x++)
                    {
                        if (dp[i][j - 1][x])
                            dp[i][j].set(x ^ grid[i][j]);
                    }
                }
            }
        }
        for (int x = 0; x < 1024; x++)
        {
            if (dp[m - 1][n - 1][x])
                return x;
        }
        return 0;
    }
};