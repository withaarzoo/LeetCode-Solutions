class Solution
{
public:
    vector<vector<int>> shiftGrid(vector<vector<int>> &grid, int k)
    {
        // Get the grid dimensions
        int m = grid.size();
        int n = grid[0].size();

        // Total number of elements
        int total = m * n;

        // Extra full rotations do not change the grid
        k %= total;

        // Create the answer grid
        vector<vector<int>> ans(m, vector<int>(n));

        // Move every element directly to its final position
        for (int i = 0; i < m; i++)
        {
            for (int j = 0; j < n; j++)
            {

                // Current position in the flattened array
                int oldIndex = i * n + j;

                // Position after shifting k times
                int newIndex = (oldIndex + k) % total;

                // Convert back to row and column
                int newRow = newIndex / n;
                int newCol = newIndex % n;

                // Place the element
                ans[newRow][newCol] = grid[i][j];
            }
        }

        // Return the shifted grid
        return ans;
    }
};