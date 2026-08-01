/**
 * @param {number[]} nums
 * @return {boolean}
 */
var predictTheWinner = function (nums) {
  const n = nums.length;

  // dp[i][j] stores the maximum score difference
  // for subarray [i...j].
  const dp = Array.from({ length: n }, () => Array(n).fill(0));

  // Base case:
  // Only one number remains.
  for (let i = 0; i < n; i++) {
    dp[i][i] = nums[i];
  }

  // Build answers for larger subarrays.
  for (let len = 2; len <= n; len++) {
    for (let i = 0; i + len - 1 < n; i++) {
      const j = i + len - 1;

      // Take the left number.
      const takeLeft = nums[i] - dp[i + 1][j];

      // Take the right number.
      const takeRight = nums[j] - dp[i][j - 1];

      // Store the better choice.
      dp[i][j] = Math.max(takeLeft, takeRight);
    }
  }

  // Player 1 wins or ties.
  return dp[0][n - 1] >= 0;
};
