class Solution {
public:
    vector<vector<int>> reverseSubmatrix(vector<vector<int>>& grid, int x, int y, int k) {
        // Loop through half of the rows of the square
        for (int i = 0; i < k / 2; i++) {
            int top = x + i;
            int bottom = x + k - 1 - i;

            // Swap elements column-wise
            for (int j = 0; j < k; j++) {
                swap(grid[top][y + j], grid[bottom][y + j]);
            }
        }
        return grid;
    }
};