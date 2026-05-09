class Solution
{
public:
    vector<vector<int>> rotateGrid(vector<vector<int>> &grid, int k)
    {
        int m = grid.size();
        int n = grid[0].size();

        // Number of layers in the matrix
        int layers = min(m, n) / 2;

        // Process every layer separately
        for (int layer = 0; layer < layers; layer++)
        {

            vector<int> nums;

            int top = layer;
            int bottom = m - layer - 1;
            int left = layer;
            int right = n - layer - 1;

            // Store top row
            for (int j = left; j <= right; j++)
            {
                nums.push_back(grid[top][j]);
            }

            // Store right column
            for (int i = top + 1; i <= bottom - 1; i++)
            {
                nums.push_back(grid[i][right]);
            }

            // Store bottom row
            for (int j = right; j >= left; j--)
            {
                nums.push_back(grid[bottom][j]);
            }

            // Store left column
            for (int i = bottom - 1; i >= top + 1; i--)
            {
                nums.push_back(grid[i][left]);
            }

            int len = nums.size();

            // Remove extra full rotations
            int rotate = k % len;

            // Rotated version of current layer
            vector<int> rotated(len);

            // Left rotation
            for (int i = 0; i < len; i++)
            {
                rotated[i] = nums[(i + rotate) % len];
            }

            int idx = 0;

            // Put values back into top row
            for (int j = left; j <= right; j++)
            {
                grid[top][j] = rotated[idx++];
            }

            // Put values back into right column
            for (int i = top + 1; i <= bottom - 1; i++)
            {
                grid[i][right] = rotated[idx++];
            }

            // Put values back into bottom row
            for (int j = right; j >= left; j--)
            {
                grid[bottom][j] = rotated[idx++];
            }

            // Put values back into left column
            for (int i = bottom - 1; i >= top + 1; i--)
            {
                grid[i][left] = rotated[idx++];
            }
        }

        return grid;
    }
};