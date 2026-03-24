/**
 * @param {number[][]} grid
 * @return {number[][]}
 */
var constructProductMatrix = function (grid) {
  const MOD = 12345;
  const n = grid.length;
  const m = grid[0].length;

  const ans = Array.from({ length: n }, () => Array(m).fill(1));

  let prefix = 1;
  for (let i = 0; i < n; i++) {
    for (let j = 0; j < m; j++) {
      ans[i][j] = prefix;
      prefix = (prefix * grid[i][j]) % MOD;
    }
  }

  let suffix = 1;
  for (let i = n - 1; i >= 0; i--) {
    for (let j = m - 1; j >= 0; j--) {
      ans[i][j] = (ans[i][j] * suffix) % MOD;
      suffix = (suffix * grid[i][j]) % MOD;
    }
  }

  return ans;
};
