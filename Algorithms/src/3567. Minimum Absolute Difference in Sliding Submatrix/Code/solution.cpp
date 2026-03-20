class Solution
{
public:
    vector<vector<int>> minAbsDiff(vector<vector<int>> &grid, int k)
    {
        int m = grid.size();
        int n = grid[0].size();
        vector<vector<int>> ans(m - k + 1, vector<int>(n - k + 1, 0));

        for (int i = 0; i + k <= m; i++)
        {
            for (int j = 0; j + k <= n; j++)
            {
                vector<int> vals;
                vals.reserve(k * k);

                // Collect all values from the current k x k submatrix
                for (int r = i; r < i + k; r++)
                {
                    for (int c = j; c < j + k; c++)
                    {
                        vals.push_back(grid[r][c]);
                    }
                }

                sort(vals.begin(), vals.end());

                int best = INT_MAX;

                // Check only consecutive different values
                for (int x = 1; x < (int)vals.size(); x++)
                {
                    if (vals[x] != vals[x - 1])
                    {
                        best = min(best, vals[x] - vals[x - 1]);
                    }
                }

                ans[i][j] = (best == INT_MAX ? 0 : best);
            }
        }

        return ans;
    }
};