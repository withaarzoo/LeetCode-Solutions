class Solution
{
public:
    int numMagicSquaresInside(vector<vector<int>> &grid)
    {
        int rows = grid.size();
        int cols = grid[0].size();
        int count = 0;

        // Traverse every possible 3x3 subgrid
        for (int i = 0; i + 2 < rows; i++)
        {
            for (int j = 0; j + 2 < cols; j++)
            {
                if (isMagic(grid, i, j))
                {
                    count++;
                }
            }
        }
        return count;
    }

private:
    bool isMagic(vector<vector<int>> &g, int r, int c)
    {
        // Center must be 5
        if (g[r + 1][c + 1] != 5)
            return false;

        bool seen[10] = {false};

        // Check numbers 1 to 9 and uniqueness
        for (int i = r; i < r + 3; i++)
        {
            for (int j = c; j < c + 3; j++)
            {
                int val = g[i][j];
                if (val < 1 || val > 9 || seen[val])
                    return false;
                seen[val] = true;
            }
        }

        // Check rows and columns
        for (int i = 0; i < 3; i++)
        {
            if (g[r + i][c] + g[r + i][c + 1] + g[r + i][c + 2] != 15)
                return false;
            if (g[r][c + i] + g[r + 1][c + i] + g[r + 2][c + i] != 15)
                return false;
        }

        // Check diagonals
        if (g[r][c] + g[r + 1][c + 1] + g[r + 2][c + 2] != 15)
            return false;
        if (g[r][c + 2] + g[r + 1][c + 1] + g[r + 2][c] != 15)
            return false;

        return true;
    }
};
