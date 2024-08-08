/**
 * Generates a spiral order of coordinates in a matrix of given rows and columns,
 * starting from the specified initial position.
 *
 * @param {number} rows - The number of rows in the matrix.
 * @param {number} cols - The number of columns in the matrix.
 * @param {number} rStart - The starting row position.
 * @param {number} cStart - The starting column position.
 * @return {number[][]} - The list of coordinates in spiral order.
 */
var spiralMatrixIII = function (rows, cols, rStart, cStart) {
  // Initialize the result array to store the coordinates in spiral order.
  let result = [];

  // Define the four possible directions: right, down, left, up.
  // Each direction is represented as a pair of changes in row and column indices.
  let directions = [
    [0, 1],
    [1, 0],
    [0, -1],
    [-1, 0],
  ];

  // Start with 1 step to move in the current direction.
  let steps = 1;

  // Initialize the current direction index to 0 (right).
  let d = 0;

  // Set the initial position to the starting row and column.
  let r = rStart,
    c = cStart;

  // Add the starting position to the result.
  result.push([r, c]);

  // Loop until we have filled all cells in the matrix.
  while (result.length < rows * cols) {
    // Perform the movement in the current direction twice.
    for (let i = 0; i < 2; ++i) {
      // Move 'steps' times in the current direction.
      for (let j = 0; j < steps; ++j) {
        // Update the row and column based on the current direction.
        r += directions[d][0];
        c += directions[d][1];

        // Check if the new position is within the bounds of the matrix.
        if (r >= 0 && r < rows && c >= 0 && c < cols) {
          // If within bounds, add the position to the result.
          result.push([r, c]);
        }
      }
      // Change the direction to the next one in the sequence.
      d = (d + 1) % 4;
    }
    // Increase the number of steps after completing two directions.
    ++steps;
  }

  // Return the list of coordinates in spiral order.
  return result;
};
