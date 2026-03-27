var areSimilar = function (mat, k) {
  let m = mat.length;
  let n = mat[0].length;

  k = k % n;

  for (let i = 0; i < m; i++) {
    for (let j = 0; j < n; j++) {
      let newCol;

      if (i % 2 === 0) {
        newCol = (j + k) % n;
      } else {
        newCol = (j - k + n) % n;
      }

      if (mat[i][j] !== mat[i][newCol]) {
        return false;
      }
    }
  }

  return true;
};
