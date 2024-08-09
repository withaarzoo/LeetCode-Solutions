class Solution {
    public int numMagicSquaresInside(int[][] grid) {
        // Get the number of rows and columns in the grid
        int rows = grid.length;
        int cols = grid[0].length;

        // Initialize a counter to keep track of the number of magic squares found
        int count = 0;

        // Loop through the grid, but only up to rows-2 and cols-2 since a 3x3 square
        // needs space
        for (int i = 0; i < rows - 2; i++) {
            for (int j = 0; j < cols - 2; j++) {
                // Check if the 3x3 grid starting at (i, j) is a magic square
                if (isMagicSquare(grid, i, j)) {
                    // If it is, increment the counter
                    count++;
                }
            }
        }
        // Return the total count of magic squares found
        return count;
    }

    private boolean isMagicSquare(int[][] grid, int r, int c) {
        // Create a boolean array to track numbers 1-9 and ensure each appears exactly
        // once
        boolean[] vals = new boolean[10];

        // Iterate over the 3x3 grid starting from (r, c)
        for (int i = 0; i < 3; i++) {
            for (int j = 0; j < 3; j++) {
                // Extract the current number from the grid
                int num = grid[r + i][c + j];

                // If the number is out of bounds (not between 1 and 9) or already seen, it's
                // not a magic square
                if (num < 1 || num > 9 || vals[num])
                    return false;

                // Mark the number as seen
                vals[num] = true;
            }
        }

        // Check the sum of all rows, columns, and diagonals; they all must equal 15 in
        // a 3x3 magic square
        return (grid[r][c] + grid[r][c + 1] + grid[r][c + 2] == 15 && // Top row
                grid[r + 1][c] + grid[r + 1][c + 1] + grid[r + 1][c + 2] == 15 && // Middle row
                grid[r + 2][c] + grid[r + 2][c + 1] + grid[r + 2][c + 2] == 15 && // Bottom row
                grid[r][c] + grid[r + 1][c] + grid[r + 2][c] == 15 && // Left column
                grid[r][c + 1] + grid[r + 1][c + 1] + grid[r + 2][c + 1] == 15 && // Middle column
                grid[r][c + 2] + grid[r + 1][c + 2] + grid[r + 2][c + 2] == 15 && // Right column
                grid[r][c] + grid[r + 1][c + 1] + grid[r + 2][c + 2] == 15 && // Main diagonal
                grid[r][c + 2] + grid[r + 1][c + 1] + grid[r + 2][c] == 15); // Secondary diagonal
    }
}
