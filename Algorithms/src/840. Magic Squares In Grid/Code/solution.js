/**
 * @param {number[][]} grid - A 2D array representing the grid of numbers.
 * @return {number} - The number of 3x3 magic squares found in the grid.
 */
var numMagicSquaresInside = function (grid) {
  // Get the number of rows and columns in the grid
  const rows = grid.length;
  const cols = grid[0].length;

  // Initialize a counter to keep track of the number of magic squares found
  let count = 0;

  /**
   * Helper function to check if the 3x3 grid starting at (r, c) is a magic square.
   * A magic square is a 3x3 grid where the numbers 1-9 appear exactly once
   * and the sums of the rows, columns, and diagonals all equal 15.
   *
   * @param {number} r - The starting row of the 3x3 grid.
   * @param {number} c - The starting column of the 3x3 grid.
   * @return {boolean} - True if the 3x3 grid is a magic square, otherwise false.
   */
  const isMagicSquare = (r, c) => {
    // Create an array to track the numbers 1-9 and ensure each appears exactly once
    const vals = Array(10).fill(false);

    // Iterate over the 3x3 grid starting at (r, c)
    for (let i = 0; i < 3; i++) {
      for (let j = 0; j < 3; j++) {
        const num = grid[r + i][c + j];
        // Check if the number is between 1 and 9 and hasn't appeared before
        if (num < 1 || num > 9 || vals[num]) return false;
        // Mark the number as seen
        vals[num] = true;
      }
    }

    // Check all rows, columns, and diagonals for a sum of 15
    return (
      grid[r][c] + grid[r][c + 1] + grid[r][c + 2] === 15 && // First row
      grid[r + 1][c] + grid[r + 1][c + 1] + grid[r + 1][c + 2] === 15 && // Second row
      grid[r + 2][c] + grid[r + 2][c + 1] + grid[r + 2][c + 2] === 15 && // Third row
      grid[r][c] + grid[r + 1][c] + grid[r + 2][c] === 15 && // First column
      grid[r][c + 1] + grid[r + 1][c + 1] + grid[r + 2][c + 1] === 15 && // Second column
      grid[r][c + 2] + grid[r + 1][c + 2] + grid[r + 2][c + 2] === 15 && // Third column
      grid[r][c] + grid[r + 1][c + 1] + grid[r + 2][c + 2] === 15 && // Diagonal from top-left to bottom-right
      grid[r][c + 2] + grid[r + 1][c + 1] + grid[r + 2][c] === 15
    ); // Diagonal from top-right to bottom-left
  };

  // Iterate over the grid to find all possible 3x3 subgrids
  for (let i = 0; i < rows - 2; i++) {
    for (let j = 0; j < cols - 2; j++) {
      // Check if the 3x3 grid starting at (i, j) is a magic square
      if (isMagicSquare(i, j)) {
        // Increment the count if a magic square is found
        count++;
      }
    }
  }

  // Return the total count of magic squares found
  return count;
};
