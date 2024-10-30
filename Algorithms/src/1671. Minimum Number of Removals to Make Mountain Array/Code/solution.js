/**
 * @param {number[]} nums
 * @return {number}
 */
var minimumMountainRemovals = function (nums) {
  const n = nums.length;
  const LIS = Array(n).fill(1),
    LDS = Array(n).fill(1);

  // Compute LIS for each index
  for (let i = 0; i < n; ++i) {
    for (let j = 0; j < i; ++j) {
      if (nums[i] > nums[j]) {
        LIS[i] = Math.max(LIS[i], LIS[j] + 1);
      }
    }
  }

  // Compute LDS from each index
  for (let i = n - 1; i >= 0; --i) {
    for (let j = n - 1; j > i; --j) {
      if (nums[i] > nums[j]) {
        LDS[i] = Math.max(LDS[i], LDS[j] + 1);
      }
    }
  }

  let maxMountainLength = 0;

  // Find the maximum mountain length
  for (let i = 1; i < n - 1; ++i) {
    if (LIS[i] > 1 && LDS[i] > 1) {
      // Valid peak
      maxMountainLength = Math.max(maxMountainLength, LIS[i] + LDS[i] - 1);
    }
  }

  return n - maxMountainLength;
};
