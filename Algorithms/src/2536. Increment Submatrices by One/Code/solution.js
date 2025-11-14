/**
 * @param {number} n
 * @param {number[][]} queries
 * @return {number[][]}
 */
var rangeAddQueries = function (n, queries) {
  // diff: (n+1) x (n+1)
  const diff = Array.from({ length: n + 1 }, () => new Array(n + 1).fill(0));

  // Apply each query as 4 updates
  for (const q of queries) {
    const [r1, c1, r2, c2] = q;
    diff[r1][c1] += 1;
    diff[r1][c2 + 1] -= 1;
    diff[r2 + 1][c1] -= 1;
    diff[r2 + 1][c2 + 1] += 1;
  }

  // Build the result with 2D prefix sums
  const res = Array.from({ length: n }, () => new Array(n).fill(0));
  for (let i = 0; i < n; ++i) {
    for (let j = 0; j < n; ++j) {
      const up = i > 0 ? diff[i - 1][j] : 0;
      const left = j > 0 ? diff[i][j - 1] : 0;
      const diag = i > 0 && j > 0 ? diff[i - 1][j - 1] : 0;
      diff[i][j] = diff[i][j] + up + left - diag;
      res[i][j] = diff[i][j];
    }
  }
  return res;
};
