class Solution
{
public:
    int numberOfPaths(vector<vector<int>> &grid, int k)
    {
        const int MOD = 1'000'000'007;
        int m = grid.size();
        int n = grid[0].size();

        // prev[j][r] = number of paths to cell (i-1, j) with sum % k == r
        // cur[j][r]  = number of paths to cell (i,   j) with sum % k == r
        vector<vector<int>> prev(n, vector<int>(k, 0));
        vector<vector<int>> cur(n, vector<int>(k, 0));

        for (int i = 0; i < m; ++i)
        {
            // reset current row for each new i
            for (int j = 0; j < n; ++j)
                fill(cur[j].begin(), cur[j].end(), 0);

            for (int j = 0; j < n; ++j)
            {
                int val = grid[i][j] % k;

                // starting cell
                if (i == 0 && j == 0)
                {
                    cur[0][val] = 1;
                    continue;
                }

                // get paths from top
                if (i > 0)
                {
                    for (int r = 0; r < k; ++r)
                    {
                        if (prev[j][r] == 0)
                            continue;
                        int nr = (r + val) % k;
                        cur[j][nr] = (cur[j][nr] + prev[j][r]) % MOD;
                    }
                }

                // get paths from left
                if (j > 0)
                {
                    for (int r = 0; r < k; ++r)
                    {
                        if (cur[j - 1][r] == 0)
                            continue;
                        int nr = (r + val) % k;
                        cur[j][nr] = (cur[j][nr] + cur[j - 1][r]) % MOD;
                    }
                }
            }

            // current row becomes previous row for next iteration
            prev.swap(cur);
        }

        // bottom-right cell in last processed row
        return prev[n - 1][0];
    }
};
