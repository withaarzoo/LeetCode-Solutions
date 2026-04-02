/**
 * @param {number[][]} coins
 * @return {number}
 */
var maximumAmount = function (coins) {
  const m = coins.length;
  const n = coins[0].length;
  const NEG = -1e9;

  const dp = Array.from({ length: m }, () =>
    Array.from({ length: n }, () => Array(3).fill(NEG)),
  );

  // Starting cell
  if (coins[0][0] >= 0) {
    dp[0][0][0] = coins[0][0];
  } else {
    dp[0][0][0] = coins[0][0];
    dp[0][0][1] = 0;
  }

  for (let i = 0; i < m; i++) {
    for (let j = 0; j < n; j++) {
      for (let k = 0; k <= 2; k++) {
        if (dp[i][j][k] === NEG) continue;

        // Move Down
        if (i + 1 < m) {
          const val = coins[i + 1][j];

          dp[i + 1][j][k] = Math.max(dp[i + 1][j][k], dp[i][j][k] + val);

          if (val < 0 && k < 2) {
            dp[i + 1][j][k + 1] = Math.max(dp[i + 1][j][k + 1], dp[i][j][k]);
          }
        }

        // Move Right
        if (j + 1 < n) {
          const val = coins[i][j + 1];

          dp[i][j + 1][k] = Math.max(dp[i][j + 1][k], dp[i][j][k] + val);

          if (val < 0 && k < 2) {
            dp[i][j + 1][k + 1] = Math.max(dp[i][j + 1][k + 1], dp[i][j][k]);
          }
        }
      }
    }
  }

  return Math.max(
    dp[m - 1][n - 1][0],
    dp[m - 1][n - 1][1],
    dp[m - 1][n - 1][2],
  );
};
