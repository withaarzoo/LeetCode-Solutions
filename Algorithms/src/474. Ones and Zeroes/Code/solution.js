/**
 * @param {string[]} strs
 * @param {number} m
 * @param {number} n
 * @return {number}
 */
var findMaxForm = function (strs, m, n) {
  // dp[z][o] = max strings using at most z zeros and o ones
  const dp = Array.from({ length: m + 1 }, () => Array(n + 1).fill(0));

  for (const s of strs) {
    let z = 0,
      o = 0;
    for (const ch of s) ch === "0" ? z++ : o++;

    // Backwards iteration to enforce 0/1 usage
    for (let i = m; i >= z; i--) {
      for (let j = n; j >= o; j--) {
        dp[i][j] = Math.max(dp[i][j], dp[i - z][j - o] + 1);
      }
    }
  }
  return dp[m][n];
};
