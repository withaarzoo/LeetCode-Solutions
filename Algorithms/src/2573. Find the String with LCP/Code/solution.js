/**
 * @param {number[][]} lcp
 * @return {string}
 */
var findTheString = function (lcp) {
  const n = lcp.length;

  const group = new Array(n).fill(-1);
  let curGroup = 0;

  // Build groups
  for (let i = 0; i < n; i++) {
    if (group[i] === -1) {
      if (curGroup === 26) return "";

      group[i] = curGroup++;

      for (let j = i + 1; j < n; j++) {
        if (lcp[i][j] > 0) {
          group[j] = group[i];
        }
      }
    }
  }

  // Build answer string
  const chars = new Array(n);
  for (let i = 0; i < n; i++) {
    chars[i] = String.fromCharCode(97 + group[i]);
  }

  const ans = chars.join("");

  // Verify using DP
  const dp = Array.from({ length: n + 1 }, () => new Array(n + 1).fill(0));

  for (let i = n - 1; i >= 0; i--) {
    for (let j = n - 1; j >= 0; j--) {
      if (ans[i] === ans[j]) {
        dp[i][j] = 1 + dp[i + 1][j + 1];
      }
    }
  }

  // Compare matrices
  for (let i = 0; i < n; i++) {
    for (let j = 0; j < n; j++) {
      if (dp[i][j] !== lcp[i][j]) {
        return "";
      }
    }
  }

  return ans;
};
