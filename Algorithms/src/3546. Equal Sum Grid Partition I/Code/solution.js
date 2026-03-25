/**
 * @param {number[][]} grid
 * @return {boolean}
 */
var canPartitionGrid = function (grid) {
  let m = grid.length,
    n = grid[0].length;

  let total = 0;

  // Step 1: Total sum
  for (let row of grid) {
    for (let val of row) {
      total += val;
    }
  }

  // Step 2: Odd check
  if (total % 2 !== 0) return false;

  let target = total / 2;

  // Step 3: Horizontal cut
  let rowSum = 0;
  for (let i = 0; i < m - 1; i++) {
    for (let j = 0; j < n; j++) {
      rowSum += grid[i][j];
    }
    if (rowSum === target) return true;
  }

  // Step 4: Column sums
  let colSum = new Array(n).fill(0);
  for (let j = 0; j < n; j++) {
    for (let i = 0; i < m; i++) {
      colSum[j] += grid[i][j];
    }
  }

  // Step 5: Vertical cut
  let curr = 0;
  for (let j = 0; j < n - 1; j++) {
    curr += colSum[j];
    if (curr === target) return true;
  }

  return false;
};
