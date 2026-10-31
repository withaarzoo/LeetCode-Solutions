/**
 * @param {number[][]} grid
 * @return {number}
 */
var uniquePaths = function (grid) {
  const m = grid.length,
    n = grid[0].length;
  const MOD = 1000000007;
  const landR = Array.from({ length: m }, () =>
    Array.from({ length: n }, () => Array(2)),
  );
  const landC = Array.from({ length: m }, () =>
    Array.from({ length: n }, () => Array(2)),
  );
  for (let x = m - 1; x >= 0; --x) {
    for (let y = n - 1; y >= 0; --y) {
      for (let d = 0; d < 2; ++d) {
        if (grid[x][y] === 0) {
          landR[x][y][d] = x;
          landC[x][y][d] = y;
        } else {
          const nx = d === 0 ? x + 1 : x;
          const ny = d === 0 ? y : y + 1;
          const nd = 1 - d;
          if (nx < 0 || nx >= m || ny < 0 || ny >= n) {
            landR[x][y][d] = -1;
            landC[x][y][d] = -1;
          } else {
            landR[x][y][d] = landR[nx][ny][nd];
            landC[x][y][d] = landC[nx][ny][nd];
          }
        }
      }
    }
  }
  const dp = Array.from({ length: m }, () => Array(n).fill(0));
  dp[0][0] = 1;
  for (let i = 0; i < m; ++i) {
    for (let j = 0; j < n; ++j) {
      if (dp[i][j] === 0) continue;
      if (j + 1 < n) {
        const nr = landR[i][j + 1][0];
        const nc = landC[i][j + 1][0];
        if (nr !== -1) dp[nr][nc] = (dp[nr][nc] + dp[i][j]) % MOD;
      }
      if (i + 1 < m) {
        const nr = landR[i + 1][j][1];
        const nc = landC[i + 1][j][1];
        if (nr !== -1) dp[nr][nc] = (dp[nr][nc] + dp[i][j]) % MOD;
      }
    }
  }
  return dp[m - 1][n - 1];
};
