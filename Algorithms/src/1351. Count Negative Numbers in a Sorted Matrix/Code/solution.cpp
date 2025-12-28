class Solution {
public:
    int countNegatives(vector<vector<int>>& grid) {
        int rows = grid.size();
        int cols = grid[0].size();
        
        int r = 0;
        int c = cols - 1;
        int count = 0;

        // Start from top-right corner
        while (r < rows && c >= 0) {
            if (grid[r][c] < 0) {
                // All elements below this are also negative
                count += (rows - r);
                c--; // move left
            } else {
                r++; // move down
            }
        }
        return count;
    }
};
