/**
 * @param {number[]} values
 * @return {number}
 */
var minScoreTriangulation = function (values) {
  const n = values.length;
  if (n < 3) return 0;
  const dp = Array.from({ length: n }, () => Array(n).fill(0));

  for (let len = 3; len <= n; len++) {
    for (let i = 0; i + len - 1 < n; i++) {
      const j = i + len - 1;
      let best = Number.MAX_SAFE_INTEGER;
      for (let k = i + 1; k < j; k++) {
        const cost = dp[i][k] + dp[k][j] + values[i] * values[k] * values[j];
        if (cost < best) best = cost;
      }
      dp[i][j] = best;
    }
  }
  return dp[0][n - 1];
};
