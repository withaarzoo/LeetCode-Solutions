function minCost(grid: number[][]): number {
  const m = grid.length,
    n = grid[0].length;
  const dp: Uint8Array[][] = Array.from({ length: m }, () =>
    Array.from({ length: n }, () => new Uint8Array(1024)),
  );
  dp[0][0][grid[0][0]] = 1;
  for (let i = 0; i < m; i++) {
    for (let j = 0; j < n; j++) {
      if (i > 0) {
        for (let x = 0; x < 1024; x++) {
          if (dp[i - 1][j][x]) dp[i][j][x ^ grid[i][j]] = 1;
        }
      }
      if (j > 0) {
        for (let x = 0; x < 1024; x++) {
          if (dp[i][j - 1][x]) dp[i][j][x ^ grid[i][j]] = 1;
        }
      }
    }
  }
  for (let x = 0; x < 1024; x++) {
    if (dp[m - 1][n - 1][x]) return x;
  }
  return 0;
}
