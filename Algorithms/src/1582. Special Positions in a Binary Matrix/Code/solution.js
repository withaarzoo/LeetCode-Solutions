/**
 * @param {number[][]} mat
 * @return {number}
 */
var numSpecial = function (mat) {
  const m = mat.length;
  const n = mat[0].length;

  const rowCount = new Array(m).fill(0);
  const colCount = new Array(n).fill(0);

  // Count 1s in rows and columns
  for (let i = 0; i < m; i++) {
    for (let j = 0; j < n; j++) {
      if (mat[i][j] === 1) {
        rowCount[i]++;
        colCount[j]++;
      }
    }
  }

  let special = 0;

  // Check special positions
  for (let i = 0; i < m; i++) {
    for (let j = 0; j < n; j++) {
      if (mat[i][j] === 1 && rowCount[i] === 1 && colCount[j] === 1) {
        special++;
      }
    }
  }

  return special;
};
