/**
 * @param {number[][]} grid
 * @param {number} k
 * @return {number[][]}
 */
var shiftGrid = function (grid, k) {
  // Get grid dimensions
  const m = grid.length;
  const n = grid[0].length;

  // Total elements
  const total = m * n;

  // Remove unnecessary complete rotations
  k %= total;

  // Create the answer grid
  const ans = Array.from({ length: m }, () => Array(n));

  // Move every element directly
  for (let i = 0; i < m; i++) {
    for (let j = 0; j < n; j++) {
      // Flatten the current position
      const oldIndex = i * n + j;

      // New position after shifting
      const newIndex = (oldIndex + k) % total;

      // Convert back to row and column
      const newRow = Math.floor(newIndex / n);
      const newCol = newIndex % n;

      // Place the value
      ans[newRow][newCol] = grid[i][j];
    }
  }

  return ans;
};
