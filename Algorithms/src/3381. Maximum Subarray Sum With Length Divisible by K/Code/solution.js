/**
 * @param {number[]} nums
 * @param {number} k
 * @return {number}
 */
var maxSubarraySum = function (nums, k) {
  const n = nums.length;
  const INF = 4e18; // large number within JS safe integer range
  const minPref = new Array(k).fill(INF);

  let prefix = 0;
  let ans = -INF;

  // prefix index 0 has sum = 0 and remainder 0
  minPref[0] = 0;

  for (let i = 0; i < n; i++) {
    prefix += nums[i];
    const rem = (i + 1) % k; // prefix index is i+1

    if (minPref[rem] !== INF) {
      ans = Math.max(ans, prefix - minPref[rem]);
    }

    if (prefix < minPref[rem]) {
      minPref[rem] = prefix;
    }
  }

  return ans;
};
