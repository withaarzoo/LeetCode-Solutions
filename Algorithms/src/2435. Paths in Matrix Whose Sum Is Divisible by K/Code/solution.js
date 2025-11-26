/**
 * @param {number[][]} grid
 * @param {number} k
 * @return {number}
 */
var numberOfPaths = function (grid, k) {
  const MOD = 1_000_000_007;
  const m = grid.length;
  const n = grid[0].length;

  // Use normal numbers; sums are always reduced modulo MOD
  const makeRow = () => Array.from({ length: n }, () => Array(k).fill(0));
  let prev = makeRow();
  let cur = makeRow();

  for (let i = 0; i < m; i++) {
    // reset current row
    for (let j = 0; j < n; j++) cur[j].fill(0);

    for (let j = 0; j < n; j++) {
      const val = grid[i][j] % k;

      // starting cell
      if (i === 0 && j === 0) {
        cur[0][val] = 1;
        continue;
      }

      // from top
      if (i > 0) {
        for (let r = 0; r < k; r++) {
          if (prev[j][r] === 0) continue;
          const nr = (r + val) % k;
          cur[j][nr] = (cur[j][nr] + prev[j][r]) % MOD;
        }
      }

      // from left
      if (j > 0) {
        for (let r = 0; r < k; r++) {
          if (cur[j - 1][r] === 0) continue;
          const nr = (r + val) % k;
          cur[j][nr] = (cur[j][nr] + cur[j - 1][r]) % MOD;
        }
      }
    }

    // swap rows
    [prev, cur] = [cur, prev];
  }

  return prev[n - 1][0];
};
