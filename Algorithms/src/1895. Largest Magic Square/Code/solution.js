var largestMagicSquare = function (grid) {
  const m = grid.length;
  const n = grid[0].length;

  const row = Array.from({ length: m }, () => Array(n + 1).fill(0));
  const col = Array.from({ length: m + 1 }, () => Array(n).fill(0));

  for (let i = 0; i < m; i++) {
    for (let j = 0; j < n; j++) {
      row[i][j + 1] = row[i][j] + grid[i][j];
      col[i + 1][j] = col[i][j] + grid[i][j];
    }
  }

  for (let k = Math.min(m, n); k >= 2; k--) {
    for (let i = 0; i + k <= m; i++) {
      for (let j = 0; j + k <= n; j++) {
        let target = row[i][j + k] - row[i][j];
        let ok = true;

        for (let r = i; r < i + k && ok; r++)
          if (row[r][j + k] - row[r][j] !== target) ok = false;

        for (let c = j; c < j + k && ok; c++)
          if (col[i + k][c] - col[i][c] !== target) ok = false;

        let d1 = 0,
          d2 = 0;
        for (let x = 0; x < k; x++) {
          d1 += grid[i + x][j + x];
          d2 += grid[i + x][j + k - 1 - x];
        }

        if (ok && d1 === target && d2 === target) return k;
      }
    }
  }
  return 1;
};
