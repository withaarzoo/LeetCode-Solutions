class Solution
{
public:
    vector<int> getBiggestThree(vector<vector<int>> &grid)
    {
        int m = grid.size();
        int n = grid[0].size();

        set<int> sums;

        for (int r = 0; r < m; r++)
        {
            for (int c = 0; c < n; c++)
            {

                // size 0 rhombus (single cell)
                sums.insert(grid[r][c]);

                int maxSize = min({r, c, m - 1 - r, n - 1 - c});

                for (int k = 1; k <= maxSize; k++)
                {
                    int sum = 0;

                    int x = r - k, y = c;

                    // top -> right
                    for (int i = 0; i < k; i++)
                    {
                        sum += grid[x + i][y + i];
                    }

                    // right -> bottom
                    for (int i = 0; i < k; i++)
                    {
                        sum += grid[r + i][c + k - i];
                    }

                    // bottom -> left
                    for (int i = 0; i < k; i++)
                    {
                        sum += grid[r + k - i][c - i];
                    }

                    // left -> top
                    for (int i = 0; i < k; i++)
                    {
                        sum += grid[r - i][c - k + i];
                    }

                    sums.insert(sum);
                }
            }
        }

        vector<int> res(sums.begin(), sums.end());
        sort(res.rbegin(), res.rend());

        if (res.size() > 3)
            res.resize(3);

        return res;
    }
};