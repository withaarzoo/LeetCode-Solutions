/**
 * @param {number[]} nums
 * @param {number} k
 * @param {number} numOperations
 * @return {number}
 */
var maxFrequency = function (nums, k, numOperations) {
  const n = nums.length;
  if (n === 0) return 0;
  nums.sort((a, b) => a - b);

  // build frequency map
  const freq = new Map();
  for (const x of nums) freq.set(x, (freq.get(x) || 0) + 1);

  let ans = 1;

  // helper binary searches
  const lowerBound = (arr, target) => {
    let l = 0,
      r = arr.length;
    while (l < r) {
      const mid = (l + r) >> 1;
      if (arr[mid] < target) l = mid + 1;
      else r = mid;
    }
    return l;
  };
  const upperBound = (arr, target) => {
    let l = 0,
      r = arr.length;
    while (l < r) {
      const mid = (l + r) >> 1;
      if (arr[mid] <= target) l = mid + 1;
      else r = mid;
    }
    return l;
  };

  // Case A: existing values as target
  for (const [v, already] of freq.entries()) {
    const lowVal = v - k;
    const highVal = v + k;
    const L = lowerBound(nums, lowVal);
    const R = upperBound(nums, highVal);
    const totalInRange = R - L;
    const need = totalInRange - already;
    const canFix = Math.min(need, numOperations);
    ans = Math.max(ans, already + canFix);
  }

  // Case B: sliding window with 2*k
  let l = 0;
  for (let r = 0; r < n; ++r) {
    while (l <= r && nums[r] - nums[l] > 2 * k) l++;
    const w = r - l + 1;
    ans = Math.max(ans, Math.min(w, numOperations));
  }

  return ans;
};
