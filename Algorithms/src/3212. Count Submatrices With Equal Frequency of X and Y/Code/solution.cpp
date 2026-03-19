class Solution
{
public:
    int numberOfSubmatrices(vector<vector<char>> &grid)
    {
        int n = grid.size(), m = grid[0].size();

        // Rolling arrays to save space
        vector<vector<int>> sum(2, vector<int>(m + 1, 0));
        vector<vector<int>> countX(2, vector<int>(m + 1, 0));

        int ans = 0;

        for (int i = 0; i < n; i++)
        {
            int cur = i % 2;
            int prev = 1 - cur;

            for (int j = 0; j < m; j++)
            {
                int val = (grid[i][j] == 'X') ? 1 : (grid[i][j] == 'Y' ? -1 : 0);
                int isX = (grid[i][j] == 'X') ? 1 : 0;

                // 2D prefix sum formula
                sum[cur][j + 1] = val + sum[cur][j] + sum[prev][j + 1] - sum[prev][j];

                countX[cur][j + 1] = isX + countX[cur][j] + countX[prev][j + 1] - countX[prev][j];

                // Check condition
                if (sum[cur][j + 1] == 0 && countX[cur][j + 1] > 0)
                {
                    ans++;
                }
            }
        }

        return ans;
    }
};