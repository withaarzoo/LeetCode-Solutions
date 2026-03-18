var countSubmatrices = function (grid, k) {
  let m = grid.length,
    n = grid[0].length;
  let count = 0;

  for (let i = 0; i < m; i++) {
    for (let j = 0; j < n; j++) {
      if (i > 0) grid[i][j] += grid[i - 1][j];
      if (j > 0) grid[i][j] += grid[i][j - 1];
      if (i > 0 && j > 0) grid[i][j] -= grid[i - 1][j - 1];

      if (grid[i][j] <= k) count++;
    }
  }

  return count;
};
