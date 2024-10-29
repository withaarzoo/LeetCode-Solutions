/**
 * @param {number[][]} grid
 * @return {number}
 */
var maxMoves = function (grid) {
  const m = grid.length,
    n = grid[0].length;
  const dp = Array.from({ length: m }, () => Array(n).fill(0));
  let maxMoves = 0;

  for (let col = n - 2; col >= 0; col--) {
    for (let row = 0; row < m; row++) {
      if (row > 0 && grid[row][col] < grid[row - 1][col + 1]) {
        dp[row][col] = Math.max(dp[row][col], dp[row - 1][col + 1] + 1);
      }
      if (grid[row][col] < grid[row][col + 1]) {
        dp[row][col] = Math.max(dp[row][col], dp[row][col + 1] + 1);
      }
      if (row < m - 1 && grid[row][col] < grid[row + 1][col + 1]) {
        dp[row][col] = Math.max(dp[row][col], dp[row + 1][col + 1] + 1);
      }
    }
  }

  for (let row = 0; row < m; row++) {
    maxMoves = Math.max(maxMoves, dp[row][0]);
  }
  return maxMoves;
};
