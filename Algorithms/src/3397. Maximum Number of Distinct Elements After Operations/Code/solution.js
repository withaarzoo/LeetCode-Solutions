/**
 * @param {number[]} nums
 * @param {number} k
 * @return {number}
 */
var maxDistinctElements = function (nums, k) {
  const n = nums.length;
  const intervals = new Array(n);
  for (let i = 0; i < n; ++i) {
    intervals[i] = [nums[i] - k, nums[i] + k];
  }
  intervals.sort((a, b) => {
    if (a[1] !== b[1]) return a[1] - b[1];
    return a[0] - b[0];
  });

  let lastAssigned = Number.MIN_SAFE_INTEGER / 4;
  let ans = 0;
  for (const [l, r] of intervals) {
    const assigned = Math.max(l, lastAssigned + 1);
    if (assigned <= r) {
      ans++;
      lastAssigned = assigned;
    }
  }
  return ans;
};
