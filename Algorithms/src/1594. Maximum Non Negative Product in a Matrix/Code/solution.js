var maxProductPath = function (grid) {
  const m = grid.length,
    n = grid[0].length;
  const MOD = 1e9 + 7;

  const maxDp = Array.from({ length: m }, () => Array(n).fill(0));
  const minDp = Array.from({ length: m }, () => Array(n).fill(0));

  maxDp[0][0] = minDp[0][0] = grid[0][0];

  for (let i = 1; i < m; i++) {
    maxDp[i][0] = minDp[i][0] = maxDp[i - 1][0] * grid[i][0];
  }

  for (let j = 1; j < n; j++) {
    maxDp[0][j] = minDp[0][j] = maxDp[0][j - 1] * grid[0][j];
  }

  for (let i = 1; i < m; i++) {
    for (let j = 1; j < n; j++) {
      let val = grid[i][j];

      let a = maxDp[i - 1][j] * val;
      let b = minDp[i - 1][j] * val;
      let c = maxDp[i][j - 1] * val;
      let d = minDp[i][j - 1] * val;

      maxDp[i][j] = Math.max(a, b, c, d);
      minDp[i][j] = Math.min(a, b, c, d);
    }
  }

  let res = maxDp[m - 1][n - 1];
  return res < 0 ? -1 : res % MOD;
};
