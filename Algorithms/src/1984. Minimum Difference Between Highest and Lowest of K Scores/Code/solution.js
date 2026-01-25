/**
 * @param {number[]} nums
 * @param {number} k
 * @return {number}
 */
var minimumDifference = function (nums, k) {
  // If k is 1, difference is always 0
  if (k === 1) return 0;

  // Step 1: Sort the array
  nums.sort((a, b) => a - b);

  let minDiff = Infinity;

  // Step 2: Sliding window
  for (let i = 0; i + k - 1 < nums.length; i++) {
    let diff = nums[i + k - 1] - nums[i];
    minDiff = Math.min(minDiff, diff);
  }

  return minDiff;
};
