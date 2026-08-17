/**
 * @param {number[]} stoneValue
 * @return {number}
 */
var stoneGameV = function (stoneValue) {
  const n = stoneValue.length;

  // prefix[i] stores the sum of the first i stones.
  // This lets me get any interval sum in O(1).
  const prefix = new Array(n + 1).fill(0);

  for (let i = 0; i < n; i++) {
    // Build prefix sums so sum(l, r) can be calculated instantly.
    prefix[i + 1] = prefix[i] + stoneValue[i];
  }

  // dp[l][r] stores Alice's maximum score for interval [l, r].
  const dp = Array.from({ length: n }, () => new Array(n).fill(0));

  // leftBest[l][r] stores the best score when the left part is kept.
  const leftBest = Array.from({ length: n }, () => new Array(n).fill(0));

  // rightBest[l][r] stores the best score when the right part is kept.
  const rightBest = Array.from({ length: n }, () => new Array(n).fill(0));

  // leftPtr[l] is the last split with leftSum <= rightSum.
  const leftPtr = new Array(n);

  // rightPtr[l] is the first split with leftSum >= rightSum.
  const rightPtr = new Array(n);

  for (let i = 0; i < n; i++) {
    // A single stone has no game score, but it can be a valid side.
    leftBest[i][i] = stoneValue[i];
    rightBest[i][i] = stoneValue[i];

    // Start with no valid split for the left boundary.
    leftPtr[i] = i - 1;

    // The first possible split starts at i.
    rightPtr[i] = i;
  }

  // Process intervals from short to long so all smaller states
  // are already available when I calculate the current state.
  for (let len = 2; len <= n; len++) {
    for (let l = 0; l + len <= n; l++) {
      const r = l + len - 1;

      // Calculate the total sum of [l, r].
      const total = prefix[r + 1] - prefix[l];

      // Move leftPtr while leftSum <= rightSum.
      while (leftPtr[l] + 1 <= r - 1) {
        const k = leftPtr[l] + 1;
        const leftSum = prefix[k + 1] - prefix[l];

        // 2 * leftSum <= total means leftSum <= rightSum.
        if (2 * leftSum > total) {
          break;
        }

        leftPtr[l]++;
      }

      // Move rightPtr until leftSum >= rightSum.
      while (rightPtr[l] <= r - 1) {
        const k = rightPtr[l];
        const leftSum = prefix[k + 1] - prefix[l];

        // Stop at the first split where the left side
        // becomes at least as large as the right side.
        if (2 * leftSum >= total) {
          break;
        }

        rightPtr[l]++;
      }

      let best = 0;

      // If the left side is smaller or equal, Alice keeps it.
      if (leftPtr[l] >= l) {
        best = leftBest[l][leftPtr[l]];
      }

      // If the right side is smaller or equal, Alice keeps it.
      if (rightPtr[l] <= r - 1) {
        best = Math.max(best, rightBest[rightPtr[l] + 1][r]);
      }

      // Store the maximum score for this interval.
      dp[l][r] = best;

      // Use this interval as a possible left side of a future split.
      leftBest[l][r] = Math.max(leftBest[l][r - 1], dp[l][r] + total);

      // Use this interval as a possible right side of a future split.
      rightBest[l][r] = Math.max(rightBest[l + 1][r], dp[l][r] + total);
    }
  }

  // Return the answer for the complete array.
  return dp[0][n - 1];
};
