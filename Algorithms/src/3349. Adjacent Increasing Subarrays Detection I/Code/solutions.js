/**
 * @param {number[]} nums
 * @param {number} k
 * @return {boolean}
 */
var hasIncreasingSubarrays = function (nums, k) {
  const n = nums.length;
  if (2 * k > n) return false;

  // nextInc[i] = number of consecutive increasing adjacent pairs starting at i
  const nextInc = new Array(n).fill(0);
  for (let i = n - 2; i >= 0; --i) {
    if (nums[i] < nums[i + 1]) nextInc[i] = nextInc[i + 1] + 1;
    else nextInc[i] = 0;
  }

  const need = k - 1;
  for (let i = 0; i + 2 * k <= n; ++i) {
    if (nextInc[i] >= need && nextInc[i + k] >= need) return true;
  }
  return false;
};
