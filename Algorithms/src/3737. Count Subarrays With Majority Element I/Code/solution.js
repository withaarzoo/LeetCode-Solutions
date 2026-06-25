/**
 * @param {number[]} nums
 * @param {number} target
 * @return {number}
 */
var countMajoritySubarrays = function (nums, target) {
  const n = nums.length;
  let ans = 0;

  // Try every possible starting index
  for (let left = 0; left < n; left++) {
    let countTarget = 0;

    // Extend the subarray
    for (let right = left; right < n; right++) {
      // Update target frequency
      if (nums[right] === target) countTarget++;

      // Current subarray length
      const len = right - left + 1;

      // Check majority condition
      if (2 * countTarget > len) ans++;
    }
  }

  return ans;
};
