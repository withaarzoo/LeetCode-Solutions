/**
 * @param {number[][]} grid
 * @param {number} k
 * @return {number[][]}
 */
var rotateGrid = function (grid, k) {
  let m = grid.length;
  let n = grid[0].length;

  // Number of layers
  let layers = Math.min(m, n) / 2;

  for (let layer = 0; layer < layers; layer++) {
    let nums = [];

    let top = layer;
    let bottom = m - layer - 1;
    let left = layer;
    let right = n - layer - 1;

    // Store top row
    for (let j = left; j <= right; j++) {
      nums.push(grid[top][j]);
    }

    // Store right column
    for (let i = top + 1; i <= bottom - 1; i++) {
      nums.push(grid[i][right]);
    }

    // Store bottom row
    for (let j = right; j >= left; j--) {
      nums.push(grid[bottom][j]);
    }

    // Store left column
    for (let i = bottom - 1; i >= top + 1; i--) {
      nums.push(grid[i][left]);
    }

    let len = nums.length;

    // Effective rotations only
    let rotate = k % len;

    let rotated = new Array(len);

    // Left rotation
    for (let i = 0; i < len; i++) {
      rotated[i] = nums[(i + rotate) % len];
    }

    let idx = 0;

    // Fill top row
    for (let j = left; j <= right; j++) {
      grid[top][j] = rotated[idx++];
    }

    // Fill right column
    for (let i = top + 1; i <= bottom - 1; i++) {
      grid[i][right] = rotated[idx++];
    }

    // Fill bottom row
    for (let j = right; j >= left; j--) {
      grid[bottom][j] = rotated[idx++];
    }

    // Fill left column
    for (let i = bottom - 1; i >= top + 1; i--) {
      grid[i][left] = rotated[idx++];
    }
  }

  return grid;
};
