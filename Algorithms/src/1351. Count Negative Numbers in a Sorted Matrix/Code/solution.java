class Solution {
    public int countNegatives(int[][] grid) {
        int rows = grid.length;
        int cols = grid[0].length;

        int r = 0;
        int c = cols - 1;
        int count = 0;

        // Start from top-right corner
        while (r < rows && c >= 0) {
            if (grid[r][c] < 0) {
                count += (rows - r);
                c--; // move left
            } else {
                r++; // move down
            }
        }
        return count;
    }
}
