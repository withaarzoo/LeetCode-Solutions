/**
 * @param {number[][]} grid
 * @return {number}
 */
var countNegatives = function (grid) {
  let rows = grid.length;
  let cols = grid[0].length;

  let r = 0;
  let c = cols - 1;
  let count = 0;

  // Start from top-right corner
  while (r < rows && c >= 0) {
    if (grid[r][c] < 0) {
      count += rows - r;
      c--; // move left
    } else {
      r++; // move down
    }
  }
  return count;
};
