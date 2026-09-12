function maximumWeight(intervals: number[][]): number[] {
  const n = intervals.length; // Number of intervals.
  const K = 4; // At most four intervals may be selected.

  // Each item stores left, right, weight, and original index.
  const a: number[][] = intervals.map((x, i) => [x[0], x[1], x[2], i]);

  // Sort by right endpoint to make predecessor searching possible.
  a.sort((x, y) => x[1] - y[1]);

  // Save the right endpoints for binary search.
  const ends: number[] = a.map((x) => x[1]);

  // A DP state contains the maximum score and its original indices.
  type State = {
    score: number;
    ids: number[];
  };

  // dp[k][i] = best result using exactly k intervals from first i intervals.
  const dp: Array<Array<State | null>> = Array.from({ length: K + 1 }, () =>
    Array<State | null>(n + 1).fill(null),
  );

  // Choosing zero intervals always gives score 0.
  for (let i = 0; i <= n; i++) {
    dp[0][i] = { score: 0, ids: [] };
  }

  // Checks whether candidate a is better than candidate b.
  const better = (a: State | null, b: State | null): boolean => {
    // A valid state beats an unreachable state.
    if (a === null) return false;
    if (b === null) return true;

    // Higher score has priority.
    if (a.score !== b.score) return a.score > b.score;

    // Equal scores are resolved lexicographically.
    const len = Math.min(a.ids.length, b.ids.length);

    for (let i = 0; i < len; i++) {
      if (a.ids[i] !== b.ids[i]) {
        return a.ids[i] < b.ids[i];
      }
    }

    // A shorter prefix is lexicographically smaller.
    return a.ids.length < b.ids.length;
  };

  // Finds the first right endpoint >= target.
  const lowerBound = (length: number, target: number): number => {
    let lo = 0;
    let hi = length;

    // Standard binary search for the first incompatible interval.
    while (lo < hi) {
      const mid = lo + Math.floor((hi - lo) / 2);

      // End >= target is not allowed because touching overlaps.
      if (ends[mid] >= target) {
        hi = mid;
      } else {
        // End < target is compatible.
        lo = mid + 1;
      }
    }

    // This position is also the count of compatible intervals.
    return lo;
  };

  // Process every interval after sorting.
  for (let i = 1; i <= n; i++) {
    const left = a[i - 1][0]; // Current interval's left endpoint.
    const weight = a[i - 1][2]; // Current interval's weight.
    const index = a[i - 1][3]; // Original index before sorting.

    // Find the compatible prefix before the current interval.
    const p = lowerBound(i - 1, left);

    // Try selecting exactly k intervals.
    for (let k = 1; k <= K; k++) {
      // Option 1: skip the current interval.
      dp[k][i] = dp[k][i - 1];

      // Option 2: take the current interval.
      const prev = dp[k - 1][p];

      if (prev !== null) {
        // Copy the old indices so other DP states stay unchanged.
        const ids = [...prev.ids, index];

        // Keep original indices sorted for lexicographical comparison.
        ids.sort((x, y) => x - y);

        // Build the candidate formed by taking this interval.
        const take: State = {
          score: prev.score + weight,
          ids,
        };

        // Keep the better candidate.
        if (better(take, dp[k][i])) {
          dp[k][i] = take;
        }
      }
    }
  }

  // The best result can contain from 1 to 4 intervals.
  let answer: State | null = null;

  // Compare every possible number of chosen intervals.
  for (let k = 1; k <= K; k++) {
    if (better(dp[k][n], answer)) {
      answer = dp[k][n];
    }
  }

  // 'answer' must exist because n >= 1.
  return answer!.ids;
}
