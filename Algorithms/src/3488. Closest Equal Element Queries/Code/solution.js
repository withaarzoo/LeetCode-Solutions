/**
 * @param {number[]} nums
 * @param {number[]} queries
 * @return {number[]}
 */
var solveQueries = function (nums, queries) {
  const n = nums.length;

  // Store all indices for every value
  const positions = new Map();

  for (let i = 0; i < n; i++) {
    if (!positions.has(nums[i])) {
      positions.set(nums[i], []);
    }
    positions.get(nums[i]).push(i);
  }

  // answer[i] = minimum circular distance for index i
  const answer = new Array(n).fill(-1);

  // Process each group
  for (const pos of positions.values()) {
    const m = pos.length;

    if (m === 1) continue;

    for (let i = 0; i < m; i++) {
      const curr = pos[i];

      const prev = pos[(i - 1 + m) % m];
      const next = pos[(i + 1) % m];

      let distPrev = Math.abs(curr - prev);
      distPrev = Math.min(distPrev, n - distPrev);

      let distNext = Math.abs(curr - next);
      distNext = Math.min(distNext, n - distNext);

      answer[curr] = Math.min(distPrev, distNext);
    }
  }

  // Build final result
  return queries.map((idx) => answer[idx]);
};
