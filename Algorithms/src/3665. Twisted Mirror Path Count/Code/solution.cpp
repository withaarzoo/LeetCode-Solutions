class Solution
{
public:
    int uniquePaths(vector<vector<int>> &grid)
    {
        int m = grid.size(), n = grid[0].size();
        const int MOD = 1e9 + 7;
        vector<vector<array<int, 2>>> landR(m, vector<array<int, 2>>(n));
        vector<vector<array<int, 2>>> landC(m, vector<array<int, 2>>(n));
        for (int x = m - 1; x >= 0; --x)
        {
            for (int y = n - 1; y >= 0; --y)
            {
                for (int d = 0; d < 2; ++d)
                {
                    if (grid[x][y] == 0)
                    {
                        landR[x][y][d] = x;
                        landC[x][y][d] = y;
                    }
                    else
                    {
                        int nx = (d == 0) ? x + 1 : x;
                        int ny = (d == 0) ? y : y + 1;
                        int nd = 1 - d;
                        if (nx < 0 || nx >= m || ny < 0 || ny >= n)
                        {
                            landR[x][y][d] = -1;
                            landC[x][y][d] = -1;
                        }
                        else
                        {
                            landR[x][y][d] = landR[nx][ny][nd];
                            landC[x][y][d] = landC[nx][ny][nd];
                        }
                    }
                }
            }
        }
        vector<vector<int>> dp(m, vector<int>(n, 0));
        dp[0][0] = 1;
        for (int i = 0; i < m; ++i)
        {
            for (int j = 0; j < n; ++j)
            {
                if (dp[i][j] == 0)
                    continue;
                if (j + 1 < n)
                {
                    int nr = landR[i][j + 1][0];
                    int nc = landC[i][j + 1][0];
                    if (nr != -1)
                        dp[nr][nc] = (dp[nr][nc] + dp[i][j]) % MOD;
                }
                if (i + 1 < m)
                {
                    int nr = landR[i + 1][j][1];
                    int nc = landC[i + 1][j][1];
                    if (nr != -1)
                        dp[nr][nc] = (dp[nr][nc] + dp[i][j]) % MOD;
                }
            }
        }
        return dp[m - 1][n - 1];
    }
};