/**
 * @param {number[]} nums
 * @param {number} k
 * @return {number}
 */
var firstStableIndex = function (nums, k) {
  const n = nums.length;

  // suffixMin[i] stores the minimum value from i to the end.
  const suffixMin = new Array(n);

  // The last suffix contains only the last element.
  suffixMin[n - 1] = nums[n - 1];

  // Build suffix minimums from right to left.
  for (let i = n - 2; i >= 0; i--) {
    // Store the smaller value between nums[i]
    // and the minimum of the suffix to its right.
    suffixMin[i] = Math.min(nums[i], suffixMin[i + 1]);
  }

  // Store the largest value seen from index 0 to the current index.
  let prefixMax = nums[0];

  // Check indices from left to right.
  for (let i = 0; i < n; i++) {
    // Update the maximum value in nums[0..i].
    prefixMax = Math.max(prefixMax, nums[i]);

    // Calculate the instability score at index i.
    const instability = prefixMax - suffixMin[i];

    // Return immediately because this is the first valid index.
    if (instability <= k) {
      return i;
    }
  }

  // No stable index exists.
  return -1;
};
