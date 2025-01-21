/**
 * @param {number[][]} grid
 * @return {number}
 */
var gridGame = function (grid) {
  const n = grid[0].length;
  const topSuffix = Array(n).fill(0);
  const bottomPrefix = Array(n).fill(0);

  // Calculate suffix sum for the top row
  topSuffix[n - 1] = grid[0][n - 1];
  for (let i = n - 2; i >= 0; --i) {
    topSuffix[i] = topSuffix[i + 1] + grid[0][i];
  }

  // Calculate prefix sum for the bottom row
  bottomPrefix[0] = grid[1][0];
  for (let i = 1; i < n; ++i) {
    bottomPrefix[i] = bottomPrefix[i - 1] + grid[1][i];
  }

  // Find the minimum maximum points Robot 2 can collect
  let result = Infinity;
  for (let i = 0; i < n; ++i) {
    const top = i + 1 < n ? topSuffix[i + 1] : 0;
    const bottom = i > 0 ? bottomPrefix[i - 1] : 0;
    result = Math.min(result, Math.max(top, bottom));
  }

  return result;
};
