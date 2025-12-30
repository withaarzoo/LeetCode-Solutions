var numMagicSquaresInside = function (grid) {
  let rows = grid.length,
    cols = grid[0].length;
  let count = 0;

  for (let i = 0; i + 2 < rows; i++) {
    for (let j = 0; j + 2 < cols; j++) {
      if (isMagic(grid, i, j)) count++;
    }
  }
  return count;
};

function isMagic(g, r, c) {
  if (g[r + 1][c + 1] !== 5) return false;

  let seen = new Array(10).fill(false);

  for (let i = r; i < r + 3; i++) {
    for (let j = c; j < c + 3; j++) {
      let val = g[i][j];
      if (val < 1 || val > 9 || seen[val]) return false;
      seen[val] = true;
    }
  }

  for (let i = 0; i < 3; i++) {
    if (g[r + i][c] + g[r + i][c + 1] + g[r + i][c + 2] !== 15) return false;
    if (g[r][c + i] + g[r + 1][c + i] + g[r + 2][c + i] !== 15) return false;
  }

  if (g[r][c] + g[r + 1][c + 1] + g[r + 2][c + 2] !== 15) return false;
  if (g[r][c + 2] + g[r + 1][c + 1] + g[r + 2][c] !== 15) return false;

  return true;
}
