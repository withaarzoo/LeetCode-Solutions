/**
 * @param {number[][]} grid
 * @param {number} k
 * @return {number}
 */
var maxPathScore = function (grid, k) {
  const m = grid.length; // Number of rows in the grid.
  const n = grid[0].length; // Number of columns in the grid.
  const NEG = -1e9; // Sentinel for impossible states.

  // prev[j][c] = best score at column j in the previous row using exact cost c.
  let prev = Array.from({ length: n }, () => Array(k + 1).fill(NEG));

  for (let i = 0; i < m; i++) {
    // Rebuild the current row from scratch so old values do not leak into new answers.
    let curr = Array.from({ length: n }, () => Array(k + 1).fill(NEG));

    for (let j = 0; j < n; j++) {
      const gain = grid[i][j]; // Score gained by taking this cell.
      const need = gain > 0 ? 1 : 0; // Cost spent by this cell: 0 for 0, 1 for 1/2.

      // A path to (i, j) cannot spend more than i + j budget points.
      const limit = Math.min(k, i + j);

      // The start cell is fixed and always has value 0.
      if (i === 0 && j === 0) {
        curr[0][0] = 0; // Zero score, zero cost at the start.
        continue;
      }

      for (let c = need; c <= limit; c++) {
        let best = NEG;

        // From above: use the completed previous row.
        if (i > 0 && prev[j][c - need] !== NEG) {
          best = Math.max(best, prev[j][c - need] + gain);
        }

        // From left: use the state already computed in this row.
        if (j > 0 && curr[j - 1][c - need] !== NEG) {
          best = Math.max(best, curr[j - 1][c - need] + gain);
        }

        curr[j][c] = best; // Store the best exact-cost value for this cell.
      }
    }

    prev = curr; // Move the current row into prev for the next iteration.
  }

  let ans = NEG; // Best score among all allowed costs at the end.
  for (let c = 0; c <= k; c++) {
    ans = Math.max(ans, prev[n - 1][c]);
  }

  return ans < 0 ? -1 : ans; // If nothing is reachable, return -1.
};
