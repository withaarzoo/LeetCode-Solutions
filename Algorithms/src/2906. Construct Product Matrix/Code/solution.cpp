class Solution
{
public:
    vector<vector<int>> constructProductMatrix(vector<vector<int>> &grid)
    {
        const int MOD = 12345;
        int n = grid.size();
        int m = grid[0].size();

        vector<vector<int>> ans(n, vector<int>(m, 1));

        long long prefix = 1;
        for (int i = 0; i < n; i++)
        {
            for (int j = 0; j < m; j++)
            {
                ans[i][j] = (int)prefix;
                prefix = (prefix * grid[i][j]) % MOD;
            }
        }

        long long suffix = 1;
        for (int i = n - 1; i >= 0; i--)
        {
            for (int j = m - 1; j >= 0; j--)
            {
                ans[i][j] = (int)((1LL * ans[i][j] * suffix) % MOD);
                suffix = (suffix * grid[i][j]) % MOD;
            }
        }

        return ans;
    }
};