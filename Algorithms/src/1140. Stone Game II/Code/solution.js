/**
 * @param {number[]} piles - Array representing the number of stones in each pile.
 * @return {number} - The maximum number of stones the first player can collect.
 */
var stoneGameII = function (piles) {
  const n = piles.length;

  // Initialize a DP table with dimensions (n+1) x (n+1) filled with 0s.
  // dp[i][M] represents the maximum number of stones the first player can get
  // starting from pile i with M as the maximum number of piles they can take.
  const dp = Array.from({ length: n + 1 }, () => Array(n + 1).fill(0));

  // Initialize a suffix sum array to store the sum of stones from pile i to the end.
  // suffixSum[i] represents the total number of stones from pile i to the last pile.
  const suffixSum = Array(n + 1).fill(0);

  // Calculate the suffix sums from the last pile to the first pile.
  for (let i = n - 1; i >= 0; --i) {
    suffixSum[i] = suffixSum[i + 1] + piles[i];
  }

  // Fill the DP table starting from the last pile to the first pile.
  for (let i = n - 1; i >= 0; --i) {
    for (let M = 1; M <= n; ++M) {
      // Calculate the best outcome for the current position (i, M)
      // by trying to take X piles (where 1 <= X <= 2 * M)
      for (let X = 1; X <= 2 * M && i + X <= n; ++X) {
        // The first player takes X piles and the second player is left with the remaining piles.
        // dp[i][M] is the maximum stones the first player can collect, which is
        // the total stones from i to the end (suffixSum[i]) minus the best response
        // of the second player (dp[i + X][Math.max(M, X)]).
        dp[i][M] = Math.max(dp[i][M], suffixSum[i] - dp[i + X][Math.max(M, X)]);
      }
    }
  }

  // Return the maximum number of stones the first player can get starting from the first pile with M=1.
  return dp[0][1];
};
