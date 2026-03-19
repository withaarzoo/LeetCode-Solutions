var numberOfSubmatrices = function (grid) {
  let n = grid.length,
    m = grid[0].length;

  let sum = Array.from({ length: 2 }, () => Array(m + 1).fill(0));
  let countX = Array.from({ length: 2 }, () => Array(m + 1).fill(0));

  let ans = 0;

  for (let i = 0; i < n; i++) {
    let cur = i % 2;
    let prev = 1 - cur;

    for (let j = 0; j < m; j++) {
      let val = grid[i][j] === "X" ? 1 : grid[i][j] === "Y" ? -1 : 0;
      let isX = grid[i][j] === "X" ? 1 : 0;

      sum[cur][j + 1] = val + sum[cur][j] + sum[prev][j + 1] - sum[prev][j];

      countX[cur][j + 1] =
        isX + countX[cur][j] + countX[prev][j + 1] - countX[prev][j];

      if (sum[cur][j + 1] === 0 && countX[cur][j + 1] > 0) {
        ans++;
      }
    }
  }

  return ans;
};
