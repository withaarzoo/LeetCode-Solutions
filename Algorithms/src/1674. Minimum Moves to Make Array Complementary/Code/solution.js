/**
 * @param {number[]} nums
 * @param {number} limit
 * @return {number}
 */
var minMoves = function (nums, limit) {
  const n = nums.length;

  // Difference array
  const diff = new Array(2 * limit + 2).fill(0);

  // Process all pairs
  for (let i = 0; i < n / 2; i++) {
    let a = Math.min(nums[i], nums[n - 1 - i]);
    let b = Math.max(nums[i], nums[n - 1 - i]);

    // Range where only 1 move is needed
    diff[a + 1] -= 1;
    diff[b + limit + 1] += 1;

    // Exact sum where 0 moves are needed
    diff[a + b] -= 1;
    diff[a + b + 1] += 1;
  }

  const pairs = Math.floor(n / 2);

  // Initially assume every pair needs 2 moves
  let current = pairs * 2;

  let answer = Number.MAX_SAFE_INTEGER;

  // Build prefix sums
  for (let sum = 2; sum <= 2 * limit; sum++) {
    current += diff[sum];

    answer = Math.min(answer, current);
  }

  return answer;
};
