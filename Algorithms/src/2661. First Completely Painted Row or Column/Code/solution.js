/**
 * @param {number[]} arr
 * @param {number[][]} mat
 * @return {number}
 */
var firstCompleteIndex = function (arr, mat) {
  const m = mat.length,
    n = mat[0].length;
  const position = new Map();
  const rowCount = Array(m).fill(0);
  const colCount = Array(n).fill(0);

  // Map matrix values to their positions
  for (let i = 0; i < m; i++) {
    for (let j = 0; j < n; j++) {
      position.set(mat[i][j], [i, j]);
    }
  }

  // Iterate through the array and simulate painting
  for (let i = 0; i < arr.length; i++) {
    const [row, col] = position.get(arr[i]);
    rowCount[row]++;
    colCount[col]++;

    if (rowCount[row] === n || colCount[col] === m) {
      return i;
    }
  }
  return -1;
};
