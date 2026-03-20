/**
 * @param {number[][]} grid
 * @param {number} k
 * @return {number[][]}
 */
var minAbsDiff = function (grid, k) {
  const m = grid.length;
  const n = grid[0].length;
  const ans = Array.from({ length: m - k + 1 }, () => Array(n - k + 1).fill(0));

  for (let i = 0; i + k <= m; i++) {
    for (let j = 0; j + k <= n; j++) {
      const vals = [];

      // Collect all values from the current k x k submatrix
      for (let r = i; r < i + k; r++) {
        for (let c = j; c < j + k; c++) {
          vals.push(grid[r][c]);
        }
      }

      vals.sort((a, b) => a - b);

      let best = Infinity;

      // Check only consecutive different values
      for (let x = 1; x < vals.length; x++) {
        if (vals[x] !== vals[x - 1]) {
          best = Math.min(best, vals[x] - vals[x - 1]);
        }
      }

      ans[i][j] = best === Infinity ? 0 : best;
    }
  }

  return ans;
};
