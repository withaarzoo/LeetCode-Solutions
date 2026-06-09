/**
 * @param {number[]} nums
 * @param {number} k
 * @return {number}
 */
var maxTotalValue = function (nums, k) {
  // Track global minimum and maximum
  let mn = Infinity;
  let mx = -Infinity;

  // Find minimum and maximum element
  for (const num of nums) {
    mn = Math.min(mn, num);
    mx = Math.max(mx, num);
  }

  // Best subarray value
  const best = mx - mn;

  // Choose it k times
  return best * k;
};
