/**
 * @param {number[][]} grid
 * @return {number}
 */
var countServers = function (grid) {
  const rows = grid.length;
  const cols = grid[0].length;
  const rowCount = new Array(rows).fill(0);
  const colCount = new Array(cols).fill(0);

  // First pass: Count servers in each row and column
  for (let i = 0; i < rows; i++) {
    for (let j = 0; j < cols; j++) {
      if (grid[i][j] === 1) {
        rowCount[i]++;
        colCount[j]++;
      }
    }
  }

  // Second pass: Count communicable servers
  let count = 0;
  for (let i = 0; i < rows; i++) {
    for (let j = 0; j < cols; j++) {
      if (grid[i][j] === 1 && (rowCount[i] > 1 || colCount[j] > 1)) {
        count++;
      }
    }
  }
  return count;
};
