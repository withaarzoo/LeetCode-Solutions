/**
 * @param {number} poured
 * @param {number} query_row
 * @param {number} query_glass
 * @return {number}
 */
var champagneTower = function (poured, query_row, query_glass) {
  const dp = Array.from({ length: 101 }, () => Array(101).fill(0));
  dp[0][0] = poured;

  for (let r = 0; r <= query_row; r++) {
    for (let c = 0; c <= r; c++) {
      if (dp[r][c] > 1.0) {
        let overflow = (dp[r][c] - 1.0) / 2.0;

        dp[r + 1][c] += overflow;
        dp[r + 1][c + 1] += overflow;

        dp[r][c] = 1.0;
      }
    }
  }

  return Math.min(1.0, dp[query_row][query_glass]);
};
