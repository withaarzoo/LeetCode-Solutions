var reverseSubmatrix = function (grid, x, y, k) {
  // Loop half rows
  for (let i = 0; i < Math.floor(k / 2); i++) {
    let top = x + i;
    let bottom = x + k - 1 - i;

    // Swap columns
    for (let j = 0; j < k; j++) {
      let temp = grid[top][y + j];
      grid[top][y + j] = grid[bottom][y + j];
      grid[bottom][y + j] = temp;
    }
  }
  return grid;
};
