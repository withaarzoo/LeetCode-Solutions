class Solution
{
public:
    int largestMagicSquare(vector<vector<int>> &grid)
    {
        int m = grid.size();
        int n = grid[0].size();

        // Row and column prefix sums
        vector<vector<int>> row(m, vector<int>(n + 1, 0));
        vector<vector<int>> col(m + 1, vector<int>(n, 0));

        for (int i = 0; i < m; i++)
        {
            for (int j = 0; j < n; j++)
            {
                row[i][j + 1] = row[i][j] + grid[i][j];
                col[i + 1][j] = col[i][j] + grid[i][j];
            }
        }

        int maxSize = min(m, n);

        for (int k = maxSize; k >= 2; k--)
        {
            for (int i = 0; i + k <= m; i++)
            {
                for (int j = 0; j + k <= n; j++)
                {

                    int target = row[i][j + k] - row[i][j];
                    bool ok = true;

                    // Check rows
                    for (int r = i; r < i + k && ok; r++)
                    {
                        if (row[r][j + k] - row[r][j] != target)
                            ok = false;
                    }

                    // Check columns
                    for (int c = j; c < j + k && ok; c++)
                    {
                        if (col[i + k][c] - col[i][c] != target)
                            ok = false;
                    }

                    // Diagonals
                    int d1 = 0, d2 = 0;
                    for (int x = 0; x < k; x++)
                    {
                        d1 += grid[i + x][j + x];
                        d2 += grid[i + x][j + k - 1 - x];
                    }

                    if (ok && d1 == target && d2 == target)
                        return k;
                }
            }
        }
        return 1;
    }
};
