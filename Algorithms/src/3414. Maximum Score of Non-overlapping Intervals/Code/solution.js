/**
 * @param {number[][]} intervals
 * @return {number[]}
 */
var maximumWeight = function (intervals) {
  const n = intervals.length; // Number of intervals.
  const K = 4; // At most four intervals can be selected.

  // Add the original index because sorting would otherwise lose it.
  const a = intervals.map((x, i) => [x[0], x[1], x[2], i]);

  // Sort by right endpoint so compatible intervals become a prefix.
  a.sort((x, y) => x[1] - y[1]);

  // Store all right endpoints for binary search.
  const ends = a.map((x) => x[1]);

  // dp[k][i] stores the best result using exactly k intervals
  // from the first i sorted intervals.
  const dp = Array.from({ length: K + 1 }, () => Array(n + 1).fill(null));

  // Zero intervals always give score 0 and an empty selection.
  for (let i = 0; i <= n; i++) {
    dp[0][i] = { score: 0, ids: [] };
  }

  // Compare two DP candidates.
  const better = (a, b) => {
    // A real candidate beats an empty/unreachable state.
    if (a === null) return false;
    if (b === null) return true;

    // Larger score is always better.
    if (a.score !== b.score) return a.score > b.score;

    // For equal scores, compare original indices lexicographically.
    const len = Math.min(a.ids.length, b.ids.length);

    for (let i = 0; i < len; i++) {
      if (a.ids[i] !== b.ids[i]) {
        return a.ids[i] < b.ids[i];
      }
    }

    // If one array is a prefix of the other, shorter is smaller.
    return a.ids.length < b.ids.length;
  };

  // Find the first end >= target among the first 'length' intervals.
  const lowerBound = (length, target) => {
    let lo = 0;
    let hi = length;

    // Binary search for the first incompatible endpoint.
    while (lo < hi) {
      const mid = lo + Math.floor((hi - lo) / 2);

      // End >= target cannot be used, so search left.
      if (ends[mid] >= target) {
        hi = mid;
      } else {
        // End < target is compatible, so search right.
        lo = mid + 1;
      }
    }

    // This is the number of compatible previous intervals.
    return lo;
  };

  // Process intervals in sorted order.
  for (let i = 1; i <= n; i++) {
    const [left, , weight, originalIndex] = a[i - 1];

    // Find how many earlier intervals are compatible.
    const p = lowerBound(i - 1, left);

    // Try every possible number of selected intervals.
    for (let k = 1; k <= K; k++) {
      // Option 1: skip the current interval.
      dp[k][i] = dp[k][i - 1];

      // Option 2: take it, if k - 1 intervals can be chosen before it.
      if (dp[k - 1][p] !== null) {
        // Copy the old indices because DP states must not be mutated.
        const ids = [...dp[k - 1][p].ids, originalIndex];

        // The final comparison must use sorted original indices.
        ids.sort((x, y) => x - y);

        // Build the candidate created by taking this interval.
        const take = {
          score: dp[k - 1][p].score + weight,
          ids,
        };

        // Keep the better of skipping and taking.
        if (better(take, dp[k][i])) {
          dp[k][i] = take;
        }
      }
    }
  }

  // The answer may contain 1 through 4 intervals.
  let answer = null;

  // Compare every possible count.
  for (let k = 1; k <= K; k++) {
    if (better(dp[k][n], answer)) {
      answer = dp[k][n];
    }
  }

  // Return only the original indices.
  return answer.ids;
};
