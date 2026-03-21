class Solution {
    public int[][] reverseSubmatrix(int[][] grid, int x, int y, int k) {
        // Loop through half rows
        for (int i = 0; i < k / 2; i++) {
            int top = x + i;
            int bottom = x + k - 1 - i;

            // Swap column-wise
            for (int j = 0; j < k; j++) {
                int temp = grid[top][y + j];
                grid[top][y + j] = grid[bottom][y + j];
                grid[bottom][y + j] = temp;
            }
        }
        return grid;
    }
}