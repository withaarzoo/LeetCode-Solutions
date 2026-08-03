/**
 * @param {number[]} stoneValue
 * @return {string}
 */
var stoneGameIII = function (stoneValue) {
  const n = stoneValue.length;

  // dp[i] stores the best score difference from index i
  const dp = new Array(n + 1).fill(0);

  // Build DP from right to left
  for (let i = n - 1; i >= 0; i--) {
    dp[i] = -Infinity;
    let sum = 0;

    // Try taking 1, 2 and 3 stones
    for (let j = i; j < Math.min(n, i + 3); j++) {
      // Current collected score
      sum += stoneValue[j];

      // Choose the move with maximum score difference
      dp[i] = Math.max(dp[i], sum - dp[j + 1]);
    }
  }

  // Return the final result
  if (dp[0] > 0) return "Alice";
  if (dp[0] < 0) return "Bob";
  return "Tie";
};
