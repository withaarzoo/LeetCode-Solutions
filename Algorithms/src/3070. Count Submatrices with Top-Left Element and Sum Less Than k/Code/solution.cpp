class Solution
{
public:
    int countSubmatrices(vector<vector<int>> &grid, int k)
    {
        int m = grid.size(), n = grid[0].size();
        int count = 0;

        // Convert grid into prefix sum matrix
        for (int i = 0; i < m; i++)
        {
            for (int j = 0; j < n; j++)
            {

                // Add top
                if (i > 0)
                    grid[i][j] += grid[i - 1][j];

                // Add left
                if (j > 0)
                    grid[i][j] += grid[i][j - 1];

                // Remove double counted area
                if (i > 0 && j > 0)
                    grid[i][j] -= grid[i - 1][j - 1];

                // Check condition
                if (grid[i][j] <= k)
                    count++;
            }
        }

        return count;
    }
};