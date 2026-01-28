/**
 * @param {number[][]} grid
 * @param {number} k
 * @return {number}
 */
var minCost = function (grid, k) {
  const m = grid.length,
    n = grid[0].length;
  const INF = Number.MAX_SAFE_INTEGER / 4;

  // dp (no teleports)
  let dp = Array.from({ length: m }, () => Array(n).fill(INF));
  dp[0][0] = 0;
  for (let i = 0; i < m; i++) {
    for (let j = 0; j < n; j++) {
      if (i > 0) dp[i][j] = Math.min(dp[i][j], dp[i - 1][j] + grid[i][j]);
      if (j > 0) dp[i][j] = Math.min(dp[i][j], dp[i][j - 1] + grid[i][j]);
    }
  }

  // prepare sorted cells by value desc
  const cells = [];
  for (let i = 0; i < m; i++)
    for (let j = 0; j < n; j++) cells.push([grid[i][j], i, j]);
  cells.sort((a, b) => b[0] - a[0]);

  for (let step = 0; step < k; step++) {
    const start = Array.from({ length: m }, () => Array(n).fill(INF));
    let runningMin = INF;
    let idx = 0;
    while (idx < cells.length) {
      const val = cells[idx][0];
      let j = idx;
      let minGroup = INF;
      while (j < cells.length && cells[j][0] === val) {
        const [_, ii, jj] = cells[j];
        minGroup = Math.min(minGroup, dp[ii][jj]);
        j++;
      }
      runningMin = Math.min(runningMin, minGroup);
      for (let p = idx; p < j; p++) {
        const [_, ii, jj] = cells[p];
        start[ii][jj] = Math.min(dp[ii][jj], runningMin);
      }
      idx = j;
    }

    const dp2 = Array.from({ length: m }, () => Array(n).fill(INF));
    for (let i = 0; i < m; i++) {
      for (let j = 0; j < n; j++) {
        if (start[i][j] < dp2[i][j]) dp2[i][j] = start[i][j];
        if (i + 1 < m && dp2[i][j] < INF)
          dp2[i + 1][j] = Math.min(dp2[i + 1][j], dp2[i][j] + grid[i + 1][j]);
        if (j + 1 < n && dp2[i][j] < INF)
          dp2[i][j + 1] = Math.min(dp2[i][j + 1], dp2[i][j] + grid[i][j + 1]);
      }
    }
    dp = dp2;
  }

  return dp[m - 1][n - 1];
};
