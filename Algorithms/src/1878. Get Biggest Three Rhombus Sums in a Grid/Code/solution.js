var getBiggestThree = function (grid) {
  let m = grid.length;
  let n = grid[0].length;
  let set = new Set();

  for (let r = 0; r < m; r++) {
    for (let c = 0; c < n; c++) {
      set.add(grid[r][c]);

      let maxSize = Math.min(r, c, m - 1 - r, n - 1 - c);

      for (let k = 1; k <= maxSize; k++) {
        let sum = 0;

        for (let i = 0; i < k; i++) sum += grid[r - k + i][c + i];

        for (let i = 0; i < k; i++) sum += grid[r + i][c + k - i];

        for (let i = 0; i < k; i++) sum += grid[r + k - i][c - i];

        for (let i = 0; i < k; i++) sum += grid[r - i][c - k + i];

        set.add(sum);
      }
    }
  }

  let res = Array.from(set);
  res.sort((a, b) => b - a);

  return res.slice(0, 3);
};
