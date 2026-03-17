var largestSubmatrix = function (matrix) {
  let m = matrix.length,
    n = matrix[0].length;

  // Step 1: Build heights
  for (let i = 1; i < m; i++) {
    for (let j = 0; j < n; j++) {
      if (matrix[i][j] === 1) {
        matrix[i][j] += matrix[i - 1][j];
      }
    }
  }

  let maxArea = 0;

  // Step 2 & 3
  for (let i = 0; i < m; i++) {
    let row = [...matrix[i]].sort((a, b) => b - a);

    for (let j = 0; j < n; j++) {
      let area = row[j] * (j + 1);
      maxArea = Math.max(maxArea, area);
    }
  }

  return maxArea;
};
