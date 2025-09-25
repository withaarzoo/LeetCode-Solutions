/**
 * @param {number[][]} triangle
 * @return {number}
 */
var minimumTotal = function (triangle) {
  const n = triangle.length;
  // copy last row as dp
  let dp = triangle[n - 1].slice();

  // bottom-up
  for (let i = n - 2; i >= 0; --i) {
    for (let j = 0; j <= i; ++j) {
      dp[j] = triangle[i][j] + Math.min(dp[j], dp[j + 1]);
    }
  }
  return dp[0];
};
