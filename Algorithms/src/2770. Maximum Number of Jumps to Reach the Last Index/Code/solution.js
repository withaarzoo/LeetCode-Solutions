/**
 * @param {number[]} nums
 * @param {number} target
 * @return {number}
 */
var maximumJumps = function (nums, target) {
  const n = nums.length;

  // dp[i] = maximum jumps needed to reach index i
  const dp = new Array(n).fill(-1);

  // Starting position
  dp[0] = 0;

  // Try every current index
  for (let i = 0; i < n; i++) {
    // Skip unreachable indices
    if (dp[i] === -1) continue;

    // Try every next index
    for (let j = i + 1; j < n; j++) {
      // Calculate difference
      const diff = nums[j] - nums[i];

      // Check valid jump
      if (diff >= -target && diff <= target) {
        // Update maximum jumps
        dp[j] = Math.max(dp[j], dp[i] + 1);
      }
    }
  }

  // Final answer
  return dp[n - 1];
};
